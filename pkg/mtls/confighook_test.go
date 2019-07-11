/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mtls

import (
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"sofastack.io/sofa-mosn/pkg/api/v2"
	"sofastack.io/sofa-mosn/pkg/mtls/certtool"
)

func pass(resp *http.Response, err error) bool {
	if err != nil {
		return false
	}
	ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return true
}
func fail(resp *http.Response, err error) bool {
	if err != nil {
		return true
	}
	ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return false
}

// test ConfigHooks
// define VerifyPeerCertificate, verify common name instead of san, ignore keyusage
type testConfigHooks struct {
	defaultConfigHooks
	Name           string
	Root           *x509.CertPool
	PassCommonName string
}

// over write
func (hook *testConfigHooks) VerifyPeerCertificate() func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
	return hook.verifyPeerCertificate
}

func (hook *testConfigHooks) GetX509Pool(caIndex string) (*x509.CertPool, error) {
	// usually, the certpool should be created with the caIndex
	root := certtool.GetRootCA()
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM([]byte(root.CertPem))
	hook.Root = pool
	return hook.Root, nil
}

// verifiedChains is always nil
func (hook *testConfigHooks) verifyPeerCertificate(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
	var certs []*x509.Certificate
	for _, asn1Data := range rawCerts {
		cert, err := x509.ParseCertificate(asn1Data)
		if err != nil {
			return err
		}
		certs = append(certs, cert)
	}
	intermediates := x509.NewCertPool()
	for _, cert := range certs[1:] {
		intermediates.AddCert(cert)
	}
	opts := x509.VerifyOptions{
		Roots:         hook.Root,
		Intermediates: intermediates,
	}
	leaf := certs[0]
	_, err := leaf.Verify(opts)
	if err != nil {
		return err
	}
	if leaf.Subject.CommonName != hook.PassCommonName {
		return errors.New("common name miss match")
	}
	return nil
}

const testType = "test"

type testConfigHooksFactory struct{}

func (f *testConfigHooksFactory) CreateConfigHooks(config map[string]interface{}) ConfigHooks {
	c := make(map[string]string)
	for k, v := range config {
		if s, ok := v.(string); ok {
			c[strings.ToLower(k)] = s
		}
	}
	return &testConfigHooks{
		defaultConfigHooks: defaultConfigHooks{},
		Name:               c["name"],
		PassCommonName:     c["cn"],
	}
}

func TestMain(m *testing.M) {
	Register(testType, &testConfigHooksFactory{})
	m.Run()
}

// TestTLSExtensionsVerifyClient tests server allow request with certificate's common name is client only
func TestTLSExtensionsVerifyClient(t *testing.T) {
	// Server
	extendVerify := map[string]interface{}{
		"name": "server",
		"cn":   "client",
	}
	serverInfo := &certInfo{
		CommonName: extendVerify["name"].(string),
		Curve:      "RSA",
	}
	serverConfig, err := serverInfo.CreateCertConfig()
	if err != nil {
		t.Errorf("create server certificate error %v", err)
		return
	}
	serverConfig.RequireClientCert = true
	serverConfig.VerifyClient = true
	serverConfig.Type = testType
	serverConfig.ExtendVerify = extendVerify
	filterChains := []v2.FilterChain{
		{
			TLSContexts: []v2.TLSConfig{
				*serverConfig,
			},
		},
	}
	lc := &v2.Listener{}
	lc.FilterChains = filterChains
	ctxMng, err := NewTLSServerContextManager(lc)
	if err != nil {
		t.Errorf("create context manager failed %v", err)
		return
	}
	server := MockServer{
		Mng: ctxMng,
		t:   t,
	}
	server.GoListenAndServe(t)
	defer server.Close()
	time.Sleep(time.Second) //wait server start
	testCases := []struct {
		Info *certInfo
		Pass func(resp *http.Response, err error) bool
	}{
		{
			Info: &certInfo{
				CommonName: extendVerify["cn"].(string),
				Curve:      serverInfo.Curve,
			},
			Pass: pass,
		},
		{
			Info: &certInfo{
				CommonName: "invalid client",
				Curve:      serverInfo.Curve,
			},
			Pass: fail,
		},
	}
	for i, tc := range testCases {
		cfg, err := tc.Info.CreateCertConfig()
		cfg.ServerName = "127.0.0.1"
		if err != nil {
			t.Errorf("#%d create client certificate error %v", i, err)
			continue
		}
		cltMng, err := NewTLSClientContextManager(cfg)
		if err != nil {
			t.Errorf("#%d create client context manager failed %v", i, err)
			continue
		}

		resp, err := MockClient(t, server.Addr, cltMng)
		if !tc.Pass(resp, err) {
			t.Errorf("#%d verify failed", i)
		}
	}

}

// TestTestTLSExtensionsVerifyServer tests client accept server response with cerificate's common name is server only
func TestTestTLSExtensionsVerifyServer(t *testing.T) {
	extendVerify := map[string]interface{}{
		"name": "client",
		"cn":   "server",
	}
	clientInfo := &certInfo{
		CommonName: extendVerify["name"].(string),
		Curve:      "RSA",
	}

	testCases := []struct {
		Info *certInfo
		Pass func(resp *http.Response, err error) bool
	}{
		{
			Info: &certInfo{
				CommonName: extendVerify["cn"].(string),
				Curve:      clientInfo.Curve,
				DNS:        "www.pass.com",
			},
			Pass: pass,
		},
		{
			Info: &certInfo{
				CommonName: "invalid server",
				Curve:      clientInfo.Curve,
				DNS:        "www.fail.com",
			},
			Pass: fail,
		},
	}
	var filterChains []v2.FilterChain
	for i, tc := range testCases {
		cfg, err := tc.Info.CreateCertConfig()
		if err != nil {
			t.Errorf("#%d %v", i, err)
			return
		}
		fc := v2.FilterChain{
			TLSContexts: []v2.TLSConfig{
				*cfg,
			},
		}
		filterChains = append(filterChains, fc)
	}
	lc := &v2.Listener{}
	lc.FilterChains = filterChains
	ctxMng, err := NewTLSServerContextManager(lc)
	if err != nil {
		t.Errorf("create context manager failed %v", err)
		return
	}
	server := MockServer{
		Mng: ctxMng,
		t:   t,
	}
	server.GoListenAndServe(t)
	defer server.Close()
	time.Sleep(time.Second) //wait server start
	clientConfig, err := clientInfo.CreateCertConfig()
	if err != nil {
		t.Errorf("create client certificate error %v", err)
		return
	}
	clientConfig.Type = testType
	clientConfig.ExtendVerify = extendVerify
	for i, tc := range testCases {
		clientConfig.ServerName = tc.Info.DNS
		cltMng, err := NewTLSClientContextManager(clientConfig)
		if err != nil {
			t.Errorf("create client context manager failed %v", err)
			return
		}

		resp, err := MockClient(t, server.Addr, cltMng)
		if !tc.Pass(resp, err) {
			t.Errorf("#%d verify failed", i)
		}
	}
	// insecure skip will skip even if it is registered
	skipConfig := &v2.TLSConfig{
		Status:       true,
		Type:         clientConfig.Type,
		CACert:       clientConfig.CACert,
		CertChain:    clientConfig.CertChain,
		PrivateKey:   clientConfig.PrivateKey,
		InsecureSkip: true,
	}
	for i, tc := range testCases {
		skipConfig.ServerName = tc.Info.DNS
		skipMng, err := NewTLSClientContextManager(skipConfig)
		if err != nil {
			t.Errorf("create client context manager failed %v", err)
			return
		}
		resp, err := MockClient(t, server.Addr, skipMng)
		// ignore the case, must be pass
		if !pass(resp, err) {
			t.Errorf("#%d skip verify failed", i)
		}
	}

}

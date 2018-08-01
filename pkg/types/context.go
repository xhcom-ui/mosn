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

package types

// ContextKey type
type ContextKey string

// Context key types
const (
	ContextKeyStreamID                   ContextKey = "StreamId"
	ContextKeyConnectionID               ContextKey = "ConnectionId"
	ContextKeyListenerPort               ContextKey = "ListenerPort"
	ContextKeyListenerName               ContextKey = "ListenerName"
	ContextKeyListenerStatsNameSpace     ContextKey = "ListenerStatsNameSpace"
	ContextKeyNetworkFilterChainFactory  ContextKey = "NetworkFilterChainFactory"
	ContextKeyStreamFilterChainFactories ContextKey = "StreamFilterChainFactory"
	ContextKeyConnectionCodecMapPool     ContextKey = "ConnectionCodecMapPool"
	ContextKeyConnectionBytesBufferPool  ContextKey = "ConnectionBytesBufferPool    "
	ContextKeyLogger                     ContextKey = "Logger"
	ContextKeyAccessLogs                 ContextKey = "AccessLogs"
	ContextOriRemoteAddr                 ContextKey = "OriRemoteAddr"
	ContextKeyAcceptChan                 ContextKey = "ContextKeyAcceptChan"
	ContextKeyAcceptBuffer               ContextKey = "ContextKeyAcceptBuffer"
)

const (
	// GlobalStatsNamespace is the stats namesapce
	GlobalStatsNamespace = ""
)

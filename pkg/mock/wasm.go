// Code generated by MockGen. DO NOT EDIT.
// Source: ../types/wasm.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	v2 "mosn.io/mosn/pkg/config/v2"
	types "mosn.io/mosn/pkg/types"
	reflect "reflect"
)

// MockWasmManager is a mock of WasmManager interface.
type MockWasmManager struct {
	ctrl     *gomock.Controller
	recorder *MockWasmManagerMockRecorder
}

// MockWasmManagerMockRecorder is the mock recorder for MockWasmManager.
type MockWasmManagerMockRecorder struct {
	mock *MockWasmManager
}

// NewMockWasmManager creates a new mock instance.
func NewMockWasmManager(ctrl *gomock.Controller) *MockWasmManager {
	mock := &MockWasmManager{ctrl: ctrl}
	mock.recorder = &MockWasmManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmManager) EXPECT() *MockWasmManagerMockRecorder {
	return m.recorder
}

// AddOrUpdateWasm mocks base method.
func (m *MockWasmManager) AddOrUpdateWasm(wasmConfig v2.WasmPluginConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrUpdateWasm", wasmConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrUpdateWasm indicates an expected call of AddOrUpdateWasm.
func (mr *MockWasmManagerMockRecorder) AddOrUpdateWasm(wasmConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateWasm", reflect.TypeOf((*MockWasmManager)(nil).AddOrUpdateWasm), wasmConfig)
}

// GetWasmPluginWrapperByName mocks base method.
func (m *MockWasmManager) GetWasmPluginWrapperByName(pluginName string) types.WasmPluginWrapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWasmPluginWrapperByName", pluginName)
	ret0, _ := ret[0].(types.WasmPluginWrapper)
	return ret0
}

// GetWasmPluginWrapperByName indicates an expected call of GetWasmPluginWrapperByName.
func (mr *MockWasmManagerMockRecorder) GetWasmPluginWrapperByName(pluginName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWasmPluginWrapperByName", reflect.TypeOf((*MockWasmManager)(nil).GetWasmPluginWrapperByName), pluginName)
}

// UninstallWasmPluginByName mocks base method.
func (m *MockWasmManager) UninstallWasmPluginByName(pluginName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallWasmPluginByName", pluginName)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallWasmPluginByName indicates an expected call of UninstallWasmPluginByName.
func (mr *MockWasmManagerMockRecorder) UninstallWasmPluginByName(pluginName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallWasmPluginByName", reflect.TypeOf((*MockWasmManager)(nil).UninstallWasmPluginByName), pluginName)
}

// MockWasmPluginWrapper is a mock of WasmPluginWrapper interface.
type MockWasmPluginWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockWasmPluginWrapperMockRecorder
}

// MockWasmPluginWrapperMockRecorder is the mock recorder for MockWasmPluginWrapper.
type MockWasmPluginWrapperMockRecorder struct {
	mock *MockWasmPluginWrapper
}

// NewMockWasmPluginWrapper creates a new mock instance.
func NewMockWasmPluginWrapper(ctrl *gomock.Controller) *MockWasmPluginWrapper {
	mock := &MockWasmPluginWrapper{ctrl: ctrl}
	mock.recorder = &MockWasmPluginWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmPluginWrapper) EXPECT() *MockWasmPluginWrapperMockRecorder {
	return m.recorder
}

// GetPlugin mocks base method.
func (m *MockWasmPluginWrapper) GetPlugin() types.WasmPlugin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlugin")
	ret0, _ := ret[0].(types.WasmPlugin)
	return ret0
}

// GetPlugin indicates an expected call of GetPlugin.
func (mr *MockWasmPluginWrapperMockRecorder) GetPlugin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlugin", reflect.TypeOf((*MockWasmPluginWrapper)(nil).GetPlugin))
}

// GetConfig mocks base method.
func (m *MockWasmPluginWrapper) GetConfig() v2.WasmPluginConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(v2.WasmPluginConfig)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockWasmPluginWrapperMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockWasmPluginWrapper)(nil).GetConfig))
}

// RegisterPluginHandler mocks base method.
func (m *MockWasmPluginWrapper) RegisterPluginHandler(pluginHandler types.WasmPluginHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterPluginHandler", pluginHandler)
}

// RegisterPluginHandler indicates an expected call of RegisterPluginHandler.
func (mr *MockWasmPluginWrapperMockRecorder) RegisterPluginHandler(pluginHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterPluginHandler", reflect.TypeOf((*MockWasmPluginWrapper)(nil).RegisterPluginHandler), pluginHandler)
}

// Update mocks base method.
func (m *MockWasmPluginWrapper) Update(config v2.WasmPluginConfig, plugin types.WasmPlugin) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", config, plugin)
}

// Update indicates an expected call of Update.
func (mr *MockWasmPluginWrapperMockRecorder) Update(config, plugin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWasmPluginWrapper)(nil).Update), config, plugin)
}

// MockWasmPluginHandler is a mock of WasmPluginHandler interface.
type MockWasmPluginHandler struct {
	ctrl     *gomock.Controller
	recorder *MockWasmPluginHandlerMockRecorder
}

// MockWasmPluginHandlerMockRecorder is the mock recorder for MockWasmPluginHandler.
type MockWasmPluginHandlerMockRecorder struct {
	mock *MockWasmPluginHandler
}

// NewMockWasmPluginHandler creates a new mock instance.
func NewMockWasmPluginHandler(ctrl *gomock.Controller) *MockWasmPluginHandler {
	mock := &MockWasmPluginHandler{ctrl: ctrl}
	mock.recorder = &MockWasmPluginHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmPluginHandler) EXPECT() *MockWasmPluginHandlerMockRecorder {
	return m.recorder
}

// OnConfigUpdate mocks base method.
func (m *MockWasmPluginHandler) OnConfigUpdate(config v2.WasmPluginConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnConfigUpdate", config)
}

// OnConfigUpdate indicates an expected call of OnConfigUpdate.
func (mr *MockWasmPluginHandlerMockRecorder) OnConfigUpdate(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnConfigUpdate", reflect.TypeOf((*MockWasmPluginHandler)(nil).OnConfigUpdate), config)
}

// OnPluginStart mocks base method.
func (m *MockWasmPluginHandler) OnPluginStart(plugin types.WasmPlugin) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnPluginStart", plugin)
}

// OnPluginStart indicates an expected call of OnPluginStart.
func (mr *MockWasmPluginHandlerMockRecorder) OnPluginStart(plugin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPluginStart", reflect.TypeOf((*MockWasmPluginHandler)(nil).OnPluginStart), plugin)
}

// OnPluginDestroy mocks base method.
func (m *MockWasmPluginHandler) OnPluginDestroy(plugin types.WasmPlugin) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnPluginDestroy", plugin)
}

// OnPluginDestroy indicates an expected call of OnPluginDestroy.
func (mr *MockWasmPluginHandlerMockRecorder) OnPluginDestroy(plugin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPluginDestroy", reflect.TypeOf((*MockWasmPluginHandler)(nil).OnPluginDestroy), plugin)
}

// MockWasmPlugin is a mock of WasmPlugin interface.
type MockWasmPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockWasmPluginMockRecorder
}

// MockWasmPluginMockRecorder is the mock recorder for MockWasmPlugin.
type MockWasmPluginMockRecorder struct {
	mock *MockWasmPlugin
}

// NewMockWasmPlugin creates a new mock instance.
func NewMockWasmPlugin(ctrl *gomock.Controller) *MockWasmPlugin {
	mock := &MockWasmPlugin{ctrl: ctrl}
	mock.recorder = &MockWasmPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmPlugin) EXPECT() *MockWasmPluginMockRecorder {
	return m.recorder
}

// PluginName mocks base method.
func (m *MockWasmPlugin) PluginName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PluginName")
	ret0, _ := ret[0].(string)
	return ret0
}

// PluginName indicates an expected call of PluginName.
func (mr *MockWasmPluginMockRecorder) PluginName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PluginName", reflect.TypeOf((*MockWasmPlugin)(nil).PluginName))
}

// GetPluginConfig mocks base method.
func (m *MockWasmPlugin) GetPluginConfig() v2.WasmPluginConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPluginConfig")
	ret0, _ := ret[0].(v2.WasmPluginConfig)
	return ret0
}

// GetPluginConfig indicates an expected call of GetPluginConfig.
func (mr *MockWasmPluginMockRecorder) GetPluginConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginConfig", reflect.TypeOf((*MockWasmPlugin)(nil).GetPluginConfig))
}

// GetVmConfig mocks base method.
func (m *MockWasmPlugin) GetVmConfig() v2.WasmVmConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVmConfig")
	ret0, _ := ret[0].(v2.WasmVmConfig)
	return ret0
}

// GetVmConfig indicates an expected call of GetVmConfig.
func (mr *MockWasmPluginMockRecorder) GetVmConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVmConfig", reflect.TypeOf((*MockWasmPlugin)(nil).GetVmConfig))
}

// EnsureInstanceNum mocks base method.
func (m *MockWasmPlugin) EnsureInstanceNum(num int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureInstanceNum", num)
	ret0, _ := ret[0].(int)
	return ret0
}

// EnsureInstanceNum indicates an expected call of EnsureInstanceNum.
func (mr *MockWasmPluginMockRecorder) EnsureInstanceNum(num interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureInstanceNum", reflect.TypeOf((*MockWasmPlugin)(nil).EnsureInstanceNum), num)
}

// InstanceNum mocks base method.
func (m *MockWasmPlugin) InstanceNum() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceNum")
	ret0, _ := ret[0].(int)
	return ret0
}

// InstanceNum indicates an expected call of InstanceNum.
func (mr *MockWasmPluginMockRecorder) InstanceNum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceNum", reflect.TypeOf((*MockWasmPlugin)(nil).InstanceNum))
}

// GetInstance mocks base method.
func (m *MockWasmPlugin) GetInstance() types.WasmInstanceWrapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstance")
	ret0, _ := ret[0].(types.WasmInstanceWrapper)
	return ret0
}

// GetInstance indicates an expected call of GetInstance.
func (mr *MockWasmPluginMockRecorder) GetInstance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstance", reflect.TypeOf((*MockWasmPlugin)(nil).GetInstance))
}

// ReleaseInstance mocks base method.
func (m *MockWasmPlugin) ReleaseInstance(instanceWrapper types.WasmInstanceWrapper) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReleaseInstance", instanceWrapper)
}

// ReleaseInstance indicates an expected call of ReleaseInstance.
func (mr *MockWasmPluginMockRecorder) ReleaseInstance(instanceWrapper interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseInstance", reflect.TypeOf((*MockWasmPlugin)(nil).ReleaseInstance), instanceWrapper)
}

// Exec mocks base method.
func (m *MockWasmPlugin) Exec(f func(types.WasmInstanceWrapper) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Exec", f)
}

// Exec indicates an expected call of Exec.
func (mr *MockWasmPluginMockRecorder) Exec(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockWasmPlugin)(nil).Exec), f)
}

// SetCpuLimit mocks base method.
func (m *MockWasmPlugin) SetCpuLimit(cpu int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCpuLimit", cpu)
}

// SetCpuLimit indicates an expected call of SetCpuLimit.
func (mr *MockWasmPluginMockRecorder) SetCpuLimit(cpu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCpuLimit", reflect.TypeOf((*MockWasmPlugin)(nil).SetCpuLimit), cpu)
}

// SetMemLimit mocks base method.
func (m *MockWasmPlugin) SetMemLimit(mem int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMemLimit", mem)
}

// SetMemLimit indicates an expected call of SetMemLimit.
func (mr *MockWasmPluginMockRecorder) SetMemLimit(mem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMemLimit", reflect.TypeOf((*MockWasmPlugin)(nil).SetMemLimit), mem)
}

// Clear mocks base method.
func (m *MockWasmPlugin) Clear() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Clear")
}

// Clear indicates an expected call of Clear.
func (mr *MockWasmPluginMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockWasmPlugin)(nil).Clear))
}

// MockWasmInstanceWrapper is a mock of WasmInstanceWrapper interface.
type MockWasmInstanceWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockWasmInstanceWrapperMockRecorder
}

// MockWasmInstanceWrapperMockRecorder is the mock recorder for MockWasmInstanceWrapper.
type MockWasmInstanceWrapperMockRecorder struct {
	mock *MockWasmInstanceWrapper
}

// NewMockWasmInstanceWrapper creates a new mock instance.
func NewMockWasmInstanceWrapper(ctrl *gomock.Controller) *MockWasmInstanceWrapper {
	mock := &MockWasmInstanceWrapper{ctrl: ctrl}
	mock.recorder = &MockWasmInstanceWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmInstanceWrapper) EXPECT() *MockWasmInstanceWrapperMockRecorder {
	return m.recorder
}

// GetExportsFunc mocks base method.
func (m *MockWasmInstanceWrapper) GetExportsFunc(funcName string) (types.WasmFunction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExportsFunc", funcName)
	ret0, _ := ret[0].(types.WasmFunction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExportsFunc indicates an expected call of GetExportsFunc.
func (mr *MockWasmInstanceWrapperMockRecorder) GetExportsFunc(funcName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExportsFunc", reflect.TypeOf((*MockWasmInstanceWrapper)(nil).GetExportsFunc), funcName)
}

// GetExportsMem mocks base method.
func (m *MockWasmInstanceWrapper) GetExportsMem(memName string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExportsMem", memName)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExportsMem indicates an expected call of GetExportsMem.
func (mr *MockWasmInstanceWrapperMockRecorder) GetExportsMem(memName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExportsMem", reflect.TypeOf((*MockWasmInstanceWrapper)(nil).GetExportsMem), memName)
}

// GetMemory mocks base method.
func (m *MockWasmInstanceWrapper) GetMemory(addr, size uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemory", addr, size)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemory indicates an expected call of GetMemory.
func (mr *MockWasmInstanceWrapperMockRecorder) GetMemory(addr, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemory", reflect.TypeOf((*MockWasmInstanceWrapper)(nil).GetMemory), addr, size)
}

// PutMemory mocks base method.
func (m *MockWasmInstanceWrapper) PutMemory(addr, size uint64, content []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutMemory", addr, size, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutMemory indicates an expected call of PutMemory.
func (mr *MockWasmInstanceWrapperMockRecorder) PutMemory(addr, size, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutMemory", reflect.TypeOf((*MockWasmInstanceWrapper)(nil).PutMemory), addr, size, content)
}

// GetByte mocks base method.
func (m *MockWasmInstanceWrapper) GetByte(addr uint64) (byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByte", addr)
	ret0, _ := ret[0].(byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByte indicates an expected call of GetByte.
func (mr *MockWasmInstanceWrapperMockRecorder) GetByte(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByte", reflect.TypeOf((*MockWasmInstanceWrapper)(nil).GetByte), addr)
}

// PutByte mocks base method.
func (m *MockWasmInstanceWrapper) PutByte(addr uint64, b byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutByte", addr, b)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutByte indicates an expected call of PutByte.
func (mr *MockWasmInstanceWrapperMockRecorder) PutByte(addr, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutByte", reflect.TypeOf((*MockWasmInstanceWrapper)(nil).PutByte), addr, b)
}

// GetUint32 mocks base method.
func (m *MockWasmInstanceWrapper) GetUint32(addr uint64) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUint32", addr)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUint32 indicates an expected call of GetUint32.
func (mr *MockWasmInstanceWrapperMockRecorder) GetUint32(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUint32", reflect.TypeOf((*MockWasmInstanceWrapper)(nil).GetUint32), addr)
}

// PutUint32 mocks base method.
func (m *MockWasmInstanceWrapper) PutUint32(addr uint64, value uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutUint32", addr, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutUint32 indicates an expected call of PutUint32.
func (mr *MockWasmInstanceWrapperMockRecorder) PutUint32(addr, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutUint32", reflect.TypeOf((*MockWasmInstanceWrapper)(nil).PutUint32), addr, value)
}

// Malloc mocks base method.
func (m *MockWasmInstanceWrapper) Malloc(size int32) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Malloc", size)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Malloc indicates an expected call of Malloc.
func (mr *MockWasmInstanceWrapperMockRecorder) Malloc(size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Malloc", reflect.TypeOf((*MockWasmInstanceWrapper)(nil).Malloc), size)
}

// RegisterFunc mocks base method.
func (m *MockWasmInstanceWrapper) RegisterFunc(namespace, funcName string, f interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterFunc", namespace, funcName, f)
}

// RegisterFunc indicates an expected call of RegisterFunc.
func (mr *MockWasmInstanceWrapperMockRecorder) RegisterFunc(namespace, funcName, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterFunc", reflect.TypeOf((*MockWasmInstanceWrapper)(nil).RegisterFunc), namespace, funcName, f)
}

// GetData mocks base method.
func (m *MockWasmInstanceWrapper) GetData() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetData indicates an expected call of GetData.
func (mr *MockWasmInstanceWrapperMockRecorder) GetData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockWasmInstanceWrapper)(nil).GetData))
}

// SetData mocks base method.
func (m *MockWasmInstanceWrapper) SetData(data interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetData", data)
}

// SetData indicates an expected call of SetData.
func (mr *MockWasmInstanceWrapperMockRecorder) SetData(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetData", reflect.TypeOf((*MockWasmInstanceWrapper)(nil).SetData), data)
}

// Acquire mocks base method.
func (m *MockWasmInstanceWrapper) Acquire(data interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Acquire", data)
}

// Acquire indicates an expected call of Acquire.
func (mr *MockWasmInstanceWrapperMockRecorder) Acquire(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Acquire", reflect.TypeOf((*MockWasmInstanceWrapper)(nil).Acquire), data)
}

// Release mocks base method.
func (m *MockWasmInstanceWrapper) Release() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Release")
}

// Release indicates an expected call of Release.
func (mr *MockWasmInstanceWrapperMockRecorder) Release() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockWasmInstanceWrapper)(nil).Release))
}

// MockWasmVM is a mock of WasmVM interface.
type MockWasmVM struct {
	ctrl     *gomock.Controller
	recorder *MockWasmVMMockRecorder
}

// MockWasmVMMockRecorder is the mock recorder for MockWasmVM.
type MockWasmVMMockRecorder struct {
	mock *MockWasmVM
}

// NewMockWasmVM creates a new mock instance.
func NewMockWasmVM(ctrl *gomock.Controller) *MockWasmVM {
	mock := &MockWasmVM{ctrl: ctrl}
	mock.recorder = &MockWasmVMMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmVM) EXPECT() *MockWasmVMMockRecorder {
	return m.recorder
}

// Init mocks base method.
func (m *MockWasmVM) Init() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init")
}

// Init indicates an expected call of Init.
func (mr *MockWasmVMMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockWasmVM)(nil).Init))
}

// NewModule mocks base method.
func (m *MockWasmVM) NewModule(wasmBytes []byte) types.WasmModule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewModule", wasmBytes)
	ret0, _ := ret[0].(types.WasmModule)
	return ret0
}

// NewModule indicates an expected call of NewModule.
func (mr *MockWasmVMMockRecorder) NewModule(wasmBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewModule", reflect.TypeOf((*MockWasmVM)(nil).NewModule), wasmBytes)
}

// MockWasmModule is a mock of WasmModule interface.
type MockWasmModule struct {
	ctrl     *gomock.Controller
	recorder *MockWasmModuleMockRecorder
}

// MockWasmModuleMockRecorder is the mock recorder for MockWasmModule.
type MockWasmModuleMockRecorder struct {
	mock *MockWasmModule
}

// NewMockWasmModule creates a new mock instance.
func NewMockWasmModule(ctrl *gomock.Controller) *MockWasmModule {
	mock := &MockWasmModule{ctrl: ctrl}
	mock.recorder = &MockWasmModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmModule) EXPECT() *MockWasmModuleMockRecorder {
	return m.recorder
}

// Init mocks base method.
func (m *MockWasmModule) Init() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init")
}

// Init indicates an expected call of Init.
func (mr *MockWasmModuleMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockWasmModule)(nil).Init))
}

// NewInstance mocks base method.
func (m *MockWasmModule) NewInstance() types.WasmInstance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewInstance")
	ret0, _ := ret[0].(types.WasmInstance)
	return ret0
}

// NewInstance indicates an expected call of NewInstance.
func (mr *MockWasmModuleMockRecorder) NewInstance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewInstance", reflect.TypeOf((*MockWasmModule)(nil).NewInstance))
}

// MockWasmInstance is a mock of WasmInstance interface.
type MockWasmInstance struct {
	ctrl     *gomock.Controller
	recorder *MockWasmInstanceMockRecorder
}

// MockWasmInstanceMockRecorder is the mock recorder for MockWasmInstance.
type MockWasmInstanceMockRecorder struct {
	mock *MockWasmInstance
}

// NewMockWasmInstance creates a new mock instance.
func NewMockWasmInstance(ctrl *gomock.Controller) *MockWasmInstance {
	mock := &MockWasmInstance{ctrl: ctrl}
	mock.recorder = &MockWasmInstanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmInstance) EXPECT() *MockWasmInstanceMockRecorder {
	return m.recorder
}

// GetExportsFunc mocks base method.
func (m *MockWasmInstance) GetExportsFunc(funcName string) (types.WasmFunction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExportsFunc", funcName)
	ret0, _ := ret[0].(types.WasmFunction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExportsFunc indicates an expected call of GetExportsFunc.
func (mr *MockWasmInstanceMockRecorder) GetExportsFunc(funcName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExportsFunc", reflect.TypeOf((*MockWasmInstance)(nil).GetExportsFunc), funcName)
}

// GetExportsMem mocks base method.
func (m *MockWasmInstance) GetExportsMem(memName string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExportsMem", memName)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExportsMem indicates an expected call of GetExportsMem.
func (mr *MockWasmInstanceMockRecorder) GetExportsMem(memName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExportsMem", reflect.TypeOf((*MockWasmInstance)(nil).GetExportsMem), memName)
}

// GetMemory mocks base method.
func (m *MockWasmInstance) GetMemory(addr, size uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemory", addr, size)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemory indicates an expected call of GetMemory.
func (mr *MockWasmInstanceMockRecorder) GetMemory(addr, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemory", reflect.TypeOf((*MockWasmInstance)(nil).GetMemory), addr, size)
}

// PutMemory mocks base method.
func (m *MockWasmInstance) PutMemory(addr, size uint64, content []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutMemory", addr, size, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutMemory indicates an expected call of PutMemory.
func (mr *MockWasmInstanceMockRecorder) PutMemory(addr, size, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutMemory", reflect.TypeOf((*MockWasmInstance)(nil).PutMemory), addr, size, content)
}

// GetByte mocks base method.
func (m *MockWasmInstance) GetByte(addr uint64) (byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByte", addr)
	ret0, _ := ret[0].(byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByte indicates an expected call of GetByte.
func (mr *MockWasmInstanceMockRecorder) GetByte(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByte", reflect.TypeOf((*MockWasmInstance)(nil).GetByte), addr)
}

// PutByte mocks base method.
func (m *MockWasmInstance) PutByte(addr uint64, b byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutByte", addr, b)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutByte indicates an expected call of PutByte.
func (mr *MockWasmInstanceMockRecorder) PutByte(addr, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutByte", reflect.TypeOf((*MockWasmInstance)(nil).PutByte), addr, b)
}

// GetUint32 mocks base method.
func (m *MockWasmInstance) GetUint32(addr uint64) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUint32", addr)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUint32 indicates an expected call of GetUint32.
func (mr *MockWasmInstanceMockRecorder) GetUint32(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUint32", reflect.TypeOf((*MockWasmInstance)(nil).GetUint32), addr)
}

// PutUint32 mocks base method.
func (m *MockWasmInstance) PutUint32(addr uint64, value uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutUint32", addr, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutUint32 indicates an expected call of PutUint32.
func (mr *MockWasmInstanceMockRecorder) PutUint32(addr, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutUint32", reflect.TypeOf((*MockWasmInstance)(nil).PutUint32), addr, value)
}

// Malloc mocks base method.
func (m *MockWasmInstance) Malloc(size int32) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Malloc", size)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Malloc indicates an expected call of Malloc.
func (mr *MockWasmInstanceMockRecorder) Malloc(size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Malloc", reflect.TypeOf((*MockWasmInstance)(nil).Malloc), size)
}

// RegisterFunc mocks base method.
func (m *MockWasmInstance) RegisterFunc(namespace, funcName string, f interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterFunc", namespace, funcName, f)
}

// RegisterFunc indicates an expected call of RegisterFunc.
func (mr *MockWasmInstanceMockRecorder) RegisterFunc(namespace, funcName, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterFunc", reflect.TypeOf((*MockWasmInstance)(nil).RegisterFunc), namespace, funcName, f)
}

// GetData mocks base method.
func (m *MockWasmInstance) GetData() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetData indicates an expected call of GetData.
func (mr *MockWasmInstanceMockRecorder) GetData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockWasmInstance)(nil).GetData))
}

// SetData mocks base method.
func (m *MockWasmInstance) SetData(data interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetData", data)
}

// SetData indicates an expected call of SetData.
func (mr *MockWasmInstanceMockRecorder) SetData(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetData", reflect.TypeOf((*MockWasmInstance)(nil).SetData), data)
}

// MockWasmFunction is a mock of WasmFunction interface.
type MockWasmFunction struct {
	ctrl     *gomock.Controller
	recorder *MockWasmFunctionMockRecorder
}

// MockWasmFunctionMockRecorder is the mock recorder for MockWasmFunction.
type MockWasmFunctionMockRecorder struct {
	mock *MockWasmFunction
}

// NewMockWasmFunction creates a new mock instance.
func NewMockWasmFunction(ctrl *gomock.Controller) *MockWasmFunction {
	mock := &MockWasmFunction{ctrl: ctrl}
	mock.recorder = &MockWasmFunctionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWasmFunction) EXPECT() *MockWasmFunctionMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockWasmFunction) Call(args ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Call", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call.
func (mr *MockWasmFunctionMockRecorder) Call(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockWasmFunction)(nil).Call), args...)
}

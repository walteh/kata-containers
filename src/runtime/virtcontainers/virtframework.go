//
// Copyright (c) 2023 Apple Inc.
//
// SPDX-License-Identifier: Apache-2.0
//

//go:build darwin
// +build darwin

package virtcontainers

import (
	"context"
	"fmt"
	"runtime"

	hv "github.com/kata-containers/kata-containers/src/runtime/pkg/hypervisors"
	"github.com/kata-containers/kata-containers/src/runtime/virtcontainers/types"
)

// virtFramework is a Hypervisor interface implementation for Darwin Virtualization.framework.
type virtFramework struct{}

func unimplemented() error {
	callerInfo, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(callerInfo).Name()
	return fmt.Errorf("virtframework hypervisor: %s is not implemented yet", funcName)
}

func (vfw *virtFramework) CreateVM(ctx context.Context, id string, network Network, hypervisorConfig *HypervisorConfig) error {
	return unimplemented()
}

func (vfw *virtFramework) StartVM(ctx context.Context, timeout int) error {
	return unimplemented()
}

// If wait is set, don't actively stop the sandbox:
// just perform cleanup.
func (vfw *virtFramework) StopVM(ctx context.Context, waitOnly bool) error {
	return unimplemented()
}

func (vfw *virtFramework) PauseVM(ctx context.Context) error {
	return unimplemented()
}

func (vfw *virtFramework) SaveVM() error {
	return unimplemented()
}

func (vfw *virtFramework) ResumeVM(ctx context.Context) error {
	return unimplemented()
}

func (vfw *virtFramework) AddDevice(ctx context.Context, devInfo interface{}, devType DeviceType) error {
	return unimplemented()
}

func (vfw *virtFramework) HotplugAddDevice(ctx context.Context, devInfo interface{}, devType DeviceType) (interface{}, error) {
	return nil, unimplemented()
}

func (vfw *virtFramework) HotplugRemoveDevice(ctx context.Context, devInfo interface{}, devType DeviceType) (interface{}, error) {
	return nil, unimplemented()
}

func (vfw *virtFramework) ResizeMemory(ctx context.Context, memMB uint32, memoryBlockSizeMB uint32, probe bool) (uint32, MemoryDevice, error) {
	return 0, MemoryDevice{}, unimplemented()
}

func (vfw *virtFramework) ResizeVCPUs(ctx context.Context, vcpus uint32) (uint32, uint32, error) {
	return 0, 0, unimplemented()
}

func (vfw *virtFramework) GetVMConsole(ctx context.Context, sandboxID string) (string, string, error) {
	return "", "", unimplemented()
}

func (vfw *virtFramework) Disconnect(ctx context.Context) {
	panic(unimplemented())
}

func (vfw *virtFramework) Capabilities(ctx context.Context) types.Capabilities {
	panic(unimplemented())
	return types.Capabilities{}
}

func (vfw *virtFramework) HypervisorConfig() HypervisorConfig {
	panic(unimplemented())
	return HypervisorConfig{}
}

func (vfw *virtFramework) GetThreadIDs(ctx context.Context) (VcpuThreadIDs, error) {
	var vcpuInfo VcpuThreadIDs

	vcpuInfo.vcpus = make(map[int]int)

	panic(unimplemented())
	return vcpuInfo, nil
}

func (vfw *virtFramework) Cleanup(ctx context.Context) error {
	panic(unimplemented())
}

func (vfw *virtFramework) GetTotalMemoryMB(ctx context.Context) uint32 {
	panic(unimplemented())
	return 0
}

func (vfw *virtFramework) setConfig(config *HypervisorConfig) error {
	panic(unimplemented())
}

func (vfw *virtFramework) GetPids() []int {
	panic(unimplemented())
}

func (vfw *virtFramework) GetVirtioFsPid() *int {
	panic(unimplemented())
}

func (vfw *virtFramework) fromGrpc(ctx context.Context, hypervisorConfig *HypervisorConfig, j []byte) error {
	panic(unimplemented())
}

func (vfw *virtFramework) toGrpc(ctx context.Context) ([]byte, error) {
	panic(unimplemented())
}

func (vfw *virtFramework) Check() error {
	panic(unimplemented())
}

func (vfw *virtFramework) Save() hv.HypervisorState {
	panic(unimplemented())
}

func (vfw *virtFramework) Load(hv.HypervisorState) {
	panic(unimplemented())
}

func (vfw *virtFramework) GenerateSocket(id string) (interface{}, error) {
	return nil, unimplemented()
}

func (vfw *virtFramework) IsRateLimiterBuiltin() bool {
	panic(unimplemented())
	return false
}

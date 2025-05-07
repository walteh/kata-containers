// Copyright (c) 2021 Apple Inc.
//
// SPDX-License-Identifier: Apache-2.0
//

package virtcontainers

import (
	"github.com/kata-containers/kata-containers/src/runtime/virtcontainers/types"
	"github.com/kata-containers/kata-containers/src/runtime/virtcontainers/utils"
)

func generateVMSocket(id string, vmStogarePath string) (interface{}, error) {
	vhostFd, contextID, err := utils.FindContextID()
	if err != nil {
		return nil, err
	}

	return types.VSock{
		VhostFd:   vhostFd,
		ContextID: contextID,
		Port:      uint32(vSockPort),
	}, nil
}

func init() {
	RegisterHypervisor(QemuHypervisor, func() Hypervisor { return &qemu{} })
	RegisterHypervisor(FirecrackerHypervisor, func() Hypervisor { return &firecracker{} })
	RegisterHypervisor(ClhHypervisor, func() Hypervisor { return &cloudHypervisor{} })
	RegisterHypervisor(StratovirtHypervisor, func() Hypervisor { return &stratovirt{} })
	RegisterHypervisor(DragonballHypervisor, func() Hypervisor { return &mockHypervisor{} })
	RegisterHypervisor(RemoteHypervisor, func() Hypervisor { return &remoteHypervisor{} })
	RegisterHypervisor(MockHypervisor, func() Hypervisor { return &mockHypervisor{} })
}

// Copyright (c) 2023 Apple Inc.
//
// SPDX-License-Identifier: Apache-2.0
//

package virtcontainers

const (
	QemuCCWVirtio = "qemu-ccw-virtio"
)

func init() {
	RegisterHypervisor(MockHypervisor, func() Hypervisor { return &mockHypervisor{} })
}

// NewHypervisor returns a hypervisor from a hypervisor type.

func availableGuestProtection() (guestProtection, error) {
	return noneProtection, nil
}

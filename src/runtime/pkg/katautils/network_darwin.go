// Copyright (c) 2022 Apple Inc.
//
// SPDX-License-Identifier: Apache-2.0
//

package katautils

import (
	vc "github.com/kata-containers/kata-containers/src/runtime/virtcontainers"
)

func EnterNetNS(networkID string, cb func() error) error {
	return nil
}

func SetupNetworkNamespace(config *vc.NetworkConfig) error {
	return nil
}

func cleanupNetNS(netNSPath string) error {
	return nil
}

const (
	netNsMountType    = "nsfs"
	mountTypeFieldIdx = 8
	mountDestIdx      = 4
)

// getNetNsFromBindMount returns the network namespace for the bind-mounted path
func getNetNsFromBindMount(nsPath string, procMountFile string) (string, error) {
	return "", nil
}

// hostNetworkingRequested checks if the network namespace requested is the
// same as the current process.
func hostNetworkingRequested(configNetNs string) (bool, error) {
	return false, nil
}

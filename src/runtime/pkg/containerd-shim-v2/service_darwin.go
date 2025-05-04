//go:build darwin

package containerdshim

// setupMntNs is a stub implementation for Darwin
// Darwin doesn't support Linux-specific mount namespaces and mount flags
// This function exists to provide API compatibility with the Linux implementation
func setupMntNs() error {
	// No-op implementation for Darwin
	// Darwin doesn't support mount namespaces like Linux
	return nil
}

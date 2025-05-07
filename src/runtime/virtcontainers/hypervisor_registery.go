package virtcontainers

import (
	"sync"

	"gitlab.com/tozd/go/errors"
)

var registeredHypervisorsMutex sync.Mutex
var registeredHypervisors = map[HypervisorType]func() Hypervisor{}

func RegisterHypervisor(hType HypervisorType, hypervisor func() Hypervisor) {
	registeredHypervisorsMutex.Lock()
	defer registeredHypervisorsMutex.Unlock()

	registeredHypervisors[hType] = hypervisor
}

func NewHypervisor(hType HypervisorType) (Hypervisor, error) {
	registeredHypervisorsMutex.Lock()
	defer registeredHypervisorsMutex.Unlock()

	hypervisor, ok := registeredHypervisors[hType]
	if !ok {
		return nil, errors.Errorf("Unknown hypervisor type %s", hType)
	}

	return hypervisor(), nil
}

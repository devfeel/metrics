package metrics

import (
	"sync"
)

// Registry
type Registry interface {

	// GetCounter get counter by the given name or nil if none is registered.
	GetCounter(string) Counter

	// UnregisterCounter delete the counter with the given name.
	UnregisterCounter(string)

	// GetOrRegister Gets an existing metric or registers the given one.
	GetOrRegister(string, interface{}) interface{}

	// Register the given metric under the given name.
	Register(string, interface{}) error

}


// StandardRegistry The standard implementation of a Registry
type StandardRegistry struct {
	counters *sync.Map
}

// Create a new registry.
func NewRegistry() Registry {
	return &StandardRegistry{counters: new(sync.Map)}
}

// GetCounter get counter by the given name or nil if none is registered.
func (r *StandardRegistry) GetCounter(name string) Counter {
	counter, ok := r.counters.Load(name)
	if !ok{
		return nil
	}else{
		return counter.(Counter)
	}
}

// UnregisterCounter delete the counter with the given name.
func (r *StandardRegistry) UnregisterCounter(name string) {
	r.counters.Delete(name)
}

// GetOrRegister Gets an existing metric or creates and registers a new one.
func (r *StandardRegistry) GetOrRegister(name string, i interface{}) interface{} {
	switch i.(type) {
		case Counter:
			if metric, ok := r.counters.Load(name); ok {
				return metric
			}else{
				r.register(name, i)
			}
	}
	return i
}

// Register the given metric under the given name.
func (r *StandardRegistry) Register(name string, i interface{}) error {
	return r.register(name, i)
}


// register register metric, if exists same name's metrics, replace it
func (r *StandardRegistry) register(name string, i interface{}) error {
	switch i.(type) {
	case Counter:
		r.counters.Store(name, i)
	}
	return nil
}
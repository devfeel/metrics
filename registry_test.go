package metrics

import (
	"testing"
)

var globalRegistry Registry

func init(){
	globalRegistry = NewRegistry()
}

func TestStandardRegistry_Register(t *testing.T) {
	counter := NewCounter()
	globalRegistry.Register("default", counter)
	counterGet := globalRegistry.GetCounter("default")
	counterNil := globalRegistry.GetCounter("nil")
	if counterGet == nil{
		t.Error("get counter is nil but want counter")
		return
	}

	if counterNil != nil{
		t.Error("get counter is not nil but want nil")
		return
	}

	t.Log("register success")


}

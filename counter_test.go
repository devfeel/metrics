package metrics

import "testing"

func TestStandardCounter_IncAndDec(t *testing.T) {
	counter := NewCounter()
	counter.Inc(1)
	counter.Inc(2)
	counter.Dec(2)
	if counter.Count() != 1{
		t.Error("Inc not match")
	}

	t.Log("Inc test success")
}

package coins

import "testing"

func TestPermFn(t *testing.T) {
	fn := PermsFn()

	expected := []int{1, 2, 3, 5, 7, 11, 15, 22, 30, 42}

	for i, v := range expected {

		if res := fn(); res != v {
			t.Errorf("PermsFn for %d coins: expected %d but was %d", i+1, v, res)
		}
	}
}

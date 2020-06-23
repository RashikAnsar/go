package primefactor

import (
	"fmt"
	"testing"
)


func TestPrimeFactor(t *testing.T) {
	if fmt.Sprintf("%v", PrimeFactor(23)) != `[23]` {
		t.Errorf("23")
	}

	if fmt.Sprintf("%v", PrimeFactor(100)) != `[2 2 5 5]` {
		t.Errorf("100")
	}
}
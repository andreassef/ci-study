import "testing"

func TestSoma(t *testing.T) {
	if Soma(1, 2) != 3 {
		t.Errorf("Soma(1, 2) != 3")
	}
}
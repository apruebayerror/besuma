package besuma

import "testing"

func TestSuma(t *testing.T) {
	got := Suma(10)
	want := 20
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

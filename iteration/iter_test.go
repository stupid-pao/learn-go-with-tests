package iterator

import "testing"

func Repeate(str string) string {
	repeated := ""
	for i := 0; i < 5; i++ {
		repeated += str
	}
	return repeated
}

func TestRepeat(t *testing.T) {
	repeated := Repeate("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q', but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeate("a")
	}
}

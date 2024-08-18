package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Valid repeat", func(t *testing.T) {
		repeated := Repeat("a", 5)
        expected := "aaaaa"
        if repeated != expected {
            t.Errorf("expected %q, but repeated %q", repeated, expected)
        }
	})
	t.Run("Caller to specify repeat", func(t *testing.T) {
		repeated := Repeat("a", 3)
        expected := "aaa"
        if repeated != expected {
            t.Errorf("expected %q, but repeated %q", repeated, expected)
        }
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
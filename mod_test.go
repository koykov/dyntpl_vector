package vector_inspector

import "testing"

func TestMod(t *testing.T) {
	t.Run("coalesce", testStage)
}

func BenchmarkMod(b *testing.B) {
	b.Run("coalesce", benchStage)
}

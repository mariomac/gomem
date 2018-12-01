package donut

import "testing"

func BenchmarkValue(b *testing.B) {
	preferences := RndPreferences()
	for i := 0; i < b.N; i++ {
		_ = ScoreVal(RndVal(), preferences)
	}
}

func BenchmarkPointers(b *testing.B) {
	preferences := RndPreferences()
	for i := 0; i < b.N; i++ {
		_ = ScorePtr(RndPtr(), &preferences)
	}
}

func BenchmarkValToPointer(b *testing.B) {
	preferences := RndPreferences()
	for i := 0; i < b.N; i++ {
		val := RndVal()
		_ = ScorePtr(&val, &preferences)
	}
}


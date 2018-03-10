package diff

import "testing"

func BenchmarkCompare(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Compare("example/left.xml", "example/right.xml")
	}
}

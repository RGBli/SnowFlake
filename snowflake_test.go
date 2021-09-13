package snowflake

import (
	"testing"
	"time"
)

func BenchmarkNextId(b *testing.B) {
	sf := NewSnowFlake(3, 2, 5, time.Date(2021, 6, 3, 0, 0, 0, 0, time.Local).Unix())
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sf.NextId()
		}
	})
}

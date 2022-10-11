package hashcash

import (
	"fmt"
	"testing"
)

var benchTargetBits = []uint64{4, 8, 12, 16, 20, 24}

func BenchmarkHashcash_Solve(b *testing.B) {
	for _, bits := range benchTargetBits {
		b.Run(fmt.Sprintf("target_bits_%d", bits), func(b *testing.B) {
			b.StopTimer()
			pow := NewPOW(bits)
			challenge := pow.Challenge()
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				pow.Solve(challenge)
			}
		})
	}
}

func BenchmarkHashcash_Verify(b *testing.B) {
	for _, bits := range benchTargetBits {
		b.Run(fmt.Sprintf("target_bits_%d", bits), func(b *testing.B) {
			b.StopTimer()
			pow := NewPOW(bits)
			challenge := pow.Challenge()
			answer := pow.Solve(challenge)
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				_ = pow.Verify(challenge, answer)
			}
		})
	}
}

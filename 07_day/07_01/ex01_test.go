package main

import (
	"fmt"
	"testing"
)

func BenchmarkMinCoins(b *testing.B) {
	allSum := []int{247, 554, 1108, 2216, 4432, 8864}
	want := []int{100, 50, 10, 5, 1}
	for _, sum := range allSum {
		b.Run(fmt.Sprintf("benchmark-sum-%d", sum), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				minCoins(sum, want)
			}
		})
	}
}

func BenchmarkMinCoins2(b *testing.B) {
	allSum := []int{247, 554, 1108, 2216, 4432, 8864}
	want := []int{100, 50, 10, 5, 1}
	for _, sum := range allSum {
		b.Run(fmt.Sprintf("benchmark-sum-%d", sum), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				minCoins2(sum, want)
			}
		})
	}
}

/*
go test -bench=. -cpuprofile cpu.out
go tool pprof cpu.out
  output=top10.txt
  top 10
  exit
*/

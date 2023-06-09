package main

import (
	"reflect"
	"testing"
	"fmt"
)

func TestMinCoins(t *testing.T) {
	want := []int{10, 1, 1, 1}

	t.Run("no sort coins", func(t *testing.T) {
		got := minCoins(13, []int{10, 5, 1})
		if len(got) != len(want) {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), len(want))
		}
	})
	t.Run("dubbles in coins", func(t *testing.T) {
		got := minCoins(13, []int{1, 5, 5, 10})
		if len(got) != len(want) {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), len(want))
		} else {
			for i := 0; i < len(got); i++ {
				if !reflect.DeepEqual(got[i], want[i]) {
					t.Errorf("Result was incorrect, got[%d]: %d, want[%d]: %d\n", i, len(got), i, len(want))
				}
			}
		}
	})
	t.Run("negative val", func(t *testing.T) {
		got := minCoins(-13, []int{1, 5, 10})
		if len(got) != 0 {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), 0)
		}
	})
	t.Run("zero val", func(t *testing.T) {
		got := minCoins(0, []int{1, 5, 10})
		// t.Deadline()
		if len(got) != 0 {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), 0)
		}
	})
	t.Run("double 3", func(t *testing.T) {
		got := minCoins(6, []int{1, 3, 4})
		if len(got) != 2 {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), 2)
		}
	})
	// t.Run("negative coins", func(t *testing.T) {
	// 	got := minCoins(13, []int{-1, -5, -10})
	// 	// if len(got) != 0 {
	// 	// 	t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), 0)
	// 	// }
	// 	t.Log("got: ", got)
	// })
}

func TestMinCoins2(t *testing.T) {
	want := []int{10, 1, 1, 1}

	t.Run("no sort array", func(t *testing.T) {
		got := minCoins2(13, []int{10, 5, 1})
		if len(got) != len(want) {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), len(want))
		}
	})
	t.Run("dubbles in array", func(t *testing.T) {
		got := minCoins2(13, []int{1, 5, 5, 10})
		if len(got) != len(want) {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), len(want))
		} else {
			for i := 0; i < len(got); i++ {
				if got[i] != want[i] {
					t.Errorf("Result was incorrect, got[%d]: %d, want[%d]: %d\n", i, len(got), i, len(want))
				}
			}
		}
	})
	t.Run("negative val", func(t *testing.T) {
		got := minCoins2(-13, []int{1, 5, 10})
		if len(got) != 0 {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), 0)
		}
	})
	t.Run("zero val", func(t *testing.T) {
		got := minCoins2(0, []int{1, 5, 10})
		if len(got) != 0 {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), 0)
		}
	})
	t.Run("negative coins", func(t *testing.T) {
		got := minCoins2(13, []int{-1, -5, -10})
		if len(got) != 0 {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), 0)
		}
	})
	t.Run("double 3", func(t *testing.T) {
		got := minCoins2(6, []int{1, 3, 4})
		if len(got) != 2 {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), 2)
		}
	})
}

func TestMinCoinsNew(t *testing.T) {
	want := []int{10, 1, 1, 1}

	t.Run("no sort array", func(t *testing.T) {
		got := minCoinsNew(13, []int{10, 5, 1})
		if len(got) != len(want) {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), len(want))
		}
	})
	t.Run("dubbles in array", func(t *testing.T) {
		got := minCoinsNew(13, []int{1, 5, 5, 10})
		if len(got) != len(want) {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), len(want))
		} else {
			for i := 0; i < len(got); i++ {
				if got[i] != want[i] {
					t.Errorf("Result was incorrect, got[%d]: %d, want[%d]: %d\n", i, len(got), i, len(want))
				}
			}
		}
	})
	t.Run("negative val", func(t *testing.T) {
		got := minCoinsNew(-13, []int{1, 5, 10})
		if len(got) != 0 {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), 0)
		}
	})
	t.Run("zero val", func(t *testing.T) {
		got := minCoinsNew(0, []int{1, 5, 10})
		if len(got) != 0 {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), 0)
		}
	})
	t.Run("negative coins", func(t *testing.T) {
		got := minCoinsNew(13, []int{-1, -5, -10})
		if len(got) != 0 {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), 0)
		}
	})
	t.Run("double 3", func(t *testing.T) {
		got := minCoinsNew(6, []int{1, 3, 4})
		if len(got) != 2 {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), 2)
		}
	})
}

func BenchmarkMinCoins(b *testing.B) {
	allSum := []int{247, 554, 1108, 2216, 4432, 8864}
	want := []int{100, 50, 10, 5, 1}
	for _, sum := range allSum {
		b.Run(fmt.Sprintf("benchmark-sum-%d", sum), func(b *testing.B){
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
		b.Run(fmt.Sprintf("benchmark-sum-%d", sum), func(b *testing.B){
				for i := 0; i < b.N; i++ {
					minCoins2(sum, want)
				}
		})
	}
}

func BenchmarkMinCoins2Optimized(b *testing.B) {
	allSum := []int{247, 554, 1108, 2216, 4432, 8864}
	want := []int{100, 50, 10, 5, 1}
	for _, sum := range allSum {
		b.Run(fmt.Sprintf("benchmark-sum-%d", sum), func(b *testing.B){
				for i := 0; i < b.N; i++ {
					minCoins2Optimized(sum, want)
				}
		})
	}
}

func BenchmarkMinCoins2New(b *testing.B) {
	allSum := []int{247, 554, 1108, 2216, 4432, 8864}
	want := []int{100, 50, 10, 5, 1}
	for _, sum := range allSum {
		b.Run(fmt.Sprintf("benchmark-sum-%d", sum), func(b *testing.B){
				for i := 0; i < b.N; i++ {
					minCoinsNew(sum, want)
				}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestMinCoins(t *testing.T) {
	want := []int{10, 1, 1, 1}

	t.Run("valid data", func(t *testing.T) {
		got := minCoins(13, []int{1, 5, 10})
		if len(got) != len(want) {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), len(want))
		} else {
			for i := 0; i < len(got); i++ {
				if !reflect.DeepEqual(got[i], want[i]) {
					t.Errorf("Result was incorrect, got[%d]: %d, want[%d]: %d\n", i, got[i], i, want[i])
				}
			}
		}
	})
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
					t.Errorf("Result was incorrect, got[%d]: %d, want[%d]: %d\n", i, got[i], i, want[i])
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
	t.Run("valid data", func(t *testing.T) {
		got := minCoins2(13, []int{1, 5, 10})
		if len(got) != len(want) {
			t.Errorf("Result was incorrect, lenght got: %d, lenght want: %d\n", len(got), len(want))
		} else {
			for i := 0; i < len(got); i++ {
				if !reflect.DeepEqual(got[i], want[i]) {
					t.Errorf("Result was incorrect, got[%d]: %d, want[%d]: %d\n", i, got[i], i, want[i])
				}
			}
		}
	})
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
}

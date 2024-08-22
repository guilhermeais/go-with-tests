package sum

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any size of numbers (slice)", func(t *testing.T) {
		// array
		numbers := []int{10, 20, 30}

		got := Sum(numbers)
		want := 60

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

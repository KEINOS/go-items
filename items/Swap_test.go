package items_test

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/KEINOS/go-items/items"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
//  Examples
// ----------------------------------------------------------------------------

func ExampleTItems_Swap() {
	list := items.New("s", "t", "u", "d", "y")

	list.Swap(1, 3) // Swap t and d = [s d u t y]
	list.Swap(0, 1) // Swap s and d = [d s u t y]
	list.Swap(1, 2) // Swap s and u = [d u s t y]

	fmt.Println(list.Items)
	// Output: [d u s t y]
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestSwap(t *testing.T) {
	list := items.New("a", "b", "c", "d", "e")

	{
		err := list.Swap(0, 4) // Swap a and e

		assert.NoError(t, err, "in-range swap should not return an error")
		assert.Equal(t, "e", list.Items[0], "the first item should be swapped")
		assert.Equal(t, "a", list.Items[4], "the last item should be swapped")
	}
	{
		err := list.Swap(-1, 1) // Negative index

		assert.Error(t, err, "negative index should return an error")
	}
	{
		err := list.Swap(0, 100) // Index out of range

		assert.Error(t, err, "overshoot range should return an error")
	}
}

// ----------------------------------------------------------------------------
//  Benchmark Tests
// ----------------------------------------------------------------------------

func BenchmarkSwap_reflect(b *testing.B) {
	items := genDataNthIntSlice(b, b.N)
	swapF := reflect.Swapper(items)
	index := 0

	fmt.Println("Refeclt: ", b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		//nolint:gosec // using of weak random number generator here is ok
		index = rand.Intn(b.N)
		swapF(0, index)
	}
}

func BenchmarkSwap_index(b *testing.B) {
	items := genDataNthIntSlice(b, b.N)
	index := 0

	fmt.Println("Index: ", b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		//nolint:gosec // using of weak random number generator here is ok
		index = rand.Intn(b.N)
		items[0], items[index] = items[index], items[0]
	}
}

// ----------------------------------------------------------------------------
//  Helper Functions
// ----------------------------------------------------------------------------

func genDataNthIntSlice(b *testing.B, n int) []int {
	b.Helper()

	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}

	return data
}

package items_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/KEINOS/go-items/items"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
//  Example
// ----------------------------------------------------------------------------

func ExampleTItems_Sort() {
	// Create a new list with random order items
	list := items.New("4", "5", "3", "2", "1")

	// Sort function for comparing items
	myIsLessThen := func(a, b int) bool {
		// Note that the items are interface. Re-cast them to string then convert to int
		numA, _ := strconv.Atoi(list.Items[a].(string))
		numB, _ := strconv.Atoi(list.Items[b].(string))

		return numA < numB
	}

	// Sort the items with the custom function
	list.Sort(myIsLessThen)

	// Print the sorted items
	fmt.Println(list.Items)
	// Output: [1 2 3 4 5]
}

// Example to sort own defined struct items but limited to a range.
func ExampleTItems_SortRange() {
	type MyItem struct {
		ID string
	}

	list := items.New()

	// Add items to the list but in random order
	list.Append(MyItem{ID: "4"})
	list.Append(MyItem{ID: "1"})
	list.Append(MyItem{ID: "3"})
	list.Append(MyItem{ID: "7"})
	list.Append(MyItem{ID: "8"})
	list.Append(MyItem{ID: "5"})
	list.Append(MyItem{ID: "6"})
	list.Append(MyItem{ID: "9"})
	list.Append(MyItem{ID: "10"})
	list.Append(MyItem{ID: "2"})

	fmt.Println("Before sort:", list.Items) // => [{4} {1} {3} {7} {8} {5} {6} {9} {10} {2}]

	// Sort function for comparing items
	myIsLessThen := func(a, b int) bool {
		// Note that the items are interface.
		// So you need to re-cast to the item type then convert to int
		numA, _ := strconv.Atoi(list.Items[a].(MyItem).ID)
		numB, _ := strconv.Atoi(list.Items[b].(MyItem).ID)

		return numA < numB
	}

	// Sort only the first 5 items from index 0
	if err := list.SortRange(0, 5, myIsLessThen); err != nil {
		fmt.Fprintf(os.Stderr, "error sorting: %v", err)
		os.Exit(1)
	}

	fmt.Println("After sort :", list.Items) // => [{1} {3} {4} {7} {8} {5} {6} {9} {10} {2}]
	// Output:
	// Before sort: [{4} {1} {3} {7} {8} {5} {6} {9} {10} {2}]
	// After sort : [{1} {3} {4} {7} {8} {5} {6} {9} {10} {2}]
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestSortRange(t *testing.T) {
	list := items.New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	// Range. From index 0 to 5 items (index 0-4)
	indexFrom := 0
	numItemsToSort := 5

	list.Shuffle()
	t.Logf("Shuffled: %v\n", list.Items)

	expectLastHalf := fmt.Sprintln(list.Items[numItemsToSort:])

	myIsLessThen := func(a, b int) bool {
		t.Logf("a:b = %v:%v = %#v:%#v, %#v", a, b, list.Items[a], list.Items[b], list.Items)

		return list.Items[a].(int) < list.Items[b].(int)
	}

	if err := list.SortRange(indexFrom, numItemsToSort, myIsLessThen); err != nil {
		t.Fatalf("Error sorting: %v", err)
	}

	t.Logf("Sorted: %v\n", list.Items)

	actualLastHalf := fmt.Sprintln(list.Items[numItemsToSort:])

	assert.Equal(t, expectLastHalf, actualLastHalf)

	// Check first half
	previous := -1
	for i := 0; i < numItemsToSort; i++ {
		assert.Greater(t, list.Items[i].(int), previous, "it should be sorted")
		previous, _ = list.Items[i].(int)
	}
}

func TestSortRange_out_of_range(t *testing.T) {
	list := items.New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	dummyFunc := func(a, b int) bool {
		return false
	}

	{
		indexFrom := -1
		numItemsToSort := 5

		err := list.SortRange(indexFrom, numItemsToSort, dummyFunc)

		assert.Error(t, err, "negative index should return error")
	}
	{
		indexFrom := len(list.Items)
		numItemsToSort := 5

		err := list.SortRange(indexFrom, numItemsToSort, dummyFunc)

		assert.Error(t, err, "out-of-range index should return error")
	}
	{
		indexFrom := 5
		numItemsToSort := 5 // index 5 is the 6th item so 4 items are left

		err := list.SortRange(indexFrom, numItemsToSort, dummyFunc)

		assert.Error(t, err, "out-of-range item num should return error")
	}
	{
		indexFrom := 5
		numItemsToSort := 4

		err := list.SortRange(indexFrom, numItemsToSort, dummyFunc)

		assert.NoError(t, err, "in-range index should return error")
	}
}

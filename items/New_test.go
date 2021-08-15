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
//  Examples
// ----------------------------------------------------------------------------

// Example of basic method usage.
func ExampleNew_basic_usage() {
	// Create a new list of items with initial values
	list := items.New("a", "b", "c")

	/* Add to the bottom */

	// Push adds the arg as a new item to the end of the list
	list.Push("d") // {"a", "b", "c", "d"}

	// Append is an alias of Push
	list.Append("e") // {"a", "b", "c", "d", "e"}

	/* Add to the top */

	// Unshift adds new item to the beginning of the list
	list.Unshift("f") // {"f", "a", "b", "c", "d", "e"}

	// Prepend is an alias of Unshift
	list.Prepend("g") // {"g", "f", "a", "b", "c", "d", "e"}

	/* Insert */

	// InsertArter adds new item after the specified index
	list.InsertAfter("h", 1) // {"g", "f", "h", "a", "b", "c", "d", "e"}

	// InsertBefore adds new item before the specified index
	list.InsertBefore("i", 1) // {"g", "i", "f", "h", "a", "b", "c", "d", "e"}

	/* Remove */

	// Delete removes the item at the specified index
	list.Delete(1) // {"g", "f", "h", "a", "b", "c", "d", "e"}

	// Pop removes the last item from the list and returns that item
	popOut := list.Pop()          // {"g", "f", "h", "a", "b", "c", "d"}
	fmt.Println("popOut", popOut) // => popOut e

	// Shift removes the first item from the list and returns that item
	shiftOut := list.Shift()           // {"f", "h", "a", "b", "c", "d"}
	fmt.Println("shiftOut:", shiftOut) // => shiftOut: g

	/* Manage */

	// Swap the item between the specified index and the next one
	if err := list.Swap(1, 5); err != nil { // {"f", "d", "a", "b", "c", "h"}
		fmt.Fprintf(os.Stderr, "error while swapping: %v", err)
		os.Exit(1)
	}

	fmt.Println(list.Items) // => [f d a b c h]
	// Output:
	// popOut e
	// shiftOut: g
	// [f d a b c h]
}

// Example of Sort() method.
func ExampleNew_sort() {
	// Define item type.
	type demoItem struct {
		ID  string // Note that ID is a string
		Num int    // Of course, you can use any type
	}

	// Create empty list
	list := items.New()

	// Push items to the list randomly
	list.Push(demoItem{ID: "1", Num: 1})
	list.Push(demoItem{ID: "5", Num: 2})
	list.Push(demoItem{ID: "4", Num: 3})
	list.Push(demoItem{ID: "2", Num: 4})
	list.Push(demoItem{ID: "3", Num: 5})

	fmt.Println("Before Sort: ", list.Items) // [{1 1} {5 2} {4 3} {2 4} {3 5}]

	// Define IsLessThan function.
	// This function is used to sort the list by ID fields. It returns true if
	// the first item is less than the second item. The given args a and b are
	// the the indexs of the item.
	myIsLessThen := func(a, b int) bool {
		// Since the items are in interface{}, you need to re-cast the item's type
		// to get the value from the field name.
		valA := list.Items[a].(demoItem).ID
		valB := list.Items[b].(demoItem).ID

		// String -> Int
		numA, _ := strconv.Atoi(valA)
		numB, _ := strconv.Atoi(valB)

		return numA < numB
	}

	// Sort the list.
	// You need to define the range of the list and the function to sort the list.
	list.Sort(myIsLessThen)

	fmt.Println("After Sort by ID: ", list.Items) // [{1 1} {2 4} {3 5} {4 3} {5 2}]

	// Sort the list by Num field.
	list.Sort(func(a, b int) bool {
		// Re-cast type and get Num field
		valA := list.Items[a].(demoItem).Num
		valB := list.Items[b].(demoItem).Num

		return valA < valB
	})

	// Note that the order of the items were sorted by Num field.
	fmt.Println("After Sort by Num: ", list.Items) // [{1 1} {5 2} {4 3} {2 4} {3 5}]

	// Output:
	// Before Sort:  [{1 1} {5 2} {4 3} {2 4} {3 5}]
	// After Sort by ID:  [{1 1} {2 4} {3 5} {4 3} {5 2}]
	// After Sort by Num:  [{1 1} {5 2} {4 3} {2 4} {3 5}]
}

// Example of user defined type items.
func ExampleNew_original_item() {
	// Define your strtuct of the item
	type demoItem struct {
		MyName string
	}

	// Create item
	item1 := demoItem{MyName: "demo1"}
	item2 := demoItem{MyName: "demo2"}

	// Create list with item1 for the first item
	list := items.New(item1)

	// Push item2 to the list
	list.Append(item2)

	// To access item field, re-cast the original item's type
	result1 := list.Items[0].(demoItem).MyName
	fmt.Println(result1) // "demo1"

	// To access item field, re-cast the original item's type
	result2 := list.Items[1].(demoItem).MyName
	fmt.Println(result2) // "demo2"
	// Output:
	// demo1
	// demo2
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestNew_capacity(t *testing.T) {
	list := items.New("a", "b", "c", "d", "e", "f", "g", "h", "i", "j")

	lenList := len(list.Items) // => 10
	capList := cap(list.Items) // => 10

	assert.Equal(t, lenList, capList, "capacity should be same as it's length")
}

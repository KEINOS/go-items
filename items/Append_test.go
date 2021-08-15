package items_test

import (
	"fmt"
	"testing"

	"github.com/KEINOS/go-items/items"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
//  Examples
// ----------------------------------------------------------------------------

func ExampleTItems_Append() {
	list := items.New("a", "b")

	list.Append("c")

	for _, item := range list.Items {
		fmt.Println(item)
	}
	// Output:
	// a
	// b
	// c
}

func ExampleTItems_Push() {
	list := items.New("a", "b")

	list.Push("c")

	for _, item := range list.Items {
		fmt.Println(item)
	}
	// Output:
	// a
	// b
	// c
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestAppend(t *testing.T) {
	list := items.New("a")

	list.Append("b") // it should cover Push as well
	list.Append("c") // it should cover Push as well

	t.Logf("List item: %#v\n", list)

	assert.Contains(t, list.Items, "a")

	expect := []interface{}{"a", "b", "c"}
	actual := list.Items

	assert.Equal(t, expect, actual)
}

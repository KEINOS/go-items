package items_test

import (
	"fmt"
	"testing"

	"github.com/KEINOS/go-items/items"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func ExampleTItems_Prepend() {
	list := items.New("a", "b")

	list.Prepend("c") // adds to the top

	for _, item := range list.Items {
		fmt.Println(item)
	}
	// Output:
	// c
	// a
	// b
}

func ExampleTItems_Unshift() {
	list := items.New("a", "b")

	list.Unshift("c") // adds to the top

	for _, item := range list.Items {
		fmt.Println(item)
	}
	// Output:
	// c
	// a
	// b
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestPrepend(t *testing.T) {
	list := items.New("a", "b")

	list.Prepend("c") // It should cover Unshift as well

	t.Logf("List item: %#v\n", list)

	assert.Contains(t, list.Items, "a")

	expect := []interface{}{"c", "a", "b"}
	actual := list.Items

	assert.Equal(t, expect, actual)
}

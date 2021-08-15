package items_test

import (
	"fmt"
	"testing"

	"github.com/KEINOS/go-items/items"
	"github.com/stretchr/testify/assert"
)

func ExampleTItems_Shift() {
	list := items.New("a", "b", "c")

	shiftedItem := list.Shift()

	fmt.Println("Shifted item:", shiftedItem)
	fmt.Println("Left overs  :", list.Items)
	// Output:
	// Shifted item: a
	// Left overs  : [b c]
}

func TestShipt_empty_item(t *testing.T) {
	list := items.New()

	shiftedItem := list.Shift()
	t.Logf("Returned: %#v", shiftedItem)

	assert.Nil(t, shiftedItem, "shifting an empty list should return nil")
}

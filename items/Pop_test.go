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

func ExampleTItems_Pop() {
	list := items.New("a", "b", "c")

	poppedOut := list.Pop()

	fmt.Println("Popped out:", poppedOut)
	fmt.Println("Left overs:", list.Items)
	// Output:
	// Popped out: c
	// Left overs: [a b]
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestPop_empty_item(t *testing.T) {
	list := items.New()

	poppedOut := list.Pop()
	t.Logf("Returned: %#v", poppedOut)

	assert.Nil(t, poppedOut, "popping out from an empty list should return nil")
}

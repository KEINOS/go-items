package items_test

import (
	"fmt"

	"github.com/KEINOS/go-items/items"
)

func ExampleTItems_Delete() {
	list := items.New("a", "b", "c", "d")

	list.Delete(1)  // Remove index 1 = "b"
	list.Delete(10) // Remove undefined index(should do nothing)

	fmt.Println("Left overs:", list.Items)
	// Output:
	// Left overs: [a c d]
}

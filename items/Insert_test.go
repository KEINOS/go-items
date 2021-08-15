package items_test

import (
	"fmt"

	"github.com/KEINOS/go-items/items"
)

func ExampleTItems_InsertAfter() {
	list := items.New("1", "2", "3") // [1 2 3], Len = 3

	list.InsertAfter("4", 0)                  // Insert "4" after index 0 = "1"
	fmt.Println("Insert 4 to 0:", list.Items) // [1 4 2 3]

	list.InsertAfter("5", 1)                  // Insert "5" after index 1 = "4"
	fmt.Println("Insert 5 to 1:", list.Items) // [1 4 5 2 3]

	list.InsertAfter("6", -1)                  // Negative index will add to the top
	fmt.Println("Insert 6 to -1:", list.Items) // [6 1 4 5 2 3]

	list.InsertAfter("7", 10)                  // Oversize index will add to the bottom
	fmt.Println("Insert 7 to 10:", list.Items) // [6 1 4 5 2 3 7]

	// Output:
	// Insert 4 to 0: [1 4 2 3]
	// Insert 5 to 1: [1 4 5 2 3]
	// Insert 6 to -1: [6 1 4 5 2 3]
	// Insert 7 to 10: [6 1 4 5 2 3 7]
}

func ExampleTItems_InsertBefore() {
	list := items.New("1", "2", "3") // [1 2 3], Len = 3

	list.InsertBefore("4", 0)                 // Insert "4" before index 0 = "1"
	fmt.Println("Insert 4 to 0:", list.Items) // [4 1 2 3]

	list.InsertBefore("5", 1)                 // Insert "5" before index 1 = "1"
	fmt.Println("Insert 5 to 1:", list.Items) // [4 5 1 2 3]

	list.InsertBefore("6", -1)                 // Negative index will add to the top
	fmt.Println("Insert 6 to -1:", list.Items) // [6 4 5 1 2 3]

	list.InsertBefore("7", 10)                 // Oversize index will add to the bottom
	fmt.Println("Insert 7 to 10:", list.Items) // [6 4 5 1 2 3 7]

	// Output:
	// Insert 4 to 0: [4 1 2 3]
	// Insert 5 to 1: [4 5 1 2 3]
	// Insert 6 to -1: [6 4 5 1 2 3]
	// Insert 7 to 10: [6 4 5 1 2 3 7]
}

package items

import (
	"math/rand"
	"reflect"
	"time"
)

// Shuffle the items.
func (i *TItems) Shuffle() {
	refValue := reflect.ValueOf(i.Items)
	lenItems := refValue.Len()

	//nolint:gosec // use weak random but fast since it's not used for cryptographic purposes
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	swapper := reflect.Swapper(i.Items)

	randomizer.Shuffle(lenItems, func(i, j int) {
		swapper(i, j)
	})

	i.Sorted = true
}

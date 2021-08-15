package items

// InsertAfter inserts the given item after the given index.
func (i *TItems) InsertAfter(item interface{}, index int) {
	i.InsertBefore(item, index+1)
}

// InsertBefore inserts the given item before the given index.
func (i *TItems) InsertBefore(item interface{}, index int) {
	lenItem := len(i.Items)

	if 0 <= index && index < lenItem-1 {
		// Create enough capacity for the new item
		i.Items = append(i.Items[:index+1], i.Items[index:]...)
		// Set the new item
		i.Items[index] = item

		return
	}

	if index > lenItem-1 {
		i.Push(item) // Add the item to the end of the list

		return
	}

	if index < 0 {
		i.Unshift(item) // Add the item to the beginning of the list

		return
	}
}

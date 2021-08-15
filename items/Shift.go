package items

// Shift the item from the list. It will return the deleted first item from the list.
func (i *TItems) Shift() interface{} {
	if len(i.Items) == 0 {
		return nil
	}

	shifted := i.Items[0] // Save the shifted item

	i.Items = i.Items[1:] // Shift and update

	return shifted
}

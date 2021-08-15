package items

// Delete a single item of the given index from the list.
func (i *TItems) Delete(index int) {
	if index > len(i.Items)-1 || index < 0 {
		return
	}

	i.Items = append(i.Items[:index], i.Items[index+1:]...)
}

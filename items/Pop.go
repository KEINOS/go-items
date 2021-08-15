package items

// Pop outs the last item from the list. It will return the last item which was removed from the list.
func (i *TItems) Pop() interface{} {
	if len(i.Items) == 0 {
		return nil
	}

	lastItem := i.Items[len(i.Items)-1] // The item to be popped

	i.Items = i.Items[:len(i.Items)-1] // Pops the last item from the list

	return lastItem
}

package items

// Append adds the item to the bottom of the list which is an alias of Push().
func (i *TItems) Append(item interface{}) {
	i.Push(item)
}

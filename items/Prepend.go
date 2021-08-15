package items

// Prepend adds the item to the top of the list which is an alias of Unshift().
func (i *TItems) Prepend(item interface{}) {
	i.Unshift(item)
}

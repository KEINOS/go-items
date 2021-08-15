package items

/* The Append_test.go should cover this method as well */

// Push the given item into the bottom of the list.
func (i *TItems) Push(item interface{}) {
	i.Items = append(i.Items, item)
}

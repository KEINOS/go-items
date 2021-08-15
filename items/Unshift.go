package items

/* The Prepend_test.go should cover this method as well */

// Unshift adds the item to the top of the list.
func (i *TItems) Unshift(item interface{}) {
	i.Items = append([]interface{}{item}, i.Items...)
}

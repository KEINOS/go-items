package items

// New returns the pointer to newly created TItems object.
func New(initial ...interface{}) *TItems {
	tmp := new(TItems)

	if lenInput := len(initial); lenInput == 0 {
		return tmp
	}

	tmp.Items = append(tmp.Items, initial...)

	return tmp
}

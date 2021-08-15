package items

import "golang.org/x/xerrors"

// Swap items between index of a and b.
func (i *TItems) Swap(a int, b int) error {
	lenItem := len(i.Items)

	if a < 0 || a >= lenItem || b < 0 || b >= lenItem {
		return xerrors.New("Index out of range")
	}

	i.Items[a], i.Items[b] = i.Items[b], i.Items[a]

	return nil
}

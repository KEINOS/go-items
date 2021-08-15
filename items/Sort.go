package items

import (
	"sort"

	"golang.org/x/xerrors"
)

// Sort thelist using the given isLessThan compare function.
//
// The isLessThan function should return true if the first argument is less than the second argument.
//
// 	isLessThan := func(a, b int) bool {
// 		return yourItems[a] < yourItems[b]
// 	}
func (i *TItems) Sort(isLessThan func(int, int) bool) {
	sort.Slice(i.Items, isLessThan)
}

// SortRange sorts the number of items "numItems" from index "indexFrom"
// using the given "isLessThan" compare function.
//
// The isLessThan function should return true if the first argument is less
// than the second argument.
//
// 	isLessThan := func(a, b int) bool {
// 		return yourItems[a] < yourItems[b]
// 	}
func (i *TItems) SortRange(indexFrom int, numItems int, isLessThan func(int, int) bool) error {
	lenList := len(i.Items)

	if indexFrom < 0 || indexFrom > lenList-1 || indexFrom+1+numItems > lenList {
		return xerrors.Errorf(
			"out of range. Index From-To: %v-%v, but item len was: %v\n",
			indexFrom, numItems-1, lenList,
		)
	}

	sort.Slice(i.Items[indexFrom:indexFrom+numItems], isLessThan)

	return nil
}

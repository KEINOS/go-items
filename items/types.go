package items

// TItems is the structure that holds all the items.
type TItems struct {
	// Items is the list of items.
	Items []interface{} `json:"items"`
	// Sorted is a flag that indicates if the items are sorted.
	Sorted bool `json:"sorted"`
}

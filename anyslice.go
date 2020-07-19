package mtype

// Built-in type.
type AnySlice []interface{}

// Set value.
func (as *AnySlice) Set(args string, value interface{}) {
	link := Set(args, value, *as)
	*as = link.Reassignment().([]interface{})
}

// Get value.
func (as *AnySlice) Get(args ...string) *AnyValue {
	return Get(*as, args...)
}

// Check this value is in array or not.
func (as *AnySlice) InArray(value interface{}) bool {
	for _, v := range *as {
		if v == value {
			return true
		}
	}
	return false
}

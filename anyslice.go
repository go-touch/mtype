package mtype

// Built-in type.
type AnySlice []interface{}

// Select value.
func (as *AnySlice) Select(args ...string) *AnyValue {
	return Select(*as, args...)
}

// Insert | Update value.
func (as *AnySlice) Modify(args string, value interface{}) {
	link := Modify(args, value, *as)
	*as = link.Reassignment().([]interface{})
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

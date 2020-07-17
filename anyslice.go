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
package mtype

// Built-in type.
type AnyMap map[string]interface{}

// Select value.
func (am *AnyMap) Select(args ...string) *AnyValue {
	return Select(*am, args...)
}

// Insert | Update value.
func (am *AnyMap) Modify(args string, value interface{}) {
	link := Modify(args, value, *am)
	*am = link.Reassignment().(map[string]interface{})
}

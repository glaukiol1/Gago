package lang

// this file holds the functions specific to
// the lang.String type
// find the definition in types.go

func (str *TypeString) len() int {
	return len(str.Value.(string))
}

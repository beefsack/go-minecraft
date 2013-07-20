package nbt

const (
	TAG_CLOSE = iota
	TAG_BYTE
	TAG_SHORT
	TAG_INT
	TAG_LONG
	TAG_FLOAT
	TAG_DOUBLE
	TAG_BYTE_ARRAY
	TAG_STRING
	TAG_LIST
	TAG_COMPOUND
	TAG_INT_ARRAY
)

type List struct {
	ListType int
	Items    []interface{}
}

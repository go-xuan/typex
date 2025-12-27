package typex

// JSON JSON类型
type JSON interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

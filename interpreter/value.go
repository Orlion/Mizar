package interpreter

type mizarValueType string

const (
	mizarValueTypeString mizarValueType = "string"
	mizarValueTypeNumber mizarValueType = "number"
)

type mizarValue struct {
	v *value
	t mizarValueType
}

type value struct {
	str    string
	number int64
}

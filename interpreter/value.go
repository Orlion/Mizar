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
	number int
}

type StatementResultType string

const (
	StatementResultTypeNormal   StatementResultType = "normal"
	StatementResultTypeReturn   StatementResultType = "return"
	StatementResultTypeBreak    StatementResultType = "break"
	StatementResultTypeContinue StatementResultType = "continue"
)

type StatementResult struct {
	T           StatementResultType
	ReturnValue *mizarValue
}

type NativeFunc func(args []*mizarValue) (mval *mizarValue, err error)

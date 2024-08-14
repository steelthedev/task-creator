package types

type LogError struct {
	Msg   string
	Args  string
	Error error
}

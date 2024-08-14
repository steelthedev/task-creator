package data

type Validate interface {
	IsEmpty(text string) (bool, error)
}

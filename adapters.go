package mocks

type ExternalAdapter interface {
	Do(in string) (out string, err error)
}

type InternalAdapter interface {
	Do() (bool, error)
}

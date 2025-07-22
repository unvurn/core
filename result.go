package core

type Result interface {
	As(any) error
}

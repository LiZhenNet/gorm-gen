package core

type Generator interface {
	Init(ctx Context) error
	Exec(ctx Context) error
	AfterExec(ctx Context) error
}

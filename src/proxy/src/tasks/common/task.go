package common

type TaskInterface interface {
	Start() error
	Stop() error
}
type Task struct {
}

func (t Task) Start() error {
	panic("implement me")
}

func (t Task) Stop() error {
	panic("implement me")
}

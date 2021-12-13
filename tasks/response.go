package tasks

type IResponse interface {
	Handle(response []byte)
}

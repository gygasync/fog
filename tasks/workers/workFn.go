package workers

type IWorkFn interface {
	Fn() func(b []byte)
}

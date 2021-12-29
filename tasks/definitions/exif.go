package definitions

import (
	"fmt"
	"fog/tasks"
)

type ExifWorkGroup struct {
	workFn func([]byte) []byte
	wGroup tasks.IWorkerGroup
}

// func GetExifWorkFn(wGroup tasks.IWorkerGroup) *ExifWorkGroup {
// 	fn := func(b []byte) {
// 		fmt.Println(b)
// 		go time.Sleep(time.Millisecond * 300)
// 		fmt.Println("END")
// 		wGroup.Respond([]byte("SUCCESS"))
// 	}
// 	return &ExifWorkGroup{workFn: fn, wGroup: wGroup}
// }

func (f *ExifWorkGroup) Fn() func([]byte) []byte {
	return f.workFn
}

type ExifWorkHandler struct {
	internal string
}

func GetExifWorkHandler() *ExifWorkHandler {
	return &ExifWorkHandler{internal: "TEMP"}
}

func (f *ExifWorkHandler) Handle(response []byte) {
	fmt.Println("Response from job ")
}

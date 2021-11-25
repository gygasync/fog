package workers

import "fmt"

type Exif struct {
	workFn func([]byte)
}

func GetExifFn() *Exif {
	fn := func(b []byte) { fmt.Println(b) }
	return &Exif{workFn: fn}
}

func (f *Exif) Fn() func([]byte) {
	return f.workFn
}

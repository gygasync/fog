package workers

type Exif struct {
	workFn func([]byte) []byte
}

func GetExifFn() *Exif {
	fn := func(b []byte) []byte {
		return b
	}
	return &Exif{workFn: fn}
}

func (f *Exif) Fn() func([]byte) []byte {
	return f.workFn
}

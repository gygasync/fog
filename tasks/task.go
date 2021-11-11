package tasks

type ITask interface {
	GetType() string
	GetBytes() []byte
}

type ExifTask struct {
	data []byte
}

func NewExifTask(data []byte) *ExifTask {
	return &ExifTask{data: data}
}

func (e *ExifTask) GetType() string {
	return "exif"
}

func (e *ExifTask) GetBytes() []byte {
	return e.data
}

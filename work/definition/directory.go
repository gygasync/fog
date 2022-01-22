package definition

type DirectoryWorkDefinition struct {
	DirPath string
}

func NewDirectoryWork(path string) *Work {
	workDefinition := &DirectoryWorkDefinition{DirPath: path}
	return NewWorkDefinition("directory", workDefinition)
}

package models

type Reference struct {
	Id   int64
	Tag  string // Reference to Tag
	Item string // Reference to either File or Directory
}

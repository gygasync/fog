package models

type TagToTag struct {
	Id     int64
	Source string // Reference to Tag
	Target string // Reference to Tag
}

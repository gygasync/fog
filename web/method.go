package web

type Method int8

const (
	GET Method = iota
	POST
)

func (m Method) String() string {
	switch m {
	case GET:
		return "GET"
	case POST:
		return "POST"
	}

	return ""
}

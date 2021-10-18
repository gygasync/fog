package learning

import "rsc.io/quote"

func Getquote() string {
	return quote.Go()
}

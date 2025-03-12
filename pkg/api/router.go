package api

import "fmt"

type Route struct {
	Method      string
	Path        string
	HandlerName string
}

type Routes []Route

func (rs Routes) Print() {
	lm := 3
	lp := 0

	for _, r := range rs {
		if len(r.Method) > lm {
			lm = len(r.Method)
		}

		if len(r.Path) > lp {
			lp = len(r.Path)
		}
	}

	for _, r := range rs {
		fmt.Printf("Uzone | route | %*s - %*s\n", lm, r.Method, lp, r.Path)
	}
}

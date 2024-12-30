package tool

import "github.com/loveuer/uzone/pkg/log"

func Must(errs ...error) {
	for _, err := range errs {
		if err != nil {
			log.New().Panic(err.Error())
		}
	}
}

package console

import (
	"fmt"
	"log"
)

func PanicIf(cond bool, format string, args ...any) {
	if cond {
		log.Panicf(format, args)
	}
}

func PanicIfError(err error, format string, args ...any) {
	if err != nil {
		log.Panicf("%s: %+v", fmt.Sprintf(format, args), err)
	}
}

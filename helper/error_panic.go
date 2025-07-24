package helper

import "fmt"

func ErrorPanic(desc string, err error) {
	if err != nil {
		panic(fmt.Sprintf("%s: %v", desc, err))
	}
}

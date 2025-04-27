package cling

import (
	"fmt"
	"log"
)

type Output struct {
	prefix string
}

func NewOutput(prefix string) *Output {
	return &Output{
		prefix: prefix,
	}
}

func (o *Output) Log(message string) {
	if o.prefix != "" {
		message = fmt.Sprintf("[%s] %s", o.prefix, message)
	}

	log.Println(message)
}

func (o *Output) LogIf(message string, condition bool) {
	if !condition {
		return
	}

	o.Log(message)
}

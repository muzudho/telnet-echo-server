package main

import (
	"fmt"

	"github.com/reiver/go-telnet"
)

func server() error {
	return telnet.ListenAndServe(fmt.Sprintf(":%d", 5555), handler{})
}

type handler struct{}

func (h handler) ServeTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	println("-- connect --")

	var buffer [1]byte
	p := buffer[:]

	for {
		n, err := r.Read(p)
		if n > 0 {
			bytes := p[:n]
			print(string(bytes))
		}

		if nil != err {
			break
		}
	}

	println("-- disconnect --")
}

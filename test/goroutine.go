package main

import (
	"fmt"
	"github.com/goperate/convert/core/goroutine"
	"log"
	"runtime/debug"
	"time"
)

func main() {
	obj := goroutine.NewGO(3).RecoverFunc(func(e interface{}) {
		log.Println(e)
		debug.PrintStack()
	})
	for {
		obj.Run(func() {
			time.Sleep(time.Second)
			fmt.Println(time.Now())
			a, b := 0, 0
			fmt.Println(a / b)
		})
	}
}

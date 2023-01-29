// Code generated by hertz generator.

package main

import (
	"toy-tok/middleware"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()
	h.Use(middleware.GlobalErrorHandler)
	register(h)
	h.Spin()
}

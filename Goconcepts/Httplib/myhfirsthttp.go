/*
func ListenAndServe(addr string, handler Handler) error
we need address and handler to start with

addr: ip:port , handler Handler object

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

type ResponseWriter interface {
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(statusCode int)
}

*/
package main

import (
	"fmt"
	"io"
	"net/http"
)

type HttpHandler struct{}

//
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("This is My Home page in Golang.")
	io.WriteString(res, "Hello")
	fmt.Fprint(res, " World! ")
	res.Write(data)
}

func main() {
	fmt.Println("This is My Home page in Golang")
	handler := HttpHandler{}
	http.ListenAndServe(":4300", handler)
}

package main

import (
	"net/http"
)

type Handler interface {
	ServeHttp(http.ResponseWriter, *http.Request)
}

func main() {
	http.Handle("/", new(testHandler))

	http.ListenAndServe(":5000", nil)
}

type testHandler struct {
	http.Handler
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "YOur Request Path is " + req.URL.Path
	w.Write([]byte(str))
}
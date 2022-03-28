package main

import (
	"fmt"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.Handle("/hello", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`hello`))
	}))
	m.Handle("/rank", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`rank`))
	}))
	m.Handle("/history/xiaohei", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`xiaohei`))
	}))
	m.Handle("/history", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		qp := request.URL.Query()
		name := qp.Get("name")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(fmt.Sprintf(`%s: %s的历史`, request.Method, name)))
	}))
	http.ListenAndServe(":8080", m)
}

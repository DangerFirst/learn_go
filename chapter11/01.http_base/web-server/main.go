package main

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

func main() {
	http.ListenAndServe(":8088", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Body == nil {
			writer.Write([]byte("no body"))
			return
		}
		data, _ := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		encoded := base64.StdEncoding.EncodeToString(data)
		writer.Write(append(data, []byte(encoded)...))
	}))
}

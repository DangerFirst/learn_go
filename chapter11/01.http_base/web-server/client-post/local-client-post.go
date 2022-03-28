package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	r := strings.NewReader("fooooo")
	//bytes.NewReader([]byte(""))
	resp, err := http.Post("http://localhost:8088", "*/*", r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("resp:", string(data))
}

package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		key := "q"
		value := req.FormValue(key)
		io.WriteString(res, "Query String - Name :"+value)
	})

	http.ListenAndServe(":8080", nil)
}

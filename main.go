package main

import (
	"fmt"
	"unicode/utf8"
	"net/http"
)

func trimFirst(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html lang=\"en_US\"><head><meta charset=\"UTF-8\"/><title>Hello %s</title></head><body><div>Hello, %s!</div></body></html>\n", trimFirst(r.URL.Path), trimFirst(r.URL.Path))
	})

	http.ListenAndServe(":80", nil)
}


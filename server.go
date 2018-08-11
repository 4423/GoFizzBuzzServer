package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func fizzbuzz(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "Fizz Buzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("", "")
	r.ParseForm()
	s := strings.Split(r.URL.Path, "/")
	n, err := strconv.Atoi(s[len(s)-1])
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "invalid number")
	} else {
		fmt.Fprintf(w, fizzbuzz(n))
	}
}

func main() {
	http.HandleFunc("/fizzbuzz/", handler)
	http.ListenAndServe(":8080", nil)
}

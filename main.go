package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func fizzbuzz(num int) string {
	if num%15 == 0 {
		return fmt.Sprintf("Fizz Buzz")
	} else if num%3 == 0 {
		return fmt.Sprintf("Fizz")
	} else if num%5 == 0 {
		return fmt.Sprintf("Buzz")
	} else {
		return fmt.Sprintf("%d", num)
	}
}

func fizzbuzz_handler(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(strings.Split(r.RequestURI, "/")[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid number")
	} else {
		fmt.Fprintf(w, fizzbuzz(num))
	}
}

func main() {
	http.HandleFunc("/fizzbuzz/", fizzbuzz_handler)
	http.ListenAndServe(":8080", nil)
}

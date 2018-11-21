package main

import (
	"fmt"
	"net/http"
)

func poemHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Poem coming soon\n")
}

func main() {
	http.HandleFunc("/poem", poemHandler)
	http.ListenAndServe(":8088", nil)
}

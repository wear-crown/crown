package main

import (
	"crown/crown"
	"fmt"
	"net/http"
)

func main() {
	r := crown.New()
	
	r.GET("/index", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.Run(":8080")
}

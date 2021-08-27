package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(`:9080`, http.FileServer(http.Dir(`.`)))
}

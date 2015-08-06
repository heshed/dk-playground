package main
import (
	"net/http"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
	http.ListenAndServe("0.0.0.0:8888", nil)
}
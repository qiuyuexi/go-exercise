package main

import (
	"fmt"
	"net/http"
)

func handle(r http.ResponseWriter, w *http.Request) {
	fmt.Println("hello")
}

type mux struct {

}

func (mux *mux) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Println(r)
	return
}


func main() {
	http.HandleFunc("/hello",handle)
	//http.ListenAndServe(":9090",nil);
	http.ListenAndServe(":9090",&mux{});
}

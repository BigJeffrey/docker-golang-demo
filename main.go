package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func tempereatureHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	values := r.URL.Query()

	fmt.Println(string(body))
	for k, v := range values {
		fmt.Println(k, v)
	}
}

func main() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/templog", tempereatureHandler)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

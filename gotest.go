package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Request struct {
	Input string
}

func helloWorld(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello, world!\n"))
}
func helloInput(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("proxessing input!")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var req Request
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}

	if req.Input == "" {
		w.Write([]byte("\n\n" + "hello, empty string"))
	} else {
		w.Write([]byte("\n\n" + "hello, " + req.Input))
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", helloWorld)
	r.HandleFunc("/products", helloInput)
	r.HandleFunc("/dresses/{type}", dressInput)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func dressInput(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//fmt.Print(vars)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("\n\n" + "hello, " + vars["type"]))
}

func formatRequest(r *http.Request) string {

	var request []string

	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)

	request = append(request, fmt.Sprintf("Host: %v", r.Host))

	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}

	return strings.Join(request, "\n")
}


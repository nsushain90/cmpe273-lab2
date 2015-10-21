package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

type Values struct {
	Name string `json:"name"`
}

type JsonValue struct {
	Greeting string `json:"greeting"`
}

func hello(response http.ResponseWriter, request *http.Request, p httprouter.Params) {

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	var a Values
	err = json.Unmarshal(body, &a)
	if err != nil {
		panic(err)
	}
	b := JsonValue{"Hello " + a.Name + "!"}
	json, err := json.Marshal(&b)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the text diagnostics to the client.
	fmt.Fprintf(response, "%v", string(json))
}
func main() {
	mux := httprouter.New()
	mux.POST("/hello/", hello)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

type Values struct {
	Name string `json:"name"`
}

type JsonValue struct {
	Greeting string `json:"greeting"`
}

func hello(response http.ResponseWriter, request *http.Request, p httprouter.Params) {

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	var a Values
	err = json.Unmarshal(body, &a)
	if err != nil {
		panic(err)
	}
	b := JsonValue{"Hello " + a.Name + "!"}
	json, err := json.Marshal(&b)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the text diagnostics to the client.
	fmt.Fprintf(response, "%v", string(json))
}
func main() {
	mux := httprouter.New()
	mux.POST("/hello/", hello)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()

}

package service

import (

	"net/http"
	"math/rand"
	"fmt"
	"encoding/json"

)
type Result struct{
	MaCH string
	Price int
}
func ServeHTTP1(w http.ResponseWriter, r *http.Request){
	response,err := getRandom1()
	if err!=nil{
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}
func ServeHTTP2(w http.ResponseWriter, r *http.Request){
	response,err := getRandom2()
	if err!=nil{
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}
func getRandom1() ([]byte, error){
	k := rand.Intn(100)
	m := Result{"ch1",k}
	return json.Marshal(m)
}
func getRandom2() ([]byte, error){
	k := rand.Intn(100)
	m := Result{"ch2",k}
	return json.Marshal(m)
}
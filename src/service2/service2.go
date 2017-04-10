package main

import(
	"net/http"
	"math/rand"
	"fmt"
	"encoding/json"
)
type Result struct{
	MaCH string
	Price int
}
func ServeHTTP2(w http.ResponseWriter, r *http.Request){
	response,err := getRandom2()
	if err!=nil{
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}
func getRandom2() ([]byte, error){
	k := rand.Intn(100)
	m := Result{"ch2",k}
	return json.Marshal(m)
}
func main(){
	http.HandleFunc("localhost:8081/ch2",ServeHTTP2)
	http.ListenAndServe("localhost:8081", nil)
}
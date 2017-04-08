package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"service"
	"data"
	"time"
)
type Result struct{
	MaCH string
	Price int
}
type Info struct{
	link string
	price int
}
var temp  map[string]string
func getService1() []byte{
	url1 := "http://localhost:8000/ch1"
	res, err := http.Get(url1)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return body
}
func getService2() []byte{
	url2 := "http://localhost:8000/ch2"
	res, err := http.Get(url2)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}
	return body
}

func main(){
	go func(){
		for ;;{
			temp = data.GetData()
			fmt.Println(temp)
			time.Sleep(300000*time.Millisecond)
			fmt.Println("Updated!")
		}

	}()
	http.HandleFunc("localhost:8000/ch1", service.ServeHTTP1)
	http.HandleFunc("localhost:8000/ch2",service.ServeHTTP2)
	http.HandleFunc("/",Server)
	http.ListenAndServe("localhost:8000",nil)
}
func Server(w http.ResponseWriter, r *http.Request){
	var link map[string]string
	link = temp
	var m1,m2 Result
	var m Info
	err := json.Unmarshal(getService1(),&m1)
	if err!=nil{
		panic(err)
	}
	err1 := json.Unmarshal(getService2(),&m2)
	if err1!=nil{
		panic(err1)
	}
	if m1.Price >= m2.Price{
		m = Info{ link["ch2"], m2.Price }
		fmt.Fprint(w,m)
	}else{
		m = Info{link["ch1"], m1.Price}
		fmt.Fprint(w,m)
	}

}
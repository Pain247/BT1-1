package data
import (
	"testing"
	"fmt"
)

func TestGetData(t *testing.T){
	if GetData()== nil{
		t.Error(`GetData()=false`)
	}
	fmt.Println(GetData())
}

func BenchmarkGetData(b *testing.B) {
	for i:=0;i<b.N;i++{
		GetData()
	}
}
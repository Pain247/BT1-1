package data
import (
	"testing"
)

func TestGetData(t *testing.T){
	if GetData()== nil{
		t.Error(`GetData()=false`)
	}
}

func BenchmarkGetData(b *testing.B) {
	for i:=0;i<b.N;i++{
		GetData()
	}
}
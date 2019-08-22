package BestLogFramework

import "testing"

func TestStart(t *testing.T){
	if Start() != "Hello"{
		t.Fail()
	}
}

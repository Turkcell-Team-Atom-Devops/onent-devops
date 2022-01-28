package app

import (	
	"testing"	
		
)

func TestRecycleSinglePool(t *testing.T) {
	var r IPool = Pool{Name: "GitSolution"}
	result, err := r.RecycleSingleAppPool()
	if !result || err != nil {
		t.Fatalf(err.Error())
	}
}

func TestRecyleIISAll(t *testing.T) {
	var r IPool = Pool{}
	result, err := r.IISReset()
	if !result || err != nil {		
		t.Fatalf(err.Error())
	}
}


func TestGetAllPool(t *testing.T) {
	var r IPool = Pool{}
	_, err := r.GetAll()
	if err != nil {		
		t.Fatalf(err.Error())
	}
}
package model

import (
	"fmt"
	"github.com/jeevatkm/go-model"
	"testing"
)

type SampleStruct struct {
	BookCount int
	BookCode  string
}

type SampleStruct2 struct {
	BookCount int
	BookCode  string
}

func Test_SetMethod(t *testing.T) {

	src := SampleStruct{
		BookCount: 10,
		BookCode:  "test_book",
	}

	err := model.Set(&src, "BookCount", 100)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	fmt.Println(src)
}

func Test_GetMethod(t *testing.T) {

	src := SampleStruct{
		BookCount: 10,
		BookCode:  "test_book",
	}
	value, err := model.Get(&src, "BookCount")

	if err != nil {
		t.Errorf("error: %v", err)
	}
	fmt.Println(value)

	fieldKind, _ := model.Kind(src, "BookCount")
	fmt.Println("Field kind:", fieldKind)

	fields, _ := model.Fields(src)
	fmt.Println("Fields:", fields)
}

func Test_CopyMethod(t *testing.T) {
	src := SampleStruct{
		BookCount: 10,
		BookCode:  "test_book",
	}

	dest := SampleStruct2{}

	errs := model.Copy(&dest, src)
	fmt.Println("Errors:", errs)

	fmt.Printf("\nSource: %#v\n", src)
	fmt.Printf("\nDestination: %#v\n", dest)
}

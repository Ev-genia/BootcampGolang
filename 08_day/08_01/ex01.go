package main

import (
	"fmt"
	"log"
	"reflect"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

type StructReader interface {
	ReadWrite()
}

func (u UnknownPlant) ReadWrite() {
	s := reflect.ValueOf(&u).Elem()
	typeOfU := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		tagExist := typeOfU.Field(i).Tag.Get("color_scheme")
		if tagExist != "" {
			fmt.Printf("%s(color_scheme=%v): %v,\n", typeOfU.Field(i).Name, tagExist, f.Interface())
		} else {
			fmt.Printf("%s: %v,\n", typeOfU.Field(i).Name, f.Interface())
		}
	}
}

func (a AnotherUnknownPlant) ReadWrite() {
	s := reflect.ValueOf(&a).Elem()
	typeOfU := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		tagExist := typeOfU.Field(i).Tag.Get("unit")
		if tagExist != "" {
			fmt.Printf("%s(unit=%v): %v,\n", typeOfU.Field(i).Name, tagExist, f.Interface())
		} else {
			fmt.Printf("%s: %v,\n", typeOfU.Field(i).Name, f.Interface())
		}
	}
}

func getReader(empty interface{}) StructReader {
	switch reflect.TypeOf(empty) {
	case reflect.TypeOf(UnknownPlant{}):
		return &UnknownPlant{}
	case reflect.TypeOf(AnotherUnknownPlant{}):
		return &AnotherUnknownPlant{}
	default:
		log.Fatal("Undefined type of struct\n")
	}
	return nil
}

func main() {
	uPlant := UnknownPlant{"first", "second", 10}
	getReader(uPlant)
	uPlant.ReadWrite()
	fmt.Println()
	aUPlant := AnotherUnknownPlant{42, "school21", 21}
	getReader(aUPlant)
	aUPlant.ReadWrite()
}

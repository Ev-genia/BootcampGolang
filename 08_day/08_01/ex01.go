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
		tagExist := typeOfU.Field(i).Tag.Get("unit")
		if tagExist != "" {
			fmt.Printf("%s(unit=%v): %v,\n", typeOfU.Field(i).Name, tagExist, f.Interface())
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
		tagExist := typeOfU.Field(i).Tag.Get("color_scheme")
		if tagExist != "" {
			fmt.Printf("%s(color_scheme=%v): %v,\n", typeOfU.Field(i).Name, tagExist, f.Interface())
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

// 	// TypeOf() возвращает reflect.Type переменной в пустой интерфейс.
// func TypeOf(i interface{}) Type

// // Interface вернёт значение v как interface{}.
// func (v Value) Interface() interface{}

func main() {
	// u := UnknownPlant{}
	// var empty interface{}
	// empty = u
	// ut := reflect.TypeOf(u)
	// field := ut.Field(0)
	// fmt.Println(field.Tag.Get("color_scheme"))

	uPlant := UnknownPlant{"first", "second", 10}
	getReader(uPlant)
	uPlant.ReadWrite()
	fmt.Println()
	aUPlant := AnotherUnknownPlant{42, "school21", 21}
	getReader(aUPlant)
	aUPlant.ReadWrite()
	// s := reflect.ValueOf(&t).Elem()
	// typeOfT := s.Type()
	// for i := 0; i < s.NumField(); i++ {
	// 	f := s.Field(i)
	// 	fmt.Printf("i: %d, typeOfT.Field(i).Name: %s, f.Type(): %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	// }
	// s.Field(0).SetString("forty")
	// s.Field(1).SetString("two")
	// s.Field(2).SetInt(42)
	// fmt.Println("t is now: ", t)

	// var x UnknownPlant
	// d := reflect.TypeOf(x)
	// fmt.Println("tag: ", d.Field(2).Tag.Get("color_scheme")) //tag:  rgb
	// f, k := d.Field(2).Tag.Lookup("color_scheme")
	// fmt.Println("f: ", f, ", k: ", k)                            //f:  rgb , k:  true
	// fmt.Printf("color_scheme=%v\n", f)                           //color_scheme=rgb
	// fmt.Printf("%v=%v\n", d.Field(2).Tag.Get("color_scheme"), k) //rgb=true

	// fmt.Println("type: ", reflect.TypeOf(u))                 //type:  main.UnknownPlant
	// fmt.Println("value:", reflect.ValueOf(u.Color))          //value: 0
	// fmt.Println("value:", reflect.ValueOf(u.Color).String()) //value: <int Value>
	// fmt.Println("value:", reflect.ValueOf(u.Color).Type())   //value: int
	// fmt.Println()

	// v := reflect.ValueOf(u.Color)
	// fmt.Println("type:", v.Type())                       //type: int
	// fmt.Println("kind is int:", v.Kind() == reflect.Int) //kind is int: true
	// fmt.Println("value: ", v.Int())                      //value: 0
	// fmt.Println()

	// x := v.Interface().(int)
	// fmt.Println("x: ", x) //x:  0

	// fmt.Println()
	// fmt.Println("x: ", v.Interface()) //x:  0

	// var i int = 10
	// fmt.Println()
	// // u.Color = 10
	// // y := reflect.ValueOf(u.Color)
	// y := reflect.ValueOf(i)
	// fmt.Println("settability of y: ", y.CanSet()) //settability of y:  false
	// fmt.Println()
	// p := reflect.ValueOf(&y)
	// fmt.Println("type of p: ", p.Type())          //type of p:  *reflect.Value
	// fmt.Println("settability of p: ", p.CanSet()) //settability of p:  false
	// fmt.Println()
	// e := p.Elem()
	// fmt.Println("settability of e: ", e.CanSet()) //settability of e:  true
	// e.SetInt(42)
	// fmt.Println(e.Interface())

}

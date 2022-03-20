package main

import (
	"encoding/base64"
	"fmt"
	"sort"
)

const (
	DNA = "DNA"
	XX  = "XX"
)

var (
	var1 = "var1"
	var2 = "var2"
)

func swap(a, b string) (string, int) {
	return a + b, 6
}
func myfun(a string, arg ...int) {
	fmt.Println(a, arg)
	fmt.Printf("%T", arg)
}

func main() {
	var a, b string
	a, b = "aaa", "bbb"
	//a = "dsf"
	//b = "df"
	//c := "aaaa"
	fmt.Println(a, b)
	//fmt.Println(DNA,XX)
	if a = "66"; a != "" {
		fmt.Println("a != kong")
	}
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println(swap("a", "b"))
	myfun("66", 1, 2, 5)

	arr := []int{1, 2, 34, 5}
	arr1 := arr
	fmt.Println(arr)
	for k, _ := range arr {
		fmt.Println(k)
	}
	fmt.Println(&arr, &arr1)
	//println(&arr,&arr1)
	mp := map[int]string{1: "111", 2: "23232"}
	fmt.Println(mp[2])
	str := "dsgd"
	fmt.Println(str[1:])

	var a1 int
	var a2 float32
	a1 = 3
	a2 = 3
	fmt.Printf("%T,%T \n", a1, a2)
	//fmt.Println(reflect.TypeOf(a1))
	//fmt.Println(reflect.TypeOf(a2))

	var a3, a4 float32
	a3 = 1
	a4 = 1.00000001
	fmt.Println(a3 == a4)

	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(people)
	//fmt.Println(sort.Sort(PersonArr(people)))
	sort.Sort(PersonArr(people))
	fmt.Println(people)
	//ints := []int{1,8,4,6,3,66,4}
	ints := []int{1, 6, 2, 3, 4}
	strs := []string{"ac2", "ac3", "ac1", "b"}
	//sort.Ints(ints)
	fmt.Println(sort.SearchInts(ints, 99))
	sort.Strings(strs)
	fmt.Println(strs)
	fmt.Println(sort.IntsAreSorted(ints))
	fmt.Println("------------------------")
	decoded, _ := base64.StdEncoding.DecodeString("MTIzNDU2")
	fmt.Println(string(decoded))
}

type Person struct {
	Name string
	Age  int
}
type PersonArr []Person

func (a PersonArr) Len() int           { return len(a) }
func (a PersonArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PersonArr) Less(i, j int) bool { return a[i].Age > a[j].Age }

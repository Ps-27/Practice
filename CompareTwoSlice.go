package main

import (
	"bytes"
	"fmt"
)

func main() {
	Arr := []byte{'p', 's', 'p', 's'}
	fmt.Println(string(Arr))
	s1 := Arr[:2]
	s2 := Arr[2:]
	fmt.Println(string(s1), " ", string(s2))
	// fmt.Printf("%T ", s1,)
	fmt.Printf("%T is datatype of %s\n", s1, s1)
	fmt.Printf("%T is datatype of %s\n", s2, s2)
	res := bytes.Compare(s1, s2)
	if res == 0 {
		fmt.Println("Equal slice")
	} else {
		fmt.Println("not equal slice")
	}

	a := []string{"prity", "sinha", "prity", "snha"}
	slice1 := a[:2]
	slice2 := a[2:]
	d := diff(slice1, slice2)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(d)
	if len(d) == 0 {
		fmt.Println("Equal slice")
	} else {
		fmt.Println("not equal slice")
	}

}

func diff(a, b []string) []string {
	m := make(map[string]struct{}, len(b))
	for _, x := range b {
		m[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := m[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

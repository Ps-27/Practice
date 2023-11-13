package main

import (
	"fmt"
	"sort"
)

type intervalArr [][]int

func (intervals intervalArr) Len() int {
	return len(intervals)
}
func (intervals intervalArr) Less(i, j int) bool {
	return intervals[i][0] < intervals[j][0]
}

func (intervals intervalArr) Swap(i, j int) {
	intervals[i], intervals[j] = intervals[j], intervals[i]
}

func merge(arr [][]int) [][]int {
	//inteface of given type
	a := intervalArr(arr)

	// sort on the bases of start time
	sort.Sort(a)

	fmt.Println("sorted :", a)

	//merge overlap using end and start index

	var output [][]int

	CurrentStartIndex := a[0][0]
	CurrentEndIndex := a[0][1]

	for j := 1; j < len(a); j++ {
		// if end element is >= then setting new end element
		if CurrentEndIndex >= a[j][0] {
			CurrentEndIndex = a[j][1]
		} else {

			// if end ele is < then simply adding both start and end into output and setting new start and end
			output = append(output, []int{CurrentStartIndex, CurrentEndIndex})
			CurrentStartIndex = a[j][0]
			CurrentEndIndex = a[j][1]
		}

	}
	output = append(output, []int{CurrentStartIndex, CurrentEndIndex})
	return output
}

type interArr [][]int

func (n interArr) Len() int {
	return len(n)
}
func (n interArr) Less(i, j int) bool {
	return n[i][0] < n[j][0]
}
func (n interArr) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func merge1(a [][]int) [][]int {

	//inteface of taken type
	b := interArr(a)

	//sorting
	sort.Sort(b)

	//sorted array on 1st value
	fmt.Println("sorted b :", b)

	//merge overlap using index
	curIndex := b[0][0]
	endIndex := b[0][1]
	res := [][]int{}
	for i := 1; i < len(b); i++ {
		if endIndex > b[i][0] {
			endIndex = b[i][1]
		} else {
			res = append(res, []int{curIndex, endIndex})
			curIndex = b[i][0]
			endIndex = b[i][1]
		}
		// res = append(res, []int{curIndex,endIndex})

	}
	res = append(res, []int{curIndex, endIndex})

	return res
}

func main() {
	arr := [][]int{{1, 4}, {2, 4}, {3, 5}, {6, 8}, {8, 10}, {9, 12}, {14, 15}}
	arr = merge(arr)

	fmt.Println(arr)

	arr2 := [][]int{{1, 4}, {2, 4}, {3, 5}, {6, 8}, {8, 10}, {9, 12}, {14, 15}}
	arr2 = merge1(arr2)
	fmt.Println(arr2)
}

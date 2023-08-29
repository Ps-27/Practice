package main

import (
	"fmt"
	"sort"
)
type intervalsArr [][]int

func (intervals intervalsArr)Len()int{
	return len(intervals)
}
func (intervals intervalsArr)Less(i,j int)bool{
	return intervals[i][0] < intervals[j][0]
}

func (intervals intervalsArr)Swap(i,j int){
	intervals[i],intervals[j]=intervals[j],intervals[i]
}

func merge(intervals [][]int)[][]int{
	a:=intervalsArr(intervals)
	//sort on the basis of start time
	sort.Sort(a)
	// SortedInterval:=[][]int(a)
	fmt.Println(a)
	//merge overlap using end and start index

	var output [][]int

	CstartIndex:=a[0][0]
	CEndIndex:=a[0][1]
	for j:=1;j<len(a);j++{
		if CEndIndex >= a[j][0] {
			CEndIndex=a[j][1]
		}else{
			output=append(output,[]int{CstartIndex,CEndIndex})
			CstartIndex=a[j][0]
			CEndIndex=a[j][1]
		}
	}
	output=append(output,[]int{CstartIndex,CEndIndex})

	return output
}


func main(){
	a:=[][]int{{1, 4}, {8, 10}, {9, 12}, {3, 5}}
	a=merge(a)
	fmt.Println(a)
}

package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'RemainderSorting' function below (and other code for sorting if needed).
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY strArr as parameter.
 */
type Rmodulo []RemainderModulo

type RemainderModulo struct {
	str     string
	len     int
	modulo3 int
}

func (r Rmodulo) Len() int {
	return len(r)
}
func (r Rmodulo) Less(i, j int) bool {
	if r[i].modulo3 == r[j].modulo3 {
		fmt.Println("modulo same for ", r[i], r[j])
		if r[i].str < r[j].str {
			return r[i].str < r[j].str
		}
	}
	return r[i].modulo3 < r[j].modulo3
}
func (r Rmodulo) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func RemainderSorting(strArr []string) []string {
	Rm := make([]RemainderModulo, len(strArr))
	fmt.Println(strArr[0])
	for i := 0; i < len(strArr); i++ {
		Rm[i].str = strArr[i]
		Rm[i].len = len(strArr[i])
		Rm[i].modulo3 = (Rm[i].len) % 3
	}

	fmt.Println("Without sorting:", Rm)
	sort.Sort(Rmodulo(Rm))
	RmAlpha := make([]RemainderModulo, 0)
	for i := 1; i < len(Rm); i++ {
		if Rm[i].modulo3 == Rm[i-1].modulo3 {
			RmAlpha = append(RmAlpha, Rm[i])
		}
	}

	sort.Sort(Rmodulo(RmAlpha))

	res := []string{}
	for i := 0; i < len(Rm); i++ {
		res = append(res, Rm[i].str)
	}
	return res
}

func main() {

	var strArr []string

	// for i := 0; i < int(strArrCount); i++ {
	// 	strArrItem := readLine(reader)
	// 	strArr = append(strArr, strArrItem)
	// }
	strArr = []string{"prity", "pooja", "a", "ab", "abc"}
	result := RemainderSorting(strArr)
	fmt.Println(result)
	// for i, resultItem := range result {
	// 	fmt.Fprintf(writer, "%s", resultItem)

	// 	if i != len(result)-1 {
	// 		fmt.Fprintf(writer, "\n")
	// 	}
	// }

	// fmt.Fprintf(writer, "\n")

	// writer.Flush()
}

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
// 	checkError(err)

// 	defer stdout.Close()

// 	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

// 	strArrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)

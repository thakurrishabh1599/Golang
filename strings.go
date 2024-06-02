// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// 	// a := "go guru"
// 	// 	// var b string
// 	// 	// b = "I love coding"
// 	// 	// fmt.Println(a + b)
// 	// 	// d := "some \nstring"
// 	// 	// fmt.Println("Some \n", d)
// 	// 	// c := `go guru
// 	// 	// scsacdcdvd\nvcvdv`
// 	// 	// fmt.Println(c)
// 	// 	// fmt.Println(d[0])
// 	// 	// f := 'e'
// 	// 	// fmt.Println(f)
// 	// 	// testStr := "test"
// 	// 	// for index, char := range testStr {
// 	// 	// 	fmt.Println("index", index, "bytes", char, "characters", string(char))
// 	// 	// }
// 	// 	// fmt.Println("type assertion", []byte(testStr))
// 	// str1 := "@@@@Some@ @string+@+@"
// 	// fmt.Println(strings.Trim(str1, "@"))
// 	// fmt.Println(strings.TrimLeft(str1, "@"))
// 	// fmt.Println(strings.TrimRight(str1, "@"))
// 	// str2 := "                        some string           "
// 	// fmt.Println(str2, "Before Trim")
// 	// fmt.Println(strings.TrimSpace(str2))

// 	// //Trim suffix
// 	// str3 := "Hello friends"
// 	// fmt.Println(strings.TrimSuffix(str3, "friends"))

// 	// //Trim Prefix
// 	// str4 := "Hello friends"
// 	// fmt.Println(strings.TrimPrefix(str4, "Hello"))
// 	//###############SPLIT STRINGS##################
// 	// str1 := "Welcome,to my class"
// 	// fmt.Println("Before split:-", str1)
// 	// strArr := strings.Split(str1, ",")
// 	// fmt.Println("After Split:-", strArr, len(strArr), strArr[1])
// 	// strArr2 := strings.SplitAfter(str1, ",")
// 	// fmt.Println("After SplitAfter:-", strArr2, strArr2[0])
// 	// strArr3 := strings.SplitAfterN(str1, ",", 2)
// 	// fmt.Println("After SplitAfter:-", strArr3, strArr3[0])
// 	// fmt.Println(strings.ToUpper(str1)) //Actuall string
// 	// fmt.Println(strings.ToLower(str1))
// 	// fmt.Println(strings.ToTitle(str1)) //Create copy of string

// 	//################COMPARE STRINGS#########################
// 	str1 := "abc"
// 	str2 := "ab"
// 	fmt.Println("Comparison", str1 == str2)
// 	result1 := "ABC" > "xyz" ///works based on lexical order-> that is ascii value
// 	fmt.Println("Result 1: ", result1)
// 	result2 := "XYZ" > "abc" ///works based on lexical order-> that is ascii value
// 	fmt.Println("Result 2: ", result2)

// 	/* return 0 if str1==str2
// 	return 1 if str1> str2
// 	return 2 if str1> str2
// 	*/
// 	fmt.Println(strings.Compare("gfg", "xyz"))
// 	fmt.Println(strings.Compare("Abc", "Abc"))
// 	fmt.Println(strings.Compare("Ayz", "ABC"))
// 	fmt.Println(strings.Compare("BooKS", "BooKs"))

// }

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// fmt.Println("enter you name:")
	// // var txtInput string
	// txtInput := ""

	// fmt.Scan(&txtInput)
	// fmt.Println("" + txtInput)

	age := 2
	fmt.Println((reflect.TypeOf(age)))

	//convert
	height := "100"
	var h int
	// var error
	h, err := strconv.Atoi(height)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(h)

	//แปลงค่า ทศนิยม
	water := "12.30"
	ml, err := strconv.ParseFloat(water, 64)
	fmt.Println(ml)
	fmt.Println("==============")


	// ต่อ string
	txtShow := fmt.Sprintf("water = %f, hegiht = %d", ml, h)

	fmt.Print(txtShow)


	// loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// switch case
	score := 89
	switch score {
	case 80:
		{
			println("A")
		}
	case 70:
		{
			println("B")
		}
	}

	// if else statement
	if score >= 80 {
		fmt.Println("A")
	} else if score < 80 && score >= 70 {
		fmt.Println("B")
	} else {
		fmt.Println("ไม่ผ่าน")
	}

	
}

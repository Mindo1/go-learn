package main

import (
	"fmt"
	"time"
)

// GO Routine & Channel

func Ping1S(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Ping: ", i)
		time.Sleep(1 * time.Second)
	}
	c <- 20 //channel ตัวแปรที่รับเข้ามาใน func
}

func SendNoti5S(ch chan string) {
	fmt.Println("send noti")
	time.Sleep(5 * time.Second)
	fmt.Println("send noti")
	ch <- "suces"
}

func main() {
	c := make(chan int)
	ch := make(chan string)
	go Ping1S(c) //run function con-current ทำงานไปพร้อมกัน โดยไม่ต้องรอให้บรรทัดไหนเสร็จก่อน
	go SendNoti5S(ch)
	time.Sleep(10 * time.Second)

	pingValue, SendNotiValue := <-c, <-ch
	fmt.Println(pingValue, SendNotiValue) //print ค่า ch ที่บอกว่า function ทำงานเสร็จเมื่อไหร่
	fmt.Println("completed!")
}

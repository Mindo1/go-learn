package main

import (
	"fmt"
	"time"
)

func IsInSyetem(username string) bool {
	return true
}

func GetUserDetail(username string) (int, string) {

	return 201, "manager"
}

func GetDeparture(username string, departure *string) {
	if username != "" {
		*departure = "home"
	}
}

func CheckLogin(username string, password string) {
	if IsInSyetem(username) {
		fmt.Println("found user in system")
		GetUserDetail(username)
		departure := ""
		GetDeparture(username, &departure)
		fmt.Println(departure)
	}
}

func LogEnd() {
	time.Now()
	fmt.Println("completed program")
	fmt.Println(time.Now())

}

func GetMember() {
	fmt.Println("please wait...")
	time.Sleep(3 * time.Second)
}

func CheckServerResponse() {
	fmt.Println("check server time")
	time.Sleep(3 * time.Second)
	panic("server error")
}

func main() {
	defer func() {
		if r := recover(); r != nil { //print log panic ที่อ่านง่ายมาแสดง
			fmt.Println("recover")
			fmt.Println(r)
		}
	}()
	GetMember()
	CheckServerResponse()
}

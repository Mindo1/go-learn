package main

import "fmt"

func GetMember() {
	fmt.Println("please wait...")
}

func IsInSyetem(username string) bool {
	return true
}

func GetUserDetail(username string) (int, string) {
	
	return 201, "manager"
}

func GetDeparture(username string, departure *string) {
	if username != ""{
		*departure = "home"
	}
}

func CheckLogin (username string, password string) {
	if IsInSyetem(username) {
		fmt.Println("found user in system")
		GetUserDetail(username)
		departure := ""
		GetDeparture(username, &departure)
		fmt.Println(departure)
	}
}

func main(){
	GetMember()
}
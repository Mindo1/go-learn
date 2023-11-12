package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Map, struct,
func main() {
	fmt.Println("xxx")
	// การ map
	user := map[string]string{} //ประกาศให้เป็น object ว่างๆ

	user["username"] = "mind"
	user["password"] = "xxxxxx"
	fmt.Println(user)
	fmt.Println(user["username"])

	userProfile := UserProfile{
		Firstname: "matchima",
		Lastname:  "klom",
		Age:       20,
	}
	userProfile.Address.PostCode = "40000"
	fmt.Println(userProfile)
	fmt.Println(userProfile.Firstname)

	//ทำเป็น string ต่อกัน
	fmt.Println(userProfile.ToFullDesc())

	// แปลง struct เป็น json
	byteTxtJson, err := json.MarshalIndent(userProfile, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteTxtJson)) //แปลง จาก byte to string

	// เขียน string หลายบรรทัดใช้ ``
	dataJson := `{
	"firstname": "ouddy",
	"lastname": "moodeng",
	"Age": 12,
	"Height": 130,
	"Address": {
	 "Address": "BKK",
	 "PostCode": "40000"
	},
	"Bill": {
	 "BillAddress": "BKK"
	}
   }`

	// การแปลง json แปลงเป็น struct
	var oudProfile UserProfile
	err = json.Unmarshal([]byte(dataJson), &oudProfile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("======== Json to struct")
	fmt.Println(oudProfile)
}

type Address struct {
	Address  string
	PostCode string
}

// ปกติ ประกาศใน struct ตัวแรกต้องเป็นตัวพิมพ์ใหญ่
type UserProfile struct {
	Firstname string `json:"firstname"` // เปลี่ยนชื่อ field ในไฟล์ json
	Lastname  string `json:"lastname"`
	Age       int
	Height    float32
	// วิธีการ ref หา อีก struct นึง
	Address Address

	// วิธีการ ประกาศแบบ anonymous struct
	Bill struct {
		BillAddress string
	}
}

// ทำเป็น
func (u UserProfile) ToFullDesc() string {
	return fmt.Sprintf("%s %s %d %s", u.Firstname, u.Lastname, u.Age, u.Address.Address)
}

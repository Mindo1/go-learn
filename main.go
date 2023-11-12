package main

import "fmt"


// Interface

type SaveLogInterface interface {
	Savelog()
}

func Savelog(st SaveLogInterface) {
	st.Savelog()
}

type TransferBBL struct {
	Name string
}

func (tBBL *TransferBBL) Savelog() {
	tBBL.Name = "test"
	fmt.Println("Save to database ", tBBL.Name)
}

type TransferKTB struct {
	Name string
}

func (tKTB TransferKTB) Savelog() {
	fmt.Println("Save to database ", tKTB.Name)
}

func main() {
	transferA := TransferBBL{Name: "BBL"}
	transferB := TransferKTB{Name: "KTB"}
	Savelog(&transferA)
	Savelog(transferB)
}

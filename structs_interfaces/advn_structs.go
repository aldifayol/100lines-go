package structsinterfaces

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street string `json:"my_street"`
	City   string `json:"my_city"`
}

type Employee struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	IsRemote bool   `json:"is_remote"`
	Address
}

func AdvanceStruct() {
	address1 := Address{
		Street: "Jalan Tanah Putih",
		City:   "Manado",
	}

	employee1 := Employee{
		Name:     "Aldi",
		Age:      37,
		IsRemote: true,
		Address:  address1,
	}

	jsonData, _ := json.Marshal(employee1)
	fmt.Println(string(jsonData))
}

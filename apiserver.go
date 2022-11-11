package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type User struct {
	Id        uint32  `json:"id,string,omitempty"`
	Money     float64 `json:"money,string"`
	InProcess bool
	Market    Marketplace
}
type Users struct {
	Users map[uint32]*User
}

type Marketplace struct {
	UserId    uint32  `json:"userid,string,omitempty"`
	ServiceId uint32  `json:"serviceid,string,omitempty"`
	Cost      float64 `json:"cost,string,omitempty"`
	OrderID   uint32  `json:"orderid,string,omitempty"`
}

type moneyTrade struct {
	Idsrc  uint32  `json:"idsrc,string,omitempty"`
	Iddst  uint32  `json:"iddst,string,omitempty"`
	Amount float64 `json:"amount,string,omitempty"`
}

type idUsr struct {
	Id uint32 `json:"id,string,omitempty"`
}

func (u *Users) InitUsers() {
	u.Users = make(map[uint32]*User)
}
func (u *Users) AddUser() {
	rand.Seed(time.Now().UnixNano())
	id := rand.Uint32()
	if _, present := u.Users[id]; !present {
		u.Users[id] = &User{Id: id, Money: 0}
	} else {
		id = rand.Uint32()
		u.Users[id] = &User{Id: id, Money: 0}
	}
}
func (u *Users) Donate(amount float64, id uint32) {
	for k, v := range u.Users {

		if k == id {
			v.Money += amount
			fmt.Println(u)
		}
	}

}

func (u *Users) getInfo(id uint32) float64 {

	for k, v := range u.Users {
		if k == id {
			return v.Money
		}
	}
	return 0
}
func (u *Users) Trade(amount float64, dst uint32) {
	for k, v := range u.Users {
		if k == dst {
			v.Money += amount
		}
	}

}
func WriteCSV(objMarket Marketplace) {
	fmt.Println("otladka")
	file, err := os.Create("results.csv")
	fmt.Println(file)
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()
	data := [][]string{
		{fmt.Sprint("userID: ", objMarket.UserId), fmt.Sprint("orderID: ", objMarket.OrderID), fmt.Sprint("serviceID: ", objMarket.ServiceId), fmt.Sprint("cost: ", objMarket.Cost)},
	}
	writer := csv.NewWriter(file)
	err = writer.WriteAll(data)
	if err != nil {
		log.Println("Cannot write to CSV file:", err)
	}
	defer writer.Flush()

}

// Code generated by paladin. DO NOT EDIT.

package dbc

import (
	"encoding/json"
	"fmt"
	"os"
)

type Shop struct {
	Id         int64
	Name       string
	Desc       string
	ProductId  string
	Flag       int32
	DollarCent int64
	Icon       string
	Gold       int64
	Diamond    int64
	GoldBonus  int64
	Pack       int64
}

var tblShop map[int64]*Shop

func GetShop(id int64) *Shop {
	return tblShop[id]
}

func GetAllShop() map[int64]*Shop {
	return tblShop
}

func init() {
	file, err := os.Open("data/shop.json")
	if err != nil {
		fmt.Println("fail to open!!", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tblShop)
	if err != nil {
		fmt.Println("fail to decode!!", err)
		return
	}
}

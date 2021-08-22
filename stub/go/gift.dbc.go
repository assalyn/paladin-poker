// Code generated by paladin. DO NOT EDIT.

package dbc

import (
	"encoding/json"
	"fmt"
	"os"
)

type Gift struct {
	Id          int64
	Name        string
	Charm       int64
	CostGold    int64
	CostDiamond int64
	Priority    int64
}

var tblGift map[int64]*Gift

func GetGift(id int64) *Gift {
	return tblGift[id]
}

func GetAllGift() map[int64]*Gift {
	return tblGift
}

func init() {
	file, err := os.Open("data/gift.json")
	if err != nil {
		fmt.Println("fail to open!!", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tblGift)
	if err != nil {
		fmt.Println("fail to decode!!", err)
		return
	}
}
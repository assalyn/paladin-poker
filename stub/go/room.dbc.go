// Code generated by paladin. DO NOT EDIT.

package dbc

import (
	"encoding/json"
	"fmt"
	"os"
)

type Room struct {
	Id         int64
	Name       string
	Level      int64
	RoomType   int64
	GameType   int64
	BootAmount int64
	EnterMin   int64
	EnterMax   int64
	MaxBlind   int32
	ChaalLimit int64
	PotLimit   int64
	TipAmount  int64
}

var tblRoom map[int64]*Room

func GetRoom(id int64) *Room {
	return tblRoom[id]
}

func GetAllRoom() map[int64]*Room {
	return tblRoom
}

func init() {
	file, err := os.Open("data/room.json")
	if err != nil {
		fmt.Println("fail to open!!", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tblRoom)
	if err != nil {
		fmt.Println("fail to decode!!", err)
		return
	}
}

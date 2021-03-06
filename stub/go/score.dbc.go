// Code generated by paladin. DO NOT EDIT.

package dbc

import (
	"encoding/json"
	"fmt"
	"os"
)

type Score struct {
	Id           int64
	Name         string
	Dan          int64
	SubDan       int64
	DanStar      int64
	TotalScore   int64
	PromoteScore int64
	LevelScore   int64
}

var tblScore map[int64]*Score

func GetScore(id int64) *Score {
	return tblScore[id]
}

func GetAllScore() map[int64]*Score {
	return tblScore
}

func init() {
	file, err := os.Open("data/score.json")
	if err != nil {
		fmt.Println("fail to open!!", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tblScore)
	if err != nil {
		fmt.Println("fail to decode!!", err)
		return
	}
}

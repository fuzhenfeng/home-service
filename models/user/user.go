package user

import (
	"encoding/json"
	"fmt"
	"github.com/alphadose/haxmap"
)

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
}

func Set(user User) {
	jsonStr, errs := json.Marshal(user)
	if errs != nil {
		fmt.Println("json error:", errs)
		return
	}
	mep := haxmap.New[int, string]()
	mep.Set(user.Id, string(jsonStr))
}

func Get(userId int) (user User) {
	mep := haxmap.New[int, string]()
	jsonStr, ok := mep.Get(userId)
	if ok {
		fmt.Println("map error")
		return
	}
	temp := User{}
	errs := json.Unmarshal([]byte(jsonStr), &temp)
	if errs != nil {
		fmt.Println("json error:", errs)
		return
	}
	return temp
}

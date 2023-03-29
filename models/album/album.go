package album

import (
	"encoding/json"
	"fmt"
	"github.com/alphadose/haxmap"
)

type Album struct {
	id  int
	url string
}

type Albums struct {
	id     int
	url    string
	albums []Album
}

var mep = haxmap.New[int, string]()

func Set(album Album) {
	jsonStr, errs := json.Marshal(album)
	if errs != nil {
		fmt.Println("json error:", errs)
		return
	}
	mep.Set(album.id, string(jsonStr))
}

func Get(pictureId int) (album Album) {
	jsonStr, ok := mep.Get(pictureId)
	if ok {
		fmt.Println("map error")
		return
	}
	temp := Album{}
	errs := json.Unmarshal([]byte(jsonStr), &temp)
	if errs != nil {
		fmt.Println("json error:", errs)
		return
	}
	return temp
}

func GetByPageNo(userId int, pageNo int) (album []Albums) {
	jsonStr, ok := mep.Get(userId)
	if ok {
		fmt.Println("map error")
		return
	}
	var temp []Albums
	errs := json.Unmarshal([]byte(jsonStr), &temp)
	if errs != nil {
		fmt.Println("json error:", errs)
		return
	}
	return temp[pageNo:10]
}

func Post(userId int, albums []Albums) {
	jsonStr, ok := mep.Get(userId)
	if ok {
		fmt.Println("map error")
		return
	}
	var temp []Albums
	json.Unmarshal([]byte(jsonStr), &temp)
	newAlbums := append(temp, albums...)
	newJsonStr, errs := json.Marshal(newAlbums)
	if errs != nil {
		fmt.Println("json error:", errs)
		return
	}
	mep.Set(userId, string(newJsonStr))
}

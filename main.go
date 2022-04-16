package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type People struct {
	Data []string `json:"data"`
}

func main() {
	file, err := ioutil.ReadFile("./people.json")
	if err != nil {
		fmt.Println("ファイルが見つかりませんでした。")
	}

	var p People
	json.Unmarshal(file, &p)

	var names []string
	for _, v := range p.Data {
		names = append(names, v)
	}

	// ランダム性のリセット
	rand.Seed(time.Now().UnixNano())
	for i := range names {
		j := rand.Intn(i + 1)
		names[i], names[j] = names[j], names[i]
	}

	fmt.Println(names[0])
}

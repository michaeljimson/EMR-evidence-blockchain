package tojson

/* package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)

	}
	err = ioutil.WriteFile("test.json", b, os.ModeAppend)
	if err != nil {
		return
	}
	var data interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println("error:", err)

	}
	fmt.Println("data", data)
	dataJson := data.(map[string]interface{})["Colors"]
	fmt.Println("dataJson", dataJson)
	b11, err := json.Marshal(dataJson)
	if err != nil {
		fmt.Println("error:", err)
	}
	ioutil.WriteFile("test11.json", b11, os.ModeAppend)
	if err != nil {
		return
	}
}
*/

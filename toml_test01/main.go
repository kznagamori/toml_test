package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Item struct {
	ID   string `toml:"id"`
	Name string `toml:"name"`
}

type Config struct {
	Title string `toml:"title"`
	Items []Item `toml:"items"`
}

func main() {
	var conf Config
	if _, err := toml.DecodeFile(".\\config.toml", &conf); err != nil {
		fmt.Println("Error reading TOML file:", err)
		return
	}

	// ハッシュテーブルの構築
	itemsMap := make(map[string]Item)
	for _, item := range conf.Items {
		itemsMap[item.ID] = item
	}

	// 特定のIDを検索
	idToFind := "456"
	if item, ok := itemsMap[idToFind]; ok {
		fmt.Printf("Found item: %s with ID: %s\n", item.Name, item.ID)
	} else {
		fmt.Println("Item not found")
	}
}

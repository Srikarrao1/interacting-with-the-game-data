package main

import (
	"bufio"
	"fmt"
	"os"
	"encoding/json"
	"strconv"
	"strings"
)

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

func main() {


	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item
		genre string
	}
	type jsonGame struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Genre string `json:"genre"`
		Price int    `json:"price"`
	}

	
	var decodable []jsonGame

			if err := json.Unmarshal([]byte(data), &decodable); err != nil {
				fmt.Println("Sorry there's a problem:", err)
				return
			}	

	var games []game

	for _, dg := range decodable {
		games = append(games, game{item{dg.ID, dg.Name, dg.Price}, dg.Genre})
	}
	
	

	// games := []game{
	// 	{
	// 		item:  item{id: 1, name: "god of war", price: 50},
	// 		genre: "action adventure",
	// 	},
	// 	{item: item{id: 2, name: "x-com2", price: 30}, genre: "strategy"},
	// 	{item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
	// }

	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}

	fmt.Printf("srikar's game store has %d games.\n\n", len(games))

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
		> list : lists all the games
		> id   : queries all the games listed
		> save : exports the data to json and quits
		> quit : quits
		`)

		if !in.Scan() {
			break
		}

		cmd := strings.Fields(in.Text())

		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit":
			fmt.Println("bye!")
		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n", g.id, g.name, "("+g.genre+")", g.price)
			}
		case "id":
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}
			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}
			g, ok := byID[id]

			if ok {
				fmt.Printf("#%d: %-15q %-20s $%d", g.id, g.name, "("+g.genre+")", g.price)
			} else {
				fmt.Printf("sory no id for this game")
			}
		case "save":

			// load the data into the encodable game values

			

			var encodable []jsonGame 
				for _, g := range games {
					encodable = append(encodable, jsonGame{g.id, g.name, g.genre, g.price})
			
				}
				out, err := json.MarshalIndent(encodable, "", "\t")
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}
				fmt.Println(string(out))
				return
		}
	}

}

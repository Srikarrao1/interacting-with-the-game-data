package main

import (
	"bufio"
	"fmt"
	"os"
	"encoding/json"
	"strconv"
	"strings"
)



func main() {



	
	// var decodable []jsonGame

	// 		if err := json.Unmarshal([]byte(data), &decodable); err != nil {
	// 			fmt.Println("Sorry there's a problem:", err)
	// 			return
	// 		}	

	
	
	

	// games := []game{
	// 	{
	// 		item:  item{id: 1, name: "god of war", price: 50},
	// 		genre: "action adventure",
	// 	},
	// 	{item: item{id: 2, name: "x-com2", price: 30}, genre: "strategy"},
	// 	{item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
	// }

	games := load()
	
	byID := indexByID(games)

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

			if !ok {
				fmt.Println("sorry no Id for the game")
				continue
			}
			printGame(g)
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

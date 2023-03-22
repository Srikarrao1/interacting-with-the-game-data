package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	games := []game{
		{
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		{item: item{id: 2, name: "x-com2", price: 30}, genre: "strategy"},
		{item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
	}

	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}

	fmt.Printf("Inanc's game store has %d games.\n\n", len(games))

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
		> list : lists all the games
		> id   : queries all the games listed
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
		}
	}

}

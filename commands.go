package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)



func cmdRun(input string, games []game, byID map[int]game) bool {
	fmt.Println()

	cmd := strings.Fields(input)
	if len(cmd) == 0 {
		return true
	}

	switch cmd[0] {
	case "quit":
		return cmdQuit()

	case "list":
		return cmdList(games)

	case "id":
		return cmdByID(cmd, games, byID)

	case "save":
		return cmdSave(games)

	}
	return true

}

func cmdQuit() bool {
	fmt.Println("bye!")
	return false

}

func cmdList(games []game) bool {
	for _, g := range games {
		printGame(g)
		
	}
	return true

}

func cmdByID(cmd []string, games []game, byID map[int]game) bool {

	if len(cmd) != 2 {
		fmt.Println("wrong Id")
		return true
	}

	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong Id")
		return true
	}
	g, ok := byID[id]
	if !ok {
		fmt.Println("sorry, I don't have the game")
		return true
	}
	printGame(g)
	return true

}

func cmdSave(games []game) bool {
	var jg []jsonGame

	for _, g := range games {
		jg = append(jg, jsonGame{g.id, g.name, g.genre, g.price})
	}

	out, err := json.MarshalIndent(jg, "", "\t")
	if err != nil {
		fmt.Println("Error: can't save because", err)
		return true
	}
	fmt.Println(string(out))
	return false

}

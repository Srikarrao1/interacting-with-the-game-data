package main

import (
	"fmt"
	

)

func main() {
	phones := map[string]string {
		"bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"greco": "034989698793",
	}
	products := map[int]bool {
		617841573: true,
		879401371: false,
		576872813: true,
	}

	multiPhones := map[string][]string {
		"bowen": {"202-555-0179"},
		"dulin": {"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"greco": {"03489940240", "03489900120"},
	}
	basket := map[int]map[int]int {
		100: {617841573: 4, 576872813: 2},
		101: {576872813: 5, 657473833: 20},
		102: {},
	}

	who, phone := "dulin", "N/A"

	if v, ok := phones[who]; ok {
		phone = v
	}
	fmt.Printf("%s's phone number is %s\n", who, phone)

	id, status := 879401371, "Available"

	if !products[id] {
		status = "Not" + status
	}
	fmt.Printf("Product ID #%d is %s\n", id, status)

	who, phone = "greco", "N/A"

	if phones := multiPhones[who]; len(phones) >= 2 {
		phone = phones[1]
	}
	fmt.Printf("%s's second phone number is %s", who, phone )

	cid, pid := 101, 576872813
	fmt.Printf("Customer #%d is going to buy %d from product ID #%d.\n", cid, basket[cid][pid], pid)
}
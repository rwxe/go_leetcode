package main

import (
	"fmt"
	"src/leetcode"
)

func main() {
	fmt.Println(leetcode.AdventureCamp([]string{"leet->code", "leet->code->Campsite->Leet", "leet->code->leet->courier"}))
	fmt.Println(leetcode.AdventureCamp([]string{"Alice->Dex", "", "Dex"}))
	fmt.Println(leetcode.AdventureCamp([]string{"", "Gryffindor->Slytherin->Gryffindor", "Hogwarts->Hufflepuff->Ravenclaw"}))
}

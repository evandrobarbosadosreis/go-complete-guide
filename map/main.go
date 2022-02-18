package main

import "fmt"

func main() {

	// var colors map[strnig]string

	// colors := New(map[string]string)

	colors := map[string]string{
		"red":   "#4ds20c",
		"green": "#eed125",
		"white": "#ffffff",
	}

	// colors["blue"] = "#dsa451"

	// delete(colors, "blue")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex color for", color, "is", hex)
	}
}

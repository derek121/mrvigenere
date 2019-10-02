package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

//https://en.wikipedia.org/wiki/Vigenère_cipher

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--usage" {
		usage()
		os.Exit(0)
	}

	if len(os.Args) == 2 && os.Args[1] == "--printtable" {
		table := make_table()
		print_table(table)
		os.Exit(0)
	}

	if len(os.Args) != 4 {
		usage()
		os.Exit(1)
	}

	table := make_table()
	// print_table(table)

	// keyword := "LEMON"
	// message := "ATTACKATDAWN"
	keyword := os.Args[2]
	message := os.Args[3]

	if os.Args[1] == "--encode" {
		encoded := encode(table, message, keyword)
		// fmt.Printf("Message: %s. Encoded: %s\n", message, encoded)
		fmt.Println(encoded)
	} else if os.Args[1] == "--decode" {
		decoded := decode(table, message, keyword)
		// fmt.Printf("Message: %s. Decoded: %s\n", message, decoded)
		fmt.Println(decoded)
	} else {
		log.Fatalf("usage args: {--encode|--decode} key message")
	}
}

func encode(table [26][26]int, message string, keyword string) string {
	var b bytes.Buffer

	for i := 0; i < len(message); i++ {
		keywordChar := keyword[(i % len(keyword))]
		plainChar := message[i]
		b.WriteByte(byte(table[keywordChar-'A'][plainChar-'A']))
	}

	return b.String()
}

func decode(table [26][26]int, encoded string, keyword string) string {
	var b bytes.Buffer

	for i := 0; i < len(encoded); i++ {
		keywordChar := keyword[(i % len(keyword))]
		encodedChar := encoded[i]

		// Determine how far the encoded char (in the table) is from the key char (the row index)
		// We can then use that offset from 'A' for the answer
		var diff byte
		if encodedChar >= keywordChar {
			diff = encodedChar - keywordChar
		} else {
			diff = ('Z' - keywordChar) + (encodedChar - 'A') + 1
		}

		decodedChar := 'A' + diff
		b.WriteByte(byte(decodedChar))
	}

	return b.String()
}

func make_table() [26][26]int {
	table := [26][26]int{}
	for row := 0; row < 26; row++ {
		for col := 0; col < 26; col++ {
			table[row][col] = ((row + col) % 26) + 'A'
		}
	}

	return table
}

func print_table(table [26][26]int) {
	print_col_header()

	for row := 0; row < 26; row++ {
		fmt.Printf("%c|", row+'A')
		for col := 0; col < 26; col++ {
			fmt.Printf("%c ", table[row][col])
		}
		fmt.Println()
	}
}

func print_col_header() {
	fmt.Printf(" |")
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", 'A'+i)
	}
	fmt.Println()

	for i := 0; i < 53; i++ {
		fmt.Printf("-")
	}
	fmt.Println()
}

func usage() {
	log.Printf("usage args: {{--encode|--decode} key message} | --printtable | --usage")
}

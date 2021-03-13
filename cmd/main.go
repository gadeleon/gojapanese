package main

import (
	"bufio"
	"fmt"
	"github.com/gadeleon/gojapanese/words"
	"math/rand"
	"os"
	"strings"
	"time"
)
// HardCode 3 Nouns for now, MVP!


func main() {

	reader  := bufio.NewReader(os.Stdin)
	n := []gojapanese.Noun{
		{"石"},
		{"本"},
		{"木"},
	}
	fmt.Println(n)
	// Pick a rando noun from n
	rand.Seed(time.Now().Unix())
	q := n[rand.Intn(len(n))]
	// Display, Ask for Input
	fmt.Printf("What's the negative form of: %s", q)
	text, _ := reader.ReadString('\n')
	// Eval Input
	text = strings.TrimSuffix(text, "\n")
	if !q.IsNegative(text) {
		fmt.Printf("NAh..... --%s-- does not match --%s--\n", text, q.Negative())
		return
	}
	fmt.Printf("Correct!!!!!! --%s-- matches  --%s--\n", text, q.Negative())

//TODO: Abstract Question to function
//TODO: Test question function
//TODO: Make Simple Webserver
//TODO: Have WebServer Handle Question Function/Eval/Etc
//TODO: Get words from Wanikani
//TODO: Write Grammar Engine
//TODO: Madlib Generator
//TODO: Picture Engine (Google Image Search??? Safe Searh...On? I'm Feeling Lucky?)



}





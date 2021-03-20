package main

import (
	"bufio"
	"fmt"
	"github.com/gadeleon/gojapanese/proctor"

	"os"
	"strings"
)
// HardCode 3 Nouns for now, MVP!


func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	return text
}

func main() {


    pr := gojapanese.NewProctor()
    fmt.Println("Here have a question.  Press (ctrl-alt-esq-w-p-r) to escape")
	fmt.Println("Type next to get a new question.")

    //text, _ := reader.ReadString('\n')
	//text = strings.TrimSuffix(text, "\n")


	text := ""
	fmt.Println(pr.Prompt)
	for {
		switch text {
		case "next":
			pr.GetQuestion()
			fmt.Println(pr.Prompt)
			text = GetInput()
		case "quit":
			fmt.Println("Quitters never prosper")
			//break
			os.Exit(0)
		default:
			//fmt.Println("B Final answer",text)
			text = GetInput()
		}



	}
		//
		//fmt.Println("Under Construction. Thank you for using go Japanese.")
		//pr.GetQuestion()

	}

	//n := []gojapanese.Noun{
	//	{"石"},
	//	{"本"},
	//	{"木"},
	//}
	//fmt.Println(n)
	//// Pick a rando noun from n
	//rand.Seed(time.Now().Unix())
	//q := n[rand.Intn(len(n))]
	//// Display, Ask for Input
	//fmt.Printf("What's the negative form of: %s", q)
	//text, _ := reader.ReadString('\n')
	//// Eval Input
	//

	//if !q.IsNegative(text) {
	//	fmt.Printf("NAh..... --%s-- does not match --%s--\n", text, q.Negative())
	//	return
	//}
	//fmt.Printf("Correct!!!!!! --%s-- matches  --%s--\n", text, q.Negative())


//TODO: Abstract Question to function
//TODO: Test question function
//TODO: Make Simple Webserver
//TODO: Have WebServer Handle Question Function/Eval/Etc
//TODO: Get words from Wanikani
//TODO: Ask webapp to give me words to use
//TODO: Write Grammar Engine
//TODO: Madlib Generator
//TODO: Picture Engine (Google Image Search??? Safe Searh...On? I'm Feeling Lucky?)








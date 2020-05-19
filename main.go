package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	user, _  := user.Current()
	fmt.Println("Hi!", user.Name)
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
			os.Exit(1)
		} else {
			fmt.Println("Posible options:")
			fmt.Println("\t - hi")
		}
	}
}

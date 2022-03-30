package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var text2 string

func main() {
	fmt.Println("Вводим, не задерживаем! ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println(" \n go \n")

	fmt.Println(text, " \n")

	text2 = strings.Trim(text, "\n")

	fmt.Println(text2, " \n")

}

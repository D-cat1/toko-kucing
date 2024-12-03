package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func HandleLongInput(text *string) {
	reader := bufio.NewReader(os.Stdin)
	dataInput, _ := reader.ReadString('\n')
	*text = strings.TrimSpace(dataInput)
}

func Clrscr() {
	fmt.Print("\033[H\033[2J")
}

func border(typeBound string, text string, length int) string {
	var merger string
	if len(text) != 0 {
		text = " " + text + " "
	}
	length -= len(text)

	var leftBorder = length / 2
	var rightBorder = length - leftBorder

	for i := 0; i < leftBorder; i++ {
		merger += typeBound
	}
	merger += text
	for i := 0; i < rightBorder; i++ {
		merger += typeBound
	}
	return merger
}

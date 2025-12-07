package main

import (
	"bufio"
	"fmt"
	"os"
	mycheck "sprint-2/custom-error"
	reflect_spell "sprint-2/reflect-spell"
	"strings"
)

// Function for Task #1:
func CustomError() {
	for {
		fmt.Printf("Укажите строку (q для выхода): ")
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = strings.TrimRight(ret, "\n")
		if ret == `q` {
			break
		}
		if err = mycheck.MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(`Строка прошла проверку`)
		}
	}
}

func main() {

	// task1: custom-error
	// CustomError()

	// task2: reflect-spell
	reflect_spell.ReflectSpell()
}

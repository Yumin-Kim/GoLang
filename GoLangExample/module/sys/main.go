package main

import (
	"fmt"

	"github.com/Yumin-Kim/GoLang/GoLangExample/module/greeting"
)

func main(){
	message := greeting.Hello("john")
	fmt.Println(message)
}
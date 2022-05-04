package main

import (
	"GoLangExample/usepkg/custompkg"
	"fmt"

	"github.com/guptarohit/asciigraph"
)

func main(){
	custompkg.PrintCustomPkg();

	data := []float64{3,4,5,6,7,8,9,10,2,3,4,5}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
package testlib

import "fmt"

var pop map[string]string;

func init(){
	pop = make(map[string]string)
	pop["Adele"] = "Hello World"
	pop["alicia Keys"] = "Fallin'"
	pop["John Lefend"] ="All of Me"
} 

func GetMusic(signer string) string{
	return pop[signer]
}

func getKeys(){
	for _ , kv := range pop{
		fmt.Println(kv)
	}
}
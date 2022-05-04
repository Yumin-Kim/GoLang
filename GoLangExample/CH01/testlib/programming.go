package testlib

import "fmt"

var langs map[int]string;

func init(){
	langs = make(map[int]string)
	langs[0] = "CLang"
	langs[1] = "CPPLang"
	langs[2] = "GoLang"
	langs[3] = "JavaScriptLang"
}

func GetLang(key int) string{
	return langs[key]
}

func SetLangName(key int , changeName string){
	langs[key] = changeName;
}

func GetKeys(){
	for _ , kv := range langs{
		fmt.Println(kv);
	}
}
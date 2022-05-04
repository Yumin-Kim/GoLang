//https://velog.io/@comdori-web/Go-package%EC%99%80-module
package main

import (
	"fmt"
)

// 원형 정의
type Calculator func(int , int) int;
const global = 10;

func returnString() string{
	hello := "Hello";
		var (
		str1 = `Hello World
		Hello World
		Hello World
		Hello WorldHello WorldHello WorldHello WorldHello WorldHello World` 
		str2 = "Hello01" 
		str3 = "Hello01" 
	);
	println(str1,str2,str3)
	return hello;
}
// ./ch01.go:7:2: too many arguments to return
//         have (string)
//         want ()
func typeConverslon(){
	var i int = 100;
	var u uint = uint(i);
	var f float32 = float32(i);
	println(u , f)
	str := "ABC";
	str3 := "a";
	bytes := []byte(str);
	bytes2 := []byte(str3);
	str2 := string(bytes);
	fmt.Println(bytes2);
	fmt.Println(bytes)
	println(&bytes)
	fmt.Println(bytes,bytes2,str2)
}

func condition(){
	b := 10;
	if b > 10{
		println("b > 10")
	}else if b == 10{
		println("b == 10")
	}else if b < 10{
		println("b < 10")
	}

}

func sum(nums ...int) (count int, total int) {
    // s := 0      // 합계
    // count := 0  // 요소 갯수
    // for _, n := range nums {
    //     s += n
    //     count++
    // }
    // return count, s
	for _, n := range nums {
        total += n
    }
    count = len(nums)
    return
}

func calc(f func(int ,int) int , a int , b int ) (result int){
	result = f(a,b);
	return
}
// 일급 함수으로써 변수에 함수 할당 가능
// https://velog.io/@mgm-dev/%EC%9D%BC%EA%B8%89%ED%95%A8%EC%88%98%EB%9E%80-%EB%AC%B4%EC%97%87%EC%9D%B8%EA%B0%80
func firstClassFunc(){
	add := func (i int , j int ) int {
			return i + j;
		}	
	r1 := calc(add,10,20)
	println(r1)

	r2 := calc(func(x int , y int)(int){return x -y},10,20)
	println(r2)
}

func calcType(f Calculator , a int , b int)(int){
	result := f(a,b);
	return result;
}

func nextValue() func () (int) {
	i := 0;
	return func ()(int)  {
		i++;
		return i;
	}
}
// 일급함수를 사용할 수 있기에 가능한 동작
func closureFunc(){
	next := nextValue()
	println(next())
	println(next())
	println(next())
	
	anotherNext := nextValue()
	println(anotherNext())
	println(anotherNext())
}

func arrayFunc(){
	var a [3]int;
	var a1 = [3]int{2,4,6}; // 정적 할당
	var a2 = [...]int{3,6,9}; // 동적 할당
	var multiArray [3][4]int;
	var multiArray1 = [2][3] int {
		{1,2,3},
		{4,5,6},
	}
	multiArray[0][1] = 1;
	a[0] = 1;
	a[1] = 2;
	a[2] = 3;
	fmt.Println(a);
	fmt.Println(a1[2])
	fmt.Println(a2[2])
	fmt.Println(multiArray[0])
	fmt.Println(multiArray1)
	println(multiArray1[0][0])
}

func mapFunc(){
	//정의
	var idMap map[int]string;
	idMap = make(map[int]string)
	//할당
	idMap[0] = "Hello1";
	idMap[1] = "Hello2";
	idMap[2] = "Hello3";
	idMap[3] = "Hello4";
	//정의 및 할당 동시에
	ticker := map[string]string{
		"Hello1" :"World1",
		"Hello2" :"World1",
		"Hello3" :"World1",
	}
	myMap := map[string]string{
		"A":"Apple",
		"B":"Banana",
		"C":"Charlie",
	}
	
	fmt.Println(idMap)
	fmt.Println(myMap)
	fmt.Println(ticker["Hello1"])
	//map에서 원하는 키 값 삭제
	delete(idMap,2)
	fmt.Println(idMap)

	//map 키 체크
	val , validKey := ticker["Hello4"];
	if !validKey {
		fmt.Println("NOt Hello4")
	}else {
		fmt.Println(val)
	}

	for key , val  := range myMap {
		fmt.Println(key , val)
	}
	
}

//https://developer.mozilla.org/ko/docs/Web/JavaScript/Closures
func main(){
	println("Main Func")
	k := 10;
	var p = &k;
	fmt.Println(*p) 
	// condition()
	count1, total := sum(1, 7, 3, 5, 9)
	println(count1, total)
	// 사용자 외부 모듈 구현
	// https://pronist.dev/86
	// song :=  testlib.GetMusic("Adele");
	// lang := testlib.GetLang(1);
	// fmt.Println(lang)
	// fmt.Println(song)
	// mapFunc()
	// arrayFunc()
	// result := calcType(func (x int , y int)(int){return x + y},30,100)
	// println(result)
	// typeConverslon();
	// println(hello+"");
}


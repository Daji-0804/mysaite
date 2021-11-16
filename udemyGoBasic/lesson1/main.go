package main

//format package printとか
import (
	"fmt"
	"time"
)

//Hello World
/*
複数行の
コメントアウト
*/

//明示的な場合は関数外でも定義可能
var i5 int = 500

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
	fmt.Println(i5)
}

func main() {
	//改行表示
	fmt.Println("Hello world")
	fmt.Println(time.Now())
	//明示的な定義
	/*
		静的型付けのメリットでもあるので基本的に型付けする
	*/
	//var 変数名　型　＝　値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i = 300
	s3 = "Go"
	fmt.Println(i, s3)

	i = 150
	fmt.Println(i)

	//暗黙的な定義
	//変数名 :=値
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	fmt.Println(i5)

	outer()
}

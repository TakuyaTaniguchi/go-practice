# memo

## https://go-tour-jp.appspot.com/basics/4
- 外部パッケージから呼び出したメソッドは大文字で呼び出せる
```go
import "encoding/json"
json.Marshal(todo)
```

```go
// 関数　型は変数の後に書く。
func substruct(x int, y int) int {
 return x - y
}

// 型を省略もできる。tsを書いてる身からするとしたくない
func add(x, y int) int {
return x + y
}

```
## Interesting
変数名と一緒に代入する
:= (短縮変数宣言演算子)


- 小文字のフィールドはパッケージ外部に公開されない。
- 基本的にGOは値渡しであり、元の値を変えたいときは、明示してポインタを記述する"&c,*Person "

```go

func swap(x, y string) (string, string) {
return y, x
}


type Person struct {
Name    string
Age     int
Address string
sex string
}

// Person構造体を返す関数
func createPerson(name string, age int) Person {
    return Person{
    Name: name,
    Age:  age,
    Address: "東京都新宿区",
    sex: "男性",
    }
}

func changeAddress(p *Person) error {
    p.Address = "京都府京都市"
    return nil
}

func main() {
    c :=createPerson("Taro",20)
    changeAddress(&c)
    fmt.Println(c) // {Taro 20 京都府京都市 男性}
}
```


GOは言語仕様で複数の戻り値を持つことができる。
jsなどでは、一度オブジェクトにするが、Goではその必要がない。

名前付き戻り値の主なメリットは、

関数の定義を見るだけで、どんな値が返るのか分かりやすい
長い関数で、戻り値の変数に途中で値を代入しておける
空の return 文だけで、現在の名前付き変数の値が返される

```go

// Goでの複数の戻り値
func getValues() (int, string) {
    return 42, "hello"
}

a, b := getValues() // a=42, b="hello"

func getValues() (num int, text string) {
num = 42
text = "hello"
return  // 空のreturnでOK
}

```

文字列変換について
GOでは 型(値)で変換できるが、文字列はUnicodeになるのでライブラリを使う。
メソッド名がItos(int to string)ではないのは、伝統的に、C言語では、ASCIIを使っているため
```go


func main() {
    var id  int = 100
    mozi  := strconv.Itoa(id)
    fmt.Print(mozi)

}


```

変数について`var` `const`はjsとあまり変わらない。強いて言えばconstは`:=`が使えない。

### for
for 分の引数はカンマではなくセミコロン

初期化と後処理は省略できる。（whileみたいなもの）
条件式の評価が false となった場合にイテレーションを停止する

```go
func main() {
	sum := 10
	for i:=0; i<sum; i++ {
        fmt.Println(i)
	}
}


func main() {
    sum := 1
    for sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)
}

//　無限ループ

func main() {
    for {
    }
}

```

### if


少し番外編
このコードは下記のようなエラーが出る。

```go
./prog.go:9:9: too many return values
	have (int)
	want ()
./prog.go:13:10: ageLimit(10) (no value) used as value

```

```go
package main

import (
	"fmt"
)


func ageLimit(age int) {
	return age
}

func main() {
	test := ageLimit(10)
	fmt.Println(test)
}

```

エラーの読み方

```go
「too many return values」というエラーメッセージの意味は、「関数が宣言されている戻り値の数より多くの値を返そうとしている」ということです。
have は「対象のコードが実際に返そうとしている値の型」
want は「関数宣言で指定された戻り値の型」
```


動くコード

```go
func ageLimit(age int) string {
	if age > 20 {
		return "OK"
	} else {
		return "NG"
	}
	
}

func main() {
	test := ageLimit(100)
	fmt.Println(test)
}
```

条件前にステートメントをかける。
この場合`add`はif内でしか参照できないので、スコープを絞ることができる。
```go

func ageLimit(age int) string {
    if add := age; + 10 > 20 {
        fmt.Println(add)
        return "OK"
    } else {
        return "NG"
    }

}

```
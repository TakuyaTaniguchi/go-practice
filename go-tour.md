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

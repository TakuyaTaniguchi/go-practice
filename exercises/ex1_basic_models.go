/*
演習1: 基本的なGoモデルの作成

この演習では、Go言語での構造体（struct）とカスタム型の定義を学びます。
Web APIで使用するデータモデルを作成します。
*/

package main

import (
	"encoding/json"
	"fmt"
)

// 問題1: 優先度を表すカスタム型を作成してください
// 以下の3つの値を持つカスタム型を定義: LOW, MEDIUM, HIGH
// ヒント: Goでは型エイリアスと定数を使って列挙型を表現します
// TODO: ここにPriority型と定数を実装
type Priority string

const (
	// ここにPriorityLow, PriorityMedium, PriorityHighを定義してください
)

// TodoCreate型の定義
// 問題2: TodoCreate構造体を作成してください
// 以下のフィールドが必要です:
// - Title: 文字列（必須）
// - Description: 文字列（省略可能）
// - Priority: Priority型（デフォルトはMEDIUM）
// - Completed: 真偽値（デフォルトはfalse）
// ヒント: JSONタグも設定してください
// TODO: ここにTodoCreate構造体を実装
type TodoCreate struct {
	// ここにフィールドを定義してください
}

// Todo型の定義
// 問題3: Todo構造体をTodoCreateを拡張して作成してください
// TodoCreateの全フィールドに加えて、以下のフィールドが必要:
// - ID: 文字列（必須）
// ヒント: 構造体の埋め込みを使用してください
// TODO: ここにTodo構造体を実装
type Todo struct {
	// ここにIDフィールドと埋め込みを定義してください
}

// サンプル解答の確認用関数
func main() {
	// 解答を確認するために使用するコード
	// 実装したら、以下のコメントを外して実行してみてください

	/*
	// TodoCreateの作成と確認
	todoCreate := TodoCreate{
		Title:       "牛乳を買う",
		Description: "帰りにスーパーで",
		Priority:    PriorityHigh,
		Completed:   false,
	}
	
	// JSONに変換して出力
	jsonData, _ := json.MarshalIndent(todoCreate, "", "  ")
	fmt.Println("TodoCreate:")
	fmt.Println(string(jsonData))
	
	// Todoの作成と確認
	todo := Todo{
		TodoCreate: todoCreate,
		ID:         "1",
	}
	
	// JSONに変換して出力
	jsonData, _ = json.MarshalIndent(todo, "", "  ")
	fmt.Println("\nTodo:")
	fmt.Println(string(jsonData))
	*/
}
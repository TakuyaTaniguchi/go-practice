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
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
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
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Priority    Priority `json:"priority"`
	Completed   bool     `json:"completed"`
}

// Todo型の定義
// 問題3: Todo構造体をTodoCreateを拡張して作成してください
// TodoCreateの全フィールドに加えて、以下のフィールドが必要:
// - ID: 文字列（必須）
// ヒント: 構造体の埋め込みを使用してください
// TODO: ここにTodo構造体を実装
type Todo struct {
	// ここにIDフィールドと埋め込みを定義してください
	Id int `json:"id"`
	TodoCreate
}

// サンプル解答の確認用関数
func main() {
	todoCreate := TodoCreate{
		Title:       "Title",
		Description: "Description",
		Priority:    PriorityMedium,
		Completed:   false,
	}

	todo := Todo{
		TodoCreate: todoCreate,
		ID:         "1",
	}

	b, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}

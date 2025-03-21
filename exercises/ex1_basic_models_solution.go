///*
//演習1: 基本的なGoモデルの作成 (解答例)
//
//この演習では、Go言語での構造体（struct）とカスタム型の定義を学びます。
//Web APIで使用するデータモデルを作成します。
//*/
//
//package main
//
//import (
//	"encoding/json"
//	"fmt"
//)
//
//// Priority型の定義 - カスタム文字列型
//type Priority string
//
//// Priority型の定数
//const (
//	PriorityLow    Priority = "low"
//	PriorityMedium Priority = "medium"
//	PriorityHigh   Priority = "high"
//)
//
//// TodoCreate型の定義
//type TodoCreate struct {
//	Title       string   `json:"title"`
//	Description string   `json:"description,omitempty"`
//	Priority    Priority `json:"priority,omitempty"`
//	Completed   bool     `json:"completed,omitempty"`
//}
//
//// Todo型の定義 - TodoCreateを埋め込み
//type Todo struct {
//	TodoCreate
//	ID string `json:"id"`
//}
//
//// 実行して確認
//func main() {
//	// TodoCreateの作成と確認
//	todoCreate := TodoCreate{
//		Title:       "牛乳を買う",
//		Description: "帰りにスーパーで",
//		Priority:    PriorityHigh,
//		Completed:   false,
//	}
//
//	// JSONに変換して出力
//	jsonData, _ := json.MarshalIndent(todoCreate, "", "  ")
//	fmt.Println("TodoCreate:")
//	fmt.Println(string(jsonData))
//
//	// Todoの作成と確認
//	todo := Todo{
//		TodoCreate: todoCreate,
//		ID:         "1",
//	}
//
//	// JSONに変換して出力
//	jsonData, _ = json.MarshalIndent(todo, "", "  ")
//	fmt.Println("\nTodo:")
//	fmt.Println(string(jsonData))
//}
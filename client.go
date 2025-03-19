package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Todo APIとやり取りするためのサンプルクライアント

func main() {
	fmt.Println("Todo APIクライアント")
	fmt.Println("====================")
	
	baseURL := "http://localhost:8080"
	
	// Todoを作成
	todo := map[string]interface{}{
		"title":       "サンプルタスク",
		"description": "これはサンプルのタスクです",
		"priority":    "high",
		"completed":   false,
	}
	
	// POSTリクエストでTodoを作成
	todoJSON, _ := json.Marshal(todo)
	resp, err := http.Post(baseURL+"/todos", "application/json", bytes.NewBuffer(todoJSON))
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	
	// レスポンスを読み取り
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	
	// 作成されたTodoを表示
	var createdTodo map[string]interface{}
	json.Unmarshal(body, &createdTodo)
	
	fmt.Println("\n作成されたTodo:")
	printJSON(createdTodo)
	
	// 作成されたTodoのID
	id := createdTodo["id"].(string)
	
	// GETリクエストで全Todoを取得
	resp, err = http.Get(baseURL + "/todos")
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	
	// レスポンスを読み取り
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	
	// 全Todoを表示
	var allTodos []map[string]interface{}
	json.Unmarshal(body, &allTodos)
	
	fmt.Println("\n全Todos:")
	for i, t := range allTodos {
		fmt.Printf("\nTodo %d:\n", i+1)
		printJSON(t)
	}
	
	// PUTリクエストでTodoを更新
	updatedTodo := map[string]interface{}{
		"title":       "更新されたタスク",
		"description": "これは更新されたタスクです",
		"priority":    "medium",
		"completed":   true,
	}
	
	updatedJSON, _ := json.Marshal(updatedTodo)
	req, _ := http.NewRequest("PUT", baseURL+"/todos/"+id, bytes.NewBuffer(updatedJSON))
	req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	
	// レスポンスを読み取り
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	
	// 更新されたTodoを表示
	var updatedTodoResp map[string]interface{}
	json.Unmarshal(body, &updatedTodoResp)
	
	fmt.Println("\n更新されたTodo:")
	printJSON(updatedTodoResp)
	
	fmt.Println("\nNote: このクライアントを実行する前に 'go run app/main.go' でサーバーを起動してください。")
}

// JSONを整形して表示する関数
func printJSON(data interface{}) {
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(jsonData))
}
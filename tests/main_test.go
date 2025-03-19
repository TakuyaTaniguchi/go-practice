package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/gorilla/mux"
)

// これはサンプルのテストファイルです
// 実際のアプリケーションコードをテストするには、適切にインポートして実装する必要があります

func TestCreateTodoHandler(t *testing.T) {
	// このテストはサンプルとして提供されています
	// 実際のアプリケーションコードをテストする際は、適切なインポートと実装が必要です
	
	t.Skip("このテストは実装サンプルです。実際のテストを書く際に参考にしてください。")
	
	// リクエストボディを作成
	todoJSON := `{"title":"テストタスク","description":"テスト用のタスクです","priority":"high"}`
	req, err := http.NewRequest("POST", "/todos", bytes.NewBufferString(todoJSON))
	if err != nil {
		t.Fatal(err)
	}
	
	// レスポンスレコーダーを作成
	rr := httptest.NewRecorder()
	
	// ハンドラを作成（実際のアプリケーションコードをテストする場合は、アプリのハンドラを使用）
	router := mux.NewRouter()
	//router.HandleFunc("/todos", createTodoHandler).Methods("POST")
	
	// ハンドラにリクエストを送信
	router.ServeHTTP(rr, req)
	
	// レスポンスコードをチェック
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("ハンドラが不適切なステータスコードを返しました: got %v want %v", status, http.StatusCreated)
	}
	
	// レスポンスボディをチェック
	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	
	// 期待する値をチェック
	if response["title"] != "テストタスク" {
		t.Errorf("ハンドラが不適切なタイトルを返しました: got %v want %v", response["title"], "テストタスク")
	}
}
// タスクのクエリサービスインターフェース
package query

import "context"

// TaskView はクライアントに返すタスクのビューモデルです。
type TaskView struct {
	ID       string
	Name     string
	UserName string
	Labels   []string
}

// TaskQueryService はタスクのクエリサービスのインターフェースです。
type TaskQueryService interface {
	FindTaskViews(ctx context.Context, userID string) ([]TaskView, error)
}

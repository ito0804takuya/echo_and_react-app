// migrate実行時は、migrate処理をエントリーポイントにしたいためmainに
package main

import (
	"echo-api/db"
	"echo-api/model"
	"fmt"
)

// 実行コマンド: GO_ENV=dev go run migrate/migrate.go
func main() {
	// DBとの接続を生成
	dbConn := db.NewDB()

	defer fmt.Println("Success Migrated.")
	// 忘れずに接続を閉じる
	defer db.CloseDB(dbConn)

	// AutoMigrate(空の構造体のポインタ) を実行すると、その構造体に対応するテーブルが作成される
	// NOTE: AutoMigrateは新規作成のみ実行する。つまり、スキーマの変更などは、別で処理しないといけない
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
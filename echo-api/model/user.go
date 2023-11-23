package model

import "time"

type User struct {
	// json:"id" : jsonに変換するとidという小文字のキーにするように指定
	// gorm:"primaryKey : gormを使うときに主キーとして認識される
	ID        uint      `json:"id" gorm:"primaryKey`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// レスポンス用のユーザ型
type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey`
	Email string `json:"email" gorm:"unique"`
}

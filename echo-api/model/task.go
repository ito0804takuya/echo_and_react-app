package model

import "time"

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// foreignKey : Userとのリレーション
	// constraint(制約) : OnDelete:CASCADEで、Userが削除されたらTaskも削除するよう指定
	User User `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	// 1対多
	UserId uint `json:"user_id" gorm:"not null"`
}

type TaskResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

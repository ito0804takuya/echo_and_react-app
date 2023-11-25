package repository

import (
	"echo-api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	// emailが一致するユーザ情報を取得し、user引数のポインタ先のUserオブジェクトを取得したユーザ情報に置き換える
	// NOTE: ur.db.Where("email=?", email).First(user) で、SELECTは実行している。
	//       エラー判定用のerr変数を用意する処理の流れを、.Where().First().Errorでまとめて行っている
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
	  return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	// user引数のポインタ先のUserオブジェクトを、作成したユーザ情報に置き換える
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
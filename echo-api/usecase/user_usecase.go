package usecase

import (
	"echo-api/model"
	"echo-api/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

// Usecaseが（Repositoryでなく）Repositoryのインターフェースに依存する
type userUsecase struct {
	ur repository.IUserRepository
}

// Repositoryのインターフェースを注入（DI）するためのコンストラクタ
func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	// 平文のパスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}

	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}

	// レスポンス用に加工
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}

	// 入力されたパスワードと保存されている（ユーザの）パスワードが等しいかチェック
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// トークンのペイロード
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(), // トークン有効期限
	})
	// tokenStringはユーザに返却する認証情報トークン
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

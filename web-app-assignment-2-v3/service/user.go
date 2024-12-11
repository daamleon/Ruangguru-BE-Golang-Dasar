package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

type UserService interface {
	Register(user *model.User) (model.User, error)
	Login(user *model.User) (token *string, err error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepository repo.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Register(user *model.User) (model.User, error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return *user, err
	}

	if dbUser.Email != "" || dbUser.ID != 0 {
		return *user, errors.New("email already exists")
	}

	user.CreatedAt = time.Now()

	newUser, err := s.userRepo.CreateUser(*user)
	if err != nil {
		return *user, err
	}

	return newUser, nil
}

func (s *userService) Login(user *model.User) (token *string, err error) {
	// Cari pengguna berdasarkan email
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		// Asumsikan jika error bukan `nil`, pengguna tidak ditemukan
		return nil, errors.New("invalid email or password")
	}

	// Verifikasi password
	if dbUser.Password != user.Password {
		return nil, errors.New("invalid email or password")
	}

	// Buat token JWT
	jwtToken, err := generateJWT(dbUser.ID)
	if err != nil {
		return nil, err
	}

	return &jwtToken, nil
}



// Fungsi helper untuk menghasilkan token JWT
func generateJWT(userID int) (string, error) {
	// Buat klaim untuk token
	claims := &model.Claims{UserID: userID}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tanda tangani token dengan secret key
	signedToken, err := token.SignedString(model.JwtKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *userService) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	// Panggil repository untuk mendapatkan data tugas pengguna berdasarkan kategori
	userTaskCategories, err := s.userRepo.GetUserTaskCategory()
	if err != nil {
		return nil, err
	}

	return userTaskCategories, nil
}

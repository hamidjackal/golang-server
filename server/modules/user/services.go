package user

import (
	"errors"
	"server/modules/core"
	"server/modules/dbManager"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func (u UserService) New() UserService {
	return UserService{
		db: dbManager.GetDb(),
	}
}

func (u UserService) list(c *fiber.Ctx) error {

	var users []User

	result := u.db.Find(&users)

	if result.Error != nil {
		println("Result Error", result.Error)
	}

	newResp := core.SuccessResponse[User]{}.New(c)

	return newResp.List(users)

}

func (u UserService) signin(credentials *SignIn) (SignedInUser, error) {
	var signedInUser SignedInUser
	user, err := u.userByEmail(credentials.Email)
	if err != nil {
		return signedInUser, err
	}

	isValid := core.IsPasswordValid(user.Password, credentials.Password)

	if isValid == false {
		return signedInUser, errors.New("Invalid credentials")
	}

	token := u.generateToken(user.ID)
	signedInUser.User = user
	signedInUser.Token = token

	return signedInUser, nil
}

func (u UserService) generateToken(userId uint) string {
	// Create the Claims
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		println("Token Generation err", err)
	}
	return t
}

func (u UserService) userById(id string) User {
	var user User

	result := u.db.First(&user, id)

	if result.Error != nil {
		println("Result Error", result.Error)
	}

	return user
}

func (u UserService) userByEmail(email string) (User, error) {
	var user User

	result := u.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return user, errors.New("Invalid credentials")
	}

	return user, nil
}

func (u UserService) create(p *SignUp) (User, error) {
	user := User{
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Password:  p.Password,
	}

	result := u.db.Create(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (u UserService) delete(id string) error {
	result := u.db.Delete(&User{}, id)
	if result.Error != nil {
		return errors.New("Invalid id")
	}
	return nil
}

func (u UserService) getOne(id string) (User, error) {
	var user User
	result := u.db.First(&user, id)
	if result.Error != nil {
		return user, errors.New("Invalid id")
	}
	return user, nil
}

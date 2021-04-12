package handlers

import (
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/nid90/kawaii-blog-engine/models"
	"golang.org/x/crypto/bcrypt"
	// jwtware "github.com/gofiber/jwt/v2"
	"time"
)

func SignInView(ctx *fiber.Ctx) error {
	return ctx.Render("auth/new", nil, "layout/main")
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SignIn(ctx *fiber.Ctx) error {
	type SignInData struct {
		Email     string
		Password  string
	}
	signInData := new(SignInData)

	if err := ctx.BodyParser(signInData); err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}


	author, err := models.FindAuthorByEmail(signInData.Email)

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	if author == nil {
		return ctx.Render("auth/new", fiber.Map{"Error": "Author not found"}, "layout/main")
	}

	if CheckPasswordHash(signInData.Password, author.Password) == false {
		return ctx.Render("auth/new", fiber.Map{"Error": "Incorrect email/password"}, "layout/main")
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = author.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	ctx.Set("Authorization", t)
	ctx.Cookie()

	return ctx.Redirect("/posts")
}

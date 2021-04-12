package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func SignInView(ctx *fiber.Ctx) error {
	return ctx.Render("auth/new", nil, "layout/main")
}

func SignIn(ctx *fiber.Ctx) error {
	return ctx.Redirect("/posts")
}

package routes

import (
	"ubm-canteen/handlers"
	"ubm-canteen/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, sellerHandlers *handlers.SellerHandler) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":"404 not found",
		})
	});
	
	v1:= app.Group("/api/v1")
	
	auth:= v1.Group("/auth")

	//SELLER
	cctv:= auth.Group("/cctv")
	cctv.Post("/register", sellerHandlers.RegisterSeller)
	cctv.Post("/login", sellerHandlers.LoginSeller)
}

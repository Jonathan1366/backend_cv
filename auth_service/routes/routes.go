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
	// cctv.Post("/otp/send", sellerHandlers.SendOTP)
	// cctv.Post("/otp/verify", sellerHandlers.VerifyOTP) 
	
	// sellerPrivate := auth.Group("/seller", middleware.AuthMiddleware("seller"))
	// sellerPrivate.Post("/logout", sellerHandlers.LogoutSeller)
	// sellerPrivate.Put("/store/location", sellerHandlers.StoreLocSeller)

	// //S3 BUCKET
	// // seller.Post("/presignurl", Seller.GeneratePresignedUploadURL)
	
	// //USER
	// user:= auth.Group("/user")
	// user.Post("/register", userHandlers.RegisterUser)
	// user.Post("/login", userHandlers.LoginUser)
	// user.Post("/logout", userHandlers.LogoutUser)

	// //OAUTH2
	// google:= auth.Group("/google")
	// google.Post("/login", googleHandlers.GoogleSignIn)
}

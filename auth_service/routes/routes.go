package routes

import (
	"ubm-canteen/handlers"
	"ubm-canteen/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, authHandler *handlers.authHandler, cameraHandler *handlers.cameraHandler, detectionHandler *handlers.detectionHandler) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to HIGO CV API",
		})
	});
	
	v1:= app.Group("/api/v1")
	
	auth:= v1.Group("/auth")

	auth.Post("/register", authHandler.RegisterUser) // Endpoint untuk mendaftarkan user baru
	auth.Post("/login", authHandler.LoginUser)       // Endpoint untuk login dan mendapatkan JWT

	apiProtected := v1.Group("", middleware.AuthMiddleware()) // Middleware untuk memeriksa JWT

	cameras := apiProtected.Group("/cameras")
	cameras.Post("/", cameraHandler.CreateCamera)      // Menambah CCTV baru
	cameras.Get("/", cameraHandler.GetAllCameras)       // Mendapatkan daftar semua CCTV

	detections := apiProtected.Group("/detections")
	detections.Post("/", detectionHandler.CreateDetection) // Menerima & menyimpan data deteksi baru dari service AI
	detections.Get("/", detectionHandler.GetDetections)     // Mengambil riwayat data deteksi untuk ditampilkan di grafik/tabel

	// --- Rute untuk Real-time WebSocket ---
	// Rute ini juga dilindungi middleware, tapi di-handle secara khusus
	ws := v1.Group("/ws")
	ws.Get("/live", websocket.New(detectionHandler.LiveDetections, websocket.Config{

	}))
}

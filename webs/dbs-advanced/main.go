package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// initialise the database global access variable here.
var DB *gorm.DB

func main() {
	loadEnv() // load credentials

	initDB() // Initialize database, check connections if they are working perfectly.

	// Set Gin mode from env
	if mode := os.Getenv("GIN_MODE"); mode != "" {
		gin.SetMode(mode) // set test, release or debug
	}

	router := gin.Default()

	// Public routes (no authentication needed)
	public := router.Group("/")
	{
		public.GET("/health", healthCheck)
		public.GET("/albums", getAlbumsPublic) // Public endpoint to see all albums
	}

	// Protected routes (require API key)
	protected := router.Group("/")
	protected.Use(AuthMiddleware())
	{
		protected.GET("/my-albums", getMyAlbums)      // User's own albums
		protected.POST("/albums", postAlbums)
		protected.GET("/albums/:id", getAlbumByID)
		protected.PUT("/albums/:id", updateAlbum)
		protected.DELETE("/albums/:id", deleteAlbum)
		protected.GET("/me", getCurrentUser)
	}

	// Get port from env or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	router.Run(":" + port)
}

// Load environment variables
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	} else {
		log.Println("Environment variables loaded from .env file")
	}
}

// Build database connection string from env variables
func getDatabaseDSN() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if sslmode == "" {
		sslmode = "disable"
	}

	if user == "" || password == "" || dbname == "" {
		log.Fatal("Missing required database credentials: DB_USER, DB_PASSWORD, or DB_NAME")
	}

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)
}

// Initialize PostgreSQL connection
func initDB() {
	dsn := getDatabaseDSN()

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate both tables
	if err := DB.AutoMigrate(&User{}, &Album{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected and migrated successfully")
}

// Health check endpoint
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Server is running",
	})
}

// GET /me - Get current authenticated user
func getCurrentUser(c *gin.Context) {
	user, exists := GetCurrentUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Success: false,
			Error:   "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    user,
	})
}

// GET /albums - Public: Get all albums (including those without users)
func getAlbumsPublic(c *gin.Context) {
	var albums []Album

	// Get all albums and preload user if it exists
	result := DB.Preload("User").Find(&albums)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    albums,
	})
}

// GET /my-albums - Protected: Get only authenticated user's albums
func getMyAlbums(c *gin.Context) {
	userID, _ := GetCurrentUserID(c)
	var albums []Album

	// Get only albums created by this user
	result := DB.Where("user_id = ?", userID).Preload("User").Find(&albums)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    albums,
	})
}

// POST /albums - Create new album (requires authentication)
func postAlbums(c *gin.Context) {
	userID, _ := GetCurrentUserID(c)
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Set the user who created this album
	newAlbum.UserID = &userID

	result := DB.Create(&newAlbum)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   result.Error.Error(),
		})
		return
	}

	// Load user relationship
	DB.Preload("User").First(&newAlbum, newAlbum.ID)

	c.JSON(http.StatusCreated, SuccessResponse{
		Success: true,
		Data:    newAlbum,
	})
}

// GET /albums/:id - Get album by ID (user can only access their own albums)
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	userID, _ := GetCurrentUserID(c)
	var album Album

	// Find album and check ownership
	result := DB.Preload("User").First(&album, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Success: false,
			Error:   "Album not found",
		})
		return
	}

	// Check if user owns this album (or if album has no owner)
	if album.UserID != nil && *album.UserID != userID {
		c.JSON(http.StatusForbidden, ErrorResponse{
			Success: false,
			Error:   "You don't have permission to access this album",
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    album,
	})
}

// PUT /albums/:id - Update album (only if user owns it)
func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	userID, _ := GetCurrentUserID(c)
	var album Album

	// Find the album
	if err := DB.First(&album, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Success: false,
			Error:   "Album not found",
		})
		return
	}

	// Check ownership
	if album.UserID != nil && *album.UserID != userID {
		c.JSON(http.StatusForbidden, ErrorResponse{
			Success: false,
			Error:   "You don't have permission to update this album",
		})
		return
	}

	// If album has no owner, assign to current user
	if album.UserID == nil {
		album.UserID = &userID
	}

	// Bind and update
	if err := c.BindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Ensure user_id doesn't change
	album.UserID = &userID
	DB.Save(&album)

	DB.Preload("User").First(&album, album.ID)

	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    album,
	})
}

// DELETE /albums/:id - Delete album (only if user owns it)
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	userID, _ := GetCurrentUserID(c)
	var album Album

	// Find the album
	if err := DB.First(&album, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Success: false,
			Error:   "Album not found",
		})
		return
	}

	// Check ownership
	if album.UserID != nil && *album.UserID != userID {
		c.JSON(http.StatusForbidden, ErrorResponse{
			Success: false,
			Error:   "You don't have permission to delete this album",
		})
		return
	}

	result := DB.Delete(&album)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    gin.H{"message": "Album deleted successfully"},
	})
}
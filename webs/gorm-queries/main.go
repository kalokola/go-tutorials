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

// initialise the database global access avariable here.
var DB *gorm.DB

func main() {
	
	loadEnv() // load credentials

	
	initDB() // Initialize database, check connections if they are working perfectly.

	// Set Gin mode from env
	if mode := os.Getenv("GIN_MODE"); mode != "" {
		gin.SetMode(mode) // set test, release or debug
	}

	router := gin.Default()
	
	// Routes
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteAlbum)
	
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

	// Set defaults if not provided
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if sslmode == "" {
		sslmode = "disable"
	}

	// Check required fields
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

	// Auto migrate the schema
	if err := DB.AutoMigrate(&Album{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	
	log.Println("Database connected and migrated successfully")
}

// Response structs
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// Album model with GORM tags
type Album struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	Title  string  `gorm:"not null" json:"title"`
	Artist string  `gorm:"not null" json:"artist"`
	Price  float64 `gorm:"not null" json:"price"`
}

// GET /albums - Get all albums
func getAlbums(c *gin.Context) {
	var albums []Album
	
	result := DB.Find(&albums) // retrieves all albums here.
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

// POST /albums - Create new album
func postAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	result := DB.Create(&newAlbum) // create a new record given the payload sent to it
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Success: true,
		Data:    newAlbum,
	})
}

// GET /albums/:id - Get album by ID
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var album Album

	result := DB.First(&album, id) // retriving the first album
	if result.Error != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Success: false,
			Error:   "Album not found",
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    album,
	})
}

// PUT /albums/:id - Update album
func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	var album Album

	if err := DB.First(&album, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Success: false,
			Error:   "Album not found",
		})
		return
	}

	if err := c.BindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	DB.Save(&album) // save the entrire struct this is a good way to update  a fulll record

	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    album,
	})
}

// DELETE /albums/:id - Delete album
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	var album Album

	result := DB.Delete(&album, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Success: false,
			Error:   "Album not found",
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    gin.H{"message": "Album deleted successfully"},
	})
}
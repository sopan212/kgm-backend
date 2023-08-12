package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"kgm-backend/config"
	"kgm-backend/models"
	"net/http"
	"os"
	"path/filepath"
)

func CreateBook(c *gin.Context) {
	var book models.Books

	// Mengambil data dari request
	book.Judul = c.PostForm("judul")
	book.Penulis = c.PostForm("penulis")
	book.Penerbit = c.PostForm("penerbit")
	book.Halaman = c.PostForm("halaman")
	book.Ukuran = c.PostForm("ukuran")
	book.Harga = c.PostForm("harga")
	book.Isbn = c.PostForm("isbn")

	// Mengambil gambar dari request
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// Membaca gambar sebagai byte array
	imageBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	book.Image = imageBytes

	imageFileName := fmt.Sprintf("book_%d.jpg", book.ID)
	imagePath := filepath.Join("image", imageFileName)
	err = saveImageToFile(imagePath, imageBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menyimpan produk ke database
	config.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
func saveImageToFile(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func GetBooks(c *gin.Context) {
	var books []models.Books

	// Mencari produk berdasarkan ID
	result := config.DB.Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func EditBook(c *gin.Context) {
	var book models.Books
	var id = c.Param("id")

	result := config.DB.First(&book, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Buku tidak di temukan",
			"status":  http.StatusNotFound,
		})
		return
		c.Abort()
	}

	//mengambil data dari request form
	book.Judul = c.PostForm("judul")
	book.Penulis = c.PostForm("penulis")
	book.Penerbit = c.PostForm("penerbit")
	book.Halaman = c.PostForm("halaman")
	book.Ukuran = c.PostForm("ukuran")
	book.Harga = c.PostForm("harga")
	book.Isbn = c.PostForm("isbn")

	//mengambil data dari request (jika ada)
	file, _, err := c.Request.FormFile("image")

	if err != nil {
		defer file.Close()

		//membaca gambar dari byte array
		imageBytes, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "error",
				"error":   http.StatusInternalServerError,
			})
			return
			c.Abort()
		}
		book.Image = imageBytes
	}
	config.DB.Save(&book)
	c.JSON(http.StatusOK, gin.H{
		"success": book,
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Books

	result := config.DB.First(&book, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data tidakdi temukan",
			"status":  result.Error.Error(),
		})
		return
	}
	imageFileName := fmt.Sprintf("book_%d.jpg", book.ID)
	imagePath := filepath.Join("image", imageFileName)

	data := config.DB.Delete(&book)
	if data.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": data.Error.Error()})
		return
	}

	if len(book.Image) > 0 {
		err := os.Remove(imagePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted succes",
		"data":    book,
	})
}

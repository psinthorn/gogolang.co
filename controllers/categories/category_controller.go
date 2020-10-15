package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//
// CRUD Categories
//

// Create category and save to datanase
func Create(c *gin.Context) {
	c.JSON(http.StatusOK, "Implement me please")
}

// Read categories from database
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, "Implement get category by id please")
}

// Read all categories from database
func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, "Implement GetAll categories please")
}

// Update Category
func Update(c *gin.Context) {
	c.JSON(http.StatusOK, "Implement update category by id please")
}

// Delete category
func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, "Implement update category by id please")
}

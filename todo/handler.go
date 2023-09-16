package todo

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	repo := c.MustGet("repository").(TodoRepository)
	todos, err := repo.FindAll(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrNotFound})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	repo := c.MustGet("repository").(TodoRepository)
	var body struct {
		Title string `json: "title"`
	}

	if err := c.ShouldBindJSON(&body); err != nil || body.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidBody})
		return
	}

	now := time.Now().UTC()
	todo := Todo{
		ID:        rand.Uint64(),
		Title:     body.Title,
		Done:      false,
		CreatedAt: &now,
	}

	if err := repo.Insert(c, todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func GetTodo(c *gin.Context) {
	repo := c.MustGet("repository").(TodoRepository)

	idStr := c.Param("id")
	id, err := toUint(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidId})
		return
	}

	todo, err := repo.Find(c, id)
	if err != nil {
		if err == ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrNotFound})
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	repo := c.MustGet("repository").(TodoRepository)

	idStr := c.Param("id")
	id, err := toUint(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidId})
		return
	}

	var body struct {
		ID    uint64 `json: "id" gorm:"primary_gey"`
		Title string `json: "title"`
		Done  bool   `json: "done"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	todo, err := repo.Find(c, id)
	if err != nil {
		if err == ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": ErrNotFound})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	todo.Title = body.Title
	todo.Done = body.Done

	if err := repo.Update(c, todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	repo := c.MustGet("repository").(TodoRepository)
	idStr := c.Param("id")
	id, err := toUint(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	_, err = repo.Find(c, id)
	if err != nil {
		if err == ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		return
	}

	if err := repo.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func DeleteAllTodos(c *gin.Context) {
	repo := c.MustGet("repository").(TodoRepository)

	if err := repo.DeleteAll(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func toUint(number string) (uint64, error) {
	const base, bitSize = 10, 64
	return strconv.ParseUint(number, base, bitSize)
}

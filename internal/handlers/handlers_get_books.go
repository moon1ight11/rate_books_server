package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rate_books/internal/database"
	"strconv"
	"strings"
)

// список всех книг
func GetAllBooks(c *gin.Context) {
	// пагинация
	page_number := c.DefaultQuery("page_number", "0")
	page_size := c.DefaultQuery("page_size", "10")
	pageNumberInt, err := strconv.Atoi(page_number)
	if err != nil || pageNumberInt < 0 {
		log.Println("error page", "err:", err, "page number:", pageNumberInt)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	pageSizeInt, err := strconv.Atoi(page_size)
	if err != nil || pageSizeInt <= 0 {
		log.Println("error page", "err:", err, "page number:", pageSizeInt)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size number"})
		return
	}

	// сортировка
	sort_field := c.DefaultQuery("sort_field", "rate")
	sort_order := c.DefaultQuery("sort_order", "DESC")

	allowedSortFields := map[string]bool{
		"title": true, "author_name": true, "year_public": true, "year_read": true, "rate": true,
	}
	if !allowedSortFields[sort_field] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort field"})
		return
	}

	// фильтрация
	var whereClauses []string
	var args []interface{}
	argPos := 1

	addFilter := func(field, operator, value string) {
		if value != "" {
			whereClauses = append(whereClauses, fmt.Sprintf("%s %s $%d", field, operator, argPos))
			args = append(args, value)
			argPos++
		}
	}

	if title := c.Query("title"); title != "" {
		addFilter("rb.title", "ILIKE", "%"+title+"%")
	}

	if author := c.Query("author_name"); author != "" {
		addFilter("a.author_name", "ILIKE", "%"+author+"%")
	}

	if year_public_from := c.Query("year_public_from"); year_public_from != "" {
		if _, err := strconv.Atoi(year_public_from); err == nil {
			addFilter("rb.year_public", ">=", year_public_from)
		}
	}

	if year_public_to := c.Query("year_public_to"); year_public_to != "" {
		if _, err := strconv.Atoi(year_public_to); err == nil {
			addFilter("rb.year_public", "<=", year_public_to)
		}
	}

	if year_read_from := c.Query("year_read_from"); year_read_from != "" {
		if _, err := strconv.Atoi(year_read_from); err == nil {
			addFilter("rb.year_read", ">=", year_read_from)
		}
	}

	if year_read_to := c.Query("year_read_to"); year_read_to != "" {
		if _, err := strconv.Atoi(year_read_to); err == nil {
			addFilter("rb.year_read", "<=", year_read_to)
		}
	}

	if rate_from := c.Query("rate_from"); rate_from != "" {
		if _, err := strconv.Atoi(rate_from); err == nil {
			addFilter("rb.rate", ">=", rate_from)
		}
	}

	if rate_to := c.Query("rate_to"); rate_to != "" {
		if _, err := strconv.Atoi(rate_to); err == nil {
			addFilter("rb.rate", "<=", rate_to)
		}
	}

	where := ""
	if len(whereClauses) > 0 {
		where = "WHERE " + strings.Join(whereClauses, " AND ")
	}

	// запрос
	all_books, err := database.SelectBooks(pageNumberInt, pageSizeInt, where, sort_field, sort_order, args)
	if err != nil {
		log.Println("ошибка", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// количество книг
	AmountofBooks, err := database.SelectAmountOfBooks(where, args)
	if err != nil {
		log.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid amount of books"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"all_books": all_books, "amountOfItems": AmountofBooks})
}

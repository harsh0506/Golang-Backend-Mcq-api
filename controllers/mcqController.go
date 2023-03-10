package controllers

import (
	"fmt"
	"net/http"
	"simple-crud-api/initializers"
	"simple-crud-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func McqsCreate(c *gin.Context) {
	//get data from request body
	var Body struct {
		Question string
		O1       string
		O2       string
		O3       string
		O4       string
		Ans      string
		Sub      string
	}

	c.Bind(&Body)

	//create a mcq
	mcq := models.MCQ{Question: Body.Question, O1: Body.O1, O2: Body.O2, O3: Body.O3, O4: Body.O4, Ans: Body.Ans, Sub: Body.Sub}

	result := initializers.DB.Create(&mcq)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//reurn it
	c.JSON(200, gin.H{
		"mcq": mcq,
	})
}

func ReturnMcq(c *gin.Context) {

	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "slice bounds out of range"})
		}
	}()

	var mcqs []models.MCQ
	initializers.DB.Find(&mcqs)

	Num, _ := strconv.Atoi(c.Param("num"))
	Num2, _ := strconv.Atoi(c.Param("num2"))

	var pageSize int
	var pageNum int

	pageSize = Num
	pageNum = Num2

	// Calculate the starting index and ending index of the slice
	start := (pageNum - 1) * pageSize

	end := start + pageSize

	if end > len(mcqs) {
		end = len(mcqs)
		print(end)
	}

	// Return the slice corresponding to the current page
	currentBodies := mcqs[start:end]
	print(currentBodies)

	//reurn it
	c.JSON(200, gin.H{
		"bodies": currentBodies,
		"page":   pageNum,
		"size":   pageSize,
		"total":  len(mcqs)%10,
	})
}

func ReturnSingleMcq(c *gin.Context) {
	//get id from url
	id := c.Param("id")

	var mcq models.MCQ
	initializers.DB.First(&mcq, id)

	//reurn it
	c.JSON(200, gin.H{
		"mcq": mcq,
	})
}

func ReturnSingleMcqwITHsUB(c *gin.Context) {
	//get id from url
	sub := c.Param("sub")
	fmt.Print(sub)

	var mcq []models.MCQ
	initializers.DB.Raw("SELECT * FROM mcqs WHERE sub = ?", sub).Scan(&mcq)

	//reurn it
	c.JSON(200, gin.H{
		"mcq": mcq,
	})
}

func Update(c *gin.Context) {
	//get id from url
	id := c.Param("id")

	//get data from req.body
	var Body struct {
		Question string
		O1       string
		O2       string
		O3       string
		O4       string
		Ans      string
		Sub      string
	}

	c.Bind(&Body)

	//find post
	var mcq models.MCQ
	initializers.DB.First(&mcq, id)

	//update it
	// Update attributes with `struct`, will only update non-zero fields
	initializers.DB.Model(&mcq).Updates(models.MCQ{
		Question: Body.Question, O1: Body.O1, O2: Body.O2, O3: Body.O3, O4: Body.O4, Ans: Body.Ans, Sub: Body.Sub,
	})

	//save it
	c.JSON(200, gin.H{
		"mcq": mcq,
	})

}

func Delete(c *gin.Context) {
	//get id from url
	id := c.Param("id")

	//delete the mcq
	initializers.DB.Delete(&models.MCQ{}, id)

	//respond
	c.Status(200)
}

func GetAllTypeOFMcqs(c *gin.Context) {
	var mcqs []models.MCQ
	println(mcqs)
	initializers.DB.Find(&mcqs)

	subMap := make(map[string]bool)

	var a string = "jdbjl"
	println(a)

	fmt.Print(a)

	for _, d := range mcqs {
		subMap[d.Sub] = true
	}
	distinctSubs := make([]string, 0, len(subMap))
	for sub := range subMap {
		distinctSubs = append(distinctSubs, sub)
	}

	println(distinctSubs)

	c.JSON(200, gin.H{
		"mcq": distinctSubs,
	})

}

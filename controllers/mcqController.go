package controllers

import (
	"fmt"
	"simple-crud-api/initializers"
	"simple-crud-api/models"

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
	var mcqs []models.MCQ
	initializers.DB.Find(&mcqs)

	//reurn it
	c.JSON(200, gin.H{
		"mcqs": mcqs,
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

package controllers

import (
	"simple-crud-api/initializers"
	"simple-crud-api/models"

	"github.com/gin-gonic/gin"
)

/*

Create Lecs
Get Lecs of class based on college , dept,cource,year,sem,div
Update lecs
Remove Lecs

*/

func CreateLEc(c *gin.Context) {

	var Body struct {
		College    string
		Department string
		Cource     string
		Year       string
		Sem        string
		Div        string
		Lec        string
		Start      string
		End        string
	}

	c.Bind(&Body)

	lec := models.LEC{College: Body.College, Department: Body.Department, Cource: Body.Cource, Year: Body.Year, Sem: Body.Sem, Div: Body.Div}

	result := initializers.DB.Create(&lec)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//reurn it
	c.JSON(200, gin.H{
		"lec": lec,
	})

}

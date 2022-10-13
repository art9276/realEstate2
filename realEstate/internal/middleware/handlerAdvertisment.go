package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"realEstate/internal/db"
	"realEstate/internal/models"
	"strconv"
)

func GetAdvertisment(c *gin.Context) {
	Id := c.Query("IdAdvertisment")
	Id2, err := strconv.Atoi(Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not convert string to int",
		})
	}
	var advertisment models.Advertisment
	row := db.InitDB().QueryRow(`SELECT "IdAdvertisment", 
       "TypeAdvertisment", "Price", "TotalArea",  
       "YearOfContribution", "Address", "Description", 
       "NumberOfRooms","IsCommercial" FROM public."Advertisment" where "IdAdvertisment"=$1`, Id2)
	err2 := row.Scan(&advertisment.IdAdvertisment, &advertisment.TypeAdvertisment,
		&advertisment.Price, &advertisment.TotalArea,
		&advertisment.YearOfContribution,
		&advertisment.Address, &advertisment.Description,
		&advertisment.NumberOfRooms, &advertisment.IsCommercial)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Scan not complited",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Advertisment for ID": advertisment,
	})
}

func GetAllAdvertisment(c *gin.Context) {

	rows, err := db.InitDB().Query(`SELECT "IdAdvertisment", 
       "TypeAdvertisment", "Price", "TotalArea",  
       "YearOfContribution", "Address", "Description", 
       "NumberOfRooms","IsCommercial"
	FROM public."Advertisment"`)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Query not Scanned",
		})
	}
	defer rows.Close()
	var advertisments []models.Advertisment
	for rows.Next() {
		var advertisment models.Advertisment
		err2 := rows.Scan(&advertisment.IdAdvertisment, &advertisment.TypeAdvertisment,
			&advertisment.Price, &advertisment.TotalArea,
			&advertisment.YearOfContribution,
			&advertisment.Address, &advertisment.Description,
			&advertisment.NumberOfRooms, &advertisment.IsCommercial)
		// Exit if we get an error
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Users not scanned",
			})
		}
		advertisments = append(advertisments, advertisment)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": advertisments,
	})
}

func CreateAdvertisment(c *gin.Context) {
	a := new(models.Advertisment)
	if err := c.BindJSON(a); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "JSON non scanned",
		})
	}
	sqlStatement := `INSERT INTO public."Advertisment"(
	"TypeAdvertisment", "Price", "TotalArea", 
     "YearOfContribution", "Address", 
	"Description", "NumberOfRooms","IsCommercial")
	VALUES ($1, $2, $3, $4, $5, $6,$7,$8) `
	res, err := db.InitDB().Query(sqlStatement, a.TypeAdvertisment,
		a.Price, a.TotalArea, a.YearOfContribution,
		a.Address, a.Description, a.NumberOfRooms, a.IsCommercial)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to saved advertisment",
		})
	} else {
		fmt.Println(res)
		c.JSON(http.StatusCreated, gin.H{
			"Advertisment created": a,
		})
	}
}

func UpdateAdvertisment(c *gin.Context) {
	a := new(models.Advertisment)
	if err := c.BindJSON(a); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "JSON non scanned",
		})
	}
	sqlStatement := `UPDATE public."Advertisment"
	SET  "TypeAdvertisment"=$1, "Price"=$2, 
	"TotalArea"=$3, "YearOfContribution"=$4, "Address"=$5, 
	"Description"=$6, "NumberOfRooms"=$7, "IsCommercial"=$8
	WHERE  "IdAdvertisment"=$9`
	res, err := db.InitDB().Query(sqlStatement, a.TypeAdvertisment,
		a.Price, a.TotalArea, a.YearOfContribution, a.Address,
		a.Description, a.NumberOfRooms, a.IsCommercial, a.IdAdvertisment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to updated advertisment",
		})
	} else {
		fmt.Println(res)

		c.JSON(http.StatusOK, gin.H{
			"Advertisment updated": a,
		})
	}
}

func DeleteAdvertisment(c *gin.Context) {
	Id := c.Query("IdAdvertisment")
	Id2, err := strconv.Atoi(Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not convert string to int",
		})
	}
	sqlStatement := `DELETE FROM public."Advertisment"
	WHERE "IdAdvertisment"=$1`
	res, err2 := db.InitDB().Query(sqlStatement, Id2)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Advertisment not found",
		})
	} else {
		fmt.Println(res)

		c.JSON(http.StatusOK, "Advertisment deleted")
	}
}

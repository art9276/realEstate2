package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"realEstate/internal/db"
	"realEstate/internal/models"
	"strconv"
	"strings"
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
	query:=`SELECT "IdAdvertisment", 
       "TypeAdvertisment", "Price", "TotalArea",  
       "YearOfContribution", "Address", "Description", 
       "NumberOfRooms","IsCommercial"
	FROM public."Advertisment" where 1=1`
	//order by $1 asc Limit $2 OFFSET $3
	TypeAdvertisment:= c.Query("TypeAdvertisment")
	if TypeAdvertisment!="" {
		query =query + ` AND "TypeAdvertisment"='`+ TypeAdvertisment + `'`
	}
	PriceL:=c.Query("PriceL")
	if PriceL!="" {
		strconv.Atoi(PriceL)
		query =query + ` AND "Price" >='`+ PriceL + `'`
	}
	PriceU:=c.Query("PriceU")
	if PriceU!="" {
		strconv.Atoi(PriceU)
		query =query + ` AND "Price" <='`+ PriceU + `'`
	}
	Price:=c.Query("Price")
	if Price!="" {
		strconv.Atoi(Price)
		query =query + ` AND "Price" ='`+ Price + `'`
	}
	TotalAreaL:= c.Query("TotalAreaL")
	if TotalAreaL!="" {
		strconv.Atoi(TotalAreaL)
		query =query + ` AND "TotalArea" >='`+ TotalAreaL + `'`
	}
	TotalAreaU:= c.Query("TotalAreaU")
	if TotalAreaU !="" {
		strconv.Atoi(TotalAreaU)
		query =query + ` AND "TotalArea" <='`+ TotalAreaU + `'`
	}
	TotalArea:= c.Query("TotalArea")
	if TotalArea !="" {
		strconv.Atoi(TotalArea)
		query =query + ` AND "TotalArea" ='`+ TotalArea + `'`
	}
	YearOfContribution:=c.Query("YearOfContribution")
	if YearOfContribution!="" {
		strconv.Atoi(YearOfContribution)
		query =query + ` AND "YearOfContribution" ='`+ YearOfContribution + `'`
	}
	YearOfContributionL:=c.Query("YearOfContributionL")
	if YearOfContributionL!="" {
		strconv.Atoi(YearOfContributionL)
		query =query + ` AND "YearOfContribution" >='`+ YearOfContributionL + `'`
	}
	YearOfContributionU:=c.Query("YearOfContributionU")
	if YearOfContributionU!="" {
		strconv.Atoi(YearOfContributionU)
		query =query + ` AND "YearOfContribution" <='`+ YearOfContributionU + `'`
	}
	Address :=c.Query("Address")
	if Address !="" {
		query =query + ` AND "Address" ='`+ Address + `'`
	}
	NumberOfRooms:=c.Query("NumberOfRooms")
	if NumberOfRooms !="" {
		strconv.Atoi(NumberOfRooms)
		query =query + ` AND "NumberOfRooms" ='`+ NumberOfRooms + `'`
	}
	NumberOfRoomsL:=c.Query("NumberOfRoomsL")
	if NumberOfRooms !="" {
		strconv.Atoi(NumberOfRoomsL)
		query =query + ` AND "NumberOfRooms" >='`+ NumberOfRoomsL + `'`
	}
	NumberOfRoomsU:=c.Query("NumberOfRoomsU")
	if NumberOfRoomsU !="" {
		strconv.Atoi(NumberOfRoomsU)
		query =query + ` AND "NumberOfRooms" <='`+ NumberOfRoomsU + `'`
	}
	IsCommercial:=c.Query("IsCommercial")
	if IsCommercial !="" {
		strconv.Atoi(IsCommercial)
		query =query + ` AND "IsCommercial"='`+ IsCommercial + `'`
	}
	Start:= c.Query("Start")
	if Start==""{Start="0"
	}
	Start2, _ :=strconv.Atoi(Start)
	println(Start2)
	Limit:= c.Query("Limit")
	if Limit==""{Limit="10"
	}
	Limit2, _ :=strconv.Atoi(Limit)
	println(Limit2)
	SortBy:=c.Query("SortBy")
	if SortBy==""{SortBy="YearOfContribution"
	}
	SortOrder:=c.Query("SortOrder")
	strings.ToUpper(SortOrder)
	if SortOrder==""{query+=` order by $1 ASC Limit $2 OFFSET $3`}
	if SortOrder=="asc"{query=query + ` order by $1 ASC Limit $2 OFFSET $3`}
	if SortOrder=="desc" {query=query + ` order by $1 DESC Limit $2 OFFSET $3`}
	rows, err := db.InitDB().Query(query,SortBy,Limit2,Start2) //SortOrder
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Query not Scanned",
		})
		return
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
				"message": "Adwertisment not scanned",
			})
			return
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
		return
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
		return
	} else {
		fmt.Println(res)
		c.JSON(http.StatusCreated, gin.H{
			"Advertisment created": a,
		})
		return
	}
}

func UpdateAdvertisment(c *gin.Context) {
	a := new(models.Advertisment)
	if err := c.BindJSON(a); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "JSON non scanned",
		})
		return
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
		return
	} else {
		fmt.Println(res)
		c.JSON(http.StatusOK, gin.H{
			"Advertisment updated": a,
		})
		return
	}
}

func DeleteAdvertisment(c *gin.Context) {
	Id := c.Query("IdAdvertisment")
	Id2, err := strconv.Atoi(Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not convert string to int",
		})
		return
	}
	sqlStatement := `DELETE FROM public."Advertisment"
	WHERE "IdAdvertisment"=$1`
	res, err2 := db.InitDB().Query(sqlStatement, Id2)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Advertisment not found",
		})
		return
	} else {
		fmt.Println(res)

		c.JSON(http.StatusOK, "Advertisment deleted")
	}
	return
}

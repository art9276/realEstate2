package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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
	errval9 := validation.Validate(advertisment.IdAdvertisment,
		validation.Required,
		validation.Length(1, 6))
	if errval9 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID is not valid",
		})
		return
	}
	errval := validation.Validate(advertisment.TypeAdvertisment,
		validation.Required,
		validation.Length(10, 30))
	if errval != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Type is not valid",
		})
		return
	}
	errval2 := validation.Validate(advertisment.Price,
		validation.Required,
		validation.Length(5, 30))
	if errval2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Price is not valid",
		})
		return
	}
	errval3 := validation.Validate(advertisment.TotalArea,
		validation.Required,
		validation.Length(2, 8))
	if errval3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "TotalArea is not valid",
		})
		return
	}
	errval4 := validation.Validate(advertisment.YearOfContribution,
		validation.Required,
		validation.Length(4, 4))
	if errval4 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "YearOfContribution is not valid",
		})
		return
	}
	errval5 := validation.Validate(advertisment.Address,
		validation.Required,
		validation.Length(10, 25))
	if errval5 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Address is not valid",
		})
		return
	}
	errval6 := validation.Validate(advertisment.Description,
		validation.Required)
	if errval6 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Description is not valid",
		})
		return
	}
	errval7 := validation.Validate(advertisment.NumberOfRooms,
		validation.Required,
		validation.Length(1, 3))
	if errval7 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "NumberOfRooms is not valid",
		})
		return
	}
	errval8 := validation.Validate(advertisment.IsCommercial,
		validation.Required,
		validation.Length(1, 1))
	if errval8 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "IsCommercial is not valid",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Advertisment for ID": advertisment,
	})
	return
}

func GetAllAdvertisment(c *gin.Context) {
	query := `SELECT "IdAdvertisment", 
       "TypeAdvertisment", "Price", "TotalArea",  
       "YearOfContribution", "Address", "Description", 
       "NumberOfRooms","IsCommercial"
	FROM public."Advertisment" where 1=1`
	//start filtering function
	TypeAdvertisment := c.Query("TypeAdvertisment")
	if TypeAdvertisment != "" {
		errval := validation.Validate(TypeAdvertisment,
			validation.Required,
			validation.Length(10, 30))
		if errval != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Type is not valid",
			})
			return
		}
		query = query + ` AND "TypeAdvertisment"='` + TypeAdvertisment + `'`
	}
	PriceL := c.Query("PriceL")
	if PriceL != "" {
		strconv.Atoi(PriceL)
		errval := validation.Validate(PriceL,
			validation.Required,
			validation.Length(5, 30))
		if errval != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "PriceL is not valid",
			})
			return
		}
		query = query + ` AND "Price" >='` + PriceL + `'`
	}
	PriceU := c.Query("PriceU")
	if PriceU != "" {
		strconv.Atoi(PriceU)
		errval := validation.Validate(PriceU,
			validation.Required,
			validation.Length(5, 30))
		if errval != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "PriceU is not valid",
			})
			return
		}
		query = query + ` AND "Price" <='` + PriceU + `'`
	}
	Price := c.Query("Price")
	if Price != "" {
		strconv.Atoi(Price)
		errval := validation.Validate(Price,
			validation.Required,
			validation.Length(5, 30))
		if errval != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Price is not valid",
			})
			return
		}
		query = query + ` AND "Price" ='` + Price + `'`
	}
	TotalAreaL := c.Query("TotalAreaL")
	if TotalAreaL != "" {
		strconv.Atoi(TotalAreaL)
		errval := validation.Validate(TotalAreaL,
			validation.Required,
			validation.Length(2, 8))
		if errval != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "TotalAreaL is not valid",
			})
			return
		}
		query = query + ` AND "TotalArea" >='` + TotalAreaL + `'`
	}
	TotalAreaU := c.Query("TotalAreaU")
	if TotalAreaU != "" {
		strconv.Atoi(TotalAreaU)
		errval := validation.Validate(TotalAreaU,
			validation.Required,
			validation.Length(2, 8))
		if errval != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "TotalAreaU is not valid",
			})
			return
		}
		query = query + ` AND "TotalArea" <='` + TotalAreaU + `'`
	}
	TotalArea := c.Query("TotalArea")
	if TotalArea != "" {
		strconv.Atoi(TotalArea)
		errval := validation.Validate(TotalArea,
			validation.Required,
			validation.Length(2, 8))
		if errval != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "TotalArea is not valid",
			})
			return
		}
		query = query + ` AND "TotalArea" ='` + TotalArea + `'`
	}
	YearOfContribution := c.Query("YearOfContribution")
	if YearOfContribution != "" {
		strconv.Atoi(YearOfContribution)
		errval := validation.Validate(YearOfContribution,
			validation.Required,
			validation.Length(4, 4))
		if errval != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "YearOfContribution is not valid",
			})
			return
		}
		query = query + ` AND "YearOfContribution" ='` + YearOfContribution + `'`
	}
	YearOfContributionL := c.Query("YearOfContributionL")
	if YearOfContributionL != "" {
		strconv.Atoi(YearOfContributionL)
		errval := validation.Validate(YearOfContributionL,
			validation.Required,
			validation.Length(4, 4))
		if errval != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "YearOfContributionL is not valid",
			})
			return
		}
		query = query + ` AND "YearOfContribution" >='` + YearOfContributionL + `'`
	}
	YearOfContributionU := c.Query("YearOfContributionU")
	if YearOfContributionU != "" {
		strconv.Atoi(YearOfContributionU)
		errval := validation.Validate(YearOfContributionU,
			validation.Required,
			validation.Length(4, 4))
		if errval != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "YearOfContributionU is not valid",
			})
			return
		}
		query = query + ` AND "YearOfContribution" <='` + YearOfContributionU + `'`
	}
	Address := c.Query("Address")
	if Address != "" {
		errval := validation.Validate(Address,
			validation.Required,
			validation.Length(10, 25))
		if errval != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Address is not valid",
			})
			return
		}
		query = query + ` AND "Address" ='` + Address + `'`
	}
	NumberOfRooms := c.Query("NumberOfRooms")
	if NumberOfRooms != "" {
		strconv.Atoi(NumberOfRooms)
		errval := validation.Validate(NumberOfRooms,
			validation.Required,
			validation.Length(1, 3))
		if errval != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "NumberOfRooms is not valid",
			})
			return
		}
		query = query + ` AND "NumberOfRooms" ='` + NumberOfRooms + `'`
	}
	NumberOfRoomsL := c.Query("NumberOfRoomsL")
	if NumberOfRooms != "" {
		strconv.Atoi(NumberOfRoomsL)
		errval := validation.Validate(NumberOfRoomsL,
			validation.Required,
			validation.Length(1, 3))
		if errval != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "NumberOfRoomsL is not valid",
			})
			return
		}
		query = query + ` AND "NumberOfRooms" >='` + NumberOfRoomsL + `'`
	}
	NumberOfRoomsU := c.Query("NumberOfRoomsU")
	if NumberOfRoomsU != "" {
		strconv.Atoi(NumberOfRoomsU)
		errval := validation.Validate(NumberOfRoomsU,
			validation.Required,
			validation.Length(1, 3))
		if errval != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "NumberOfRoomsU is not valid",
			})
			return
		}
		query = query + ` AND "NumberOfRooms" <='` + NumberOfRoomsU + `'`
	}
	IsCommercial := c.Query("IsCommercial")
	if IsCommercial != "" {
		strconv.Atoi(IsCommercial)
		errval := validation.Validate(IsCommercial,
			validation.Required,
			validation.Length(1, 1))
		if errval != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "IsCommercial is not valid",
			})
			return
		}
		query = query + ` AND "IsCommercial"='` + IsCommercial + `'`
	}
	Start := c.Query("Start")
	if Start == "" {
		Start = "0"
	}
	Start2, _ := strconv.Atoi(Start)
	errval10 := validation.Validate(Start2,
		validation.Required,
		validation.Length(1, 5))
	if errval10 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Start2 is not valid",
		})
		return
	}
	Limit := c.Query("Limit")
	if Limit == "" {
		Limit = "10"
	}
	Limit2, _ := strconv.Atoi(Limit)
	errval11 := validation.Validate(Limit2,
		validation.Required,
		validation.Length(1, 5))
	if errval11 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Limit2 is not valid",
		})
		return
	}
	SortBy := c.Query("SortBy")
	errval12 := validation.Validate(SortBy,
		validation.Required,
		validation.Length(1, 5))
	if errval12 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "SortBy is not valid",
		})
		return
	}
	if SortBy == "" {
		SortBy = "YearOfContribution"
	}
	SortOrder := c.Query("SortOrder")
	errval13 := validation.Validate(SortOrder,
		validation.Required,
		validation.Length(1, 5))
	if errval13 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "SortOrder is not valid",
		})
		return
	}
	//TODO почему то даже при выполнении запроса на asc и desc сортирует по одинаковому
	if SortOrder == "" {
		query += ` order by $1 ASC Limit $2 OFFSET $3`
	}
	if SortOrder == "asc" {
		query = query + ` order by $1 ASC Limit $2 OFFSET $3`
	}
	if SortOrder == "desc" {
		query = query + ` order by $1 DESC Limit $2 OFFSET $3`
	}
	rows, err := db.InitDB().Query(query, SortBy, Limit2, Start2) //SortOrder
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
	err2 := models.ValidateAdvertismentInsert(a)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Validate error",
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
	err2 := models.ValidateAdvertismentUpdate(a)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Validate error",
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
	errval := validation.Validate(Id,
		validation.Required,
		validation.Length(1, 5),
		is.Email)
	if errval != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID not valid",
		})
		return
	}
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

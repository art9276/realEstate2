package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path/filepath"
	"realEstate/internal/db"
	"realEstate/internal/db/redis"
	"realEstate/internal/models"
	pass "realEstate/pkg/password"
	"strconv"
	"time"
)

var ctx = context.Background()

func Logout(c *gin.Context) {
	//TODO add cookie and clear whem
	c.Redirect(http.StatusSeeOther, "/")
}

// Login godoc
// @Summary Logging user
// @Description Logging user
// @ID 3
// @Accept *gin.Context
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Login and Password correct"
// @Success 303 {object} http.Redirect("/")
// @Failure 400 {object} web.APIError "User not found"
// @Router /auth/login [get]
const SecretKey = "adsad3423sdf099bcv_@sfds&8"

func Login(c *gin.Context) {
	//TODO add cookie
	Login := c.Query("Login")
	Enc_password := c.Query("Enc_password")
	var user models.User
	row := db.InitDB().QueryRow(`SELECT "Id_user","Login", "Enc_password"
 	 FROM public."Users" where "Login"=$1`, Login)
	err2 := row.Scan(&user.Id_user, &user.Login, &user.Enc_password)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
	}
	if pass.CheckPasswordHash(Enc_password, user.Enc_password) != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id_user)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Token not generated",
		})
	}
	//	cookie, _ := c.Cookie("jwt")
	//		c.SetCookie("jwt",token,3600,"/","localhost",false,true)
	//	c.Cookie(cookie)
	d := redis.InitRedis().Set(ctx, token, user.Id_user, 0)
	println(d)
	c.JSON(http.StatusOK, gin.H{
		"Token": token,
	})
}

// GetAllUser godoc
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID 4
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /users [get]
func GetAllUser(c *gin.Context) {
	// если добавлять поле date_creation,date_update
	rows, err := db.InitDB().Query(`SELECT "Id_user","Name", "Surename", 
       "Login", "Enc_password", "Telephone", "Email" ,"Date_creation","Role" FROM public."Users"`)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Query not Scanned",
		})
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err2 := rows.Scan(&user.Id_user, &user.Name, &user.Surename,
			&user.Login, &user.Enc_password, &user.Telephone, &user.Email, &user.Date_creation, &user.Role)
		// Exit if we get an error
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Users not scanned",
			})
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": users,
	})

}

// CreateUser godoc
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID 5
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /auth/register [post]
func CreateUser(c *gin.Context) {
	u := new(models.User)
	if err := c.BindJSON(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "JSON non scanned",
		})
	}
	sqlStatement := `INSERT INTO public."Users"(
			"Name", "Surename", "Login", 
			"Enc_password", "Telephone", "Email","Date_creation","Role")
			VALUES ($1, $2, $3, $4, $5, $6,$7,$8) `
	res, err := db.InitDB().Query(sqlStatement, u.Name,
		u.Surename, u.Login,
		pass.HashPassword(u.Enc_password), u.Telephone, u.Email, u.Date_creation, u.Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to saved user",
		})
	} else {
		fmt.Println(res)
		c.JSON(http.StatusCreated, gin.H{
			"User created": u,
		})
	}
}

// GetUsers godoc
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID 6
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /users/:id [get]
func GetUser(c *gin.Context) {
	Email := c.Query("Email")
	var user models.User
	row := db.InitDB().QueryRow(`SELECT "Id_user","Name", "Surename","Login", "Enc_password",
 	"Telephone", "Email", "Date_creation", "Role" FROM public."Users" where "Email"=$1`, Email)
	err := row.Scan(&user.Id_user, &user.Name, &user.Surename, &user.Login, &user.Enc_password, &user.Telephone,
		&user.Email, &user.Date_creation, &user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Scan not complited",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"User for email": user,
	})

}

// UpdateUsers godoc
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID 7
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /users/:id [put]
func UpdateUser(c *gin.Context) {
	//id := c.Param("id")
	u := new(models.User)
	if err := c.BindJSON(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "JSON non scanned",
		})
	}
	sqlStatement := `UPDATE public."Users" SET
			"Name"=$1, "Surename"=$2, "Login"=$3, 
			"Enc_password"=$4, "Telephone"=$5, "Email"=$6,"Date_creation"=$7,"Role"=$8
			Where "Email"=$9`
	res, err := db.InitDB().Query(sqlStatement, u.Name,
		u.Surename, u.Login,
		u.Enc_password, u.Telephone, u.Email, u.Date_creation, u.Role, u.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to updated user",
		})
	} else {
		fmt.Println(res)

		c.JSON(http.StatusOK, gin.H{
			"User updated": u,
		})
	}
}

// // DeleteUsers godoc
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID 8
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /users [delete]
func DeleteUser(c *gin.Context) {
	Email := c.Query("Email")
	sqlStatement := `Delete from public."Users" where "Email"=$1`
	res, err := db.InitDB().Query(sqlStatement, Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
	} else {
		fmt.Println(res)

		c.JSON(http.StatusOK, "User deleted")
	}
}

// ShowIndex godoc
// @Schemes
// @Description Welcome page
// @Tags MainPage
// @Accept json
// @Produce json
// @Success 200 {string} Welcome to site
// @Router / [get]
func ShowIndexPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to site",
	})
	return
}

// NotFound godoc
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID 9
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router  [get]
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "request not found",
	})
	return
}

// UploadFiles godoc
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID 10
// @Accept json
// @Produce json
// @Param some_id path int true "Some ID"
// @Param some_id body web.Pet true "Some ID"
// @Success 200 {string} string "Welcome to site"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /upload [post]

func UploadFiles(c *gin.Context) {
	// Parse request body as multipart form data with 32MB max memory
	file, err := c.FormFile("File")

	// The file cannot be received.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	y := fmt.Sprintf("%v", year)
	m := fmt.Sprintf("%v", month)
	d := fmt.Sprintf("%v", day)
	newFilePath := "./download/" + y + "/" + m + "/" + d + "/"
	os.MkdirAll(newFilePath, 0777)

	// The file is received, so let's save it

	if err := c.SaveUploadedFile(file, newFilePath+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
		"message": "Upload sucesfully"})
}

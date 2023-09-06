package ginuser

import (
	"app/modules/user/business"
	"app/modules/user/entity"
	"app/modules/user/repository/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	guuid "github.com/google/uuid"
)

func HandleImportUserCsv(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, header, err := c.Request.FormFile("upload")

		filename := header.Filename
		fmt.Println(header.Filename)
		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Open(filename)
		csvReader := csv.NewReader(f)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		users, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		data := createUserList(users)

		store := sql.NewSQLRepo(db)
		biz := business.NewBusiness(store)

		insertError := biz.ImportUserCSV(c.Request.Context(), data)
		if insertError != nil {
			log.Printf("Error while inserting new user into db, Reason: %v\n", insertError)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Something went wrong",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "User created Successfully",
		})
	}
}

func createUserList(data [][]string) []*entity.User {
	var userList []*entity.User

	for i, line := range data {
		if i > 0 {
			user := &entity.User{}

			for j, field := range line {
				switch j {
				case 0:
					user.FullName = field
				case 1:
					emp_code, _ := strconv.ParseInt(field, 0, 64)
					user.EmployeeCode = int(emp_code)
				case 2:
					user.Email = field
				case 3:
					user.Role = field
				}
				user.Id = guuid.New().String()
			}

			userList = append(userList, user)
		}
	}
	return userList
}

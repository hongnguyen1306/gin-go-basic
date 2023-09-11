package ginuser

import (
	"app/common"
	"app/component/app_context"
	"app/modules/user/business"
	"app/modules/user/entity"
	"app/modules/user/repository/sql"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
)

func HandleImportUserCsv(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		if requester.GetRole() != entity.SuperAdmin {
			c.JSON(http.StatusForbidden, common.NewFailResponse("Bạn không có quyền thực hiện hành động này"))
			return
		}

		file, header, _ := c.Request.FormFile("upload")

		filename := header.Filename
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

		store := sql.NewSQLRepo(appCtx.GetMainDBConnection())
		biz := business.NewBusiness(store)

		insertError := biz.ImportUserCSV(c.Request.Context(), data)
		if insertError != nil {
			c.JSON(http.StatusBadRequest, common.NewFailResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, nil, nil))
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

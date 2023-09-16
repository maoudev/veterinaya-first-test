package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/maoudev/veterinaya-go/database"
	"github.com/maoudev/veterinaya-go/models"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c *gin.Context) {
	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Error registering user: invalid data",
		})
		return
	}

	user.Id = uuid.NewString()

	encryptPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Password = string(encryptPass)

	_, err = database.DB.Exec("INSERT INTO USER(id_user, rut_user, name_user, email_user, pass_user) values(?,?,?,?,?)", user.Id, user.Rut, user.UserName, user.Email, user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user.Roles == nil {
		user.Roles = []int{1}
	}

	for _, r := range user.Roles {
		_, err = database.DB.Exec("INSERT INTO USER_ROLES(id_userrole, id_user, id_role) values(?,?,?)", uuid.NewString(), user.Id, r)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, user)
}

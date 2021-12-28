package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"symptoms_tracker/models"
	"symptoms_tracker/utils"
)

func LoginHandler(c buffalo.Context) error {
	user := models.Person{}

	email := c.Params().Get("email")
	valid := utils.CheckValidEmail(email)
	if valid != true {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"error": "Email not valid"}))
	}
	err := models.DB.Where("Email = ?", email).First(&user)
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(map[string]string{"query_error": "Error querying email " + email}))
	}

	password := c.Params().Get("password")
	valid_password := utils.CheckPasswordHash(password, user.Password)
	if valid_password != true {
		return c.Render(http.StatusOK, r.JSON(map[string]string{"auth_error": "Wrong password"}))
	}

	token, err := utils.CreateToken()
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(map[string]string{"token_error": "could not create token"}))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]string{"token": token}))
}

package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"symptoms_tracker/models"
	"symptoms_tracker/utils"
)

func RegisterHandler(c buffalo.Context) error {

	users := []models.Person{}
	email := c.Params().Get("email")
	valid := utils.CheckValidEmail(email)
	if valid != true {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"error": "Email not valid"}))
	}

	password := c.Params().Get("password")
	if len(password) == 0 {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"error": "No password given"}))
	}

	err := models.DB.Where("Email = ?", email).All(&users)

	if err != nil {
		return c.Render(http.StatusOK, r.JSON(map[string]string{"error": "Error querying email " + email}))
	}

	if len(users) != 0 {
		return c.Render(http.StatusOK, r.JSON(map[string]string{"error": "User already exists"}))
	}

	password, _ = utils.HashPassword(password)
	new_user := models.Person{
		Email:    email,
		Password: password,
		Name:     c.Param("name"),
		Age:      c.Param("age"),
		Gender:   c.Param("gender"),
	}
	err = models.DB.Create(&new_user)
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(map[string]error{"error": err}))
	}

	token, err := utils.CreateToken()
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(map[string]string{"error": "could not create token"}))
	}

	return c.Render(http.StatusOK, r.JSON(map[string]string{"token": token}))
}

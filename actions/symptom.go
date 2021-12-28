package actions

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/nulls"
	"symptoms_tracker/models"
)

func ListSymptoms(c buffalo.Context) error {
	from_str := c.Param("from")
	to_str := c.Param("to")

	person := models.Person{}
	person_id := c.Params().Get("person_id")

	err := models.DB.Eager().Where("id = ?", person_id).First(&person)
	if err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if from_str == "" && to_str == "" {
		return c.Render(200, r.JSON(map[string][]models.Symptom{"symptoms": person.Symptoms}))
	}

	symptoms := []models.Symptom{}
	from, err := strconv.ParseFloat(from_str, 64)
	if err != nil {
		return c.Error(http.StatusNotFound, err)
	}
	to, err := strconv.ParseFloat(to_str, 64)
	if err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	curr_date := time.Now()
	for _, symptom := range person.Symptoms {
		symptom_date := time.Unix(symptom.CreatedAt.Unix(), 0)
		days := math.Ceil(curr_date.Sub(symptom_date).Hours()/24) - 1
		fmt.Printf("days: %v\n", days)
		if from <= days && to >= days {
			symptoms = append(symptoms, symptom)
		}
	}
	return c.Render(200, r.JSON(map[string][]models.Symptom{"symptoms": symptoms}))

}

func CreateSymptom(c buffalo.Context) error {
	person := models.Person{}
	person_id := c.Params().Get("person_id")
	err := models.DB.Find(&person, person_id)

	if err != nil {
		return c.Render(http.StatusUnprocessableEntity, r.JSON(map[string]string{"error": "person not found"}))
	}

	symptom_content := strings.TrimSpace(strings.ToUpper(c.Params().Get("symptom")))
	symptom := models.Symptom{}
	symptom = models.Symptom{Content: nulls.NewString(symptom_content)}
	err = models.DB.Eager().Create(&symptom)

	if err != nil {
		return c.Render(http.StatusUnprocessableEntity, r.JSON(map[string]error{"error": err}))
	}

	sp := models.SymptomPerson{
		SymptomID: symptom.ID,
		PersonID:  person.ID,
	}
	err = models.DB.Create(&sp)
	if err != nil {
		return c.Render(http.StatusUnprocessableEntity, r.JSON(map[string]error{"error": err}))
	}

	return c.Render(http.StatusUnprocessableEntity, r.JSON(map[string]string{"success": "created symptom for " + person.Email}))

}

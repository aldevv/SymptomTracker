package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Person is used by pop to map your people database table to your go code.
type Person struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Email     string    `json:"username" db:"email"`
	Password  string    `json:"password" db:"password"`
	Name      string    `json:"name" db:"name"`
	Age       string    `json:"age" db:"age"`
	Gender    string    `json:"gender" db:"gender"`
	Symptoms  Symptoms  `many_to_many:"symptom_person"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p Person) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// people is not required by pop and may be deleted
type people []Person

// String is not required by pop and may be deleted
func (p people) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Person) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Name, Name: "Name"},
		&validators.StringIsPresent{Field: p.Age, Name: "Age"},
		&validators.StringIsPresent{Field: p.Gender, Name: "Gender"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Person) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Person) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

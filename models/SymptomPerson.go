package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// SymptomPerson is used by pop to map your symptompeople database table to your go code.
type SymptomPerson struct {
	ID        uuid.UUID `json:"id" db:"id"`
	SymptomID uuid.UUID `json:"symptom_id" db:"symptom_id"`
	PersonID  uuid.UUID `json:"person_id" db:"person_id"`
	Symptom   Symptom   `belongs_to:"symptom"`
	Person    Person    `belongs_to:"people"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (s SymptomPerson) TableName() string {
	return "symptom_person"
}

// String is not required by pop and may be deleted
func (s SymptomPerson) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Symptompeople is not required by pop and may be deleted
type Symptompeople []SymptomPerson

// String is not required by pop and may be deleted
func (s Symptompeople) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *SymptomPerson) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *SymptomPerson) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *SymptomPerson) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

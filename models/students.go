package models

import (
	"regexp"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
	RG   string `json:"rg" validate:"len=9, regexp=^[0-9]*$"`
}

func (s *Student) SetCPF(value string) {
	regex := regexp.MustCompile("[^0-9]")
	s.CPF = regex.ReplaceAllString(value, "")
}

func (s *Student) SetRG(value string) {
	regex := regexp.MustCompile("[^0-9]")
	s.RG = regex.ReplaceAllString(value, "")
}

func ValidateStudent(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}

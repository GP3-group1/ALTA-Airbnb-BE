package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationError(err error) string {
	reports := []string{}

	castedObject, ok := err.(validator.ValidationErrors)
	if ok {
		for _, v := range castedObject {
			switch v.Tag() {
			case "required":
				reports = append(reports, fmt.Sprintf("%s is required", v.Field()))
			case "min":
				reports = append(reports, fmt.Sprintf("%s value must be greater than %s character", v.Field(), v.Param()))
			case "max":
				reports = append(reports, fmt.Sprintf("%s value must be lower than %s character", v.Field(), v.Param()))
			case "email":
				reports = append(reports, fmt.Sprintf("%s is not valid", v.Field()))
			case "lte":
				reports = append(reports, fmt.Sprintf("%s value must be below %s", v.Field(), v.Param()))
			case "gte":
				reports = append(reports, fmt.Sprintf("%s value must be above %s", v.Field(), v.Param()))
			case "numeric":
				reports = append(reports, fmt.Sprintf("%s value must be numeric", v.Field()))
			case "url":
				reports = append(reports, fmt.Sprintf("%s value must be url", v.Field()))
			}
		}
	}
	report := strings.Join(reports, ", ")
	return report
}

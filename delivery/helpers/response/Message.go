package response

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func StatusInvalidRequest(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": message,
		"data":    nil,
	}
}

func StatusBadRequest(err error) map[string]interface{} {
	var field, tag string
	var message []string

	for _, err := range err.(validator.ValidationErrors) {
		field = fmt.Sprint(err.StructField())
		tag = fmt.Sprint(err.Tag())

		message = append(message, "field "+field+" : "+tag)
	}

	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": message,
		"data":    nil,
	}
}

func StatusBadRequestBind(err error) map[string]interface{} {
	var field [] string
	var message string
	
	for i, v := range strings.Fields(string(err.Error())) {
		if i == 1 && v == "message=Syntax" {
			message = "expected=string"
		} else if i == 1 && v == "message=Unmarshal" {
			message = "expected=string"
		} else if i == 6 {
			field = append(field, v)
		}
	}

	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": field[0]+" "+message,
		"data":    nil,
	}
}

func StatusBadRequestDuplicate(err error) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": err.Error(),
		"data":    nil,
	}
}

func StatusBadRequestUpload(err error) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": err.Error(),
		"data":    nil,
	}
}

func StatusUnauthorized(err error) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"message": err.Error(),
		"data":    nil,
	}
}

func StatusForbidden(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusForbidden,
		"message": message,
		"data":    nil,
	}
}

func StatusNotFound(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": message,
		"data":    nil,
	}
}

func StatusOK(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	}
}

func StatusCreated(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": message,
		"data":    data,
	}
}

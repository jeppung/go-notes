package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jeppung/go-notes.git/models"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func JSONBindingHandler[T comparable](ctx *gin.Context, data T) error {
	if err := ctx.BindJSON(data); err != nil {
		var e1 validator.ValidationErrors
		var e2 *json.UnmarshalTypeError

		if errors.As(err, &e1) {
			out := make([]FieldError, len(e1))
			for i, eachErr := range e1 {
				out[i] = FieldError{
					Field:   strings.ToLower(eachErr.Field()),
					Message: eachErr.Tag(),
				}
			}

			ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[[]FieldError]{
				Status: "error",
				Data:   out,
			})
			return err
		}

		if errors.As(err, &e2) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[string]{
				Status: "error",
				Data:   fmt.Sprintf("field %v must be a %v", e2.Field, e2.Type),
			})
			return err
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response[string]{
			Status: "error",
			Data:   err.Error(),
		})
		return err
	}

	return nil
}

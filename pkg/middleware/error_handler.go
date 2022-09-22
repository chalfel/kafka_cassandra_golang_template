package middleware

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/chalfel/kafka_cassandra_golang_template/pkg/exceptions"
	"github.com/chalfel/kafka_cassandra_golang_template/pkg/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			if errors.Is(err, io.EOF) {
				c.JSON(http.StatusBadRequest, gin.H{"message": "request body is empty"})
				c.Abort()
				return
			}

			switch err.Err.(type) {

			case *exceptions.AppException:
				except := err.Err.(*exceptions.AppException)
				c.JSON(except.Code, gin.H{"message": except.Message})
				c.Abort()
				return

			case validator.ValidationErrors:
				var validationErrors []exceptions.ValidationDetails
				errs := err.Err.(validator.ValidationErrors)

				for _, err := range errs {
					var vl exceptions.ValidationDetails
					vl.Field = helpers.LowerFirstChar(err.Field())
					vl.Constraint = err.ActualTag()
					vl.Value = fmt.Sprintf("%v", err.Value())
					validationErrors = append(validationErrors, vl)
				}

				c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": validationErrors})
				c.Abort()
				return
			default:
				logrus.Error("ErrorHandler: internal server error", err.Err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				c.Abort()
			}
		}
	}
}

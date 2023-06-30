package app

import (
	"net/http"

	"github.com/beego/beego/v2/core/validation"
	"github.com/cindyyangcaixia/gin-example/pkg/e"
	"github.com/gin-gonic/gin"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	valid := validation.Validation{}

	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}

	if !check {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}

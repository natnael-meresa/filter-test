package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var ErrInvalidBody = ErrInvalidUserInput.New("invalid body")

type Controller struct {
	repo *Repo
}

func NewController(repo *Repo) *Controller {
	return &Controller{repo: repo}
}

func (cr *Controller) List(c *gin.Context) {

	filterParam := NewFilterParam()

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		HandlerError(c, ErrInvalidBody)
		return
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		HandlerError(c, ErrInvalidBody)
		return
	}

	filterParam.PageNum = page
	filterParam.PageSize = size

	if column1 := c.Query("column1"); column1 != "" {
		op, value := SplitOperator(column1)
		filterParam.Filter = append(filterParam.Filter, Filter{
			Key:      "column1",
			Value:    value,
			Operator: op,
		})
	}

	if column2 := c.Query("column2"); column2 != "" {
		op, value := SplitOperator(column2)
		filterParam.Filter = append(filterParam.Filter, Filter{
			Key:      "column2",
			Value:    value,
			Operator: op,
		})
	}

	// Extract sorting fields and their directions from query parameters
	if sort := c.Query("sort"); sort != "" {
		sortFields := strings.Split(sort, ",")
		for _, field := range sortFields {
			parts := strings.SplitN(field, ":", 2) // Split field and direction
			direction := "ASC"
			if len(parts) > 1 {
				direction = strings.ToUpper(parts[1])
			}
			filterParam.Sort = append(filterParam.Sort, Sort{
				Key:   parts[0],
				Order: direction,
			})
		}
	}

	ctx := c.Request.Context()
	result, err := cr.repo.List(ctx, filterParam)
	if err != nil {
		HandlerError(c, err)
		return
	}
	SendSuccessResponse(c, http.StatusOK, result)
}

func (cr *Controller) Get(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandlerError(c, ErrInvalidBody)
		return
	}

	ctx := c.Request.Context()
	result, err := cr.repo.Get(ctx, id)
	if err != nil {
		HandlerError(c, err)
		return
	}

	SendSuccessResponse(c, http.StatusOK, result)
}

func SplitOperator(value string) (string, string) {

	split := strings.Split(value, ";")

	return split[0], split[1]
}

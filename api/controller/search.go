package controller

import (
	"net/http"

	"github.com/RrNn/detector/models"
	"github.com/labstack/echo"
)

// SearchURL godoc
// @Summary Search URL
// @Description Search the DB for any string provided
// @Produce json
// @Success 200 {object} SearchURL
// @Router /auth/search [get]
func (cont *Controller) SearchURL(c echo.Context) error {

	url := c.QueryParam("url")

	var urls []models.Url

	cont.DB.Where("link like ?", "%"+url+"%").Find(&urls)

	return c.JSON(http.StatusOK, urls)

}

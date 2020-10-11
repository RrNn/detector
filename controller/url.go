package controller

import (
  "net/http"

  "github.com/RrNn/detector/models"
  "github.com/labstack/echo"
)

func (cont *Controller) AddUrl(c echo.Context) (err error) {
  url := new(models.Url)
  if err := c.Bind(url); err != nil {
    return c.JSON(http.StatusInternalServerError, err)
  }
  if err = cont.DB.Create(&url).Error; err != nil {
    return c.JSON(http.StatusInternalServerError, err)
  }

  return err
}

func (cont *Controller) GetUrls(c echo.Context) (err error) {
  return err
}

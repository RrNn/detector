package controller

import (
  "net/http"
  "strconv"
  "time"

  "github.com/RrNn/detector/models"
  "github.com/dgrijalva/jwt-go"
  "gorm.io/gorm"

  "github.com/labstack/echo"
)

// AddURL exported
func (cont *Controller) AddURL(c echo.Context) (err error) {
  url := new(models.Url)
  token := c.Get("token").(*jwt.Token)
  claims := token.Claims.(jwt.MapClaims)
  id := claims["user_id"].(float64)

  if err := c.Bind(url); err != nil {
    return c.JSON(http.StatusInternalServerError, err.Error())
  }
  url.UserID = uint(id)

  resp, err := http.Get(url.Link)
  if err != nil {
    return c.JSON(http.StatusInternalServerError, err.Error())
  }
  check := models.Url{}
  cont.DB.Raw("SELECT * FROM urls WHERE link = ? AND user_id = ?", url.Link, uint(id)).Scan(&check)
  if check.Link != "" {
    return c.JSON(http.StatusNotAcceptable, map[string]string{
      "message": "the link already exists",
    })
  }
  if err = cont.DB.Create(&url).Error; err != nil {
    return c.JSON(http.StatusInternalServerError, err.Error())
  }

  return c.JSON(http.StatusCreated, map[string]interface{}{
    "url":                  url.Link,
    "firstCheckStatusCode": resp.StatusCode,
    "lastChecked":          time.Now(),
  })
}

// GetUrls exported
func (cont *Controller) GetUrls(c echo.Context) (err error) {
  limit := c.QueryParam("limit")
  offset := c.QueryParam("offset")

  urls := []models.Url{}

  query := cont.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
    return db.Select("ID", "Name") // mainly for not returning the password
  })
  if limit != "" {
    l, _ := strconv.Atoi(limit)
    query = query.Limit(l)
  }
  if offset != "" {
    o, _ := strconv.Atoi(offset)
    query = query.Offset(o)
  }
  query.Find(&urls)

  return c.JSON(http.StatusOK, urls)
}

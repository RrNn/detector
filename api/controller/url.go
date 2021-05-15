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

  token := c.Get("token").(*jwt.Token)
  claims := token.Claims.(jwt.MapClaims)

  id := int(claims["user_id"].(float64))

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
  query.Where("user_id = ?", id).Find(&urls)

  return c.JSON(http.StatusOK, urls)
}

// GetURL exported
func (cont *Controller) GetURL(c echo.Context) (err error) {
  id := c.Param("id")
  withpings := c.QueryParam("withpings")
  limit := c.QueryParam("limit")
  offset := c.QueryParam("offset")
  l, err := strconv.Atoi(limit)
  if err != nil {
    l = 10
  }
  o, err := strconv.Atoi(offset)
  if err != nil {
    o = 0
  }

  query := cont.DB
  if withpings != "" {
    query = query.Preload("Pings", func(db *gorm.DB) *gorm.DB {
      // return db.Select("*")
      return db.Select("UrlID", "ID", "Status", "Error", "Time").Limit(l).Offset(o).Order("ID desc")
    })
  }
  url := new(models.Url)
  query.Where("id = ?", id).Find(&url)
  return c.JSON(http.StatusOK, url)
}

// DeleteURL exported
func (cont *Controller) DeleteURL(c echo.Context) (err error) {
  id := c.Param("id")
  query := cont.DB
  var count int64
  query.Model(&models.Url{}).Where("id = ?", id).Count(&count)
  if count != 0 {
    query.Delete(&models.Url{}, id)
    return c.JSON(http.StatusAccepted, map[string]string{"message": "Link has been deleted"})
  }
  return c.JSON(http.StatusExpectationFailed, map[string]string{"message": "Either the link is already deleted or does not exist"})
}

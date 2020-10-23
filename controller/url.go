package controller

import (
  "net/http"
  "time"

  "github.com/RrNn/detector/models"
  "github.com/dgrijalva/jwt-go"

  "github.com/labstack/echo"
)

func (cont *Controller) AddUrl(c echo.Context) (err error) {
  url := new(models.Url)
  token := c.Get("token").(*jwt.Token)
  claims := token.Claims.(jwt.MapClaims)
  id := claims["user_id"].(float64)

  if err := c.Bind(url); err != nil {
    return c.JSON(http.StatusInternalServerError, err)
  }
  url.UserID = uint(id)

  if resp, err := http.Get(url.Link); err != nil {
    return c.JSON(http.StatusInternalServerError, err)
  } else {
    check := models.Url{}
    cont.DB.Raw("SELECT * FROM urls WHERE link = ? AND user_id = ?", url.Link, uint(id)).Scan(&check)
    if check.Link != "" {
      return c.JSON(http.StatusNotAcceptable, map[string]string{
        "message": "the link already exists",
      })
    }
    if err = cont.DB.Create(&url).Error; err != nil {
      return c.JSON(http.StatusInternalServerError, err)
    }

    return c.JSON(http.StatusCreated, map[string]interface{}{
      "url":                  url.Link,
      "firstCheckStatusCode": resp.StatusCode,
      "lastChecked":          time.Now(),
    })
  }

}

func (cont *Controller) GetUrls(c echo.Context) (err error) {

  return err
}

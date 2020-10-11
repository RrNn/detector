package controller

import (
  "net/http"

  "github.com/RrNn/detector/models"
  "github.com/RrNn/detector/utils"
  "github.com/labstack/echo"
)

func (cont *Controller) Register(c echo.Context) (err error) {
  user := new(models.User)

  if err := c.Bind(user); err != nil {
    return c.JSON(http.StatusInternalServerError, err)
  }
  if pass, err := utils.GenerateHashFromPassword(user.Password); err != nil {
    return c.JSON(http.StatusNotAcceptable, map[string]error{"message": err})
  } else {
    user.Password = string(pass)
  }
  if err = cont.DB.Create(&user).Error; err != nil {
    return c.JSON(http.StatusInternalServerError, err)
  }

  if t, err := utils.CreateJwtToken(user); err != nil {
    return c.JSON(http.StatusInternalServerError, err)
  } else {
    return c.JSON(http.StatusCreated, struct {
      Token string            `json:"token"`
      User  models.ReturnUser `json:"user"`
    }{
      Token: t,
      User:  models.GetUser(user),
    })
  }

}

func (cont *Controller) Login(c echo.Context) (err error) {
  user := new(models.User)
  login := new(models.LoginUserStruct)
  if err := c.Bind(login); err != nil {
    return c.JSON(http.StatusInternalServerError, err)
  }
  cont.DB.Where("email = ?", login.Email).First(&user)
  matched := utils.CompareHashAndPassword(user.Password, login.Password)

  if matched {
    if t, err := utils.CreateJwtToken(user); err != nil {
      return c.JSON(http.StatusInternalServerError, err)
    } else {
      return c.JSON(http.StatusOK, map[string]string{
        "token": t,
      })
    }
  }
  return echo.ErrUnauthorized
}

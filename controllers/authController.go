package controllers

import (
	"go-admin/database"
	"go-admin/models"
	"go-admin/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password do not match",
		})
	}

	hash, _ := utils.HashPassword(data["password"])

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		RoleId:    1,
	}

	user.SetPassword(hash)

	//save on database
	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	//get user by email
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if utils.CheckPasswordHash(data["password"], string(user.Password)) {

		//generate jwt
		token, err := utils.GenerateJwt(strconv.Itoa(int(user.Id)))

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		//add jwd in cookie
		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		}
		//send token
		c.Cookie(&cookie)

		return c.JSON(fiber.Map{
			"message": "success",
		})

	} else {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	//parse jwt
	id, _ := utils.ParseJwt(cookie)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	// do you don't have a way to remove a cookie, basically is just put time in the past, and value empty.
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})

}

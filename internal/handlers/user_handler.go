package handlers

import (
	"companypresence-api/internal/models"
	userservice "companypresence-api/internal/services/users"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{
	userService *userservice.UserService
}

func NewUserHandler(userService *userservice.UserService) *UserHandler{
	return &UserHandler{userService: userService}
}

func(h *UserHandler) CreateUser(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message" : "something went wrong while parsing model information: " + err.Error(),
		})
	}	
	
	err = h.userService.CreateUser(c.Context(), &user)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message" : "something went wrong creating the entity: " + err.Error(),
				})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "registered",
		"data" : user,
	})
}

func(h *UserHandler)GetUserById(c *fiber.Ctx) error{
	id := c.Params("id")
	user, err := h.userService.GetUserById(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong fetching the entity: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) GetUserByMail(c *fiber.Ctx) error {
	mail := c.Params("mail")
	user, err := h.userService.GetUserByMail(c.Context(), mail)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong fetching the entity: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) GetActiveUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetActiveUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong fetching the entities: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong fetching the entities: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var creds models.LoginCredentials
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid input: " + err.Error(),
		})
	}

	user, err := h.userService.ValidateCredentials(c.Context(), creds.Email, creds.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid credentials: " + err.Error(),
		})
	}

	/*token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create token: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})*/
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "logged in",
		"data" : user,
	})
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong while parsing model information: " + err.Error(),
		})
	}

	err = h.userService.UpdateUser(c.Context(), &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong updating the entity: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "updated",
		"data":    user,
	})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.userService.DeleteUser(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong deleting the entity: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "deleted",
	})
}
package controller

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"go-fiber-gorm/model/response"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(c *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return c.JSON(users)
}

func UserHandlerGetById(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user entity.User

	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": "failed",
			"data":    err.Error(),
		})
	}

	userResponse := response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    userResponse,
		// atau
		// "data": user,
	})
}

func UserHandlerCreate(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to store data",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}

func UserhandlerUpdate(c *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)

	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var user entity.User

	userId := c.Params("id")

	err := database.DB.First(&user, "id = ?", &userId).Error

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": "failed",
			"data":    err.Error(),
		})
	}

	// Update User Data
	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}

	if userRequest.Address != "" {
		user.Address = userRequest.Address
	}
	user.Phone = userRequest.Phone

	errUpdate := database.DB.Save(&user).Error

	if errUpdate != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    user,
		// "body":    userRequest,
	})
}

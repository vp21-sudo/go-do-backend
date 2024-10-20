package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/vp21-sudo/go-do-backend/db"
	"github.com/vp21-sudo/go-do-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func getValidationErrorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return fmt.Sprintf("Minimum length is %s", fieldError.Param())
	case "max":
		return fmt.Sprintf("Maximum length is %s", fieldError.Param())
	default:
		return "Invalid value"
	}
}

func GetTodos(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := db.TodoCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer cursor.Close(ctx)
	var todos []models.Todo
	if err = cursor.All(ctx, &todos); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request, ID is required."})
	}
	ctx, canel := context.WithTimeout(context.Background(), 10*time.Second)
	defer canel()

	var todo models.Todo
	err = db.TodoCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "To Do not found!"})
	}
	return c.JSON(todo)
}

func CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// add a validator to validate request body
	// Validate the Todo struct
	if err := validate.Struct(todo); err != nil {
		// Return validation errors
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = getValidationErrorMessage(err)
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}
	todo.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.TodoCollection.InsertOne(ctx, todo)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Validate the Todo struct
	if err := validate.Struct(todo); err != nil {
		// Return validation errors
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = getValidationErrorMessage(err)
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}
	update := bson.M{
		"$set": bson.M{
			"title":     todo.Title,
			"completed": todo.Completed,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := db.TodoCollection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if result.MatchedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Godo not found!"})
	}

	return c.JSON(fiber.Map{"message": "Go Do updated."})
}
func DeleteTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := db.TodoCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if result.DeletedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "To Do not found!"})
	}
	return c.JSON(fiber.Map{"message": "GoDo Dleted."})
}

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/eznxxy/go-todo/models"
	"github.com/labstack/echo"
)

type Request struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsFinish    bool   `json:"isFinish"`
}

func FetchAllTodo(ctx echo.Context) error {
	result, err := models.FetchAllTodo()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, result)
}

func CraeteTodo(ctx echo.Context) error {
	var request Request
	json.NewDecoder(ctx.Request().Body).Decode(&request)

	result, err := models.CreateTodo(request.Title, request.Description)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, result)
}

func UpdateTodo(ctx echo.Context) error {
	id := ctx.Param("id")

	var request Request
	json.NewDecoder(ctx.Request().Body).Decode(&request)

	result, err := models.UpdateTodo(request.Title, request.Description, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, result)
}

func DeleteTodo(ctx echo.Context) error {
	id := ctx.Param("id")

	result, err := models.DeleteTodo(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, result)
}

func MarkTodo(ctx echo.Context) error {
	id := ctx.Param("id")

	var request Request
	json.NewDecoder(ctx.Request().Body).Decode(&request)

	result, err := models.MarkTodo(request.IsFinish, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, result)
}

package models

import (
	"net/http"
	"strconv"

	"github.com/eznxxy/go-todo/database"
	"github.com/eznxxy/go-todo/dtos"
)

func FetchAllTodo() (Response, error) {
	var obj dtos.Todo
	var arrobj []dtos.Todo
	var res Response

	conn := database.CreateConn()

	rows, err := conn.Query("SELECT * FROM todo_items")
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err := rows.Scan(&obj.Id, &obj.Title, &obj.Description, &obj.IsFinish)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func CreateTodo(title string, description string) (Response, error) {
	var obj dtos.Todo
	var res Response

	conn := database.CreateConn()

	stmt, err := conn.Prepare("INSERT INTO todo_items (title, description) VALUES (?, ?)")
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(title, description)
	if err != nil {
		return res, err
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	obj.Id = int(insertedId)
	obj.Title = title
	obj.Description = description
	obj.IsFinish = false

	res.Status = http.StatusCreated
	res.Message = "Successfully created todo"
	res.Data = obj

	return res, nil
}

func UpdateTodo(title string, description string, id string) (Response, error) {
	var obj dtos.Todo
	var res Response

	conn := database.CreateConn()

	stmt, err := conn.Prepare("UPDATE todo_items SET title=?, description=? WHERE id=?")
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(title, description, id)
	if err != nil {
		return res, err
	}

	result, err := strconv.Atoi(id)
	obj.Id = result
	obj.Title = title
	obj.Description = description
	obj.IsFinish = false

	res.Status = http.StatusOK
	res.Message = "Successfully updated todo"
	res.Data = obj

	return res, nil
}

func DeleteTodo(id string) (Response, error) {
	var res Response

	conn := database.CreateConn()

	_, err := conn.Exec("DELETE FROM todo_items WHERE id=?", id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully deleted todo"

	return res, nil
}

func MarkTodo(isFinish bool, id string) (Response, error) {
	var res Response

	conn := database.CreateConn()

	_, err := conn.Exec("UPDATE todo_items SET isFinish=? WHERE id=?", isFinish, id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully mark todo"

	return res, nil
}

package handler

import (
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"Go-Lang-Zopsmart/datastore"
	"Go-Lang-Zopsmart/model"
)

type handler struct {
	store datastore.Student
}

func New(s datastore.Student) handler {
	return handler{store: s}
}

func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	// ctx.PathParam() returns the path parameter from HTTP request.
	ID := ctx.PathParam("ID")
	if ID == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	if _, err := validateID(ID); err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.store.GetByID(ctx, ID)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "student",
			ID:     ID,
		}
	}

	return resp, nil
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var student model.Student

	// ctx.Bind() binds the incoming data from the HTTP request to a provided interface (i).
	if err := ctx.Bind(&student); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.store.Create(ctx, &student)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("ID")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := validateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var student model.Student
	if err = ctx.Bind(&student); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	student.ID = id

	resp, err := h.store.Update(ctx, &student)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("ID")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := validateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	if err := h.store.Delete(ctx, id); err != nil {
		return nil, err
	}

	return "Deleted successfully", nil
}

func validateID(id string) (int, error) {
	res, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return res, err
}
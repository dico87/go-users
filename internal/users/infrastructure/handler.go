package infrastructure

import (
	"github.com/dico87/users/internal/common"
	"github.com/dico87/users/internal/users/application"
	"github.com/dico87/users/internal/users/domain"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	service application.Service
}

func New(service application.Service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) Create(context echo.Context) error {
	user := domain.User{}
	if err := context.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid body request")
	}

	createdUser, err := h.service.Create(user)

	if err != nil {
		return err
	}

	err = context.JSON(http.StatusOK, createdUser)

	if err != nil {
		return err
	}

	return nil
}

func (h Handler) Update(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id parameter")
	}

	user := domain.User{}
	if err := context.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid body request")
	}

	updatedUser, err := h.service.Update(uint(id), user)

	if err != nil {
		return err
	}

	err = context.JSON(http.StatusOK, updatedUser)

	if err != nil {
		return err
	}

	return nil
}

func (h Handler) FindById(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id parameter")
	}

	foundUser, err := h.service.FindById(uint(id))

	if err != nil {
		if err == common.ErrNotFoundRecord {
			return echo.NewHTTPError(http.StatusNotFound, "Id not found")
		}
		return err
	}

	err = context.JSON(http.StatusOK, foundUser)

	if err != nil {
		return err
	}

	return nil
}

func (h Handler) FindByDocument(context echo.Context) error {
	documentTypeId, err := strconv.Atoi(context.Param("documentTypeId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid documentTypeId parameter")
	}
	document := context.Param("document")

	foundUser, err := h.service.FindByDocument(uint(documentTypeId), document)

	if err != nil {
		if err == common.ErrNotFoundRecord {
			return echo.NewHTTPError(http.StatusNotFound, "Id not found")
		}
		return err
	}

	err = context.JSON(http.StatusOK, foundUser)

	if err != nil {
		return err
	}

	return nil
}

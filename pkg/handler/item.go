package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/todo-app/pkg/model"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusNotFound, "user not found")
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	var input model.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.TodoItem.Create(userId.(int), listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItem(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	items, err := h.service.TodoItem.GetAll(userId.(int), listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)

}

func (h *Handler) getItemById(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	item, err := h.service.TodoItem.GetById(userId.(int), itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusNotFound, "user not found")
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input model.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.TodoItem.Update(userId.(int), itemId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
func (h *Handler) deleteItem(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusNotFound, "user not found")
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.service.TodoItem.Delete(userId.(int), itemId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		"ok",
	})
}

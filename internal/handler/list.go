package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hostingvk4/badgerList/internal/models"
	"net/http"
	"strconv"
)

// @Summary Create list
// @Security ApiKeyAuth
// @Tags lists
// @Description create list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body models.ListDto true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var listDto models.ListDto
	if err := c.BindJSON(&listDto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	listDto.UserId = userId
	id, err := h.services.List.Create(models.ToList(listDto))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get All Lists
// @Security ApiKeyAuth
// @Tags lists
// @Description get all lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ListDto
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.services.List.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, models.ToListDTOs(lists))
}

// @Summary Get List By Id
// @Security ApiKeyAuth
// @Tags lists
// @Description get list by id
// @Param id path int true "List ID"
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ListDto
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id} [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id list not found")
		return
	}
	list, err := h.services.List.GetListById(userId, uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if list.UserId == 0 {
		newErrorResponse(c, http.StatusBadRequest, "id list not found")
		return
	}
	c.JSON(http.StatusOK, models.ToListDto(list))
}

// @Summary Update List By Id
// @Security ApiKeyAuth
// @Tags lists
// @Description update list
// @Param id path int true "List ID"
// @ID update-list
// @Accept  json
// @Produce  json
// @Param input body models.ListDto true "list info"
// @Success 200 {object} models.ListDto
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
	var listDto models.ListDto
	var list models.List
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id param error")
		return
	}
	if err = c.BindJSON(&listDto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	list.Title = listDto.Title
	list.Description = listDto.Description
	err = h.services.List.Update(userId, uint(id), list)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, models.ToListDto(list))
}

// @Summary Delete List By Id
// @Security ApiKeyAuth
// @Tags lists
// @Description delete list
// @Param id path int true "List ID"
// @ID delete-list
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id list not found")
		return
	}
	err = h.services.List.Delete(userId, uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

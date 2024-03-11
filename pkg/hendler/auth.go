package hendler

import (
	"github.com/KhilikEgor/todo-app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Hendler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Hendler) signIn(c *gin.Context) {

}

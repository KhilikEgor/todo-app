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

}

func (h *Hendler) signIn(c *gin.Context) {

}

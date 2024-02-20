package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	models *UserModel
}

func NewHandler(models *UserModel) *Handler {
	return &Handler{
		models: models,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserRequest
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.models.Insert(c.Request.Context(), &u)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) LoginUser(c *gin.Context) {
	var u LoginUserRequest
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.models.Login(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("jwt", user.AccessToken, 3600, "/", "localhost", false, true)

	res := LoginUserRes{
		Id:       user.Id,
		Username: user.Username,
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) LogoutUser(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}

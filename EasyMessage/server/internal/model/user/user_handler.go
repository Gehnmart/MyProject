package user

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	model *UserModel
}

const (
	authorizationHeader = "Authorization"
)

func InitHandler(db *sql.DB) *Handler {
	return &Handler{
		&UserModel{
			db,
		},
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user_req CreateUserReq
	err := c.ShouldBindJSON(&user_req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_res, err := h.model.CreateUser(c, user_req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, user_res)
}

func (h *Handler) LoginUser(c *gin.Context) {
	var user_req LoginUserReq
	err := c.ShouldBindJSON(&user_req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	user_res, err := h.model.LoginUser(c, user_req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, user_res)
}

func (h *Handler) IdentityUser(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Empty auth header"})
		c.Abort()
		return
	}

	header_parts := strings.Split(header, " ")
	if len(header_parts) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Empty auth header"})
		c.Abort()
		return
	}

	userId, err := h.model.ParseToken(header_parts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user token"})
		c.Abort()
		return
	}

	c.Set("userId", userId)
}

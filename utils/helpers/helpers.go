package helpers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserDeatil(c *gin.Context) (uuid.UUID, string) {
	session := sessions.Default(c)
	userId := session.Get("user_id")
	role := session.Get("role")
	return uuid.MustParse(userId.(string)), role.(string)
}

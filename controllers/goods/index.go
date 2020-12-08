package goods

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 接口概要说明
// @Description 接口详细描述信息
// @Tags 用户信息
// @Param id path int true "ID"
// @Param name query string false "name"
// @Success 200
// @Failure 400
// @Router /goods/{id} [get]
func GetOne(c *gin.Context) {
	c.JSON(http.StatusOK, "1")
}

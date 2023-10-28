package other

import (
	"github.com/gin-gonic/gin"
)

// Test App Find Stuff
// @Summary Test App Find Stuff
// @Description Find one stuff for the test app
// @Tags Test app stuff
// @Accept json
// @Produce json
// @Param user body other.Other true "Stuff data"
// @Success 200 {object} string
// @Failure 400
// @Failure 409
// @Failure 422 {object} map[string]string
// @Failure 500
// @Router /other/{id} [GET]
// @Security Bearer
func (s Other) Find(ctx *gin.Context) {
	s.controller.FindOne(ctx, &s)
}

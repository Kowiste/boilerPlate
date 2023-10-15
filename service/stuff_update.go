package service

import (
	"github.com/gin-gonic/gin"
	"test.com/model"
)

// Test App Create Stuff
// @Summary Back Office User
// @Description Create a stuff for the test app
// @Tags Test app stuff
// @Accept json
// @Produce json
// @Param user body model.Stuff true "Stuff data"
// @Success 200 {object} string
// @Failure 400
// @Failure 409
// @Failure 422 {object} map[string]string
// @Failure 500
// @Router /stuff/update/{id} [PATCH]
// @Security Bearer
func (s service) Update(c *gin.Context) {
	data := model.Stuff{}
	s.updateCore(c, &data)
}

package handler

import (
	"hamster-paas/pkg/utils/logger"

	"github.com/gin-gonic/gin"
)

func (h *HandlerServer) icpVersion(g *gin.Context) {
	resp, err := h.icpService.GetVersion()
	if err != nil {
		logger.Errorf("dfx version failed: %s", err)
		Fail(err.Error(), g)
		return
	}
	Success(resp, g)
}

func (h *HandlerServer) icpBalance(c *gin.Context) {
	_, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	resp := "icp balance"
	Success(resp, c)
}

package handler

import (
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/rpc/aline"
)

func (h *HandlerServer) getProjectList(c *gin.Context) {
	userAny, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	user := userAny.(aline.User)
	projectService, err := application.GetBean[*aline.ProjectService]("projectService")
	if err != nil {
		Fail("get project service error", c)
		return
	}
	projectIdAndNameList := projectService.GetProjectByUserId(user.Id)
	Success(projectIdAndNameList, c)
}
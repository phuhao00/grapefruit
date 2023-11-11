package login

import (
	"github.com/gin-gonic/gin"
	"grapefruit/internal/app/service/login"
	"grapefruit/internal/domain/do"
	"grapefruit/internal/domain/errwrap"
	"net/http"
)

// Login godoc
// @Schemes
// @Description  user login
// @Tags user
// @Accept json
// @Produce json
// @Param request body do.LoginReq true "login success response"
// @Success 200 {object} do.LoginRsp
// @Security JWT
// @Router /api/login [post]
// @Router /dev/api/login [post]
func Login(ctx *gin.Context) {
	var req do.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := login.LoginService.Login(req.Name, req.Pwd)
	if err != nil {
		ctx.JSON(http.StatusOK, errwrap.DefaultFailedErr(err.Error()))
	}
	ctx.JSON(http.StatusOK, errwrap.DefaultSuccessWithResponse(do.LoginRsp{}))
}

// UserRegister godoc
// @Schemes
// @Description  user register
// @Tags user
// @Accept json
// @Produce json
// @Param request body do.RegisterReq true "register success response"
// @Success 200 {object} do.RegisterRsp
// @Security JWT
// @Router /api/register [post]
// @Router /dev/api/register [post]
func UserRegister(ctx *gin.Context) {
	var req do.RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := login.LoginService.Register(req.Name, req.Pwd, req.Email)
	if err != nil {
		ctx.JSON(http.StatusOK, errwrap.DefaultFailedErr(err.Error()))
	}
	ctx.JSON(http.StatusOK, errwrap.DefaultSuccessWithResponse(do.RegisterRsp{}))
}

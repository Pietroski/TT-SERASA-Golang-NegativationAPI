package controllers

import (
	"fmt"
	legacy_service "github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/services/legacy_client"
	"github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/util"
	"github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/util/notification"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	LegacyHandlers ILegacyHandlers = &sLegacyHandlers{}

	envConfig util.SConfig
	err error
)

type ILegacyHandlers interface {
	GetNegativatedByID(ctx *gin.Context)
	ListNegativated(ctx *gin.Context)

	emitLegacyResponse(ctx *gin.Context, resp *http.Response) // legacy content
}

type sLegacyHandlers struct {}

func init() {
	envConfig, err = util.Config.LoadConfig("../..")
	if err != nil {
		fmt.Println(err)
	}
}

func (handler *sLegacyHandlers) GetNegativatedByID(ctx *gin.Context) {
	var req getNegativatedByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	// TODO: implement .env variable
	url := fmt.Sprintf("localhost:7007/v1/negativation/%d", req.ID)
	legacy_service.LegacyClient.InitCaller(http.MethodGet, url, nil)
	resp := legacy_service.LegacyClient.Call()
	handler.emitLegacyResponse(ctx, resp)
}

func (handler *sLegacyHandlers) ListNegativated(ctx *gin.Context) {
	var req listNegativated
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	page := (req.PageNumber - 1) * req.PageSize
	limit := req.PageSize

	// TODO: implement .env variable
	url := fmt.Sprintf("localhost:7007/v1/list-negativations?page=%d&limit=%d", page, limit)
	legacy_service.LegacyClient.InitCaller(http.MethodGet, url, nil)
	resp := legacy_service.LegacyClient.Call()
	handler.emitLegacyResponse(ctx, resp)
}

// TODO: create a library for this kind of response
func (handler *sLegacyHandlers) emitLegacyResponse(ctx *gin.Context, resp *http.Response) {
	if resp == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal request error",
		})
		return
	}

	if resp.StatusCode >= 500 && resp.StatusCode <= 599 {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "Legacy server error",
		})
		return
	}

	if resp.StatusCode >= 400 && resp.StatusCode <= 499 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Legacy server bad request",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp.Body)
}
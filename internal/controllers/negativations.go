package controllers

import (
	"database/sql"
	negativations "github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/services/negativation"
	"github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/util/notification"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	NegativationHandlers INegativationHandlers = &sNegativationHandlers{}
)

type INegativationHandlers interface {
	InitNegativationStore(store negativations.Store)

	Negativate(ctx *gin.Context)
	GetNegativatedByID(ctx *gin.Context)
	ListNegativated(ctx *gin.Context)
	DeleteNegativated(ctx *gin.Context)
	UpdateNegativated(ctx *gin.Context)
}

type sNegativationHandlers struct {
	Store negativations.Store
}

// InitNegativationStore should be called before any handler.
// It will panic otherwise.
func (handler *sNegativationHandlers) InitNegativationStore(store negativations.Store) {
	handler.Store = store
}

type negativateRequest struct {
	CompanyDocument  string    `json:"companyDocument" binding:"required"`
	CompanyName      string    `json:"companyName" binding:"required"`
	CustomerDocument string    `json:"customerDocument" binding:"required"`
	Value            float64   `json:"value" binding:"required"`
	Contract         string    `json:"contract" binding:"required"`
	DebtDate         time.Time `json:"debtDate" binding:"required"`
	InclusionDate    time.Time `json:"inclusionDate"`
}

// Negativate godoc
// @Summary Negativation Creation
// @Description creates a negativated client
// @ID negativate
// @Produce  json
// @Success 200 {slice} negativations.Negativations
// @Failure 400,404 {object} ErrorStruct
// @Failure 500,503 {object} ErrorStruct
// @Failure default {object} ErrorStruct
// @host localhost:8008
// @BasePath /v2
// @Router /negativate [post]
func (handler *sNegativationHandlers) Negativate(ctx *gin.Context) {
	var req negativateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	arg := negativations.NegativateParams{
		CompanyDocument:  req.CompanyDocument,
		CompanyName:      req.CompanyName,
		CustomerDocument: req.CustomerDocument,
		Value:            req.Value,
		Contract:         req.Contract,
		DebtDate:         req.DebtDate,
		InclusionDate:    req.InclusionDate,
	}

	negativation, err := handler.Store.Negativate(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	ctx.JSON(http.StatusCreated, negativation)
}

type getNegativatedByIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// GetNegativatedByID godoc
// @Summary Gets a negativated
// @Description gets a negativated by its id
// @ID get-negativated-by-id
// @Produce  json
// @Success 200 {slice} negativations.Negativations
// @Failure 400,404 {object} ErrorStruct
// @Failure 500,503 {object} ErrorStruct
// @Failure default {object} ErrorStruct
// @host localhost:8008
// @BasePath /v2
// @Router /negativated/{id} [get]
func (handler *sNegativationHandlers) GetNegativatedByID(ctx *gin.Context) {
	var req getNegativatedByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	negativation, err := handler.Store.GetNegativatedByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, notification.ClientError.Response(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, notification.ClientError.Response(err))
		return
	}

	ctx.JSON(http.StatusOK, negativation)
}

type listNegativated struct {
	PageNumber int32 `form:"page_number" biding:"required,min=1"`
	PageSize   int32 `form:"page_size" binding:"required,min=5,max=25"`
}

// ListNegativated godoc
// @Summary Lists negativated
// @Description lists negativated amounts
// @ID list-negativated
// @Produce  json
// @Success 200 {slice} negativations.Negativations
// @Failure 400,404 {object} ErrorStruct
// @Failure 500,503 {object} ErrorStruct
// @Failure default {object} ErrorStruct
// @host localhost:8008
// @BasePath /v2
// @Router /list-negativated/ [get]
func (handler *sNegativationHandlers) ListNegativated(ctx *gin.Context) {
	var req listNegativated
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	arg := negativations.ListNegativatedParams{
		Limit:  req.PageSize,
		Offset: (req.PageNumber - 1) * req.PageSize,
	}

	negativation, err := handler.Store.ListNegativated(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, notification.ClientError.Response(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	ctx.JSON(http.StatusOK, negativation)
}

type deleteNegativated struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// DeleteNegativated godoc
// @Summary Deletes a negativated
// @Description deletes a negativated by its id
// @ID delete-negativated-by-id
// @Produce  json
// @Success 200 {slice} negativations.Negativations
// @Failure 400,404 {object} ErrorStruct
// @Failure 500,503 {object} ErrorStruct
// @Failure default {object} ErrorStruct
// @host localhost:8008
// @BasePath /v2
// @Router /delete-negativated/{id} [delete]
func (handler *sNegativationHandlers) DeleteNegativated(ctx *gin.Context) {
	var req deleteNegativated
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	err := handler.Store.DeleteNegativated(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, notification.ClientError.Response(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	msg := struct {
		msg string
	}{
		msg: "deleted",
	}

	ctx.JSON(http.StatusOK, msg)
}

type updateNegativated struct {
	ID               int64     `json:"id" binding:"required"`
	CompanyDocument  string    `json:"companyDocument" binding:"required"`
	CompanyName      string    `json:"companyName" binding:"required"`
	CustomerDocument string    `json:"customerDocument" binding:"required"`
	Value            float64   `json:"value" binding:"required"`
	Contract         string    `json:"contract" binding:"required"`
	DebtDate         time.Time `json:"debtDate" binding:"required"`
	InclusionDate    time.Time `json:"inclusionDate"`
}

// UpdateNegativated godoc
// @Summary Updates negativation
// @Description updates a negativated client
// @ID update-negativated
// @Produce  json
// @Success 200 {slice} negativations.Negativations
// @Failure 400,404 {object} ErrorStruct
// @Failure 500,503 {object} ErrorStruct
// @Failure default {object} ErrorStruct
// @host localhost:8008
// @BasePath /v2
// @Router /update-negativate [put]
func (handler *sNegativationHandlers) UpdateNegativated(ctx *gin.Context) {
	var req updateNegativated
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	arg := negativations.UpdateNegativatedParams{
		ID:               req.ID,
		CompanyDocument:  req.CompanyDocument,
		CompanyName:      req.CompanyName,
		CustomerDocument: req.CustomerDocument,
		Value:            req.Value,
		Contract:         req.Contract,
		DebtDate:         req.DebtDate,
		InclusionDate:    req.InclusionDate,
	}

	updatedNegativation, err := handler.Store.UpdateNegativated(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, notification.ClientError.Response(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedNegativation)
}

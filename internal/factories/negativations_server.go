package factories

import (
	"github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/controllers"
	negativations "github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/services/negativation"
	"github.com/gin-gonic/gin"
)

var (
	Negativations iNegativationServer = &sNegativationsServer{}
)

type iNegativationServer interface {
	NewServer(store negativations.Store) *sNegativationsServer
	Start(address string) error
}

type sNegativationsServer struct {
	Store    negativations.Store
	Router   *gin.Engine
	Handlers controllers.INegativationHandlers
}

func (ns *sNegativationsServer) NewServer(store negativations.Store) *sNegativationsServer {
	server := &sNegativationsServer{
		Store: store,
		Handlers: controllers.NegativationHandlers,
	}
	controllers.NegativationHandlers.InitNegativationStore(store)

	router := gin.New()
	v2 := router.Group("/v2")

	v2.GET("/ping", controllers.Check.Ping)

	v2.POST("/negativate", server.Handlers.Negativate)
	v2.GET("/negativated/:id", server.Handlers.GetNegativatedByID)
	v2.GET("/list-negativated", server.Handlers.ListNegativated)
	v2.DELETE("/delete-negativated/:id", server.Handlers.DeleteNegativated)
	v2.PUT("/update-negativated", server.Handlers.UpdateNegativated)

	server.Router = router
	return server
}

func (ns *sNegativationsServer) Start(address string) error {
	return ns.Router.Run(address)
}

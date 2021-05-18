package factories

import (
	"github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/internal/controllers"
	"github.com/gin-gonic/gin"
)

var (
	LegacyProxy ILegacyServer = &SLegacyServer{}
)

type ILegacyServer interface {
	NewLegacyServer() *SLegacyServer
	Start(address string) error
}

type SLegacyServer struct {
	Router *gin.Engine
}

func (ns *SLegacyServer) NewLegacyServer() *SLegacyServer {
	server := &SLegacyServer{}

	router := gin.New()
	v1 := router.Group("/v1-legacy")

	v1.GET("/ping", controllers.Check.Ping)
	v1.GET("/negativated/:id", controllers.LegacyHandlers.GetNegativatedByID)
	v1.GET("/list-negativated", controllers.LegacyHandlers.ListNegativated)

	server.Router = router
	return server
}

func (ns *SLegacyServer) Start(address string) error {
	return ns.Router.Run(address)
}

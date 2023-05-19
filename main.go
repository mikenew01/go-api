package main

import (
	"dev.azure.com/ceabr/Tech/_git/alv-omega-crm-api-clientes/src/configs/logger"
	"dev.azure.com/ceabr/Tech/_git/alv-omega-crm-api-clientes/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitLog()

	router := gin.Default()
	routes.SetupRouter(router)
	router.Run(":8080")
}

package routes

import (
	"dev.azure.com/ceabr/Tech/_git/alv-omega-crm-api-clientes/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/clientes", controllers.CadastrarCliente)
}

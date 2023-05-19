package controllers

import (
	"dev.azure.com/ceabr/Tech/_git/alv-omega-crm-api-clientes/configs/logger"
	"dev.azure.com/ceabr/Tech/_git/alv-omega-crm-api-clientes/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CadastrarCliente(c *gin.Context) {
	var cliente services.ClientePontoVendaCadastro

	if err := c.ShouldBindJSON(&cliente); err != nil {
		logger.Error("ClienteController", fmt.Sprintf("cliente '%s' na loja '%d' com optin na stix '%s' recebido com erro '%s' ", cliente.MemberNumber, cliente.Loja, cliente.OptinStix, err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clienteCadastrado, _ := services.CadastrarCliente(&cliente)
	logger.Info("ClienteController", fmt.Sprintf("cliente '%s' recebido com sucesso na loja '%d' com optin na stix '%s'", cliente.MemberNumber, cliente.Loja, cliente.OptinStix))

	c.JSON(http.StatusOK, clienteCadastrado)
}

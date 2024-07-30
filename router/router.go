package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializando o Router nas configurações Default do Gin
	router := gin.Default() // Corrigido: use := para declarar e inicializar

	// Definindo uma Rota
	router.GET("/ping", func(c *gin.Context) { // Corrigido: removeu relativePath: e handlers...:
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Executando/Rodando a API
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}

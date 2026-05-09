package routes

import (
	"net/http"
	"strings"

	"github.com/Lubrum/github-actions-with-go/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() error {
	return SetupRouter().Run()
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", controllers.TodosAlunos)
	r.GET("/alunos/:id", controllers.BuscarAlunoPorID)
	r.POST("/alunos", controllers.CriarNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/alunos/", controllers.BuscaAlunoPorCPF)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(func(c *gin.Context) {
		nome := strings.Trim(c.Request.URL.Path, "/")
		if c.Request.Method == http.MethodGet && nome != "" && !strings.Contains(nome, "/") {
			c.Params = append(c.Params, gin.Param{Key: "nome", Value: nome})
			controllers.Saudacoes(c)
			return
		}

		controllers.RotaNaoEncontrada(c)
	})

	return r
}

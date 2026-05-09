package controllers

import (
	"errors"
	"net/http"

	"github.com/Lubrum/github-actions-with-go/database"
	"github.com/Lubrum/github-actions-with-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Saudacoes(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(http.StatusOK, gin.H{
		"API diz": "E ai " + nome + ", Tudo beleza?",
	})
}

func TodosAlunos(c *gin.Context) {
	alunos := []models.Aluno{}
	if err := database.DB.Find(&alunos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao listar alunos"})
		return
	}

	c.JSON(http.StatusOK, alunos)
}

func CriarNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := database.DB.Create(&aluno).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao criar aluno"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func BuscarAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	if err := database.DB.First(&aluno, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"Not Found": "Aluno não encontrado"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao buscar aluno"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	if err := database.DB.First(&aluno, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"Not Found": "Aluno não encontrado"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao buscar aluno"})
		return
	}

	if err := database.DB.Delete(&aluno).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao deletar aluno"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

func EditarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	if err := database.DB.First(&aluno, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"Not Found": "Aluno não encontrado"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao buscar aluno"})
		return
	}

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := database.DB.Save(&aluno).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao editar aluno"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	if err := database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"Not Found": "Aluno não encontrado"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao buscar aluno"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func ExibePaginaIndex(c *gin.Context) {
	alunos := []models.Aluno{}
	if err := database.DB.Find(&alunos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao listar alunos"})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/Lubrum/github-actions-with-go/controllers"
	"github.com/Lubrum/github-actions-with-go/database"
	"github.com/Lubrum/github-actions-with-go/models"
	"github.com/Lubrum/github-actions-with-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func ConectaBancoDeTeste(t *testing.T) {
	t.Helper()

	if os.Getenv("HOST") == "" || os.Getenv("USER") == "" || os.Getenv("DBNAME") == "" || os.Getenv("DBPORT") == "" {
		t.Skip("variaveis de ambiente do banco de dados nao configuradas")
	}

	require.NoError(t, database.ConectaComBancoDeDados())
}

func CriaAlunoMock(t *testing.T) {
	t.Helper()

	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	require.NoError(t, database.DB.Create(&aluno).Error)
	ID = int(aluno.ID)
}

func DeletaAlunoMock(t *testing.T) {
	t.Helper()

	var aluno models.Aluno
	require.NoError(t, database.DB.Delete(&aluno, ID).Error)
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := routes.SetupRouter()
	req := httptest.NewRequest(http.MethodGet, "/gui", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockDaResposta := `{"API diz":"E ai gui, Tudo beleza?"}`
	assert.Equal(t, mockDaResposta, resposta.Body.String())
}

func TestListaTodosOsAlunosHanlder(t *testing.T) {
	ConectaBancoDeTeste(t)
	CriaAlunoMock(t)
	defer DeletaAlunoMock(t)

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.TodosAlunos)
	req := httptest.NewRequest(http.MethodGet, "/alunos", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBucaAlunoPorCPFHandler(t *testing.T) {
	ConectaBancoDeTeste(t)
	CriaAlunoMock(t)
	defer DeletaAlunoMock(t)

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req := httptest.NewRequest(http.MethodGet, "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	ConectaBancoDeTeste(t)
	CriaAlunoMock(t)
	defer DeletaAlunoMock(t)

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscarAlunoPorID)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req := httptest.NewRequest(http.MethodGet, pathDaBusca, nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno
	require.NoError(t, json.Unmarshal(resposta.Body.Bytes(), &alunoMock))
	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletaAlunoHandler(t *testing.T) {
	ConectaBancoDeTeste(t)
	CriaAlunoMock(t)

	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req := httptest.NewRequest(http.MethodDelete, pathDeBusca, nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaUmAlunoHandler(t *testing.T) {
	ConectaBancoDeTeste(t)
	CriaAlunoMock(t)
	defer DeletaAlunoMock(t)

	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "47123456789", RG: "123456700"}
	valorJson, err := json.Marshal(aluno)
	require.NoError(t, err)
	pathParaEditar := "/alunos/" + strconv.Itoa(ID)
	req := httptest.NewRequest(http.MethodPatch, pathParaEditar, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	var alunoMockAtualizado models.Aluno
	require.NoError(t, json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado))
	assert.Equal(t, "47123456789", alunoMockAtualizado.CPF)
	assert.Equal(t, "123456700", alunoMockAtualizado.RG)
	assert.Equal(t, "Nome do Aluno Teste", alunoMockAtualizado.Nome)
}

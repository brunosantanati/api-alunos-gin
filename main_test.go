package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brunosantanati/api-go-gin/controllers"
	"github.com/brunosantanati/api-go-gin/database"
	"github.com/brunosantanati/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

/*func TestFalhador(t *testing.T) {
	t.Fatalf("Teste falhou de propósito, não se preocupe")
}*/

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/Bruno", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	/*if resposta.Code != http.StatusOK {
		t.Fatalf("Status error: valor recebido foi %d e o esperado era %d",
			resposta.Code, http.StatusOK)
	}*/
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")

	mockDaResposta := `{"API diz:":"E ai Bruno, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)

	//fmt.Println(string(respostaBody))
	//fmt.Println(mockDaResposta)

	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	//fmt.Println(resposta.Body)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brunosantanati/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

/*func TestFalhador(t *testing.T) {
	t.Fatalf("Teste falhou de propósito, não se preocupe")
}*/

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "Bruno", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	/*if resposta.Code != http.StatusOK {
		t.Fatalf("Status error: valor recebido foi %d e o esperado era %d",
			resposta.Code, http.StatusOK)
	}*/
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockDaResposta := `{"API diz:":"E ai Bruno, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)

	fmt.Println(string(respostaBody))
	fmt.Println(mockDaResposta)

	assert.Equal(t, mockDaResposta, string(respostaBody))
}

package auth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/configs"
)

func CriarToken(tenantId uint64, isCustmizavel bool) (tokenNew string, err error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["tenantId"] = tenantId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	tokenNew, err = token.SignedString([]byte(configs.SecretKey))
	return
}

func ValidarToken(r *gin.Context) (erro error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return
	}

	erro = errors.New("token inválido")
	return
}

func ExtrairTenantID(r *gin.Context) (tenantId uint64, erro error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tenantId, erro = strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["tenantId"]), 10, 64)
		if erro != nil {
			return
		}

		return
	}

	erro = errors.New("token inválido")
	return
}

func extrairToken(r *gin.Context) (token string) {
	token = r.GetHeader("Authorization")

	if len(strings.Split(token, " ")) != 2 {
		return
	}
	token = strings.Split(token, " ")[1]
	return
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}

	return configs.SecretKey, nil
}

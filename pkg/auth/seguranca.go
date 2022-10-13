package auth

import "golang.org/x/crypto/bcrypt"

func Hash(senha string) (cript []byte, err error) {
	cript, err = bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return
}

func VerificarSenha(senhaComHash, senhaString string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
	return
}

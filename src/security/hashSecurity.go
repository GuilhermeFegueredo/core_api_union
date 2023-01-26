package security

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma senha e retorna um hash dela
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerificarSenha valida se o hash condiz com a senha
func VerificarSenha(senhaComHash, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senha))
}

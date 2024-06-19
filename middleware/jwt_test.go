package middleware

import (
	"encoding/json"
	"testing"
)

func TestJwt(t *testing.T) {
	jwt := InitJwt("SecretKey")
	token := jwt.GenerateToken("qiqi") // 生成有效期为24小时的 JWT
	t.Log(token)

	claims := jwt.ParseToken(token)
	b, _ := json.Marshal(claims)
	t.Log(string(b))

	newToken := jwt.CreateTokenByOldToken(token)
	t.Log(newToken)
	newClaims := jwt.ParseToken(newToken)
	b, _ = json.Marshal(newClaims)
	t.Log(string(b))

}

func TestVerifyJwt(t *testing.T) {
	jwt := InitJwt("SecretKey")
	claims := jwt.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJTdGFyIiwic3ViIjoiMzAzMjMxNjYzNSIsImF1ZCI6WyJBbGwgcGxhdGZvcm1zIl0sImV4cCI6MTcxOTM5ODU0MCwibmJmIjoxNzE4NzkzNzQwLCJpYXQiOjE3MTg3OTM3NDAsImp0aSI6ImZmMWRhM2Q4YjZkNTY5OWZiNTNmZGU0YjcwMWVjMmQ2In0.A_oSP_S5HwEYHybBhSNu8_cbPB5KebbLOg-9tYCNAGE")
	b, _ := json.Marshal(claims)
	t.Log(string(b))
	t.Log(claims.GetSubject())
}

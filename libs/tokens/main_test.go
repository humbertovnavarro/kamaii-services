package tokens

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUserToken(t *testing.T) {
	TOKEN_SECRET = []byte{1, 0, 1, 0}
	token, err := GenerateUserToken("foo")
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestParseUserToken(t *testing.T) {
	// Good token
	TOKEN_SECRET = []byte{1, 0, 1, 0}
	tokenString, _ := GenerateUserToken("foo")
	claims, err := ParseUserToken(tokenString)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	assert.Equal(t, claims.Id, "foo")
	assert.Equal(t, claims.ExpiresAt, claims.IssuedAt+int64(TOKEN_TTL), "does not add ttl properly")
	// Junk token
	tokenString = "bar"
	claims, err = ParseUserToken(tokenString)
	assert.Error(t, err)
	assert.Nil(t, claims)

}

func TestExpiredToken(t *testing.T) {
	TOKEN_TTL = -1
	// Good token
	TOKEN_SECRET = []byte{1, 0, 1, 0}
	tokenString, _ := GenerateUserToken("foo")
	claims, err := ParseUserToken(tokenString)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "token is expired by 1s")
	assert.Nil(t, claims)
	// Junk token
	tokenString = "bar"
	claims, err = ParseUserToken(tokenString)
	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Equal(t, err.Error(), "token contains an invalid number of segments")
}

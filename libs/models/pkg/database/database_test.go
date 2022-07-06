package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDSN(t *testing.T) {
	os.Setenv("DB_SSL", "")
	os.Setenv("DB_TIMEZONE", "")
	os.Setenv("DB_HOST", "baz")
	os.Setenv("DB_USER", "buz")
	os.Setenv("DB_PASSWORD", "fiz")
	os.Setenv("DB_NAME", "buzz")
	os.Setenv("DB_PORT", "apple")
	assert.Equal(t, createDSN(), "host=baz user=buz password=fiz dbname=buzz port=apple sslmode=disabled TimeZone=America/Los_Angeles")
}

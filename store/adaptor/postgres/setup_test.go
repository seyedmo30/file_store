package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDB(t *testing.T) {
	var err error
	
	db := GetDB()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.Ping()
	assert.NoError(t, err)

}

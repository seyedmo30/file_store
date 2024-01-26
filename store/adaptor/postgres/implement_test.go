package postgres

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMetadata(t *testing.T) {
	err := CreateMetadata(context.Background())

	assert.NoError(t, err)

}

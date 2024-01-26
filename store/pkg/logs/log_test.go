package logs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {

	logger := Connect()

	assert.NotNil(t, logger)

}

// +build unit

package usecases

import (
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestPostGresClientInit will test that we can initialize the handler.
func TestPostGresClientInit(t *testing.T) {
	assert := assert.New(t)

	postgresHandler := NewPostGresClient(testutils.DBConfig)

	assert.NotNil(postgresHandler, "postgresHandler should have been initialized")
}

// TestConfigInit will test that the config was initialized.
func TestConfigInit(t *testing.T) {
	assert := assert.New(t)

	postgresHandler := NewPostGresClient(testutils.DBConfig)

	assert.NotNil(postgresHandler.cfg, "postgresHandler should have initialized a DBConfig")
	assert.EqualValues(testutils.DBPort, postgresHandler.cfg.DBPort, "DBPort in cfg should be the same as the one in testutils")
}
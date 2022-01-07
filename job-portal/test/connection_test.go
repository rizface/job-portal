package test

import (
	"github.com/stretchr/testify/assert"
	"job-portal/helper"
	"testing"
)


func TestConnection(t *testing.T) {
	db := helper.Connection()
	assert.NotNil(t,db)
}

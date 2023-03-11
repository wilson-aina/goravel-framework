package console

import (
	"testing"

	"github.com/stretchr/testify/assert"

	consolemocks "github.com/goravel/framework/contracts/console/mocks"
	"github.com/goravel/framework/support/file"
)

func TestModelMakeCommand(t *testing.T) {
	modelMakeCommand := &ModelMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	assert.Nil(t, modelMakeCommand.Handle(mockContext))
	assert.False(t, file.Exists("app/models/user.go"))

	mockContext.On("Argument", 0).Return("User").Once()
	assert.Nil(t, modelMakeCommand.Handle(mockContext))
	assert.True(t, file.Exists("app/models/user.go"))
	assert.True(t, file.Remove("app"))
}

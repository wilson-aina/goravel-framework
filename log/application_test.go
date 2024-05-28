package log

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goravel/framework/foundation/json"
	mocksconfig "github.com/goravel/framework/mocks/config"
	"github.com/goravel/framework/support/color"
)

func TestNewApplication(t *testing.T) {
	j := json.NewJson()
	app := NewApplication(nil, j)
	assert.NotNil(t, app)

	mockConfig := &mocksconfig.Config{}
	mockConfig.On("GetString", "logging.default").Return("test")
	mockConfig.On("GetString", "logging.channels.test.driver").Return("single")
	mockConfig.On("GetString", "logging.channels.test.path").Return("test")
	mockConfig.On("GetString", "logging.channels.test.level").Return("debug")
	mockConfig.On("GetBool", "logging.channels.test.print").Return(true)
	app = NewApplication(mockConfig, j)
	assert.NotNil(t, app)

	mockConfig = &mocksconfig.Config{}
	mockConfig.On("GetString", "logging.default").Return("test")
	mockConfig.On("GetString", "logging.channels.test.driver").Return("test")
	assert.Contains(t, color.CaptureOutput(func(w io.Writer) {
		assert.Nil(t, NewApplication(mockConfig, j))
	}), "Init facades.Log error: Error logging channel: test")
}

func TestApplication_Channel(t *testing.T) {
	mockConfig := &mocksconfig.Config{}
	mockConfig.On("GetString", "logging.default").Return("test")
	mockConfig.On("GetString", "logging.channels.test.driver").Return("single")
	mockConfig.On("GetString", "logging.channels.test.path").Return("test")
	mockConfig.On("GetString", "logging.channels.test.level").Return("debug")
	mockConfig.On("GetBool", "logging.channels.test.print").Return(true)
	app := NewApplication(mockConfig, json.NewJson())
	assert.NotNil(t, app)
	assert.NotNil(t, app.Channel(""))

	mockConfig.On("GetString", "logging.channels.dummy.driver").Return("daily")
	mockConfig.On("GetString", "logging.channels.dummy.path").Return("dummy")
	mockConfig.On("GetString", "logging.channels.dummy.level").Return("debug")
	mockConfig.On("GetBool", "logging.channels.dummy.print").Return(true)
	mockConfig.On("GetInt", "logging.channels.dummy.days").Return(1)
	writer := app.Channel("dummy")
	assert.NotNil(t, writer)

	mockConfig.On("GetString", "logging.channels.test2.driver").Return("test2")
	assert.Contains(t, color.CaptureOutput(func(w io.Writer) {
		assert.Nil(t, app.Channel("test2"))
	}), "Init facades.Log error: Error logging channel: test2")
}

func TestApplication_Stack(t *testing.T) {
	mockConfig := &mocksconfig.Config{}
	mockConfig.On("GetString", "logging.default").Return("test")
	mockConfig.On("GetString", "logging.channels.test.driver").Return("single")
	mockConfig.On("GetString", "logging.channels.test.path").Return("test")
	mockConfig.On("GetString", "logging.channels.test.level").Return("debug")
	mockConfig.On("GetBool", "logging.channels.test.print").Return(true)
	app := NewApplication(mockConfig, json.NewJson())
	assert.NotNil(t, app)
	assert.NotNil(t, app.Stack([]string{}))

	mockConfig.On("GetString", "logging.channels.test2.driver").Return("test2")
	assert.Contains(t, color.CaptureOutput(func(w io.Writer) {
		assert.Nil(t, app.Stack([]string{"", "test2", "daily"}))
	}), "Init facades.Log error: Error logging channel: test2")

	mockConfig.On("GetString", "logging.channels.dummy.driver").Return("daily")
	mockConfig.On("GetString", "logging.channels.dummy.path").Return("dummy")
	mockConfig.On("GetString", "logging.channels.dummy.level").Return("debug")
	mockConfig.On("GetBool", "logging.channels.dummy.print").Return(true)
	mockConfig.On("GetInt", "logging.channels.dummy.days").Return(1)
	assert.NotNil(t, app.Stack([]string{"dummy"}))
}

package r2

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestOptMethods(t *testing.T) {
	assert := assert.New(t)

	req := New("https://foo.bar.local")

	assert.Nil(OptMethod("OPTIONS")(req))
	assert.Equal("OPTIONS", req.Method)

	assert.Nil(OptGet()(req))
	assert.Equal("GET", req.Method)

	assert.Nil(OptPost()(req))
	assert.Equal("POST", req.Method)

	assert.Nil(OptPut()(req))
	assert.Equal("PUT", req.Method)

	assert.Nil(OptPatch()(req))
	assert.Equal("PATCH", req.Method)

	assert.Nil(OptDelete()(req))
	assert.Equal("DELETE", req.Method)
}

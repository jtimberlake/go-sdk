package configutil

import (
	"context"
	"testing"
	"time"

	"github.com/blend/go-sdk/assert"
)

func TestDuration(t *testing.T) {
	assert := assert.New(t)

	d := Duration(time.Second)

	ret, err := d.Duration(context.TODO())
	assert.Nil(err)
	assert.NotNil(ret)
	assert.Equal(time.Second, *ret)
}

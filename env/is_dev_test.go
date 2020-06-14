package env

import (
	"fmt"
	"testing"

	"github.com/blend/go-sdk/assert"
	"github.com/blend/go-sdk/uuid"
)

func TestIsDevlike(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		Input    string
		Expected bool
	}{
		{Input: ServiceEnvDev, Expected: false},
		{Input: ServiceEnvCI, Expected: false},
		{Input: ServiceEnvTest, Expected: false},
		{Input: ServiceEnvSandbox, Expected: false},
		{Input: ServiceEnvPreprod, Expected: true},
		{Input: ServiceEnvBeta, Expected: true},
		{Input: ServiceEnvProd, Expected: true},
		{Input: uuid.V4().String(), Expected: true},
		{Expected: true},
	}

	for _, testCase := range testCases {
		assert.Equal(!testCase.Expected, IsDevlike(testCase.Input), fmt.Sprintf("failed for: %s", testCase.Input))
	}
}

func TestIsDev(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		Input    string
		Expected bool
	}{
		{Input: ServiceEnvDev, Expected: true},
		{Input: ServiceEnvCI, Expected: false},
		{Input: ServiceEnvTest, Expected: false},
		{Input: ServiceEnvSandbox, Expected: false},
		{Input: ServiceEnvPreprod, Expected: false},
		{Input: ServiceEnvBeta, Expected: false},
		{Input: ServiceEnvProd, Expected: false},
		{Input: uuid.V4().String(), Expected: false},
		{Expected: false},
	}

	for _, testCase := range testCases {
		assert.Equal(testCase.Expected, IsDev(testCase.Input), fmt.Sprintf("failed for: %s", testCase.Input))
	}
}

package oauth

import (
	"github.com/mfirmanakbar/bookstore_utils-go/rest_errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestIsPublic(t *testing.T) {
	request, _ := http.NewRequest("", "", nil)
	request.Header.Set("X-Public", "false")

	result := IsPublic(request)
	assert.NotNil(t, result)
	assert.EqualValues(t, false, result)

	result = IsPublic(nil)
	assert.NotNil(t, result)
	assert.EqualValues(t, true, result)
}

func TestGetCallerId(t *testing.T) {
	result := GetCallerId(nil)
	assert.NotNil(t, result)
	assert.EqualValues(t, 0, result)

	request, _ := http.NewRequest("", "", nil)
	request.Header.Set("X-Caller-Id", "12345")
	result = GetCallerId(request)
	assert.NotNil(t, result)
	assert.EqualValues(t, 12345, result)

	request, _ = http.NewRequest("", "", nil)
	request.Header.Set("X-Caller-Id", "papeda-ambon")
	result = GetCallerId(request)
	assert.EqualValues(t, 0, result)
}

func TestGetClientId(t *testing.T) {
	result := GetClientId(nil)
	assert.NotNil(t, result)
	assert.EqualValues(t, 0, result)

	request, _ := http.NewRequest("", "", nil)
	request.Header.Set("X-Client-Id", "12345")
	result = GetClientId(request)
	assert.NotNil(t, result)
	assert.EqualValues(t, 12345, result)

	request, _ = http.NewRequest("", "", nil)
	request.Header.Set("X-Client-Id", "papeda-ambon")
	result = GetClientId(request)
	assert.EqualValues(t, 0, result)
}

func TestAuthenticateRequest(t *testing.T) {
	err := AuthenticateRequest(nil)
	assert.EqualValues(t, (*rest_errors.RestErr)(nil), err)
}

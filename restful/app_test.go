package restful

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	mock := httptest.NewServer(NewHandler())
	defer mock.Close()

	res, err := http.Get(mock.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("hello, world", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	mock := httptest.NewServer(NewHandler())
	defer mock.Close()

	res, err := http.Get(mock.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "Get user info")
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	mock := httptest.NewServer(NewHandler())
	defer mock.Close()

	res, err := http.Get(mock.URL + "/users/777")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "User Id : 777")
}
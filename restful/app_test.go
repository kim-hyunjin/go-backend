package restful

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
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
	assert.Contains(string(data), "No User With ID")
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	mock := httptest.NewServer(NewHandler())
	defer mock.Close()

	res, err := http.Post(mock.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"hyunjin", "last_name":"kim", "email":"hyunjin1612@gmail.com"}`))

	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	id := user.ID
	res, err = http.Get(mock.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)

	assert.Equal(http.StatusOK, res.StatusCode)
	user2 := new(User)
	err = json.NewDecoder(res.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.FirstName, user2.FirstName)

}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	mock := httptest.NewServer(NewHandler())
	defer mock.Close()

	req, _ := http.NewRequest("DELETE", mock.URL+"/users/1", nil)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User ID With: 1")

	res, err = http.Post(mock.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"hyunjin", "last_name":"kim", "email":"hyunjin1612@gmail.com"}`))

	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	req, _ = http.NewRequest("DELETE", mock.URL+"/users/1", nil)
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ = ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "User ID With: 1 Deleted!")
}

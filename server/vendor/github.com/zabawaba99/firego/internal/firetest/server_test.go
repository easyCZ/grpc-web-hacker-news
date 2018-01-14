package firetest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zabawaba99/firego/sync"
)

func TestNew(t *testing.T) {
	ft := New()
	require.NotNil(t, ft)
	assert.NotEmpty(t, ft.Secret)
}

func TestURL(t *testing.T) {
	ft := New()
	ft.Start()
	assert.Regexp(t, regexp.MustCompile(`https?://127.0.0.1:\d+`), ft.URL)

	ft.Close()
}

func TestClose(t *testing.T) {
	// ARRANGE
	ft := New()
	ft.Start()

	// ACT
	ft.Close()

	// ASSERT
	_, err := http.Get(ft.URL)
	assert.Error(t, err)
	assert.IsType(t, (*url.Error)(nil), err)
}

func TestValidJWT(t *testing.T) {
	for _, test := range []struct {
		name     string
		jwtToken string
		pass     bool
	}{
		// All tokens were generated using http://jwt.io/ with the secret "foo"
		{
			name:     "valid token",
			jwtToken: "eyJhbGciOiAiSFMyNTYiLCAidHlwIjogIkpXVCJ9.eyJ2IjowLCJkIjp7InVpZCI6IjEifSwiaWF0IjoxNDM3MTM5NTM5fQ.G7j81dhdMrJquGcy8bfKmOZOizzFBYRUMBK4CQIzX_E",
			pass:     true,
		},
		{
			name:     "valid token with exp",
			jwtToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ2IjowLCJkIjp7InVpZCI6IjEifSwiZXhwIjo5OTk5OTk5OTk5LCJpYXQiOjE0MzcxMzk1Mzl9.itryN7-cE-MDYi7I9hSFuey-AOVLSipPnZIxCGR0nOg",
			pass:     true,
		},
		{
			name:     "token with bad alg",
			jwtToken: "eyJhbGciOiJIUyIsInR5cCI6IkpXVCJ9.eyJ2IjowLCJkIjp7InVpZCI6IjEifSwiaWF0IjoxNDM3MTM5NTM5fQ.sacjjt7nrdhP3Yrp0wY8SSwPXpjhs4JMhH8s2PDrIh8",
			pass:     false,
		},
		{
			name:     "token with invalid typ",
			jwtToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXd3dUIn0.eyJ2IjowLCJkIjp7InVpZCI6IjEifSwiaWF0IjoxNDM3MTM5NTM5fQ._zpaFuWygUnsqAiws3B_l2xRjJPqo3SWm1Q65DOsNao",
			pass:     false,
		},
		{
			name:     "token expired token",
			jwtToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ2IjowLCJkIjp7InVpZCI6IjEifSwiZXhwIjoxNDM3MTM5NTM5LCJpYXQiOjE0MzcxMzk1Mzl9.FvxBGJKk32rGPv_VzJdJHMtz80_xHqO5Iccl2DkyOQs",
			pass:     false,
		},
		{
			name:     "token with invalid exp",
			jwtToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ2IjowLCJkIjp7InVpZCI6IjEifSwiZXhwIjoiOTk5OTk5OTk5OSIsImlhdCI6MTQzNzEzOTUzOX0._wISBo2CPpcYa6RMkbnKH5T4BNEMtsHlko6JEYcAEOM",
			pass:     false,
		},
		{
			name:     "token missing claim['data']",
			jwtToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ2IjowLCJpYXQiOjE0MzcxMzk1Mzl9.ihZ5BNsDbRIQA9ORgRBrDPCK-FBKlD2w32d0mSOYi6M",
			pass:     false,
		},
		{
			name:     "token missing data['uid']",
			jwtToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ2IjowLCJkIjp7fSwiaWF0IjoxNDM3MTM5NTM5fQ.ZcwtxTBdy1QWt6E9vShknDpx5gI2yIJxgy_taY9Yl-g",
			pass:     false,
		},
		{
			name:     "token invalid signature",
			jwtToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ2IjowLCJpYXQiOjE0MzcxMzk1Mzl9.oQuNcDoQ88EKTFOauENfhozDUlHX7JnRB4S-xfkyoP0",
			pass:     false,
		},
	} {
		ft := New()
		ft.Secret = "foo"
		assert.Equal(t, test.pass, ft.validJWT(test.jwtToken), test.name)
	}
}

func TestServeHTTP(t *testing.T) {
	// ARRANGE
	ft := New()
	ft.Start()

	// ACT
	req, err := http.NewRequest("GET", ft.URL+"/.json", nil)
	require.NoError(t, err)
	resp := httptest.NewRecorder()
	ft.serveHTTP(resp, req)

	// ASSERT
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestServeHTTP_MissingJSON(t *testing.T) {
	// ARRANGE
	ft := New()
	ft.Start()

	// ACT
	req, err := http.NewRequest("GET", ft.URL, nil)
	require.NoError(t, err)
	resp := httptest.NewRecorder()
	ft.serveHTTP(resp, req)

	// ASSERT
	assert.Equal(t, http.StatusForbidden, resp.Code)
	assert.Equal(t, missingJSONExtension, resp.Body.Bytes())
}

func TestServerInvalidBody(t *testing.T) {
	// ARRANGE
	ft := New()
	ft.Start()

	for _, method := range []string{"PATCH", "POST", "PUT"} {
		// ACT
		req, err := http.NewRequest(method, ft.URL+"/.json", strings.NewReader("{asd}"))
		require.NoError(t, err)
		resp := httptest.NewRecorder()
		ft.serveHTTP(resp, req)

		// ASSERT
		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Equal(t, invalidJSON, resp.Body.Bytes())
	}
}

func TestServerMissingBody(t *testing.T) {
	// ARRANGE
	ft := New()
	ft.Start()

	for _, method := range []string{"PATCH", "POST", "PUT"} {
		// ACT
		req, err := http.NewRequest(method, ft.URL+"/.json", bytes.NewReader(nil))
		require.NoError(t, err)
		resp := httptest.NewRecorder()
		ft.serveHTTP(resp, req)

		// ASSERT
		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Equal(t, missingBody, resp.Body.Bytes())
	}
}

func TestServeHTTPAuth(t *testing.T) {
	// ARRANGE
	ft := New()
	ft.Start()
	ft.RequireAuth(true)

	// ACT
	req, err := http.NewRequest("GET", ft.URL+"/.json?auth="+ft.Secret, nil)
	require.NoError(t, err)

	resp := httptest.NewRecorder()
	ft.serveHTTP(resp, req)

	// ASSERT
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestServeHTTPUnauthorized(t *testing.T) {
	// ARRANGE
	ft := New()
	ft.Start()
	ft.RequireAuth(true)

	// ACT
	req, err := http.NewRequest("GET", ft.URL+"/.json", nil)
	require.NoError(t, err)
	resp := httptest.NewRecorder()
	ft.serveHTTP(resp, req)

	// ASSERT
	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Equal(t, invalidAuth, resp.Body.Bytes())
}

func TestServeHTTPAuthJWT(t *testing.T) {
	// ARRANGE
	ft := New()
	ft.Secret = "foo"
	ft.Start()
	ft.RequireAuth(true)

	// ACT
	req, err := http.NewRequest("GET", ft.URL+"/.json?auth=eyJhbGciOiAiSFMyNTYiLCAidHlwIjogIkpXVCJ9.eyJ2IjowLCJkIjp7InVpZCI6IjEifSwiaWF0IjoxNDM3MTM5NTM5fQ.G7j81dhdMrJquGcy8bfKmOZOizzFBYRUMBK4CQIzX_E", nil)
	require.NoError(t, err)

	resp := httptest.NewRecorder()
	ft.serveHTTP(resp, req)

	// ASSERT
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestServeHTTPUnauthorizedJWT(t *testing.T) {
	// ARRANGE
	ft := New()
	ft.Start()
	ft.RequireAuth(true)

	// ACT
	req, err := http.NewRequest("GET", ft.URL+"/.json?auth=bad.jwt.nope", nil)
	require.NoError(t, err)
	resp := httptest.NewRecorder()
	ft.serveHTTP(resp, req)

	// ASSERT
	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Equal(t, invalidAuth, resp.Body.Bytes())
}

func TestServerCreate(t *testing.T) {
	// ARRANGE
	ft := New()
	ft.Start()

	// ACT
	body := `"bar"`
	req, err := http.NewRequest("POST", ft.URL+"/foo.json", strings.NewReader(body))
	require.NoError(t, err)
	resp := httptest.NewRecorder()
	ft.serveHTTP(resp, req)

	// ASSERT
	assert.Equal(t, http.StatusOK, resp.Code)

	var v map[string]string
	err = json.NewDecoder(resp.Body).Decode(&v)
	require.NoError(t, err)

	name, ok := v["name"]
	assert.True(t, ok)
	assert.NotEmpty(t, name)
}

func TestServerSet(t *testing.T) {
	// ARRANGE
	ft := New()
	ft.Start()

	// ACT
	body := `"bar"`
	req, err := http.NewRequest("PUT", ft.URL+"/foo.json", strings.NewReader(body))
	require.NoError(t, err)
	resp := httptest.NewRecorder()
	ft.serveHTTP(resp, req)

	// ASSERT
	assert.Equal(t, http.StatusOK, resp.Code)
	respBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, body, string(respBody))
}

func TestServerDel(t *testing.T) {
	// ARRANGE
	ft := New()
	ft.Start()
	path := "foo/bar"
	n := sync.NewNode("", 2)
	ft.db.add(path, n)

	// ACT
	req, err := http.NewRequest("DELETE", ft.URL+"/"+path+".json", nil)
	require.NoError(t, err)

	resp := httptest.NewRecorder()
	ft.serveHTTP(resp, req)

	// ASSERT
	assert.Equal(t, http.StatusOK, resp.Code)
	n = ft.db.get(path)
	assert.Nil(t, n)
}

func TestServerUpdate(t *testing.T) {
	// ARRANGE
	ft := New()
	ft.Start()

	path := "some/awesome/path"
	body := map[string]interface{}{
		"foo":  "bar",
		"fooy": true,
		"bar":  []interface{}{false, "lolz"},
	}
	ft.db.add(path, sync.NewNode("", body))

	// ACT
	newVal := `"notbar"`
	req, err := http.NewRequest("PATCH", ft.URL+"/some/awesome/path/foo.json", strings.NewReader(newVal))
	require.NoError(t, err)
	resp := httptest.NewRecorder()
	ft.serveHTTP(resp, req)

	// ASSERT
	assert.Equal(t, http.StatusOK, resp.Code)
	respBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, newVal, string(respBody))
}

func TestServerGet(t *testing.T) {
	// ARRANGE
	ft := New()
	ft.Start()

	path := "some/awesome/path"
	body := map[string]interface{}{
		"foo":  "bar",
		"fooy": true,
		"bar":  []interface{}{false, "lolz"},
	}
	ft.db.add(path, sync.NewNode("", body))

	b, err := json.Marshal(&body)
	require.NoError(t, err)

	// ACT
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s.json", ft.URL, path), bytes.NewReader(b))
	require.NoError(t, err)
	resp := httptest.NewRecorder()
	ft.serveHTTP(resp, req)

	// ASSERT
	assert.Equal(t, http.StatusOK, resp.Code)
	var respBody map[string]interface{}
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&respBody))

	assert.EqualValues(t, body, respBody)
}

func TestSanitizePath(t *testing.T) {
	for i, test := range []struct {
		path     string
		expected string
	}{
		{"/", ""},
		{"foo", "foo"},
		{"foo/", "foo"},
		{"/foo", "foo"},
		{"/foo/", "foo"},
		{"/foo/.json", "foo"}, // issue #6
	} {
		assert.Equal(t, test.expected, sanitizePath(test.path), "%d", i)
	}
}

func TestUnmarshal(t *testing.T) {
	v := "foo"
	jsonV := `"foo"`
	w := httptest.NewRecorder()
	r := strings.NewReader(jsonV)
	b, val, ok := unmarshal(w, r)
	assert.Equal(t, []byte(jsonV), b)
	assert.Equal(t, v, val)
	assert.True(t, ok)
}

func TestUnmarshal_MissingBody(t *testing.T) {
	w := httptest.NewRecorder()
	r := bytes.NewReader(nil)
	b, val, ok := unmarshal(w, r)
	assert.Nil(t, b)
	assert.Nil(t, val)
	assert.False(t, ok)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte(missingBody), w.Body.Bytes())
}

func TestUnmarshal_InvalidBody(t *testing.T) {
	w := httptest.NewRecorder()
	r := strings.NewReader("{asda}")
	b, val, ok := unmarshal(w, r)
	assert.Nil(t, b)
	assert.Nil(t, val)
	assert.False(t, ok)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, []byte(invalidJSON), w.Body.Bytes())
}

package firego

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShallow(t *testing.T) {
	t.Parallel()
	var (
		server = newTestServer("")
		fb     = New(server.URL, nil)
	)
	defer server.Close()

	fb.Shallow(true)
	fb.Value("")
	require.Len(t, server.receivedReqs, 1)

	req := server.receivedReqs[0]
	assert.Equal(t, shallowParam+"=true", req.URL.Query().Encode())

	fb.Shallow(false)
	fb.Value("")
	require.Len(t, server.receivedReqs, 2)

	req = server.receivedReqs[1]
	assert.Equal(t, "", req.URL.Query().Encode())
}

func TestOrderBy(t *testing.T) {
	t.Parallel()
	var (
		server = newTestServer("")
		fb     = New(server.URL, nil)
	)
	defer server.Close()

	fb.OrderBy("user_id").Value("")
	require.Len(t, server.receivedReqs, 1)

	req := server.receivedReqs[0]
	assert.Equal(t, orderByParam+"=%22user_id%22", req.URL.Query().Encode())
}

func TestEqualTo(t *testing.T) {
	t.Parallel()
	var (
		server = newTestServer("")
		fb     = New(server.URL, nil)
	)
	defer server.Close()

	fb.EqualTo("user_id").Value("")
	require.Len(t, server.receivedReqs, 1)

	req := server.receivedReqs[0]
	assert.Equal(t, equalToParam+"=%22user_id%22", req.URL.Query().Encode())
}

func TestEqualToValue(t *testing.T) {
	t.Parallel()
	var (
		server = newTestServer("")
		fb     = New(server.URL, nil)
	)
	defer server.Close()

	fb.EqualToValue(2).Value("")
	fb.EqualToValue("2").Value("")
	fb.EqualToValue(2.14).Value("")
	fb.EqualToValue("bar").Value("")
	require.Len(t, server.receivedReqs, 4)

	req := server.receivedReqs[0]
	assert.Equal(t, equalToParam+"=2", req.URL.Query().Encode())

	req = server.receivedReqs[1]
	assert.Equal(t, equalToParam+"=%222%22", req.URL.Query().Encode())

	req = server.receivedReqs[2]
	assert.Equal(t, equalToParam+"=2.14", req.URL.Query().Encode())

	req = server.receivedReqs[3]
	assert.Equal(t, equalToParam+"=%22bar%22", req.URL.Query().Encode())

}

func TestLimitToFirst(t *testing.T) {
	t.Parallel()
	var (
		server = newTestServer("")
		fb     = New(server.URL, nil)
	)
	defer server.Close()

	fb.LimitToFirst(2).Value("")
	require.Len(t, server.receivedReqs, 1)

	req := server.receivedReqs[0]
	assert.Equal(t, limitToFirstParam+"=2", req.URL.Query().Encode())
}

func TestLimitToLast(t *testing.T) {
	t.Parallel()
	var (
		server = newTestServer("")
		fb     = New(server.URL, nil)
	)
	defer server.Close()

	fb.LimitToLast(2).Value("")
	require.Len(t, server.receivedReqs, 1)

	req := server.receivedReqs[0]
	assert.Equal(t, limitToLastParam+"=2", req.URL.Query().Encode())
}

func TestStartAt(t *testing.T) {
	t.Parallel()
	var (
		server = newTestServer("")
		fb     = New(server.URL, nil)
	)
	defer server.Close()

	fb.StartAt("3").Value("")
	fb.StartAt("foo").Value("")
	require.Len(t, server.receivedReqs, 2)

	req := server.receivedReqs[0]
	assert.Equal(t, startAtParam+"=3", req.URL.Query().Encode())

	req = server.receivedReqs[1]
	assert.Equal(t, startAtParam+"=%22foo%22", req.URL.Query().Encode())
}

func TestStartAtValue(t *testing.T) {
	t.Parallel()
	var (
		server = newTestServer("")
		fb     = New(server.URL, nil)
	)
	defer server.Close()

	fb.StartAtValue(3).Value("")
	fb.StartAtValue("3").Value("")
	fb.StartAtValue(3.14).Value("")
	fb.StartAtValue("foo").Value("")
	require.Len(t, server.receivedReqs, 4)

	req := server.receivedReqs[0]
	assert.Equal(t, startAtParam+"=3", req.URL.Query().Encode())

	req = server.receivedReqs[1]
	assert.Equal(t, startAtParam+"=%223%22", req.URL.Query().Encode())

	req = server.receivedReqs[2]
	assert.Equal(t, startAtParam+"=3.14", req.URL.Query().Encode())

	req = server.receivedReqs[3]
	assert.Equal(t, startAtParam+"=%22foo%22", req.URL.Query().Encode())
}

func TestEndAt(t *testing.T) {
	t.Parallel()
	var (
		server = newTestServer("")
		fb     = New(server.URL, nil)
	)
	defer server.Close()

	fb.EndAt("4").Value("")
	fb.EndAt("theend").Value("")
	require.Len(t, server.receivedReqs, 2)

	req := server.receivedReqs[0]
	assert.Equal(t, endAtParam+"=4", req.URL.Query().Encode())

	req = server.receivedReqs[1]
	assert.Equal(t, endAtParam+"=%22theend%22", req.URL.Query().Encode())
}

func TestEndAtValue(t *testing.T) {
	t.Parallel()
	var (
		server = newTestServer("")
		fb     = New(server.URL, nil)
	)
	defer server.Close()

	fb.EndAtValue(4).Value("")
	fb.EndAtValue(3.14).Value("")
	fb.EndAtValue("4").Value("")
	fb.EndAtValue("theend").Value("")
	require.Len(t, server.receivedReqs, 4)

	req := server.receivedReqs[0]
	assert.Equal(t, endAtParam+"=4", req.URL.Query().Encode())

	req = server.receivedReqs[1]
	assert.Equal(t, endAtParam+"=3.14", req.URL.Query().Encode())

	req = server.receivedReqs[2]
	assert.Equal(t, endAtParam+"=%224%22", req.URL.Query().Encode())

	req = server.receivedReqs[3]
	assert.Equal(t, endAtParam+"=%22theend%22", req.URL.Query().Encode())
}

func TestIncludePriority(t *testing.T) {
	t.Parallel()
	var (
		server = newTestServer("")
		fb     = New(server.URL, nil)
	)
	defer server.Close()

	fb.IncludePriority(true)
	fb.Value("")
	require.Len(t, server.receivedReqs, 1)

	req := server.receivedReqs[0]
	assert.Equal(t, formatParam+"="+formatVal, req.URL.Query().Encode())

	fb.IncludePriority(false)
	fb.Value("")
	require.Len(t, server.receivedReqs, 2)

	req = server.receivedReqs[1]
	assert.Equal(t, "", req.URL.Query().Encode())
}

func TestQueryMultipleParams(t *testing.T) {
	t.Parallel()
	var (
		server = newTestServer("")
		fb     = New(server.URL, nil)
	)
	defer server.Close()

	fb.OrderBy("user_id").StartAt("7").Value("")
	require.Len(t, server.receivedReqs, 1)

	req := server.receivedReqs[0]
	assert.Equal(t, orderByParam+"=%22user_id%22&startAt=7", req.URL.Query().Encode())
}

func TestEscapeString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		value    string
		expected string
	}{
		{"foo", `"foo"`},
		{"2", `2`},
		{"false", `false`},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, escapeString(testCase.value))
	}
}

func TestEscapeParameter(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		value    interface{}
		expected string
	}{
		{"foo", `"foo"`},
		{2, `2`},
		{"3", `"3"`},
		{true, `true`},
		{"false", `"false"`},
		{3.14, `3.14`},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, escapeParameter(testCase.value))
	}
}

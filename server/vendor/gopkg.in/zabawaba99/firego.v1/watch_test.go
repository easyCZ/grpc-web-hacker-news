package firego

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zabawaba99/firego/internal/firetest"
)

func setupLargeResult() string {
	return "start" + strings.Repeat("0", 64*1024) + "end"
}

func TestWatch(t *testing.T) {
	server := firetest.New()
	server.Start()
	defer server.Close()

	fb := New(server.URL, nil)

	notifications := make(chan Event)
	err := fb.Watch(notifications)
	assert.NoError(t, err)

	l := setupLargeResult()
	server.Set("/foo", l)

	select {
	case event, ok := <-notifications:
		assert.True(t, ok)
		assert.Equal(t, "put", event.Type)
		assert.Equal(t, "/", event.Path)
		assert.Nil(t, event.Data)
	case <-time.After(250 * time.Millisecond):
		require.FailNow(t, "did not receive a notification initial notification")
	}

	select {
	case event, ok := <-notifications:
		assert.True(t, ok)
		assert.Equal(t, "/foo", event.Path)
		assert.EqualValues(t, l, event.Data)
		var v string
		assert.NoError(t, event.Value(&v))
		assert.EqualValues(t, l, v)
	case <-time.After(250 * time.Millisecond):
		require.FailNow(t, "did not receive a notification")
	}
}

func TestWatchRedirectPreservesHeader(t *testing.T) {
	t.Parallel()

	redirectServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, []string{"text/event-stream"}, req.Header["Accept"])
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
	}))
	defer redirectServer.Close()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Location", redirectServer.URL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	}))
	defer server.Close()

	fb := New(server.URL, nil)
	notifications := make(chan Event)

	err := fb.Watch(notifications)
	assert.NoError(t, err)
}

func TestWatchHeartbeatTimeout(t *testing.T) {
	var fb *Firebase

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		flusher, ok := w.(http.Flusher)
		require.True(t, ok, "streaming unsupported")

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		flusher.Flush()
		time.Sleep(2 * fb.watchHeartbeat)
	}))
	defer server.Close()

	notifications := make(chan Event)
	fb = New(server.URL, nil)
	fb.watchHeartbeat = 50 * time.Millisecond

	if err := fb.Watch(notifications); err != nil {
		t.Fatal(err)
	}

	event, ok := <-notifications
	require.True(t, ok, "notifications closed")
	assert.Equal(t, EventTypeError, event.Type, "event type doesn't match")
	assert.Empty(t, event.Path, "event path is not empty")
	assert.NotNil(t, event.Data, "event data is nil")
	assert.Implements(t, new(error), event.Data)
	t.Logf("%#v\n", event)

	_, ok = <-notifications
	require.False(t, ok, "notifications still open")
}

func TestWatchError(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		flusher, ok := w.(http.Flusher)
		require.True(t, ok, "streaming unsupported")

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		flusher.Flush()
	}))

	var (
		notifications = make(chan Event)
		fb            = New(server.URL, nil)
	)
	defer server.Close()

	if err := fb.Watch(notifications); err != nil {
		t.Fatal(err)
	}

	go server.Close()
	event, ok := <-notifications
	require.True(t, ok, "notifications closed")
	assert.Equal(t, EventTypeError, event.Type, "event type doesn't match")
	assert.Empty(t, event.Path, "event path is not empty")
	assert.NotNil(t, event.Data, "event data is nil")
	assert.Implements(t, new(error), event.Data)
}

func TestWatchAuthRevoked(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		flusher, ok := w.(http.Flusher)
		require.True(t, ok, "streaming unsupported")

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		flusher.Flush()
		fmt.Fprintf(w, "event: %s\ndata: %q\n\n", EventTypeAuthRevoked, "token expired")
	}))
	defer server.Close()

	var (
		notifications = make(chan Event)
		fb            = New(server.URL, nil)
	)

	if err := fb.Watch(notifications); err != nil {
		t.Fatal(err)
	}

	event, ok := <-notifications
	require.True(t, ok, "notifications closed")
	assert.Equal(t, EventTypeAuthRevoked, event.Type, "event type doesn't match")
	assert.Empty(t, event.Path, "event path is not empty")
	assert.Equal(t, event.Data, `"token expired"`, "event data does not match")
}

func TestWatch_Issue66(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		flusher, ok := w.(http.Flusher)
		require.True(t, ok, "streaming unsupported")

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		flusher.Flush()
	}))
	defer server.Close()

	// create an initial sse connection
	fb := New(server.URL, nil)
	notifications := make(chan Event)
	err := fb.Watch(notifications)
	require.NoError(t, err)

	// read the first connection
	<-notifications

	// close the server connection and read the error event
	go server.Close()
	<-notifications

	// call stop watching - everything should be a-ok
	fb.StopWatching()

	_, ok := <-notifications
	assert.False(t, ok)
}

func TestStopWatch(t *testing.T) {
	t.Parallel()

	server := firetest.New()
	server.Start()
	defer server.Close()

	fb := New(server.URL, nil)

	notifications := make(chan Event)
	err := fb.Watch(notifications)
	assert.NoError(t, err)

	<-notifications // get initial notification
	fb.StopWatching()
	_, ok := <-notifications
	assert.False(t, ok, "notifications should be closed")
}

package debugger_test

import (
	"context"
	"errors"
	"net"
	"regexp"
	"sync/atomic"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	config "go.octolab.org/toolkit/config/http"

	. "go.octolab.org/toolkit/cli/debugger"
)

func TestNew(t *testing.T) {
	t.Run("success initialization", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		listener, server := NewMockListener(ctrl), NewMockServer(ctrl)

		debugger, err := New(WithCustomListenerAndServer(listener, server))
		assert.NoError(t, err)
		assert.NotNil(t, debugger)
	})
	t.Run("with configuration error", func(t *testing.T) {
		debugger, err := New(WithBuiltinServer(config.Server{Address: "invalid:host"}))
		require.Error(t, err)
		assert.Contains(t, err.Error(), "debugger: configure by 1 option")
		assert.Nil(t, debugger)
	})
	t.Run("with panic", func(t *testing.T) {
		assert.Panics(t, func() { _ = Must() })
	})
}

func TestDebugger(t *testing.T) {
	var count uint32

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	shutdown := func() {
		atomic.AddUint32(&count, 1)
		cancel()
	}
	debugger := Must(WithSpecificHost("127.0.0.1:"))

	addr, success := debugger.Debug(func(err error) { require.NoError(t, err) }, shutdown)
	assert.True(t, success)
	assert.Regexp(t, regexp.MustCompile(`127.0.0.1:\d+`), addr)

	err := debugger.Stop(ctx)
	if errors.Is(err, context.Canceled) {
		t.Skip("not stable test case")
	}
	assert.NoError(t, err)
	<-ctx.Done()
	assert.True(t, atomic.LoadUint32(&count) == 1)

	t.Run("fail to debug", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		listener, server := NewMockListener(ctrl), NewMockServer(ctrl)
		listener.EXPECT().Addr().Return(&net.TCPAddr{IP: net.IP{127, 0, 0, 1}, Port: 1234})
		server.EXPECT().Serve(listener).Return(errors.New("fail to serve"))

		ctx, cancel := context.WithCancel(context.Background())
		debugger := Must(WithCustomListenerAndServer(listener, server))
		debugger.Debug(func(err error) { cancel() })
		<-ctx.Done()
	})
}

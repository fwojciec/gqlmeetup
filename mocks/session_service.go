// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/fwojciec/gqlmeetup"
	"net/http"
	"sync"
)

var (
	lockSessionServiceMockGetUser    sync.RWMutex
	lockSessionServiceMockLogin      sync.RWMutex
	lockSessionServiceMockLogout     sync.RWMutex
	lockSessionServiceMockMiddleware sync.RWMutex
)

// Ensure, that SessionServiceMock does implement gqlmeetup.SessionService.
// If this is not the case, regenerate this file with moq.
var _ gqlmeetup.SessionService = &SessionServiceMock{}

// SessionServiceMock is a mock implementation of gqlmeetup.SessionService.
//
//     func TestSomethingThatUsesSessionService(t *testing.T) {
//
//         // make and configure a mocked gqlmeetup.SessionService
//         mockedSessionService := &SessionServiceMock{
//             GetUserFunc: func(ctx context.Context) *gqlmeetup.User {
// 	               panic("mock out the GetUser method")
//             },
//             LoginFunc: func(ctx context.Context, user *gqlmeetup.User) error {
// 	               panic("mock out the Login method")
//             },
//             LogoutFunc: func(ctx context.Context) error {
// 	               panic("mock out the Logout method")
//             },
//             MiddlewareFunc: func(in1 http.Handler) http.Handler {
// 	               panic("mock out the Middleware method")
//             },
//         }
//
//         // use mockedSessionService in code that requires gqlmeetup.SessionService
//         // and then make assertions.
//
//     }
type SessionServiceMock struct {
	// GetUserFunc mocks the GetUser method.
	GetUserFunc func(ctx context.Context) *gqlmeetup.User

	// LoginFunc mocks the Login method.
	LoginFunc func(ctx context.Context, user *gqlmeetup.User) error

	// LogoutFunc mocks the Logout method.
	LogoutFunc func(ctx context.Context) error

	// MiddlewareFunc mocks the Middleware method.
	MiddlewareFunc func(in1 http.Handler) http.Handler

	// calls tracks calls to the methods.
	calls struct {
		// GetUser holds details about calls to the GetUser method.
		GetUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Login holds details about calls to the Login method.
		Login []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// User is the user argument value.
			User *gqlmeetup.User
		}
		// Logout holds details about calls to the Logout method.
		Logout []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Middleware holds details about calls to the Middleware method.
		Middleware []struct {
			// In1 is the in1 argument value.
			In1 http.Handler
		}
	}
}

// GetUser calls GetUserFunc.
func (mock *SessionServiceMock) GetUser(ctx context.Context) *gqlmeetup.User {
	if mock.GetUserFunc == nil {
		panic("SessionServiceMock.GetUserFunc: method is nil but SessionService.GetUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	lockSessionServiceMockGetUser.Lock()
	mock.calls.GetUser = append(mock.calls.GetUser, callInfo)
	lockSessionServiceMockGetUser.Unlock()
	return mock.GetUserFunc(ctx)
}

// GetUserCalls gets all the calls that were made to GetUser.
// Check the length with:
//     len(mockedSessionService.GetUserCalls())
func (mock *SessionServiceMock) GetUserCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	lockSessionServiceMockGetUser.RLock()
	calls = mock.calls.GetUser
	lockSessionServiceMockGetUser.RUnlock()
	return calls
}

// Login calls LoginFunc.
func (mock *SessionServiceMock) Login(ctx context.Context, user *gqlmeetup.User) error {
	if mock.LoginFunc == nil {
		panic("SessionServiceMock.LoginFunc: method is nil but SessionService.Login was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		User *gqlmeetup.User
	}{
		Ctx:  ctx,
		User: user,
	}
	lockSessionServiceMockLogin.Lock()
	mock.calls.Login = append(mock.calls.Login, callInfo)
	lockSessionServiceMockLogin.Unlock()
	return mock.LoginFunc(ctx, user)
}

// LoginCalls gets all the calls that were made to Login.
// Check the length with:
//     len(mockedSessionService.LoginCalls())
func (mock *SessionServiceMock) LoginCalls() []struct {
	Ctx  context.Context
	User *gqlmeetup.User
} {
	var calls []struct {
		Ctx  context.Context
		User *gqlmeetup.User
	}
	lockSessionServiceMockLogin.RLock()
	calls = mock.calls.Login
	lockSessionServiceMockLogin.RUnlock()
	return calls
}

// Logout calls LogoutFunc.
func (mock *SessionServiceMock) Logout(ctx context.Context) error {
	if mock.LogoutFunc == nil {
		panic("SessionServiceMock.LogoutFunc: method is nil but SessionService.Logout was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	lockSessionServiceMockLogout.Lock()
	mock.calls.Logout = append(mock.calls.Logout, callInfo)
	lockSessionServiceMockLogout.Unlock()
	return mock.LogoutFunc(ctx)
}

// LogoutCalls gets all the calls that were made to Logout.
// Check the length with:
//     len(mockedSessionService.LogoutCalls())
func (mock *SessionServiceMock) LogoutCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	lockSessionServiceMockLogout.RLock()
	calls = mock.calls.Logout
	lockSessionServiceMockLogout.RUnlock()
	return calls
}

// Middleware calls MiddlewareFunc.
func (mock *SessionServiceMock) Middleware(in1 http.Handler) http.Handler {
	if mock.MiddlewareFunc == nil {
		panic("SessionServiceMock.MiddlewareFunc: method is nil but SessionService.Middleware was just called")
	}
	callInfo := struct {
		In1 http.Handler
	}{
		In1: in1,
	}
	lockSessionServiceMockMiddleware.Lock()
	mock.calls.Middleware = append(mock.calls.Middleware, callInfo)
	lockSessionServiceMockMiddleware.Unlock()
	return mock.MiddlewareFunc(in1)
}

// MiddlewareCalls gets all the calls that were made to Middleware.
// Check the length with:
//     len(mockedSessionService.MiddlewareCalls())
func (mock *SessionServiceMock) MiddlewareCalls() []struct {
	In1 http.Handler
} {
	var calls []struct {
		In1 http.Handler
	}
	lockSessionServiceMockMiddleware.RLock()
	calls = mock.calls.Middleware
	lockSessionServiceMockMiddleware.RUnlock()
	return calls
}

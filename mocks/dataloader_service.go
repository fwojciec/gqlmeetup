// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/fwojciec/gqlmeetup"
	"sync"
)

var (
	lockDataLoaderServiceMockAgentGetByID        sync.RWMutex
	lockDataLoaderServiceMockAuthorListByAgentID sync.RWMutex
	lockDataLoaderServiceMockAuthorListByBookID  sync.RWMutex
	lockDataLoaderServiceMockBookListByAuthorID  sync.RWMutex
	lockDataLoaderServiceMockInitialize          sync.RWMutex
)

// Ensure, that DataLoaderServiceMock does implement gqlmeetup.DataLoaderService.
// If this is not the case, regenerate this file with moq.
var _ gqlmeetup.DataLoaderService = &DataLoaderServiceMock{}

// DataLoaderServiceMock is a mock implementation of gqlmeetup.DataLoaderService.
//
//     func TestSomethingThatUsesDataLoaderService(t *testing.T) {
//
//         // make and configure a mocked gqlmeetup.DataLoaderService
//         mockedDataLoaderService := &DataLoaderServiceMock{
//             AgentGetByIDFunc: func(ctx context.Context, id int64) (*gqlmeetup.Agent, error) {
// 	               panic("mock out the AgentGetByID method")
//             },
//             AuthorListByAgentIDFunc: func(ctx context.Context, agentID int64) ([]*gqlmeetup.Author, error) {
// 	               panic("mock out the AuthorListByAgentID method")
//             },
//             AuthorListByBookIDFunc: func(ctx context.Context, bookID int64) ([]*gqlmeetup.Author, error) {
// 	               panic("mock out the AuthorListByBookID method")
//             },
//             BookListByAuthorIDFunc: func(ctx context.Context, authorID int64) ([]*gqlmeetup.Book, error) {
// 	               panic("mock out the BookListByAuthorID method")
//             },
//             InitializeFunc: func(ctx context.Context) context.Context {
// 	               panic("mock out the Initialize method")
//             },
//         }
//
//         // use mockedDataLoaderService in code that requires gqlmeetup.DataLoaderService
//         // and then make assertions.
//
//     }
type DataLoaderServiceMock struct {
	// AgentGetByIDFunc mocks the AgentGetByID method.
	AgentGetByIDFunc func(ctx context.Context, id int64) (*gqlmeetup.Agent, error)

	// AuthorListByAgentIDFunc mocks the AuthorListByAgentID method.
	AuthorListByAgentIDFunc func(ctx context.Context, agentID int64) ([]*gqlmeetup.Author, error)

	// AuthorListByBookIDFunc mocks the AuthorListByBookID method.
	AuthorListByBookIDFunc func(ctx context.Context, bookID int64) ([]*gqlmeetup.Author, error)

	// BookListByAuthorIDFunc mocks the BookListByAuthorID method.
	BookListByAuthorIDFunc func(ctx context.Context, authorID int64) ([]*gqlmeetup.Book, error)

	// InitializeFunc mocks the Initialize method.
	InitializeFunc func(ctx context.Context) context.Context

	// calls tracks calls to the methods.
	calls struct {
		// AgentGetByID holds details about calls to the AgentGetByID method.
		AgentGetByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// AuthorListByAgentID holds details about calls to the AuthorListByAgentID method.
		AuthorListByAgentID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// AgentID is the agentID argument value.
			AgentID int64
		}
		// AuthorListByBookID holds details about calls to the AuthorListByBookID method.
		AuthorListByBookID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// BookID is the bookID argument value.
			BookID int64
		}
		// BookListByAuthorID holds details about calls to the BookListByAuthorID method.
		BookListByAuthorID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// AuthorID is the authorID argument value.
			AuthorID int64
		}
		// Initialize holds details about calls to the Initialize method.
		Initialize []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
}

// AgentGetByID calls AgentGetByIDFunc.
func (mock *DataLoaderServiceMock) AgentGetByID(ctx context.Context, id int64) (*gqlmeetup.Agent, error) {
	if mock.AgentGetByIDFunc == nil {
		panic("DataLoaderServiceMock.AgentGetByIDFunc: method is nil but DataLoaderService.AgentGetByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	lockDataLoaderServiceMockAgentGetByID.Lock()
	mock.calls.AgentGetByID = append(mock.calls.AgentGetByID, callInfo)
	lockDataLoaderServiceMockAgentGetByID.Unlock()
	return mock.AgentGetByIDFunc(ctx, id)
}

// AgentGetByIDCalls gets all the calls that were made to AgentGetByID.
// Check the length with:
//     len(mockedDataLoaderService.AgentGetByIDCalls())
func (mock *DataLoaderServiceMock) AgentGetByIDCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	lockDataLoaderServiceMockAgentGetByID.RLock()
	calls = mock.calls.AgentGetByID
	lockDataLoaderServiceMockAgentGetByID.RUnlock()
	return calls
}

// AuthorListByAgentID calls AuthorListByAgentIDFunc.
func (mock *DataLoaderServiceMock) AuthorListByAgentID(ctx context.Context, agentID int64) ([]*gqlmeetup.Author, error) {
	if mock.AuthorListByAgentIDFunc == nil {
		panic("DataLoaderServiceMock.AuthorListByAgentIDFunc: method is nil but DataLoaderService.AuthorListByAgentID was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		AgentID int64
	}{
		Ctx:     ctx,
		AgentID: agentID,
	}
	lockDataLoaderServiceMockAuthorListByAgentID.Lock()
	mock.calls.AuthorListByAgentID = append(mock.calls.AuthorListByAgentID, callInfo)
	lockDataLoaderServiceMockAuthorListByAgentID.Unlock()
	return mock.AuthorListByAgentIDFunc(ctx, agentID)
}

// AuthorListByAgentIDCalls gets all the calls that were made to AuthorListByAgentID.
// Check the length with:
//     len(mockedDataLoaderService.AuthorListByAgentIDCalls())
func (mock *DataLoaderServiceMock) AuthorListByAgentIDCalls() []struct {
	Ctx     context.Context
	AgentID int64
} {
	var calls []struct {
		Ctx     context.Context
		AgentID int64
	}
	lockDataLoaderServiceMockAuthorListByAgentID.RLock()
	calls = mock.calls.AuthorListByAgentID
	lockDataLoaderServiceMockAuthorListByAgentID.RUnlock()
	return calls
}

// AuthorListByBookID calls AuthorListByBookIDFunc.
func (mock *DataLoaderServiceMock) AuthorListByBookID(ctx context.Context, bookID int64) ([]*gqlmeetup.Author, error) {
	if mock.AuthorListByBookIDFunc == nil {
		panic("DataLoaderServiceMock.AuthorListByBookIDFunc: method is nil but DataLoaderService.AuthorListByBookID was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		BookID int64
	}{
		Ctx:    ctx,
		BookID: bookID,
	}
	lockDataLoaderServiceMockAuthorListByBookID.Lock()
	mock.calls.AuthorListByBookID = append(mock.calls.AuthorListByBookID, callInfo)
	lockDataLoaderServiceMockAuthorListByBookID.Unlock()
	return mock.AuthorListByBookIDFunc(ctx, bookID)
}

// AuthorListByBookIDCalls gets all the calls that were made to AuthorListByBookID.
// Check the length with:
//     len(mockedDataLoaderService.AuthorListByBookIDCalls())
func (mock *DataLoaderServiceMock) AuthorListByBookIDCalls() []struct {
	Ctx    context.Context
	BookID int64
} {
	var calls []struct {
		Ctx    context.Context
		BookID int64
	}
	lockDataLoaderServiceMockAuthorListByBookID.RLock()
	calls = mock.calls.AuthorListByBookID
	lockDataLoaderServiceMockAuthorListByBookID.RUnlock()
	return calls
}

// BookListByAuthorID calls BookListByAuthorIDFunc.
func (mock *DataLoaderServiceMock) BookListByAuthorID(ctx context.Context, authorID int64) ([]*gqlmeetup.Book, error) {
	if mock.BookListByAuthorIDFunc == nil {
		panic("DataLoaderServiceMock.BookListByAuthorIDFunc: method is nil but DataLoaderService.BookListByAuthorID was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		AuthorID int64
	}{
		Ctx:      ctx,
		AuthorID: authorID,
	}
	lockDataLoaderServiceMockBookListByAuthorID.Lock()
	mock.calls.BookListByAuthorID = append(mock.calls.BookListByAuthorID, callInfo)
	lockDataLoaderServiceMockBookListByAuthorID.Unlock()
	return mock.BookListByAuthorIDFunc(ctx, authorID)
}

// BookListByAuthorIDCalls gets all the calls that were made to BookListByAuthorID.
// Check the length with:
//     len(mockedDataLoaderService.BookListByAuthorIDCalls())
func (mock *DataLoaderServiceMock) BookListByAuthorIDCalls() []struct {
	Ctx      context.Context
	AuthorID int64
} {
	var calls []struct {
		Ctx      context.Context
		AuthorID int64
	}
	lockDataLoaderServiceMockBookListByAuthorID.RLock()
	calls = mock.calls.BookListByAuthorID
	lockDataLoaderServiceMockBookListByAuthorID.RUnlock()
	return calls
}

// Initialize calls InitializeFunc.
func (mock *DataLoaderServiceMock) Initialize(ctx context.Context) context.Context {
	if mock.InitializeFunc == nil {
		panic("DataLoaderServiceMock.InitializeFunc: method is nil but DataLoaderService.Initialize was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	lockDataLoaderServiceMockInitialize.Lock()
	mock.calls.Initialize = append(mock.calls.Initialize, callInfo)
	lockDataLoaderServiceMockInitialize.Unlock()
	return mock.InitializeFunc(ctx)
}

// InitializeCalls gets all the calls that were made to Initialize.
// Check the length with:
//     len(mockedDataLoaderService.InitializeCalls())
func (mock *DataLoaderServiceMock) InitializeCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	lockDataLoaderServiceMockInitialize.RLock()
	calls = mock.calls.Initialize
	lockDataLoaderServiceMockInitialize.RUnlock()
	return calls
}

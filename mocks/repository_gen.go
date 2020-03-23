// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/fwojciec/gqlmeetup"
	"sync"
)

var (
	lockRepositoryMockAgentCreate          sync.RWMutex
	lockRepositoryMockAgentDelete          sync.RWMutex
	lockRepositoryMockAgentGetByID         sync.RWMutex
	lockRepositoryMockAgentList            sync.RWMutex
	lockRepositoryMockAgentListByIDs       sync.RWMutex
	lockRepositoryMockAgentUpdate          sync.RWMutex
	lockRepositoryMockAuthorCreate         sync.RWMutex
	lockRepositoryMockAuthorDelete         sync.RWMutex
	lockRepositoryMockAuthorGetByID        sync.RWMutex
	lockRepositoryMockAuthorList           sync.RWMutex
	lockRepositoryMockAuthorListByAgentID  sync.RWMutex
	lockRepositoryMockAuthorListByAgentIDs sync.RWMutex
	lockRepositoryMockAuthorListByBookIDs  sync.RWMutex
	lockRepositoryMockAuthorUpdate         sync.RWMutex
	lockRepositoryMockBookCreate           sync.RWMutex
	lockRepositoryMockBookDelete           sync.RWMutex
	lockRepositoryMockBookGetByID          sync.RWMutex
	lockRepositoryMockBookList             sync.RWMutex
	lockRepositoryMockBookListByAuthorIDs  sync.RWMutex
	lockRepositoryMockBookUpdate           sync.RWMutex
	lockRepositoryMockUserCreate           sync.RWMutex
	lockRepositoryMockUserGetByEmail       sync.RWMutex
)

// Ensure, that RepositoryMock does implement gqlmeetup.Repository.
// If this is not the case, regenerate this file with moq.
var _ gqlmeetup.Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of gqlmeetup.Repository.
//
//     func TestSomethingThatUsesRepository(t *testing.T) {
//
//         // make and configure a mocked gqlmeetup.Repository
//         mockedRepository := &RepositoryMock{
//             AgentCreateFunc: func(ctx context.Context, data gqlmeetup.Agent) (*gqlmeetup.Agent, error) {
// 	               panic("mock out the AgentCreate method")
//             },
//             AgentDeleteFunc: func(ctx context.Context, id int64) (*gqlmeetup.Agent, error) {
// 	               panic("mock out the AgentDelete method")
//             },
//             AgentGetByIDFunc: func(ctx context.Context, id int64) (*gqlmeetup.Agent, error) {
// 	               panic("mock out the AgentGetByID method")
//             },
//             AgentListFunc: func(ctx context.Context) ([]*gqlmeetup.Agent, error) {
// 	               panic("mock out the AgentList method")
//             },
//             AgentListByIDsFunc: func(ctx context.Context, ids []int64) ([]*gqlmeetup.Agent, error) {
// 	               panic("mock out the AgentListByIDs method")
//             },
//             AgentUpdateFunc: func(ctx context.Context, id int64, data gqlmeetup.Agent) (*gqlmeetup.Agent, error) {
// 	               panic("mock out the AgentUpdate method")
//             },
//             AuthorCreateFunc: func(ctx context.Context, data gqlmeetup.Author) (*gqlmeetup.Author, error) {
// 	               panic("mock out the AuthorCreate method")
//             },
//             AuthorDeleteFunc: func(ctx context.Context, id int64) (*gqlmeetup.Author, error) {
// 	               panic("mock out the AuthorDelete method")
//             },
//             AuthorGetByIDFunc: func(ctx context.Context, id int64) (*gqlmeetup.Author, error) {
// 	               panic("mock out the AuthorGetByID method")
//             },
//             AuthorListFunc: func(ctx context.Context) ([]*gqlmeetup.Author, error) {
// 	               panic("mock out the AuthorList method")
//             },
//             AuthorListByAgentIDFunc: func(ctx context.Context, agentIDs int64) ([]*gqlmeetup.Author, error) {
// 	               panic("mock out the AuthorListByAgentID method")
//             },
//             AuthorListByAgentIDsFunc: func(ctx context.Context, agentIDs []int64) ([]*gqlmeetup.Author, error) {
// 	               panic("mock out the AuthorListByAgentIDs method")
//             },
//             AuthorListByBookIDsFunc: func(ctx context.Context, bookIDs []int64) ([]*gqlmeetup.Author, error) {
// 	               panic("mock out the AuthorListByBookIDs method")
//             },
//             AuthorUpdateFunc: func(ctx context.Context, id int64, data gqlmeetup.Author) (*gqlmeetup.Author, error) {
// 	               panic("mock out the AuthorUpdate method")
//             },
//             BookCreateFunc: func(ctx context.Context, data gqlmeetup.Book, authorIDs []int64) (*gqlmeetup.Book, error) {
// 	               panic("mock out the BookCreate method")
//             },
//             BookDeleteFunc: func(ctx context.Context, id int64) (*gqlmeetup.Book, error) {
// 	               panic("mock out the BookDelete method")
//             },
//             BookGetByIDFunc: func(ctx context.Context, id int64) (*gqlmeetup.Book, error) {
// 	               panic("mock out the BookGetByID method")
//             },
//             BookListFunc: func(ctx context.Context, limit *int, offset *int) ([]*gqlmeetup.Book, error) {
// 	               panic("mock out the BookList method")
//             },
//             BookListByAuthorIDsFunc: func(ctx context.Context, authorIDs []int64) ([]*gqlmeetup.Book, error) {
// 	               panic("mock out the BookListByAuthorIDs method")
//             },
//             BookUpdateFunc: func(ctx context.Context, id int64, data gqlmeetup.Book, authorIDs []int64) (*gqlmeetup.Book, error) {
// 	               panic("mock out the BookUpdate method")
//             },
//             UserCreateFunc: func(data gqlmeetup.User) error {
// 	               panic("mock out the UserCreate method")
//             },
//             UserGetByEmailFunc: func(ctx context.Context, email string) (*gqlmeetup.User, error) {
// 	               panic("mock out the UserGetByEmail method")
//             },
//         }
//
//         // use mockedRepository in code that requires gqlmeetup.Repository
//         // and then make assertions.
//
//     }
type RepositoryMock struct {
	// AgentCreateFunc mocks the AgentCreate method.
	AgentCreateFunc func(ctx context.Context, data gqlmeetup.Agent) (*gqlmeetup.Agent, error)

	// AgentDeleteFunc mocks the AgentDelete method.
	AgentDeleteFunc func(ctx context.Context, id int64) (*gqlmeetup.Agent, error)

	// AgentGetByIDFunc mocks the AgentGetByID method.
	AgentGetByIDFunc func(ctx context.Context, id int64) (*gqlmeetup.Agent, error)

	// AgentListFunc mocks the AgentList method.
	AgentListFunc func(ctx context.Context) ([]*gqlmeetup.Agent, error)

	// AgentListByIDsFunc mocks the AgentListByIDs method.
	AgentListByIDsFunc func(ctx context.Context, ids []int64) ([]*gqlmeetup.Agent, error)

	// AgentUpdateFunc mocks the AgentUpdate method.
	AgentUpdateFunc func(ctx context.Context, id int64, data gqlmeetup.Agent) (*gqlmeetup.Agent, error)

	// AuthorCreateFunc mocks the AuthorCreate method.
	AuthorCreateFunc func(ctx context.Context, data gqlmeetup.Author) (*gqlmeetup.Author, error)

	// AuthorDeleteFunc mocks the AuthorDelete method.
	AuthorDeleteFunc func(ctx context.Context, id int64) (*gqlmeetup.Author, error)

	// AuthorGetByIDFunc mocks the AuthorGetByID method.
	AuthorGetByIDFunc func(ctx context.Context, id int64) (*gqlmeetup.Author, error)

	// AuthorListFunc mocks the AuthorList method.
	AuthorListFunc func(ctx context.Context) ([]*gqlmeetup.Author, error)

	// AuthorListByAgentIDFunc mocks the AuthorListByAgentID method.
	AuthorListByAgentIDFunc func(ctx context.Context, agentIDs int64) ([]*gqlmeetup.Author, error)

	// AuthorListByAgentIDsFunc mocks the AuthorListByAgentIDs method.
	AuthorListByAgentIDsFunc func(ctx context.Context, agentIDs []int64) ([]*gqlmeetup.Author, error)

	// AuthorListByBookIDsFunc mocks the AuthorListByBookIDs method.
	AuthorListByBookIDsFunc func(ctx context.Context, bookIDs []int64) ([]*gqlmeetup.Author, error)

	// AuthorUpdateFunc mocks the AuthorUpdate method.
	AuthorUpdateFunc func(ctx context.Context, id int64, data gqlmeetup.Author) (*gqlmeetup.Author, error)

	// BookCreateFunc mocks the BookCreate method.
	BookCreateFunc func(ctx context.Context, data gqlmeetup.Book, authorIDs []int64) (*gqlmeetup.Book, error)

	// BookDeleteFunc mocks the BookDelete method.
	BookDeleteFunc func(ctx context.Context, id int64) (*gqlmeetup.Book, error)

	// BookGetByIDFunc mocks the BookGetByID method.
	BookGetByIDFunc func(ctx context.Context, id int64) (*gqlmeetup.Book, error)

	// BookListFunc mocks the BookList method.
	BookListFunc func(ctx context.Context, limit *int, offset *int) ([]*gqlmeetup.Book, error)

	// BookListByAuthorIDsFunc mocks the BookListByAuthorIDs method.
	BookListByAuthorIDsFunc func(ctx context.Context, authorIDs []int64) ([]*gqlmeetup.Book, error)

	// BookUpdateFunc mocks the BookUpdate method.
	BookUpdateFunc func(ctx context.Context, id int64, data gqlmeetup.Book, authorIDs []int64) (*gqlmeetup.Book, error)

	// UserCreateFunc mocks the UserCreate method.
	UserCreateFunc func(data gqlmeetup.User) error

	// UserGetByEmailFunc mocks the UserGetByEmail method.
	UserGetByEmailFunc func(ctx context.Context, email string) (*gqlmeetup.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// AgentCreate holds details about calls to the AgentCreate method.
		AgentCreate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Data is the data argument value.
			Data gqlmeetup.Agent
		}
		// AgentDelete holds details about calls to the AgentDelete method.
		AgentDelete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// AgentGetByID holds details about calls to the AgentGetByID method.
		AgentGetByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// AgentList holds details about calls to the AgentList method.
		AgentList []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// AgentListByIDs holds details about calls to the AgentListByIDs method.
		AgentListByIDs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Ids is the ids argument value.
			Ids []int64
		}
		// AgentUpdate holds details about calls to the AgentUpdate method.
		AgentUpdate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
			// Data is the data argument value.
			Data gqlmeetup.Agent
		}
		// AuthorCreate holds details about calls to the AuthorCreate method.
		AuthorCreate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Data is the data argument value.
			Data gqlmeetup.Author
		}
		// AuthorDelete holds details about calls to the AuthorDelete method.
		AuthorDelete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// AuthorGetByID holds details about calls to the AuthorGetByID method.
		AuthorGetByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// AuthorList holds details about calls to the AuthorList method.
		AuthorList []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// AuthorListByAgentID holds details about calls to the AuthorListByAgentID method.
		AuthorListByAgentID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// AgentIDs is the agentIDs argument value.
			AgentIDs int64
		}
		// AuthorListByAgentIDs holds details about calls to the AuthorListByAgentIDs method.
		AuthorListByAgentIDs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// AgentIDs is the agentIDs argument value.
			AgentIDs []int64
		}
		// AuthorListByBookIDs holds details about calls to the AuthorListByBookIDs method.
		AuthorListByBookIDs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// BookIDs is the bookIDs argument value.
			BookIDs []int64
		}
		// AuthorUpdate holds details about calls to the AuthorUpdate method.
		AuthorUpdate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
			// Data is the data argument value.
			Data gqlmeetup.Author
		}
		// BookCreate holds details about calls to the BookCreate method.
		BookCreate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Data is the data argument value.
			Data gqlmeetup.Book
			// AuthorIDs is the authorIDs argument value.
			AuthorIDs []int64
		}
		// BookDelete holds details about calls to the BookDelete method.
		BookDelete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// BookGetByID holds details about calls to the BookGetByID method.
		BookGetByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// BookList holds details about calls to the BookList method.
		BookList []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Limit is the limit argument value.
			Limit *int
			// Offset is the offset argument value.
			Offset *int
		}
		// BookListByAuthorIDs holds details about calls to the BookListByAuthorIDs method.
		BookListByAuthorIDs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// AuthorIDs is the authorIDs argument value.
			AuthorIDs []int64
		}
		// BookUpdate holds details about calls to the BookUpdate method.
		BookUpdate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
			// Data is the data argument value.
			Data gqlmeetup.Book
			// AuthorIDs is the authorIDs argument value.
			AuthorIDs []int64
		}
		// UserCreate holds details about calls to the UserCreate method.
		UserCreate []struct {
			// Data is the data argument value.
			Data gqlmeetup.User
		}
		// UserGetByEmail holds details about calls to the UserGetByEmail method.
		UserGetByEmail []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Email is the email argument value.
			Email string
		}
	}
}

// AgentCreate calls AgentCreateFunc.
func (mock *RepositoryMock) AgentCreate(ctx context.Context, data gqlmeetup.Agent) (*gqlmeetup.Agent, error) {
	if mock.AgentCreateFunc == nil {
		panic("RepositoryMock.AgentCreateFunc: method is nil but Repository.AgentCreate was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Data gqlmeetup.Agent
	}{
		Ctx:  ctx,
		Data: data,
	}
	lockRepositoryMockAgentCreate.Lock()
	mock.calls.AgentCreate = append(mock.calls.AgentCreate, callInfo)
	lockRepositoryMockAgentCreate.Unlock()
	return mock.AgentCreateFunc(ctx, data)
}

// AgentCreateCalls gets all the calls that were made to AgentCreate.
// Check the length with:
//     len(mockedRepository.AgentCreateCalls())
func (mock *RepositoryMock) AgentCreateCalls() []struct {
	Ctx  context.Context
	Data gqlmeetup.Agent
} {
	var calls []struct {
		Ctx  context.Context
		Data gqlmeetup.Agent
	}
	lockRepositoryMockAgentCreate.RLock()
	calls = mock.calls.AgentCreate
	lockRepositoryMockAgentCreate.RUnlock()
	return calls
}

// AgentDelete calls AgentDeleteFunc.
func (mock *RepositoryMock) AgentDelete(ctx context.Context, id int64) (*gqlmeetup.Agent, error) {
	if mock.AgentDeleteFunc == nil {
		panic("RepositoryMock.AgentDeleteFunc: method is nil but Repository.AgentDelete was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	lockRepositoryMockAgentDelete.Lock()
	mock.calls.AgentDelete = append(mock.calls.AgentDelete, callInfo)
	lockRepositoryMockAgentDelete.Unlock()
	return mock.AgentDeleteFunc(ctx, id)
}

// AgentDeleteCalls gets all the calls that were made to AgentDelete.
// Check the length with:
//     len(mockedRepository.AgentDeleteCalls())
func (mock *RepositoryMock) AgentDeleteCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	lockRepositoryMockAgentDelete.RLock()
	calls = mock.calls.AgentDelete
	lockRepositoryMockAgentDelete.RUnlock()
	return calls
}

// AgentGetByID calls AgentGetByIDFunc.
func (mock *RepositoryMock) AgentGetByID(ctx context.Context, id int64) (*gqlmeetup.Agent, error) {
	if mock.AgentGetByIDFunc == nil {
		panic("RepositoryMock.AgentGetByIDFunc: method is nil but Repository.AgentGetByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	lockRepositoryMockAgentGetByID.Lock()
	mock.calls.AgentGetByID = append(mock.calls.AgentGetByID, callInfo)
	lockRepositoryMockAgentGetByID.Unlock()
	return mock.AgentGetByIDFunc(ctx, id)
}

// AgentGetByIDCalls gets all the calls that were made to AgentGetByID.
// Check the length with:
//     len(mockedRepository.AgentGetByIDCalls())
func (mock *RepositoryMock) AgentGetByIDCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	lockRepositoryMockAgentGetByID.RLock()
	calls = mock.calls.AgentGetByID
	lockRepositoryMockAgentGetByID.RUnlock()
	return calls
}

// AgentList calls AgentListFunc.
func (mock *RepositoryMock) AgentList(ctx context.Context) ([]*gqlmeetup.Agent, error) {
	if mock.AgentListFunc == nil {
		panic("RepositoryMock.AgentListFunc: method is nil but Repository.AgentList was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	lockRepositoryMockAgentList.Lock()
	mock.calls.AgentList = append(mock.calls.AgentList, callInfo)
	lockRepositoryMockAgentList.Unlock()
	return mock.AgentListFunc(ctx)
}

// AgentListCalls gets all the calls that were made to AgentList.
// Check the length with:
//     len(mockedRepository.AgentListCalls())
func (mock *RepositoryMock) AgentListCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	lockRepositoryMockAgentList.RLock()
	calls = mock.calls.AgentList
	lockRepositoryMockAgentList.RUnlock()
	return calls
}

// AgentListByIDs calls AgentListByIDsFunc.
func (mock *RepositoryMock) AgentListByIDs(ctx context.Context, ids []int64) ([]*gqlmeetup.Agent, error) {
	if mock.AgentListByIDsFunc == nil {
		panic("RepositoryMock.AgentListByIDsFunc: method is nil but Repository.AgentListByIDs was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Ids []int64
	}{
		Ctx: ctx,
		Ids: ids,
	}
	lockRepositoryMockAgentListByIDs.Lock()
	mock.calls.AgentListByIDs = append(mock.calls.AgentListByIDs, callInfo)
	lockRepositoryMockAgentListByIDs.Unlock()
	return mock.AgentListByIDsFunc(ctx, ids)
}

// AgentListByIDsCalls gets all the calls that were made to AgentListByIDs.
// Check the length with:
//     len(mockedRepository.AgentListByIDsCalls())
func (mock *RepositoryMock) AgentListByIDsCalls() []struct {
	Ctx context.Context
	Ids []int64
} {
	var calls []struct {
		Ctx context.Context
		Ids []int64
	}
	lockRepositoryMockAgentListByIDs.RLock()
	calls = mock.calls.AgentListByIDs
	lockRepositoryMockAgentListByIDs.RUnlock()
	return calls
}

// AgentUpdate calls AgentUpdateFunc.
func (mock *RepositoryMock) AgentUpdate(ctx context.Context, id int64, data gqlmeetup.Agent) (*gqlmeetup.Agent, error) {
	if mock.AgentUpdateFunc == nil {
		panic("RepositoryMock.AgentUpdateFunc: method is nil but Repository.AgentUpdate was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		ID   int64
		Data gqlmeetup.Agent
	}{
		Ctx:  ctx,
		ID:   id,
		Data: data,
	}
	lockRepositoryMockAgentUpdate.Lock()
	mock.calls.AgentUpdate = append(mock.calls.AgentUpdate, callInfo)
	lockRepositoryMockAgentUpdate.Unlock()
	return mock.AgentUpdateFunc(ctx, id, data)
}

// AgentUpdateCalls gets all the calls that were made to AgentUpdate.
// Check the length with:
//     len(mockedRepository.AgentUpdateCalls())
func (mock *RepositoryMock) AgentUpdateCalls() []struct {
	Ctx  context.Context
	ID   int64
	Data gqlmeetup.Agent
} {
	var calls []struct {
		Ctx  context.Context
		ID   int64
		Data gqlmeetup.Agent
	}
	lockRepositoryMockAgentUpdate.RLock()
	calls = mock.calls.AgentUpdate
	lockRepositoryMockAgentUpdate.RUnlock()
	return calls
}

// AuthorCreate calls AuthorCreateFunc.
func (mock *RepositoryMock) AuthorCreate(ctx context.Context, data gqlmeetup.Author) (*gqlmeetup.Author, error) {
	if mock.AuthorCreateFunc == nil {
		panic("RepositoryMock.AuthorCreateFunc: method is nil but Repository.AuthorCreate was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Data gqlmeetup.Author
	}{
		Ctx:  ctx,
		Data: data,
	}
	lockRepositoryMockAuthorCreate.Lock()
	mock.calls.AuthorCreate = append(mock.calls.AuthorCreate, callInfo)
	lockRepositoryMockAuthorCreate.Unlock()
	return mock.AuthorCreateFunc(ctx, data)
}

// AuthorCreateCalls gets all the calls that were made to AuthorCreate.
// Check the length with:
//     len(mockedRepository.AuthorCreateCalls())
func (mock *RepositoryMock) AuthorCreateCalls() []struct {
	Ctx  context.Context
	Data gqlmeetup.Author
} {
	var calls []struct {
		Ctx  context.Context
		Data gqlmeetup.Author
	}
	lockRepositoryMockAuthorCreate.RLock()
	calls = mock.calls.AuthorCreate
	lockRepositoryMockAuthorCreate.RUnlock()
	return calls
}

// AuthorDelete calls AuthorDeleteFunc.
func (mock *RepositoryMock) AuthorDelete(ctx context.Context, id int64) (*gqlmeetup.Author, error) {
	if mock.AuthorDeleteFunc == nil {
		panic("RepositoryMock.AuthorDeleteFunc: method is nil but Repository.AuthorDelete was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	lockRepositoryMockAuthorDelete.Lock()
	mock.calls.AuthorDelete = append(mock.calls.AuthorDelete, callInfo)
	lockRepositoryMockAuthorDelete.Unlock()
	return mock.AuthorDeleteFunc(ctx, id)
}

// AuthorDeleteCalls gets all the calls that were made to AuthorDelete.
// Check the length with:
//     len(mockedRepository.AuthorDeleteCalls())
func (mock *RepositoryMock) AuthorDeleteCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	lockRepositoryMockAuthorDelete.RLock()
	calls = mock.calls.AuthorDelete
	lockRepositoryMockAuthorDelete.RUnlock()
	return calls
}

// AuthorGetByID calls AuthorGetByIDFunc.
func (mock *RepositoryMock) AuthorGetByID(ctx context.Context, id int64) (*gqlmeetup.Author, error) {
	if mock.AuthorGetByIDFunc == nil {
		panic("RepositoryMock.AuthorGetByIDFunc: method is nil but Repository.AuthorGetByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	lockRepositoryMockAuthorGetByID.Lock()
	mock.calls.AuthorGetByID = append(mock.calls.AuthorGetByID, callInfo)
	lockRepositoryMockAuthorGetByID.Unlock()
	return mock.AuthorGetByIDFunc(ctx, id)
}

// AuthorGetByIDCalls gets all the calls that were made to AuthorGetByID.
// Check the length with:
//     len(mockedRepository.AuthorGetByIDCalls())
func (mock *RepositoryMock) AuthorGetByIDCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	lockRepositoryMockAuthorGetByID.RLock()
	calls = mock.calls.AuthorGetByID
	lockRepositoryMockAuthorGetByID.RUnlock()
	return calls
}

// AuthorList calls AuthorListFunc.
func (mock *RepositoryMock) AuthorList(ctx context.Context) ([]*gqlmeetup.Author, error) {
	if mock.AuthorListFunc == nil {
		panic("RepositoryMock.AuthorListFunc: method is nil but Repository.AuthorList was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	lockRepositoryMockAuthorList.Lock()
	mock.calls.AuthorList = append(mock.calls.AuthorList, callInfo)
	lockRepositoryMockAuthorList.Unlock()
	return mock.AuthorListFunc(ctx)
}

// AuthorListCalls gets all the calls that were made to AuthorList.
// Check the length with:
//     len(mockedRepository.AuthorListCalls())
func (mock *RepositoryMock) AuthorListCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	lockRepositoryMockAuthorList.RLock()
	calls = mock.calls.AuthorList
	lockRepositoryMockAuthorList.RUnlock()
	return calls
}

// AuthorListByAgentID calls AuthorListByAgentIDFunc.
func (mock *RepositoryMock) AuthorListByAgentID(ctx context.Context, agentIDs int64) ([]*gqlmeetup.Author, error) {
	if mock.AuthorListByAgentIDFunc == nil {
		panic("RepositoryMock.AuthorListByAgentIDFunc: method is nil but Repository.AuthorListByAgentID was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		AgentIDs int64
	}{
		Ctx:      ctx,
		AgentIDs: agentIDs,
	}
	lockRepositoryMockAuthorListByAgentID.Lock()
	mock.calls.AuthorListByAgentID = append(mock.calls.AuthorListByAgentID, callInfo)
	lockRepositoryMockAuthorListByAgentID.Unlock()
	return mock.AuthorListByAgentIDFunc(ctx, agentIDs)
}

// AuthorListByAgentIDCalls gets all the calls that were made to AuthorListByAgentID.
// Check the length with:
//     len(mockedRepository.AuthorListByAgentIDCalls())
func (mock *RepositoryMock) AuthorListByAgentIDCalls() []struct {
	Ctx      context.Context
	AgentIDs int64
} {
	var calls []struct {
		Ctx      context.Context
		AgentIDs int64
	}
	lockRepositoryMockAuthorListByAgentID.RLock()
	calls = mock.calls.AuthorListByAgentID
	lockRepositoryMockAuthorListByAgentID.RUnlock()
	return calls
}

// AuthorListByAgentIDs calls AuthorListByAgentIDsFunc.
func (mock *RepositoryMock) AuthorListByAgentIDs(ctx context.Context, agentIDs []int64) ([]*gqlmeetup.Author, error) {
	if mock.AuthorListByAgentIDsFunc == nil {
		panic("RepositoryMock.AuthorListByAgentIDsFunc: method is nil but Repository.AuthorListByAgentIDs was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		AgentIDs []int64
	}{
		Ctx:      ctx,
		AgentIDs: agentIDs,
	}
	lockRepositoryMockAuthorListByAgentIDs.Lock()
	mock.calls.AuthorListByAgentIDs = append(mock.calls.AuthorListByAgentIDs, callInfo)
	lockRepositoryMockAuthorListByAgentIDs.Unlock()
	return mock.AuthorListByAgentIDsFunc(ctx, agentIDs)
}

// AuthorListByAgentIDsCalls gets all the calls that were made to AuthorListByAgentIDs.
// Check the length with:
//     len(mockedRepository.AuthorListByAgentIDsCalls())
func (mock *RepositoryMock) AuthorListByAgentIDsCalls() []struct {
	Ctx      context.Context
	AgentIDs []int64
} {
	var calls []struct {
		Ctx      context.Context
		AgentIDs []int64
	}
	lockRepositoryMockAuthorListByAgentIDs.RLock()
	calls = mock.calls.AuthorListByAgentIDs
	lockRepositoryMockAuthorListByAgentIDs.RUnlock()
	return calls
}

// AuthorListByBookIDs calls AuthorListByBookIDsFunc.
func (mock *RepositoryMock) AuthorListByBookIDs(ctx context.Context, bookIDs []int64) ([]*gqlmeetup.Author, error) {
	if mock.AuthorListByBookIDsFunc == nil {
		panic("RepositoryMock.AuthorListByBookIDsFunc: method is nil but Repository.AuthorListByBookIDs was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		BookIDs []int64
	}{
		Ctx:     ctx,
		BookIDs: bookIDs,
	}
	lockRepositoryMockAuthorListByBookIDs.Lock()
	mock.calls.AuthorListByBookIDs = append(mock.calls.AuthorListByBookIDs, callInfo)
	lockRepositoryMockAuthorListByBookIDs.Unlock()
	return mock.AuthorListByBookIDsFunc(ctx, bookIDs)
}

// AuthorListByBookIDsCalls gets all the calls that were made to AuthorListByBookIDs.
// Check the length with:
//     len(mockedRepository.AuthorListByBookIDsCalls())
func (mock *RepositoryMock) AuthorListByBookIDsCalls() []struct {
	Ctx     context.Context
	BookIDs []int64
} {
	var calls []struct {
		Ctx     context.Context
		BookIDs []int64
	}
	lockRepositoryMockAuthorListByBookIDs.RLock()
	calls = mock.calls.AuthorListByBookIDs
	lockRepositoryMockAuthorListByBookIDs.RUnlock()
	return calls
}

// AuthorUpdate calls AuthorUpdateFunc.
func (mock *RepositoryMock) AuthorUpdate(ctx context.Context, id int64, data gqlmeetup.Author) (*gqlmeetup.Author, error) {
	if mock.AuthorUpdateFunc == nil {
		panic("RepositoryMock.AuthorUpdateFunc: method is nil but Repository.AuthorUpdate was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		ID   int64
		Data gqlmeetup.Author
	}{
		Ctx:  ctx,
		ID:   id,
		Data: data,
	}
	lockRepositoryMockAuthorUpdate.Lock()
	mock.calls.AuthorUpdate = append(mock.calls.AuthorUpdate, callInfo)
	lockRepositoryMockAuthorUpdate.Unlock()
	return mock.AuthorUpdateFunc(ctx, id, data)
}

// AuthorUpdateCalls gets all the calls that were made to AuthorUpdate.
// Check the length with:
//     len(mockedRepository.AuthorUpdateCalls())
func (mock *RepositoryMock) AuthorUpdateCalls() []struct {
	Ctx  context.Context
	ID   int64
	Data gqlmeetup.Author
} {
	var calls []struct {
		Ctx  context.Context
		ID   int64
		Data gqlmeetup.Author
	}
	lockRepositoryMockAuthorUpdate.RLock()
	calls = mock.calls.AuthorUpdate
	lockRepositoryMockAuthorUpdate.RUnlock()
	return calls
}

// BookCreate calls BookCreateFunc.
func (mock *RepositoryMock) BookCreate(ctx context.Context, data gqlmeetup.Book, authorIDs []int64) (*gqlmeetup.Book, error) {
	if mock.BookCreateFunc == nil {
		panic("RepositoryMock.BookCreateFunc: method is nil but Repository.BookCreate was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Data      gqlmeetup.Book
		AuthorIDs []int64
	}{
		Ctx:       ctx,
		Data:      data,
		AuthorIDs: authorIDs,
	}
	lockRepositoryMockBookCreate.Lock()
	mock.calls.BookCreate = append(mock.calls.BookCreate, callInfo)
	lockRepositoryMockBookCreate.Unlock()
	return mock.BookCreateFunc(ctx, data, authorIDs)
}

// BookCreateCalls gets all the calls that were made to BookCreate.
// Check the length with:
//     len(mockedRepository.BookCreateCalls())
func (mock *RepositoryMock) BookCreateCalls() []struct {
	Ctx       context.Context
	Data      gqlmeetup.Book
	AuthorIDs []int64
} {
	var calls []struct {
		Ctx       context.Context
		Data      gqlmeetup.Book
		AuthorIDs []int64
	}
	lockRepositoryMockBookCreate.RLock()
	calls = mock.calls.BookCreate
	lockRepositoryMockBookCreate.RUnlock()
	return calls
}

// BookDelete calls BookDeleteFunc.
func (mock *RepositoryMock) BookDelete(ctx context.Context, id int64) (*gqlmeetup.Book, error) {
	if mock.BookDeleteFunc == nil {
		panic("RepositoryMock.BookDeleteFunc: method is nil but Repository.BookDelete was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	lockRepositoryMockBookDelete.Lock()
	mock.calls.BookDelete = append(mock.calls.BookDelete, callInfo)
	lockRepositoryMockBookDelete.Unlock()
	return mock.BookDeleteFunc(ctx, id)
}

// BookDeleteCalls gets all the calls that were made to BookDelete.
// Check the length with:
//     len(mockedRepository.BookDeleteCalls())
func (mock *RepositoryMock) BookDeleteCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	lockRepositoryMockBookDelete.RLock()
	calls = mock.calls.BookDelete
	lockRepositoryMockBookDelete.RUnlock()
	return calls
}

// BookGetByID calls BookGetByIDFunc.
func (mock *RepositoryMock) BookGetByID(ctx context.Context, id int64) (*gqlmeetup.Book, error) {
	if mock.BookGetByIDFunc == nil {
		panic("RepositoryMock.BookGetByIDFunc: method is nil but Repository.BookGetByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	lockRepositoryMockBookGetByID.Lock()
	mock.calls.BookGetByID = append(mock.calls.BookGetByID, callInfo)
	lockRepositoryMockBookGetByID.Unlock()
	return mock.BookGetByIDFunc(ctx, id)
}

// BookGetByIDCalls gets all the calls that were made to BookGetByID.
// Check the length with:
//     len(mockedRepository.BookGetByIDCalls())
func (mock *RepositoryMock) BookGetByIDCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	lockRepositoryMockBookGetByID.RLock()
	calls = mock.calls.BookGetByID
	lockRepositoryMockBookGetByID.RUnlock()
	return calls
}

// BookList calls BookListFunc.
func (mock *RepositoryMock) BookList(ctx context.Context, limit *int, offset *int) ([]*gqlmeetup.Book, error) {
	if mock.BookListFunc == nil {
		panic("RepositoryMock.BookListFunc: method is nil but Repository.BookList was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Limit  *int
		Offset *int
	}{
		Ctx:    ctx,
		Limit:  limit,
		Offset: offset,
	}
	lockRepositoryMockBookList.Lock()
	mock.calls.BookList = append(mock.calls.BookList, callInfo)
	lockRepositoryMockBookList.Unlock()
	return mock.BookListFunc(ctx, limit, offset)
}

// BookListCalls gets all the calls that were made to BookList.
// Check the length with:
//     len(mockedRepository.BookListCalls())
func (mock *RepositoryMock) BookListCalls() []struct {
	Ctx    context.Context
	Limit  *int
	Offset *int
} {
	var calls []struct {
		Ctx    context.Context
		Limit  *int
		Offset *int
	}
	lockRepositoryMockBookList.RLock()
	calls = mock.calls.BookList
	lockRepositoryMockBookList.RUnlock()
	return calls
}

// BookListByAuthorIDs calls BookListByAuthorIDsFunc.
func (mock *RepositoryMock) BookListByAuthorIDs(ctx context.Context, authorIDs []int64) ([]*gqlmeetup.Book, error) {
	if mock.BookListByAuthorIDsFunc == nil {
		panic("RepositoryMock.BookListByAuthorIDsFunc: method is nil but Repository.BookListByAuthorIDs was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		AuthorIDs []int64
	}{
		Ctx:       ctx,
		AuthorIDs: authorIDs,
	}
	lockRepositoryMockBookListByAuthorIDs.Lock()
	mock.calls.BookListByAuthorIDs = append(mock.calls.BookListByAuthorIDs, callInfo)
	lockRepositoryMockBookListByAuthorIDs.Unlock()
	return mock.BookListByAuthorIDsFunc(ctx, authorIDs)
}

// BookListByAuthorIDsCalls gets all the calls that were made to BookListByAuthorIDs.
// Check the length with:
//     len(mockedRepository.BookListByAuthorIDsCalls())
func (mock *RepositoryMock) BookListByAuthorIDsCalls() []struct {
	Ctx       context.Context
	AuthorIDs []int64
} {
	var calls []struct {
		Ctx       context.Context
		AuthorIDs []int64
	}
	lockRepositoryMockBookListByAuthorIDs.RLock()
	calls = mock.calls.BookListByAuthorIDs
	lockRepositoryMockBookListByAuthorIDs.RUnlock()
	return calls
}

// BookUpdate calls BookUpdateFunc.
func (mock *RepositoryMock) BookUpdate(ctx context.Context, id int64, data gqlmeetup.Book, authorIDs []int64) (*gqlmeetup.Book, error) {
	if mock.BookUpdateFunc == nil {
		panic("RepositoryMock.BookUpdateFunc: method is nil but Repository.BookUpdate was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		ID        int64
		Data      gqlmeetup.Book
		AuthorIDs []int64
	}{
		Ctx:       ctx,
		ID:        id,
		Data:      data,
		AuthorIDs: authorIDs,
	}
	lockRepositoryMockBookUpdate.Lock()
	mock.calls.BookUpdate = append(mock.calls.BookUpdate, callInfo)
	lockRepositoryMockBookUpdate.Unlock()
	return mock.BookUpdateFunc(ctx, id, data, authorIDs)
}

// BookUpdateCalls gets all the calls that were made to BookUpdate.
// Check the length with:
//     len(mockedRepository.BookUpdateCalls())
func (mock *RepositoryMock) BookUpdateCalls() []struct {
	Ctx       context.Context
	ID        int64
	Data      gqlmeetup.Book
	AuthorIDs []int64
} {
	var calls []struct {
		Ctx       context.Context
		ID        int64
		Data      gqlmeetup.Book
		AuthorIDs []int64
	}
	lockRepositoryMockBookUpdate.RLock()
	calls = mock.calls.BookUpdate
	lockRepositoryMockBookUpdate.RUnlock()
	return calls
}

// UserCreate calls UserCreateFunc.
func (mock *RepositoryMock) UserCreate(data gqlmeetup.User) error {
	if mock.UserCreateFunc == nil {
		panic("RepositoryMock.UserCreateFunc: method is nil but Repository.UserCreate was just called")
	}
	callInfo := struct {
		Data gqlmeetup.User
	}{
		Data: data,
	}
	lockRepositoryMockUserCreate.Lock()
	mock.calls.UserCreate = append(mock.calls.UserCreate, callInfo)
	lockRepositoryMockUserCreate.Unlock()
	return mock.UserCreateFunc(data)
}

// UserCreateCalls gets all the calls that were made to UserCreate.
// Check the length with:
//     len(mockedRepository.UserCreateCalls())
func (mock *RepositoryMock) UserCreateCalls() []struct {
	Data gqlmeetup.User
} {
	var calls []struct {
		Data gqlmeetup.User
	}
	lockRepositoryMockUserCreate.RLock()
	calls = mock.calls.UserCreate
	lockRepositoryMockUserCreate.RUnlock()
	return calls
}

// UserGetByEmail calls UserGetByEmailFunc.
func (mock *RepositoryMock) UserGetByEmail(ctx context.Context, email string) (*gqlmeetup.User, error) {
	if mock.UserGetByEmailFunc == nil {
		panic("RepositoryMock.UserGetByEmailFunc: method is nil but Repository.UserGetByEmail was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Email string
	}{
		Ctx:   ctx,
		Email: email,
	}
	lockRepositoryMockUserGetByEmail.Lock()
	mock.calls.UserGetByEmail = append(mock.calls.UserGetByEmail, callInfo)
	lockRepositoryMockUserGetByEmail.Unlock()
	return mock.UserGetByEmailFunc(ctx, email)
}

// UserGetByEmailCalls gets all the calls that were made to UserGetByEmail.
// Check the length with:
//     len(mockedRepository.UserGetByEmailCalls())
func (mock *RepositoryMock) UserGetByEmailCalls() []struct {
	Ctx   context.Context
	Email string
} {
	var calls []struct {
		Ctx   context.Context
		Email string
	}
	lockRepositoryMockUserGetByEmail.RLock()
	calls = mock.calls.UserGetByEmail
	lockRepositoryMockUserGetByEmail.RUnlock()
	return calls
}
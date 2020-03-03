package mocks

//go:generate moq -out repository.go -pkg mocks .. Repository
//go:generate moq -out dataloader_repository.go -pkg mocks .. DataLoaderRepository
//go:generate moq -out dataloader_service.go -pkg mocks .. DataLoaderService

package mocks

//go:generate moq -out dataloader_service_gen.go -pkg mocks .. DataLoaderService
//go:generate moq -out password_service_gen.go -pkg mocks .. PasswordService
//go:generate moq -out repository_gen.go -pkg mocks .. Repository
//go:generate moq -out session_service_gen.go -pkg mocks .. SessionService

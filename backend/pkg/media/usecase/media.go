package usecase

import "Diploma/pkg/media"

type MediaUseCaseImpl struct {
	media.Gateway
}

type MediaUseCaseModule struct {
	media.UseCase
}

func SetupMediaUseCase(gateway media.Gateway) MediaUseCaseModule {
	return MediaUseCaseModule{
		UseCase: &MediaUseCaseImpl{Gateway: gateway},
	}
}

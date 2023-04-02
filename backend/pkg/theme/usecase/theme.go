package usecase

import "Diploma/pkg/theme"

type ThemeUseCaseImpl struct {
	theme.Gateway
}

type ThemeUseCaseModule struct {
	theme.UseCase
}

func SetupThemeUseCase(gateway theme.Gateway) ThemeUseCaseModule {
	return ThemeUseCaseModule{
		UseCase: &ThemeUseCaseImpl{Gateway: gateway},
	}
}

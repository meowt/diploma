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

func (m *MediaUseCaseImpl) SetLike(userId, themeId uint) (err error) {
	return m.Gateway.SetLike(userId, themeId)
}

func (m *MediaUseCaseImpl) DeleteLike(userId, themeId uint) (err error) {
	return m.Gateway.DeleteLike(userId, themeId)
}

func (m *MediaUseCaseImpl) FollowUser(userId, followedId uint) (err error) {
	return m.Gateway.FollowUser(userId, followedId)
}

func (m *MediaUseCaseImpl) UnfollowUser(userId, followedId uint) (err error) {
	return m.Gateway.UnfollowUser(userId, followedId)
}

func (m *MediaUseCaseImpl) UpdateBackground(userId uint /*background data*/) (err error) {
	return m.Gateway.UpdateBackground(userId)
}

func (m *MediaUseCaseImpl) UpdateAvatar(userId uint /*background data*/) (err error) {
	return m.Gateway.UpdateAvatar(userId)
}

func (m *MediaUseCaseImpl) UpdateDescription(userId uint, description string) (err error) {
	return m.Gateway.UpdateDescription(userId, description)
}

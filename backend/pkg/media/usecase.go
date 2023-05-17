package media

type UseCase interface {
	SetLike(userId, themeId uint) (err error)
	DeleteLike(userId, themeId uint) (err error)
	FollowUser(userId, followedId uint) (err error)
	UnfollowUser(userId, followedId uint) (err error)
	UpdateBackground(userId uint /*background data*/) (err error)
	UpdateAvatar(userId uint /*background data*/) (err error)
	UpdateDescription(userId uint, description string) (err error)
}

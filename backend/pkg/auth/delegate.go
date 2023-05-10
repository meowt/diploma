package auth

type Delegate interface {
	SignUp(email, username, password string) (accessToken, refreshToken string, err error)
	LogIn(email, password string) (accessToken, refreshToken string, err error)
	RefreshToken(username, oldRefreshToken string) (accessToken, refreshToken string, err error)
	ParseToken(token string) (username string, err error)
}

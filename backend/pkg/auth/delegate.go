package auth

type Delegate interface {
	SignUp(email, password string) (accessToken, refreshToken string, err error)
}

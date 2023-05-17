package delegate

import (
	"strings"

	"Diploma/pkg/auth"
	"Diploma/pkg/errorPkg"
	"Diploma/pkg/models"
	"Diploma/pkg/service"

	"github.com/gin-gonic/gin"
)

type AuthDelegateImpl struct {
	auth.UseCase
	errorPkg.ErrorCreator
}

type AuthDelegateModule struct {
	auth.Delegate
}

func SetupAuthDelegate(usecase auth.UseCase) AuthDelegateModule {
	return AuthDelegateModule{
		Delegate: &AuthDelegateImpl{UseCase: usecase},
	}
}

func (au *AuthDelegateImpl) SignUp(input *models.SignUpInput) (accessToken, refreshToken string, err error) {
	user := models.UserHttp{Username: input.Username, Email: input.Email, Password: input.Password}
	hashManager := service.NewHashManager()
	user.Password, err = hashManager.HashPassword(user.Password)
	if err != nil {
		return
	}
	userUsecase := user.ToUsecase()
	return au.UseCase.SignUp(userUsecase)
}

func (au *AuthDelegateImpl) LogIn(input *models.LogInInput) (accessToken, refreshToken string, err error) {
	user := &models.UserHttp{Email: input.Email, Password: input.Password}
	userUsecase := user.ToUsecase()
	return au.UseCase.LogIn(userUsecase)
}

func (au *AuthDelegateImpl) RefreshToken(username, oldRefreshToken string) (accessToken, refreshToken string, err error) {
	return au.UseCase.RefreshToken(&models.UserUsecase{Username: username}, oldRefreshToken)
}

func (au *AuthDelegateImpl) ParseIdentity(c *gin.Context) (userIdentity *models.UserIdentity, err error) {
	header := c.GetHeader("Authorization")
	if header == "" {
		err = au.ErrorCreator.NewErrEmptyAuthHeader()
		return
	}

	_, bearerToken, ok := strings.Cut(header, " ")
	if !ok {
		err = au.ErrorCreator.NewErrStandardLibrary()
		return
	}

	userIdentity, err = au.UseCase.ParseToken(bearerToken)
	if err != nil {
		//TODO: implement custom err
		//log.Println("Parsing token error:", err.Error())
		return
	}
	return
}

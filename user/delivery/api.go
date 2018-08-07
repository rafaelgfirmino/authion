package delivery

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/rafaelgfirmino/authion/exceptions"
	"github.com/rafaelgfirmino/authion/infra/response"
	"github.com/rafaelgfirmino/authion/user/domain"
	"github.com/rafaelgfirmino/authion/user/repository"
	"github.com/rafaelgfirmino/authion/user/usecase"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func find(w http.ResponseWriter, r *http.Request) {
	userUsecase := usecase.NewUserUsecase(repository.NewMysqlUserRepository())

	switch result, err := userUsecase.FindByID(100); err {
	case exceptions.ErrorUserNotFound:
		http.Error(w, err.Error(), http.StatusNotFound)
		break
	case nil:
		w.WriteHeader(http.StatusOK)
		response.Json(result, w)
		break
	default:
		w.WriteHeader(http.StatusPreconditionFailed)
		response.Json(err, w)
	}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	userUsecase := usecase.NewUserUsecase(repository.NewMysqlUserRepository())
	ctx := r.Context()
	user := ctx.Value("user").(*domain.User)
	result, _ := userUsecase.RegisterNewUser(user)
	user.Password = ""
	response.Json(result, w)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	userUsecase := usecase.NewUserUsecase(repository.NewMysqlUserRepository())

	user := &domain.User{}
	var err error
	err = json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = userUsecase.Authenticate(user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		response.Json("Usuário ou senha inválidos", w)
		return
	}
	token := struct {
		Token string `json:"token"`
	}{generateJwtToken()}
	response.Json(token, w)
	return
}

func ConfirmationToken(w http.ResponseWriter, r *http.Request) {
	userUsecase := usecase.NewUserUsecase(repository.NewMysqlUserRepository())

	user := &domain.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch err := userUsecase.ConfirmationToken(user.ConfirmationToken); err {
	case exceptions.ErrorTokenNotFound:
		http.Error(w, err.Error(), http.StatusNotFound)
		break
	case exceptions.ErrorTryingDeleteToken, exceptions.ErrorTryingEnableUser:
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		w.WriteHeader(http.StatusOK)
		response.Json("Email de usuário confirmado com sucesso", w)
		break
	case nil:
		w.WriteHeader(http.StatusOK)
		response.Json("Email de usuário confirmado com sucesso", w)
		break
	default:
		w.WriteHeader(http.StatusPreconditionFailed)

	}
}

func generateJwtToken() string {
	path, err := os.Getwd()
	rsaPrivateFileName := "/app.rsa"
	bytes, err := ioutil.ReadFile(path + rsaPrivateFileName)
	if err != nil {
		panic("can not read the file " + rsaPrivateFileName + " on the path: " + path)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		log.Fatal("Can not do private key conversion")
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	result, err := token.SignedString(privateKey)

	if err != nil {
		log.Fatal("Can not do create token")
	}
	return result
}

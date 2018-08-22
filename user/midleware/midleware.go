package midleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rafaelgfirmino/authion/user/domain"
	"github.com/rafaelgfirmino/govalidate"
	"net/http"

	"github.com/rafaelgfirmino/SAE-Desafia/infra/response"
	"github.com/rafaelgfirmino/authion/exceptions"
	"github.com/rafaelgfirmino/authion/user/repository"
	"github.com/rafaelgfirmino/authion/user/usecase"
)

func ValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := &domain.User{}

		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		erros := govalidate.Struct(*user)
		if len(erros) > 0 {
			w.WriteHeader(http.StatusPreconditionFailed)
			response.Json(erros, w)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", user)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func UserExist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := ctx.Value("user").(*domain.User)
		userUsecase := usecase.NewUserUsecase(repository.NewMysqlUserRepository())
		result, _ := userUsecase.FindByEmail(user.Email)
		if result.Email != "" {
			w.WriteHeader(http.StatusConflict)
			responseText := fmt.Sprintf(exceptions.ErrorEmailExist.Error(), user.Email)
			response.Json(responseText, w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

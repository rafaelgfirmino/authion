package midleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oftall/authion/user/domain"
	"github.com/oftall/govalidate"
	"net/http"

	"github.com/oftall/authion/infra/response"
	"github.com/oftall/authion/exceptions"
	"github.com/oftall/authion/user/repository"
	"github.com/oftall/authion/user/usecase"
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
			response.Json(w, http.StatusPreconditionFailed, erros)
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
			responseText := fmt.Sprintf(exceptions.ErrorEmailExist.Error(), user.Email)
			response.Json(w, http.StatusConflict, responseText)
			return
		}
		next.ServeHTTP(w, r)
	})
}

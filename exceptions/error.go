package exceptions

import "errors"

var (
	ErrorUserNotFound      = errors.New("Usuário não encontrado.")
	ErrorTokenNotFound     = errors.New("O token informado não existe.")
	ErrorEmailNotFound     = errors.New("E-mail não encontrado.")
	ErrorPasswordNotFound  = errors.New("Senha incorreta.")
	ErrorTryingDeleteToken = errors.New("Erro ao tentar apagar o token de confirmação.")
	ErrorTryingEnableUser  = errors.New("Erro ao tentar habilitar o usuário.")
	ErrorEmailExist        = errors.New("Já existe um usuário com o e-mail %s cadastrado.")
)

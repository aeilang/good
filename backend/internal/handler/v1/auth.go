package v1

import (
	"net/http"

	"github.com/aeilang/backend/internal/server"
)

type Auth struct {
	authServer server.Auther
	rdbServer  server.Rediser
	mailServer server.Mialer
}

func NewAuth(auth server.Auther) *Auth {
	return &Auth{
		authServer: auth,
	}
}

const (
	pLogin          string = "login"
	pSignUp         string = "sign_up"
	pChangePassword string = "change_password"
)

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (a *Auth) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// 从requst中获取 payload

	// 验证payload合规

	// 查找该用户数据

	// 将密码哈希后进行比较

	// 生成accessToken 和 refreshToken

	// 把refreshToken保存到用户表

	// 返回accessToken 和 refreshToken

}

type SignUpPayload struct {
	UserName   string `json:"username" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6"`
	VerifyCode string `json:"verify_code" validate:"required,len=4"`
}

func (a *Auth) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	// 从r.body中获取payload

	// 验证payload合规

	// 查找是否已存在该用户

	// 从redis中取出该用户对于的验证码进行比较

	// 将用户数据保持到数据库

	// 生成accessToken 和 refreshToken

	// 将refreshToken保存到用户表

	// 返回accessToken 和 refreshToken
}

type VerifyPayload struct {
	Email   string `json:"email" validate:"required,email"`
	Purpose string `json:"purpose" validate:"required,oneof=sign_up change_password login"`
}

func (a *Auth) HandleSendVerifyCode(w http.ResponseWriter, r *http.Request) {
	// 从请求中取出payload

	// 验证payload是否合规

	// 判断purpose

	// signup 则跳过

	// change_password 或 login 则从数据库中查找该用户是否存在，如不存在这返回错误。

	// 生成4位随机验证码,并保持到redis中，key = purpose + email

	// 通过邮件发送该验证码

	// 返回Ok
}

type RefreshTokenPayload struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (a *Auth) HandleRefreshToken(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取payload

	// 验证refreshToken 是否合规

	// 解析refreshToken 判断是否过期

	// 从数据库取出该token对于的用户信息, 将两个token进行比较

	// 生成新的accessToken 和 refreshToken

	// 更改数据库中的refreshToken

	// 返回accesToken 和 refreshToken
}

type ForgetPayload struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6"`
	VerifyCode string `json:"verify_code" validate:"required,len=4"`
}

func (a *Auth) HandleForgetPassword(w http.ResponseWriter, r *http.Request) {

	// 从请求中获取payload

	// 验证该payload 是否合规

	// 从redis中取出verifyCode 并进行比较

	// 将密码进行更新

	// 取出用户数据

	// 生成accessToken 和 refreshToken

	// 更新数据库中的refreshToken

	// 返回accessToken 和 refreshToken
}

type LoginViaMail struct {
	Email      string `json:"email" validate:"required,email"`
	VerifyCode string `json:"verify_code" validate:"required,len=4"`
}

func (a *Auth) HandleLoginViaMail(w http.ResponseWriter, r *http.Request) {

	// 从请求中取出payload

	// 验证payload是否合规

	// 从redis中取出验证码,并进行比较

	// 从数据库中取出该用户数据

	// 生成accessToken 和refreshToken

	// 在数据库中更新refreshToken

	// 返回accessToken 和 refreshToken
}

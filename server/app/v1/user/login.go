package user

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
	"github.com/wenknow/gin-vue-blog/server/model/response"
	"github.com/wenknow/gin-vue-blog/server/utils/email"
	"github.com/wenknow/gin-vue-blog/server/utils/random"
	"github.com/wenknow/gin-vue-blog/server/utils/redisdb"
	"github.com/wenknow/gin-vue-blog/server/utils/verify"
	"time"
)

type LoginApi struct {
}

func (loginApi *LoginApi) Login(c *gin.Context) {
	var req request.LoginReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	var user = repository.User{}
	var err = errors.New("")
	if req.Password != "" {
		user, err = userService.CheckPwdSuccess(req.Email, req.Password)
		if err != nil {
			response.FailWithMsg(err.Error(), nil, c)
			return
		}
	} else if req.Code != "" {
		code, isOk := redisdb.Get(req.Email)
		if !isOk {
			response.FailWithMsg("验证码不存在，请重新发送验证码", nil, c)
			return
		} else if code != req.Code {
			response.FailWithMsg("验证码错误", nil, c)
			return
		}
		user, err = userService.GetUserByUsername(req.Email)
		if err == response.ErrNoRecord {
			//进行注册
			user = repository.User{
				Email:      req.Email,
				Username:   req.Email,
				Password:   req.Password,
				HeadUrl:    global.CONFIG.System.DefaultImgUrl,
				ActivateIs: true,
			}
			if newId, err := userService.AddUser(user); err != nil {
				response.FailWithMsg("添加用户出错", err, c)
				return
			} else {
				user.ID = newId
			}
		} else if err != nil {
			response.FailWithMsg(err.Error(), nil, c)
			return
		}
	} else {
		response.FailWithMsg("密码和验证码不能都为空", nil, c)
		return
	}

	//登录以后签发jwt
	j := &verify.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "wenknow",                                         // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMsg("获取token失败!", err, c)
		return
	}

	if _, isOk := redisdb.Get(user.Username); !isOk {
		redisdb.Set(user.Username, token, global.CONFIG.JWT.ExpiresTime)
	}

	response.OkWithData(response.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, c)

}

func (loginApi *LoginApi) Logout(c *gin.Context) {
	response.Ok(c)
}

func (loginApi *LoginApi) SendEmailCode(c *gin.Context) {
	var req request.SendEmailCodeReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	sendLock := req.Email + "_send_lock"
	if _, isOk := redisdb.Get(sendLock); isOk {
		response.FailWithMsg("请勿重复操作", nil, c)
		return
	} else {
		redisdb.Set(sendLock, 1, 60)
	}

	code := random.GenerateVerifyCode()
	redisdb.Set(req.Email, code, 1800)

	loginInfo := email.LoginEmailInfo{MailTo: req.Email, Code: code}
	if err := loginInfo.SendCodeEmail(); err != nil {
		response.FailWithMsg("邮件发送失败", err, c)
		return
	}
	response.Ok(c)
}

//登录页面验证码
//func (loginApi *LoginApi) Captcha(c *gin.Context) {
//	// 生成默认数字的driver
//	var store = base64Captcha.DefaultMemStore
//	driver := base64Captcha.NewDriverDigit(global.CONFIG.Captcha.ImgHeight, global.CONFIG.Captcha.ImgWidth, global.CONFIG.Captcha.KeyLong, 0.7, 80)
//	cp := base64Captcha.NewCaptcha(driver, store)
//	if id, b64s, err := cp.Generate(); err != nil {
//		response.FailWithMsg("验证码获取失败", nil, c)
//	} else {
//		response.OkWithData(response.SysCaptchaResponse{
//			CaptchaId: id,
//			PicPath:   b64s,
//		}, c)
//	}
//}

//弃用，登录和注册放一起
//func (loginApi *LoginApi) Register(c *gin.Context) {
//	var req request.RegisterReq
//	_ = c.ShouldBindJSON(&req)
//	if err := verify.Validate.Struct(req); err != nil {
//		response.FailWithMsg(verify.Translate(err), nil, c)
//		return
//	}
//	if _, err := userService.GetUserByUsername(req.Email); err != response.ErrNoRecord {
//		response.FailWithMsg("该用户已存在", nil, c)
//		return
//	}
//	//发送邮件
//	code := service.GenerateRegisterCode(req.Email)
//	confirmUrl := global.CONFIG.System.Protocol + c.Request.Host + c.Request.RequestURI + "?code=" + code
//	registerInfo := email.LoginEmailInfo{MailTo: req.Email, Url: confirmUrl}
//	if err := registerInfo.SendRegisterEmail(); err != nil {
//		response.FailWithMsg("邮件发送失败", err, c)
//		return
//	}
//
//	user := repository.User{
//		Name:       req.Name,
//		Email:      req.Email,
//		Username:   req.Email,
//		Password:   req.Password,
//		ActivateIs: false,
//	}
//	if _, err := userService.AddUser(user); err != nil {
//		response.FailWithMsg("添加用户出错", err, c)
//	} else {
//		response.OkWithData(req, c)
//	}
//}

//func (loginApi *LoginApi) RegisterConfirm(c *gin.Context) {
//	var req request.RegisterCodeReq
//	_ = c.ShouldBindQuery(&req)
//	if err := verify.Validate.Struct(req); err != nil {
//		response.FailWithMsg(verify.Translate(err), nil, c)
//		return
//	}
//	username, err := service.ValidateRegisterCode(req.Code)
//	if err != nil {
//		response.FailWithMsg(err.Error(), err, c)
//		return
//	}
//	user, err := userService.GetUserByUsername(username)
//	if err != nil {
//		response.FailWithMsg("该用户不存在", nil, c)
//		return
//	}
//	user.ActivateIs = true
//	err = userService.UpdateUser(user)
//	if err != nil {
//		response.FailWithMsg("激活账户出错", err, c)
//		return
//	}
//	response.Ok(c)
//}

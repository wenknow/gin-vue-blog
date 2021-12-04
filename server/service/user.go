package service

import (
	"encoding/json"
	"errors"
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
	"github.com/wenknow/gin-vue-blog/server/utils/encrypt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type UserService struct {
}

func (u *UserService) GetUserByUsername(username string) (user repository.User, err error) {
	err = global.DB.Where("username", username).First(&user).Error
	return
}

func (u *UserService) GetUser(cond map[string]interface{}) (user repository.User, err error) {
	err = buildCond(cond).First(&user).Error
	return
}

func (u *UserService) GetUserList(cond map[string]interface{}) (userList []repository.User, err error) {
	err = buildCond(cond).Find(&userList).Error
	return
}

func (u *UserService) GetUserByPk(id uint) (user repository.User, err error) {
	err = global.DB.First(&user, id).Error
	return
}

func (u *UserService) PagerUser(cond map[string]interface{}, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	db := buildCond(cond).Model(&repository.User{})
	var userList []repository.User
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return userList, total, err
}

func (u *UserService) AddUser(user repository.User) (uid uint, err error) {
	user.Password = GeneratePassword(user.Password)
	err = global.DB.Create(&user).Error
	return user.ID, err
}

func (u *UserService) UpdateUser(user repository.User) (err error) {
	err = global.DB.Save(&user).Error
	return
}

func (u *UserService) DelUser(user repository.User) (err error) {
	err = global.DB.Delete(&user).Error
	return
}

func (u *UserService) CheckPwdSuccess(username string, password string) (user repository.User, err error) {
	user, err = u.GetUserByUsername(username)
	if err != nil {
		return user, errors.New("用户不存在")
	}
	if !user.ActivateIs {
		return user, errors.New("请先去邮箱进行激活")
	}
	isOk, err := ValidatePassword(password, user.Password)
	if !isOk || err != nil {
		return user, errors.New("登录失败，密码错误")
	}
	return
}

func GeneratePassword(userPassword string) string {
	b, _ := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	return string(b)
}

func ValidatePassword(userPassword string, hashed string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误！")
	}
	return true, nil

}

func GenerateRegisterCode(username string) string {
	data := make(map[string]string)
	data["username"] = username
	data["time"] = strconv.FormatInt(time.Now().Unix(), 10)
	b, _ := json.Marshal(data)
	return encrypt.AesEncrypt(string(b), global.CONFIG.Encrypt.RegisterKey)
}

func ValidateRegisterCode(enData string) (username string, err error) {

	data := encrypt.AesDecrypt(enData, global.CONFIG.Encrypt.RegisterKey)
	m := make(map[string]string)
	err = json.Unmarshal([]byte(data), &m)
	if err != nil {
		return "", err
	}
	addTime, ok := m["time"]
	if !ok {
		return "", errors.New("code不合法")
	}
	addTimeInt, err := strconv.ParseInt(addTime, 10, 64)
	if err != nil {
		return "", err
	}
	if addTimeInt < time.Now().Unix()-60*60*48 {
		return "", errors.New("该验证编码已过期")
	}
	username, ok = m["username"]
	if !ok {
		return "", errors.New("code不合法")
	}
	return
}

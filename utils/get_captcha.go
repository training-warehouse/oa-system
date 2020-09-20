package utils

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type Captcha struct {
	Id   string
	BS64 string
	Code int
}

var store = base64Captcha.DefaultMemStore

func GetCaptcha() (id string, base64 string, err error) {
	rgbaColor := color.RGBA{0, 0, 0, 0}
	fonts := []string{"wqy-microhei.ttc"}

	// 生成driver（高，宽,背景文字干扰,干扰线条数,背景颜色指针）
	driver := base64Captcha.NewDriverMath(
		50, 140, 0, 2, &rgbaColor, fonts,
	)
	// 使用store和driver生成验证码实例
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, base64, err = captcha.Generate()
	return id, base64, err
}

func VerifyCaptcha(id string, ret_captcha string) bool {
	return store.Verify(id, ret_captcha, true)
}

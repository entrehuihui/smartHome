package tools

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/nfnt/resize"
	"github.com/spf13/viper"
	"golang.org/x/net/http2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MD5Password 密码MD5加密混淆
func MD5Password(password string) string {
	password = viper.GetString("Admin.key") + password
	return MD5(password)
}

// MD5 密码加密
func MD5(password string) string {
	//密码MD5加密
	p5 := md5.New()
	p5.Write([]byte(password))
	return hex.EncodeToString(p5.Sum(nil))
}

// ComputeHmac256 加密
func ComputeHmac256(message, secret *string) (string, error) {
	key := []byte(*secret)
	h := hmac.New(sha256.New, key)
	_, err := h.Write([]byte(*message))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), err
}

// URLEncode 格式化URL
func URLEncode(key, value *string) string {
	// url encode
	v := url.Values{}
	v.Add(*key, *value)
	body := v.Encode()
	return body
}

// Post post简单封装
func Post(url, code *string) ([]byte, error) {
	resp, err := http.Post(*url, "application/json", strings.NewReader(*code))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

// VerifyPasswordFormat 校验密码
func VerifyPasswordFormat(name string) bool {
	// 48 57 97 122 65 90
	// pattern := `^([0-9A-Za-z]+$){6,16}$`
	// reg := regexp.MustCompile(pattern)
	// return reg.MatchString(name)
	if len(name) < 6 || len(name) > 20 {
		return false
	}
	num := false
	to := false
	low := false
	for _, da := range name {
		if da > 96 && da < 123 {
			num = true
		} else if da > 64 && da < 91 {
			low = true
		} else if da > 47 && da < 58 {
			to = true
		}
		if num && to && low {
			break
		}
	}
	return num && to && low
}

// VerifyNameFormat 校验登录名
func VerifyNameFormat(name string) bool {
	pattern := "^[-@._a-zA-Z0-9\u4e00-\u9fa5]*$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(name)
}

// VerifyStringNumberFormat 校验数组字母
func VerifyStringNumberFormat(name string) bool {
	pattern := "^[a-zA-Z0-9]*$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(name)
}

// VerifyEmailFormat 校验邮箱
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// VerifyMobile 校验电话
func VerifyMobile(mobile string) bool {
	reg := `^1[3456789]\d{9}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(mobile)
}

// GetRequsetInfo 获取请求信息
func GetRequsetInfo(code string) string {
	return "  |" + code + "|\t\t\t  "
}

// RandString 生成随机验证码
func RandString(l int) string {
	a := fmt.Sprintf("%d", &l)
	c, _ := strconv.ParseInt(a, 10, 64)
	r := rand.New(rand.NewSource(GetTimes() + c))
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

// GetTLSConfig .
func GetTLSConfig(certPemPath, certKeyPath string) *tls.Config {
	var certKeyPair *tls.Certificate
	cert, err := ioutil.ReadFile(certPemPath)
	if err != nil {
		log.Fatal(err)
	}
	key, err := ioutil.ReadFile(certKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Println("TLS KeyPair err: ", err)
	}

	certKeyPair = &pair

	return &tls.Config{
		Certificates: []tls.Certificate{*certKeyPair},
		NextProtos:   []string{http2.NextProtoTLS},
	}
}

// ResizeJpegSize 改变传入的JPEG图片地址的图片的尺寸
func ResizeJpegSize(data []byte, pWidth, pHight uint) ([]byte, error) {
	reader := bytes.NewReader(data)
	//解码图片
	img, err := jpeg.Decode(reader)
	if err != nil {
		return nil, err
	}
	// 改变图片尺寸
	m := resize.Resize(pWidth, pHight, img, resize.Lanczos3)

	buf := make([]byte, 0)
	bufWrite := bytes.NewBuffer(buf)

	err = jpeg.Encode(bufWrite, m, nil)
	if err != nil {
		return nil, err
	}
	return bufWrite.Bytes(), nil
}

// ResizePngSize 改变传入的PNG图片地址的图片的尺寸
func ResizePngSize(data []byte, pWidth, pHight uint) ([]byte, error) {
	reader := bytes.NewReader(data)
	//解码图片
	img, err := png.Decode(reader)
	if err != nil {
		return nil, err
	}
	// 改变图片尺寸
	m := resize.Resize(pWidth, pHight, img, resize.Lanczos3)

	buf := make([]byte, 0)
	bufWrite := bytes.NewBuffer(buf)

	err = png.Encode(bufWrite, m)
	if err != nil {
		return nil, err
	}
	return bufWrite.Bytes(), nil
}

func ERROR(code codes.Code, err string) error {
	return status.Errorf(code, err)
}

func GetTimes() int64 {
	return time.Now().UnixMilli()
}

// VerifyMobLandline 校验电话和座机
func VerifyMobLandline(mobile string) bool {
	reg := `^\d{10,12}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(mobile)
}

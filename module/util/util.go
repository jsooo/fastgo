package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gosexy/to"

	"github.com/clbanning/mxj"
	"github.com/jsooo/log"
	"github.com/satori/go.uuid"

	beego_config "github.com/astaxie/beego/config"
)

type XmlCDATA struct {
	Text string `xml:",cdata"`
}

func Json2Map(jsonStr string) (s map[string]interface{}, err error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func JsonDecode(jsonStr string, structModel interface{}) error {
	decode := json.NewDecoder(strings.NewReader(jsonStr))
	err := decode.Decode(structModel)
	return err
}

func JsonEncode(structModel interface{}) (string, error) {
	jsonStr, err := json.Marshal(structModel)
	return string(jsonStr), err
}

func XmlDecode(xmlStr string) (map[string]interface{}, error) {
	xmlMap, _ := mxj.NewMapXml([]byte(xmlStr))
	retMap := make(map[string]interface{}, 0)
	keys, err := xmlMap.Elements("xml")
	if err != nil {
		return retMap, err
	}
	for _, k := range keys {
		retMap[k] = xmlMap.ValueOrEmptyForPathString("xml." + k)
	}

	return retMap, nil
}

func XmlEncode(xmlMap map[string]interface{}) (string, error) {
	mp := mxj.New()
	mp = xmlMap
	xmlStr, err := mp.Xml("xml")

	return string(xmlStr), err
}

//生成uuid
func GenUuid() string {
	uuidFunc, _ := uuid.NewV4()
	uuidStr := uuidFunc.String()
	uuidStr = strings.Replace(uuidStr, "-", "", -1)
	uuidByt := []rune(uuidStr)
	return string(uuidByt[8:24])
}

func GenShortUuid() string {
	dict := map[int]byte{
		0: '0', 1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7', 8: '8', 9: '9', 10: 'a', 11: 'b', 12: 'c',
		13: 'd', 14: 'e', 15: 'f', 16: 'g', 17: 'h', 18: 'i', 19: 'j', 20: 'k', 21: 'l', 22: 'm', 23: 'n', 24: 'o',
		25: 'p', 26: 'q', 27: 'r', 28: 's', 29: 't', 30: 'u', 31: 'v', 32: 'w', 33: 'x', 34: 'y', 35: 'z', 36: 'A',
		37: 'B', 38: 'C', 39: 'D', 40: 'E', 41: 'F', 42: 'G', 43: 'H', 44: 'I', 45: 'J', 46: 'K', 47: 'L', 48: 'M',
		49: 'N', 50: 'O', 51: 'P', 52: 'Q', 53: 'R', 54: 'S', 55: 'T', 56: 'U', 57: 'V', 58: 'W', 59: 'X', 60: 'Y',
		61: 'Z', 62: '0', 63: '1'}

	uuidFunc, _ := uuid.NewV4()
	uuidStr := uuidFunc.String()
	uuidStr = strings.Replace(uuidStr+"0000", "-", "", -1)

	var shortUuid string

	for index := 0; index < 6; index++ {
		word, _ := strconv.ParseInt(string(uuidStr[index*6:index*6+6]), 16, 0)
		binNum := fmt.Sprintf("%024b", word)
		for bit := 0; bit < 4; bit++ {
			stepBit := binNum[bit*6 : bit*6+6]
			dexNum, _ := strconv.ParseInt(stepBit, 2, 0)
			if index == 5 && bit >= 2 {
				continue
			}
			shortUuid += string(dict[int(dexNum)])
		}
	}
	return shortUuid
}

//生成交易ID
func GenTransId() string {
	return "trans_" + GenShortUuid()
}

//生成噪声字符串（复用uuid）
func GenNonceStr() string {
	return GenUuid()
}

func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

//data => aaa=data&bbb=data
func HttpPost(url string, data string) (string, error) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader(data))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetTLSConfig(certPath, keyPath string) (*tls.Config, error) {
	var _tlsConfig *tls.Config
	if _tlsConfig != nil {
		return _tlsConfig, nil
	}

	// load cert
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Errorf("load wechat keys fail: %v", err)
		return nil, err
	}

	_tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	return _tlsConfig, nil
}

func SecurePost(url string, data string, certPath string, keyPath string) (string, error) {
	tlsConfig, err := GetTLSConfig(certPath, keyPath)
	if err != nil {
		return "", err
	}

	tr := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(
		url,
		"text/xml",
		bytes.NewBuffer(to.Bytes(data)))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func HttpPostWithSignHeader(url string, data string, sign string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(to.Bytes(data)))
	req.Header.Add("X-Meisha-Sign", sign)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Timestamp() int {
	return int(time.Now().Unix())
}

func Base64Decode(sourceStr string) (string, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(sourceStr)
	if err != nil {
		return "", err
	}

	return string(decodeBytes), nil
}

func Base64Encode(sourceStr string) string {
	return base64.StdEncoding.EncodeToString([]byte(sourceStr))
}

func WxTimeStr2Timestamp(timeStr string) int {
	return TimeStr2Time(timeStr, "20060102150405")
}

func TimeStr2Time(timeStr, format string) int {
	loc, _ := time.LoadLocation("Local")                     //重要：获取时区
	theTime, _ := time.ParseInLocation(format, timeStr, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                     //转化为时间戳 类型是int64
	return int(sr)
}

//把时间戳转换成微信支付要求的时间格式：20060102030405
func GetWxTime(timestamp ...int) string {
	timestampTmp := 0
	if len(timestamp) == 0 {
		timestampTmp = Timestamp()
	} else {
		timestampTmp = timestamp[0]
	}

	tm := time.Unix(int64(timestampTmp), 0)
	return tm.Format("20060102150405")
}

//生成随机数
func Rand() float64 {
	sour := rand.New(rand.NewSource(time.Now().UnixNano()))
	return sour.Float64()
}

func RandWithMax(maxNum int) int32 {
	sour := rand.New(rand.NewSource(time.Now().UnixNano()))
	return sour.Int31n(int32(maxNum))
}

func HttpGetOriginData(getUrl string) (body []byte) {
	if len(getUrl) == 0 {
		return
	}

	log.Error(getUrl)

	response, err := http.Get(getUrl)
	if err != nil {
		log.Errorf("%v", err)
		response.Body.Close()
		return
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf("%v", err)
		response.Body.Close()
		return
	}
	response.Body.Close()
	return body
}

//校验身份证规则合法性
func VerifyIdCard(idCard string) bool {
	//位数不满18位 false
	if len(idCard) != 18 {
		return false
	}

	idCard = strings.Replace(idCard, "x", "X", -1)

	var salt = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var valiCode = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

	var sum int = 0
	for index := 0; index < 17; index++ {
		cell, _ := strconv.Atoi(string(idCard[index]))
		sum += salt[index] * cell
	}
	mod := sum % 11
	lastNum := idCard[17]
	if valiCode[mod] == lastNum {
		return true
	} else {
		//兼容大小写Xx
		if valiCode[mod] == 'x' && lastNum == 'X' {
			return true
		}
		return false
	}
}

//字符串分割
func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}

//结构体转map，tag: `data:"aliasName;required"`
func Struct2Map(obj interface{}) (data map[string]interface{}, err error) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	data = make(map[string]interface{}, 0)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("data")
		value := v.Field(i).Interface()

		if len(tag) == 0 {
			data[t.Field(i).Name] = value
		} else {
			s := strings.Split(tag, ";")
			if len(s) == 2 && s[1] == "required" && IsEmpty(value) {
				return data, errors.New(t.Field(i).Name + " 不能为空")
			} else if IsEmpty(value) {
				continue
			}
			data[s[0]] = value
		}
	}
	return
}

func IsEmpty(value interface{}) bool {
	switch value.(type) {
	case int:
		return value.(int) == 0
	case string:
		return value.(string) == ""
	case interface{}:
		return value.(interface{}) == nil
	case bool:
		return value.(bool) == false
	default:
		return false
	}
}

//判断是不是正式模式
func IsLiveMode() bool {
	//在测试环境打开mysql 调试
	iniconf, err := beego_config.NewConfig("ini", "config/app.conf")
	if err != nil {
		return false
	} else {
		if iniconf.String("runmode") == "prod" {
			return true
		} else {
			return false
		}
	}
}

//map的key周字符小写
func MapKeyFirstToLower(data map[string]interface{}) (retData map[string]interface{}) {
	retData = make(map[string]interface{}, 0)
	for key, value := range data {
		retData[StrFirstToUpper(key)] = value
	}

	return retData
}

//字符串首字母大写
func StrFirstToUpper(str string) string {
	var lowerStr string
	for y := 0; y < len(str); y++ {
		vv := []rune(str)
		if y == 0 {
			vv[y] += 32
			lowerStr += string(vv[y]) // + string(vv[i+1])
		} else {
			lowerStr += string(vv[y])
		}

	}
	return lowerStr
}

//获取两个字符串中间的字符串
func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	} else {
		n += len(start)
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

//字符串数组去重
func StringSliceUnique(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	if len(origData) == 0 {
		return nil, nil
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	if len(crypted) == 0 {
		return nil, nil
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var HttpHeader = make(map[string]interface{}, 0)

type errNo struct {
	Code    int
	Message string
}

func (c *errNo) Error() string { // 实现接口
	return c.Message
}

// MyErr 自定义错误
func MyErr(code int, Message string) error {
	return &errNo{
		Code:    code,
		Message: Message,
	}
}

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// AnyJson Json通用处理
func AnyJson(byteJson []byte) map[string]interface{} {
	resMap := make(map[string]interface{}, 0)
	if err := json.Unmarshal(byteJson, &resMap); err != nil {
		fmt.Println("解析Json失败: ", err)
	}
	return resMap
}

// HttpPost Post 请求
func HttpPost(url string, param string, paramType string) (_result []byte, _err error) {
	client := &http.Client{
		Timeout: time.Second * 20,
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(param))
	if err != nil {
		_err = err
	}
	//Post请求类型
	if paramType == "json" {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	//header 需要增加的内容
	if len(HttpHeader) > 0 {
		for k, v := range HttpHeader {
			req.Header.Add(k, v.(string))
		}
	}
	res, err := client.Do(req)
	if err != nil {
		_err = err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			_err = err
		}
	}(res.Body)
	//返回状态
	status := res.StatusCode
	if status == 200 {
		_result, err = io.ReadAll(res.Body)
		if err != nil {
			_err = err
		}
	} else {
		errMsg, _ := io.ReadAll(res.Body)
		fmt.Println("错误的返回消息: ", string(errMsg))
		_err = MyErr(status, "http错误代码:"+strconv.Itoa(status))
	}
	return _result, _err
}

// HttpGet Get 请求
func HttpGet(url string) (_result []byte, _err error) {
	client := &http.Client{
		Timeout: time.Second * 20,
	}
	req, err := client.Get(url)
	if err != nil {
		_err = err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			_err = err
		}
	}(req.Body)
	//返回状态
	statusCode := req.StatusCode
	if statusCode == 200 {
		_result, err = io.ReadAll(req.Body)
		if err != nil {
			_err = err
		}
	} else {
		_err = MyErr(statusCode, "http错误代码:"+strconv.Itoa(statusCode))
	}
	return _result, _err
}

// StringInArr 数组中是否存在字符串
func StringInArr(str string, arr []string) bool {
	for _, obj := range arr {
		if str == obj {
			return true
		}
	}
	return false
}

// ArrDel 删除数组中元素
func ArrDel(arr []string, str string) []string {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != str {
			arr[j] = arr[i]
			j++
		}
	}
	return arr[:j]
}

// ReturnOK 构造 true 返回信息
func ReturnOK(data interface{}) string {
	type returnJson struct {
		Code string      `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
	var reJson = &returnJson{}
	reJson.Code = "0"
	reJson.Msg = "ok"
	reJson.Data = data
	res, _ := json.Marshal(reJson)
	return string(res)
}

// ReturnErr 构造 false 返回信息
func ReturnErr(msg string) string {
	return `{"code": "1", "msg": "` + msg + `"}`
}

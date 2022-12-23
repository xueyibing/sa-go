package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JsonResult 返回结构
type userInfoResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}


// CounterHandler 计数器接口
func GetUserinfoHandler(w http.ResponseWriter, r *http.Request) {
	res := &userInfoResult{}

	if r.Method == http.MethodPost {

		decoder := json.NewDecoder(r.Body)
		body := make(map[string]interface{})
		if err := decoder.Decode(&body); err != nil {
			return
		}
		defer r.Body.Close()
		code, ok := body["code"]
		if !ok {
			res.ErrorMsg = "缺少 code 参数"
			return
		}
		fmt.Printf("code:%s",code)

	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

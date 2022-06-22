package responses

import (
	"github.com/gin-gonic/gin"
)

type meta_resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type meta_resperr struct {
	Error string `json:"error"`
}

type Response2 struct {
	Meta meta_resp   `json:"meta"`
	Data interface{} `json:"data"`
}

func SuccesResponses(statusCode int, metaCode int, metaMsg string, data interface{}) gin.H {
	res_meta := meta_resp{Code: metaCode, Msg: metaMsg}
	resp := Response2{Meta: res_meta, Data: data}
	return gin.H{"result": resp}
}

func ErrorResponses(statusCode int, metaCode int, metaMsg string, err string) gin.H {
	res_metad := meta_resp{Code: metaCode, Msg: metaMsg}
	deferror := meta_resperr{Error: err}
	respd := Response2{Meta: res_metad, Data: deferror}
	return gin.H{"result": respd}
}

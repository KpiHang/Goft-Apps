package middleware

import (
	"github.com/gin-gonic/gin"
)

type IndexTest struct{}

func NewIndexTest() *IndexTest {
	return &IndexTest{}
}

func (this *IndexTest) OnRequest(ctx *gin.Context) error {
	return nil
}

func (this *IndexTest) OnResponse(result interface{}) (interface{}, error) {
	if m, ok := result.(gin.H); ok {
		m["metadate"] = "index test"
	}
	return result, nil
}

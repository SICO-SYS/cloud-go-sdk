/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package wechat

import (
	"encoding/xml"
	"net/http"
	"sort"
	"strings"

	"github.com/SiCo-Ops/public"
)

type EncryptRequest struct{}
type EncryptResponse struct{}

type TextRequest struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	// text
	Content string
	// event
	Event    string
	EventKey string
	Tickey   string
	// Latitude  string
	// Longitude string
	// Precision string
	// image voice video shortvideo
	PicUrl       string
	MediaId      string
	Format       string
	Recognition  string
	ThumbMediaId string
	// location
	Location_X string
	Location_Y string
	Scale      string
	// link
	Title       string
	Description string
	Url         string
	MsgId       int64
}

type TextResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   string
	MsgType      CDATAText
	Content      CDATAText
}

type CDATAText struct {
	Text string `xml:",innerxml"`
}

func GetValidation(req *http.Request) (string, string, string) {
	nonce := req.URL.Query().Get("nonce")
	timestamp := req.URL.Query().Get("timestamp")
	signature := req.URL.Query().Get("signature")
	return nonce, timestamp, signature
}

func ValidateServer(token, nonce, timestamp, signature string) bool {
	sortparams := []string{token, nonce, timestamp}
	sort.Strings(sortparams)
	str := strings.Join(sortparams, "")
	sign := public.EncryptWithSha1(str)
	if sign == signature {
		return true
	}
	return false
}

func CDATA(v string) CDATAText {
	return CDATAText{"<![CDATA[" + v + "]]>"}
}

func Parse(data []byte) *TextRequest {
	v := &TextRequest{}
	xml.Unmarshal(data, v)
	return v
}

func Marshal(to, from, timestamp, msgtype, content string) []byte {
	v := &TextResponse{}
	v.ToUserName = CDATA(to)
	v.FromUserName = CDATA(from)
	v.CreateTime = timestamp
	v.MsgType = CDATA(msgtype)
	switch msgtype {
	case "text":
		v.Content = CDATA(content)
	default:
		v.MsgType = CDATA("text")
		v.Content = CDATA("Master is foolish, MsgType execute error")
	}
	data, _ := xml.Marshal(v)
	return data
}

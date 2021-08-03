package service

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/multi"
)

type Form struct {
	Rule    []Rule                   `json:"rule"`
	Action  string                   `json:"action"`
	Method  string                   `json:"method"`
	Title   string                   `json:"title"`
	Config  Config                   `json:"config,omitempty"`
	Headers []map[string]interface{} `json:"headers,omitempty"`
}

func (form *Form) SetAction(uri string, ctx *gin.Context) {
	form.Action = SetUrl(uri, ctx)
}

func SetUrl(uri string, ctx *gin.Context) string {
	if multi.IsAdmin(ctx) {
		return g.TENANCY_CONFIG.System.AdminPreix + uri
	} else if multi.IsTenancy(ctx) {
		return g.TENANCY_CONFIG.System.ClientPreix + uri
	}
	return ""
}

type Config struct {
}

type Rule struct {
	Title    string                   `json:"title"`
	Type     string                   `json:"type"`
	Field    string                   `json:"field"`
	Info     string                   `json:"info"`
	Value    interface{}              `json:"value"`
	Props    map[string]interface{}   `json:"props"`
	Options  []Option                 `json:"options,omitempty"`
	Control  []Control                `json:"control,omitempty"`
	Validate []map[string]interface{} `json:"validate,omitempty"`
}
type ControlRule struct {
	Title    string                   `json:"title"`
	Type     string                   `json:"type"`
	Field    string                   `json:"field"`
	Info     string                   `json:"info"`
	Value    interface{}              `json:"value"`
	Props    map[string]interface{}   `json:"props"`
	Options  []Option                 `json:"options,omitempty"`
	Validate []map[string]interface{} `json:"validate,omitempty"`
}

type Control struct {
	Value int    `json:"value"`
	Rule  []Rule `json:"rule"`
}

func (r *Rule) TransData(rule string, token []byte) {
	switch r.Type {
	case "input":
		r.Props = map[string]interface{}{
			"placeholder": "请输入" + r.Title,
			"type":        "text",
		}
	case "textarea":
		r.Props = map[string]interface{}{
			"placeholder": "请输入" + r.Title,
			"type":        "textarea",
		}
		r.Type = "input"
	case "number":
		r.Props = map[string]interface{}{
			"placeholder": "请输入" + r.Title,
		}
		r.Type = "inputNumber"
	case "radio":
		r.Props = map[string]interface{}{}
		rules := strings.Split(rule, ";")
		for _, ru := range rules {
			rus := strings.Split(ru, ":")
			if len(rus) == 2 {
				r.Options = append(r.Options, Option{Label: rus[1], Value: rus[0]})
			}
		}
	case "file":
		seitURL, _ := GetSeitURL()
		r.Props = map[string]interface{}{
			"action": fmt.Sprintf("%s/v1/admin/media/upload", seitURL),
			"data":   map[string]interface{}{},
			"headers": map[string]interface{}{
				"Authorization": "Bearer " + string(token),
			},

			"limit":      1,
			"uploadType": "file",
		}
		r.Type = "upload"
	case "image":
		r.Props = map[string]interface{}{
			"footer":    false,
			"height":    "480px",
			"maxLength": 1,
			"modal":     map[string]interface{}{"modal": false},
			"src":       "/admin/setting/uploadPicture?field=" + r.Field + "&type=1",
			"title":     "请选择" + r.Title,
			"type":      r.Type,
			"width":     "896px",
		}
		r.Type = "frame"
	}

}

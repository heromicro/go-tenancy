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

type Config struct{}

type Rule struct {
	Title    string                   `json:"title"`
	Type     string                   `json:"type"`
	Field    string                   `json:"field"`
	Info     string                   `json:"info"`
	Value    interface{}              `json:"value"`
	Props    map[string]interface{}   `json:"props"`
	Col      map[string]interface{}   `json:"col,omitempty"`
	Options  []Option                 `json:"options,omitempty"`
	Controls []Control                `json:"control,omitempty"`
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
	Value interface{} `json:"value"`
	Rule  []Rule      `json:"rule"`
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

func (rule *Rule) AddValidator(validator map[string]interface{}) *Rule {
	rule.Validate = append(rule.Validate, validator)
	return rule
}

func (rule *Rule) AddOption(opt Option) *Rule {
	rule.Options = append(rule.Options, opt)
	return rule
}

func (rule *Rule) AddControl(control Control) *Rule {
	rule.Controls = append(rule.Controls, control)
	return rule
}

func (rule *Rule) AddProps(props map[string]interface{}) *Rule {
	rule.Props = props
	return rule
}

func NewRadio(title, field, info string, value interface{}) *Rule {
	return &Rule{
		Title: title,
		Type:  "radio",
		Field: field,
		Value: value,
		Info:  info,
	}
}

func NewInput(title, field, placeholder string, value interface{}) *Rule {
	return &Rule{
		Title: title,
		Type:  "input",
		Field: field,
		Value: value,
		Props: map[string]interface{}{
			"type":        "text",
			"placeholder": placeholder,
		},
	}
}

func NewFrame(title, field, placeholder string, value interface{}) *Rule {
	return &Rule{
		Title: title,
		Type:  "frame",
		Field: field,
		Value: value,
	}
}

func NewRate(title, field string, span int64, value interface{}) *Rule {
	return &Rule{
		Title: title,
		Type:  "rate",
		Field: field,
		Value: value,
		Col: map[string]interface{}{
			"span": 8,
		},
		Props: map[string]interface{}{
			"max": 5,
		},
	}
}

func NewSwitch(title, field string, value interface{}) *Rule {
	return &Rule{
		Title: title,
		Type:  "switch",
		Field: field,
		Value: value,
		Props: map[string]interface{}{
			"activeValue":   1,
			"inactiveValue": 2,
			"inactiveText":  "关闭",
			"activeText":    "开启",
		},
	}
}

func (form *Form) AddRule(rule Rule) *Form {
	form.Rule = append(form.Rule, rule)
	return form
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

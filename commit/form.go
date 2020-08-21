package commit

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"
	"text/template"

	"github.com/spf13/viper"
	"gopkg.in/AlecAivazis/survey.v1"
)

func FillOutForm() ([]byte, error) {
	// load form
	if form, tmplText, err := loadForm(); err != nil {
		log.Printf("loadForm failed, err=%v\n", err)
		return nil, err
	} else {

		// ask the question
		answers := map[string]interface{}{}
		if err := survey.Ask(form, &answers); err != nil {
			return nil, err
		}

		// assemble the answers to commit message
		var buf bytes.Buffer
		if err := assembleMessage(&buf, tmplText, answers); err != nil {
			log.Printf("assemble failed, err=%v\n", err)
		}

		return buf.Bytes(), nil
	}
}

type FormItemOption struct {
	Name string
	Desc string
}

type FormItem struct {
	Name     string
	Desc     string
	Form     string
	Options  []*FormItemOption
	Required bool
}

type MessageConfig struct {
	Items    []*FormItem
	Template string
}

func assembleMessage(buf *bytes.Buffer, tmplText string, answers map[string]interface{}) error {
	if tmpl, err := template.New("").Parse(tmplText); err != nil {
		return err
	} else {
		for k, v := range answers {
			if vString, ok := v.(string); ok {
				answers[k] = strings.TrimSpace(vString)
			}
		}
		if err := tmpl.Execute(buf, answers); err != nil {
			return err
		}
		return nil
	}
}

func loadForm() (qs []*survey.Question, _ string, err error) {
	config := struct{ Message MessageConfig }{}
	if err := json.Unmarshal([]byte(defaultConfig), &config); err != nil {
		return nil, "", err
	}

	msgConfig := config.Message
	log.Printf("default config message tmpl: %s", msgConfig.Template)

	sub := viper.Sub("message")
	if sub == nil {
		log.Printf("no message in config file")
	} else {
		if err := sub.Unmarshal(&msgConfig); err != nil {
			log.Printf("ill message in config file, err=%v", err)
		} else {
			log.Printf("msg config from file: %v", msgConfig)
			item := msgConfig.Items[0]
			log.Printf("msg config item: %s", item.Desc)
			log.Printf("msg config template: %s", msgConfig.Template)
		}
	}

	for _, item := range msgConfig.Items {
		q := survey.Question{
			Name: item.Name,
		}
		if item.Required {
			q.Validate = survey.Required
		}
		switch item.Form {
		case "input":
			q.Prompt = &survey.Input{
				Message: item.Desc,
			}
		case "multiline":
			q.Prompt = &survey.Multiline{
				Message: item.Desc,
			}
		case "select":
			prompt := &survey.Select{
				Message:  item.Desc,
				PageSize: 8,
			}
			for _, option := range item.Options {
				prompt.Options = append(prompt.Options, option.Desc)
			}
			q.Prompt = prompt
			q.Transform = func(options []*FormItemOption) func(interface{}) interface{} {
				return func(ans interface{}) (newAns interface{}) {
					if ansString, ok := ans.(string); !ok {
						return nil
					} else {
						for _, option := range options {
							if ansString == option.Desc {
								return option.Name
							}
						}
						return nil
					}
				}
			}(item.Options)
		}
		qs = append(qs, &q)
	}

	return qs, msgConfig.Template, nil
}

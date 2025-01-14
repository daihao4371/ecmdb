package rule

import (
	"fmt"
	"github.com/Duke1616/ecmdb/internal/pkg/wechat"
	"github.com/Duke1616/enotify/notify/feishu/card"
	"github.com/ecodeclub/ekit/slice"
	"github.com/xen0n/go-workwx"
)

func GetFields(rules []Rule, provide uint8, data map[string]interface{}) []card.Field {
	ruleMap := slice.ToMap(rules, func(element Rule) string {
		return element.Field
	})

	// 拼接消息体
	num := 1
	var fields []card.Field

	// 判断不同平台的消息来源，进行处理
	switch provide {
	case SystemProvide:
		for field, value := range data {
			title := field
			val, ok := ruleMap[field]
			if ok {
				title = val.Title
			}

			fields = append(fields, card.Field{
				IsShort: true,
				Tag:     "lark_md",
				Content: fmt.Sprintf(`**%s:**\n%v`, title, value),
			})

			if num%2 == 0 {
				fields = append(fields, card.Field{
					IsShort: false,
					Tag:     "lark_md",
					Content: "",
				})
			}

			num++
		}
	case WechatProvide:
		oaData, err := wechat.Unmarshal(data)
		if err != nil {
			return nil
		}

		for _, contents := range oaData.ApplyData.Contents {
			key := contents.Title[0].Text

			switch contents.Control {
			case "Selector":
				switch contents.Value.Selector.Type {
				case "single":
					fields = append(fields, card.Field{
						IsShort: true,
						Tag:     "lark_md",
						Content: fmt.Sprintf(`**%s:**\n%v`, key, contents.Value.Selector.Options[0].Value[0].Text),
					})
				case "multi":
					value := slice.Map(contents.Value.Selector.Options, func(idx int,
						src workwx.OAContentSelectorOption) string {
						return src.Value[0].Text
					})

					fields = append(fields, card.Field{
						IsShort: true,
						Tag:     "lark_md",
						Content: fmt.Sprintf(`**%s:**\n%v`, key, value),
					})
				}
			case "Textarea":
				fields = append(fields, card.Field{
					IsShort: true,
					Tag:     "lark_md",
					Content: fmt.Sprintf(`**%s:**\n%v`, key, contents.Value.Text),
				})
			case "default":
				fmt.Println("不符合筛选规则")
			}

			if num%2 == 0 {
				fields = append(fields, card.Field{
					IsShort: false,
					Tag:     "lark_md",
					Content: "",
				})
			}

			num++
		}
	}

	return fields
}

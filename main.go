package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
)

type RuleItem struct {
	Op  string      `json:"op"`
	Val interface{} `json:"val"`
}

type Rule map[string]RuleItem

// 执行规则匹配的函数
func matchRule(rule Rule, data map[string]interface{}) bool {
	allMatch := true
	//fmt.Printf("%+v\n", rule)
	for field, ruleItem := range rule {
		value, exists := data[field]
		if !exists {
			fmt.Printf("字段 %s 不存在于输入数据中\n", field)
			allMatch = false
			return false
		}

		//fmt.Printf("%+v\n", ruleItem)

		switch ruleItem.Op {
		case "in":
			switch val := ruleItem.Val.(type) {
			case []string:
				strValue, ok := value.(string)
				if ok {
					found := false
					for _, item := range val {
						if item == strValue {
							found = true
							break
						}
					}
					if !found {
						allMatch = false
						fmt.Printf("字段 %s 的值 %s 不在 %v 中\n", field, strValue, val)
					}
				} else {
					fmt.Printf("字段 %s 的值类型不是字符串，无法进行 'in' 操作\n", field)
					allMatch = false
				}
			case []int:
				intValue, ok := value.(int)
				if ok {
					found := false
					for _, item := range val {
						if item == intValue {
							found = true
							break
						}
					}
					if !found {
						allMatch = false
						fmt.Printf("字段 %s 的值 %d 不在 %v 中\n", field, intValue, val)
					}
				} else {
					fmt.Printf("字段 %s 的值类型不是整数，无法进行 'in' 操作\n", field)
					allMatch = false
				}
			case []interface{}:
				strValue, ok := value.(string)
				if ok {
					found := false
					for _, item := range val {
						if cast.ToString(item) == strValue {
							found = true
							break
						}
					}
					if !found {
						allMatch = false
						fmt.Printf("字段 %s 的值 %s 不在 %v 中\n", field, strValue, val)
					}
				} else {
					fmt.Printf("字段 %s 的值类型不是字符串，无法进行 'in' 操作\n", field)
					allMatch = false
				}
			default:
				fmt.Printf("不支持 'in' 操作中的值类型 %T\n", ruleItem.Val)
				allMatch = false
			}
		case ">=":
			//fmt.Printf("%+v %+v \n", ruleItem.Val, value)
			if cast.ToInt(value) >= cast.ToInt(ruleItem.Val) {

			} else {
				return false
			}

		case "<=":
			if cast.ToInt(value) <= cast.ToInt(ruleItem.Val) {

			} else {
				return false
			}
		case "=":
			if cast.ToInt(ruleItem.Val) == cast.ToInt(value) {

			} else {
				return false
			}
		case "<":
			if cast.ToInt(ruleItem.Val) < cast.ToInt(value) {

			} else {
				return false
			}
		case ">":
			if cast.ToInt(ruleItem.Val) > cast.ToInt(value) {

			} else {
				return false
			}
		default:
			fmt.Printf("不支持的运算符: %s\n", ruleItem.Op)
			allMatch = false
		}
	}
	return allMatch
}

func main() {
	jsonStr := `{
        "country": {
            "op": "in",
            "val": ["cn", "us"]
        },
        "version": {
            "op": ">=",
            "val": 88888
        },
        "age": {
            "op": "<=",
            "val": 30
        }
    }`

	var rule Rule
	err := json.Unmarshal([]byte(jsonStr), &rule)
	if err != nil {
		fmt.Println("解析 JSON 出错:", err)
		return
	}

	// 示例输入数据
	data := map[string]interface{}{
		"country": "cn",
		"version": 90000,
		"age":     28,
	}

	if matchRule(rule, data) {
		fmt.Println("规则匹配成功")
	} else {
		fmt.Println("规则匹配失败")
	}
}

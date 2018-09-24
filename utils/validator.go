package utils

import (
	// "encoding/json"
	// "github.com/valyala/fasthttp"
	"fmt"
	"strconv"
	// "encoding/binary"
	// "github.com/asaskevich/govalidator"
)

type RuleItem struct {
	Type     string
	Required bool
}

type Validator struct {
	Rules map[string]RuleItem
}

type Peekrable interface {
	Peek(string) []byte
}

type ServiceParams struct {
	CommonParams map[string]interface{}
	LikeQuery    string
	Limit        int
	Offset       int
}

func (v *Validator) Validate(peekrable Peekrable) (*ServiceParams, []string) {
	var (
		pageNumber int
		pageSize   int
		errors     []string
		data       *ServiceParams = &ServiceParams{}
		p          []byte
	)
	// map validator params type Mapper
	for param, ruleItem := range v.Rules {

		p = peekrable.Peek(param)
		// fmt.Println(data)
		// fmt.Println(param, string(peekrable.Peek(param)))

		if ruleItem.Required == true && p == nil {
			errors = append(errors, param+" is required")
			continue
		}

		if p != nil {
			switch ruleItem.Type {
			case "pageNumber":
				if IsPageNumber(ruleItem.Type) {
					pageNumber, _ = strconv.Atoi(string(p))
				} else {
					errors = append(errors, param+" is invalid")
				}
				break
			case "pageSize":
				if IsPageSize(ruleItem.Type) {
					pageSize, _ = strconv.Atoi(string(p))
				} else {
					errors = append(errors, param+" is invalid")
				}
				break
			case "likeQuery":
				data.LikeQuery = string(p)
				break
			case "binary":
				if IsBinary(ruleItem.Type) {
					if string(p) == "1" {
						// data.CommonParams[param] = true
					} else {
						// data.CommonParams[param] = false
					}
				} else {
					errors = append(errors, param+" is invalid")
				}
				break
			case "number":
				if IsNumber(ruleItem.Type) {
					data.CommonParams[param] = string(p)
				} else {
					errors = append(errors, param+" is invalid")
				}
				break
			case "phoneNum":
				if IsPhoneNum(ruleItem.Type) {
					data.CommonParams[param] = string(p)
				} else {
					errors = append(errors, param+" is invalid")
				}
				break
			}
		}
	}

	if pageSize != 0 {
		data.Limit = pageSize
	} else {
		data.Limit = 10
	}

	if pageNumber != 0 {
		data.Offset = data.Limit * (pageNumber - 1)
	} else {
		data.Offset = 0
	}

	// fmt.Println(data, errors)
	fmt.Println("")
	return data, errors
}

func IsNumber(p string) bool {
	return true
}

func IsPageNumber(p interface{}) bool {
	return true
}

func IsPageSize(p interface{}) bool {
	return true
}

func IsBinary(p interface{}) bool {
	return true
}

func IsPhoneNum(p interface{}) bool {
	return true
}

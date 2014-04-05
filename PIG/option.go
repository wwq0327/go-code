package main

import (
	"fmt"
)

type Optioner interface {
	Name() string
	IsValid() bool
}

type OptionCommon struct {
	ShortName string "short option name"
	LongName  string "long option name"
}

type IntOption struct {
	OptionCommon        // 匿名字段（嵌入）
	Value, Min, Max int // 具名字段（聚合）
}

type FloatOption struct {
	Optioner         // 匿名字段，接口嵌入，需要具体的类型, 满足这个接口的都可以接入
	Value    float64 // 具名字段（聚合）
}

type StringOption struct {
	OptionCommon
	Value string
}

type GenericOption struct {
	OptionCommon
}

func (option IntOption) Name() string {
	return name(option.ShortName, option.LongName)
}

func (option IntOption) IsValid() bool {
	return option.Min <= option.Value && option.Value <= option.Max
}

func (option StringOption) Name() string {
	return name(option.ShortName, option.LongName)
}

func (option StringOption) IsValid() bool {
	if option.Value != "" {
		return true
	}
	return false
}

func name(shortName, longName string) string {
	if longName == "" {
		return shortName
	}
	return longName
}

func (option GenericOption) Name() string {
	return name(option.ShortName, option.LongName)
}

func (option GenericOption) IsValid() bool {
	return true
}

func main() {
	fileOption := StringOption{OptionCommon{"f", "file"}, "index.html"}
	topOption := IntOption{OptionCommon: OptionCommon{"t", "top"}, Max: 100}
	floatOption := FloatOption{GenericOption{OptionCommon{"s", "size"}}, 19.5}
	float2 := FloatOption{StringOption{OptionCommon{"F", "File"}, "post.html"}, 12.2}

	for _, option := range []Optioner{topOption, fileOption, floatOption, float2} {
		fmt.Print("name=", option.Name(), "--valid=", option.IsValid(), "--value=")

		switch option := option.(type) {
		case IntOption:
			fmt.Println(option.Value, "--Min=", option.Min, "--Max=", option.Max)
		case StringOption:
			fmt.Println(option.Value)
		case FloatOption:
			fmt.Println(option.Value)
		}
	}
}

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
	OptionCommon
	Value, Min, Max int
}

type FloatOption struct {
	Optioner
	Value float64
}

type GenericOption struct {
	OptionCommon
}


func (option IntOption) Name() string {

}

func (option IntOption) IsValid() bool {

}

func name(shortName, longName string) string {

}

func (option GenericOption) Name() string {

}

func (option GenericOption) IsValid() bool {

}

func main() {

}

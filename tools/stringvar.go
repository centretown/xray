package tools

import (
	"strconv"
	"strings"
)

// StringSlice - a custom type
type StringSlice []string

// Set method for StringSlice
func (ss *StringSlice) Set(value string) error {
	*ss = append(*ss, value)
	return nil
}

// String method for StringSlice
func (ss *StringSlice) String() string {
	return strings.Join(*ss, ", ")
}

type IntSlice []int

func (is *IntSlice) Set(value string) error {
	var (
		v   int
		err error
	)

	v, err = strconv.Atoi(value)
	if err == nil {
		*is = append(*is, v)
		return err
	}
	return err
}

func (is *IntSlice) String() string {
	var ss []string = make([]string, len(*is))
	for i, v := range *is {
		ss[i] += strconv.Itoa(v)
	}
	return strings.Join(ss, ", ")
}

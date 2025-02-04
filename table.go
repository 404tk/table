// Copyright 2023 404tk. All rights reserved.
// license that can be found in the LICENSE file.

// Package table produces a string that represents slice of structs data in a text table
package table

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// Output formats slice of structs data and writes to standard output.
func Output(slice interface{}) {
	coln, rows, err := parse(slice)
	if err != nil {
		log.Println("[-]", err)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(coln)
	table.SetAutoFormatHeaders(false)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.AppendBulk(rows)
	table.Render()
}

// Table formats slice of structs data and returns the resulting string.
func Table(slice interface{}) string {
	coln, rows, err := parse(slice)
	if err != nil {
		log.Println("[-]", err)
	}
	var b strings.Builder
	table := tablewriter.NewWriter(&b)
	table.SetHeader(coln)
	table.SetAutoFormatHeaders(false)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.AppendBulk(rows)
	table.Render()
	return b.String()
}

func FileOutput(filename string, slice interface{}) {
	coln, rows, err := parse(slice)
	if err != nil {
		log.Println("[-]", err)
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println("[-]", err)
	}
	defer file.Close()
	table := tablewriter.NewWriter(file)
	table.SetHeader(coln)
	table.SetAutoFormatHeaders(false)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.AppendBulk(rows)
	table.Render()
}

func parse(slice interface{}) (
	coln []string, // name of columns
	rows [][]string, // rows of content
	err error,
) {
	s, err := sliceconv(slice)
	if err != nil {
		return
	}
	check := make(map[string]int)
	var _rows []map[int]string
	for _, u := range s {
		v := reflect.ValueOf(u)
		t := reflect.TypeOf(u)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
			t = t.Elem()
		}
		if v.Kind() != reflect.Struct {
			err = errors.New("warning: table: items of slice should be on struct value")
			return
		}
		row := make(map[int]string)
		index := 0
		for n := 0; n < v.NumField(); n++ {
			if t.Field(n).PkgPath != "" {
				continue
			}
			cn := t.Field(n).Name
			ct := t.Field(n).Tag.Get("table")
			if ct == "" {
				ct = cn
			} else if ct == "-" {
				continue
			}
			cv := fmt.Sprintf("%+v", v.FieldByName(cn).Interface())
			if len(cv) > 0 {
				if !isContain(coln, ct) {
					coln = append(coln, ct)
				}
				check[ct] = index
			}

			strSlice := strings.Split(cv, "\n")
			cv = ""
			for _, s := range strSlice {
				if len(s) > 40 {
					cv += stringWrap(s, 40) + "\n"
				} else {
					cv += s + "\n"
				}
			}
			cv = strings.TrimRight(cv, "\n")

			row[index] = cv
			index += 1
		}
		_rows = append(_rows, row)
	}
	for _, r := range _rows {
		var row []string
		for _, name := range coln {
			index := check[name]
			row = append(row, r[index])
		}
		rows = append(rows, row)
	}
	return coln, rows, nil
}

func sliceconv(slice interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("warning: sliceconv: param \"slice\" should be on slice value")
	}

	l := v.Len()

	r := make([]interface{}, l)
	for i := 0; i < l; i++ {
		r[i] = v.Index(i).Interface()
	}
	return r, nil
}

func stringWrap(s string, limit int) string {
	strSlice := strings.Split(s, "")
	result := ""

	for len(strSlice) > 0 {
		if len(strSlice) >= limit {
			result += strings.Join(strSlice[:limit], "") + "\n"
			strSlice = strSlice[limit:]
		} else {
			length := len(strSlice)
			result += strings.Join(strSlice[:length], "")
			strSlice = []string{}
		}
	}

	return result
}

func isContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

package common

import (
	"regexp"
	"strings"
)

// snake string, XxYy to xx_yy , XxYY to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if !k && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// 首字母大写
func UpperFirst(s string) string {
	data := make([]byte, 0, len(s))
	d := s[0]
	if d >= 'a' && d <= 'z' {
		d = d - 32
	}
	sByte := []byte(s)
	data = append(data, d)
	data = append(data, sByte[1:]...)
	return string(data)
}

// 首字母小写
func LowerFirst(s string) string {
	data := make([]byte, 0, len(s))
	d := s[0]
	if d >= 'A' && d <= 'Z' {
		d = d + 32
	}
	sByte := []byte(s)
	data = append(data, d)
	data = append(data, sByte[1:]...)
	return string(data)
}

func Upper(s string) string {
	data := make([]byte, 0, len(s))
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if d >= 'a' && d <= 'z' {
			d = d - 32
		}
		data = append(data, d)
	}
	return string(data[:])
}

// 验证必须是字母
func MustAlpha(s string) bool {
	regex := regexp.MustCompile("^[A-Za-z]+$")
	return regex.MatchString(s)
}


func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

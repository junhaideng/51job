package utils

import "github.com/axgle/mahonia"

func Convert(src []byte) []byte {
	decoder := mahonia.NewDecoder("gbk")
	return []byte(decoder.ConvertString(string(src)))
}
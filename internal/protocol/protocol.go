package protocol

import (
	"bufio"
	"bytes"
	"errors"
	"strconv"
)

// ErrProtocolEncodeEmpty empty cmd
var (
	ErrProtocolEncodeEmpty = errors.New("encode empty cmd")
)

const (
	// ReplyStatus 状态回复
	ReplyStatus = "+"
	// ReplyError 错误回复
	ReplyError = "-"
	// ReplyInteger 整数回复
	ReplyInteger = ":"
	// ReplyBulk 批量回复
	ReplyBulk = "$"
	// ReplyMultiBulk 多条批量回复
	ReplyMultiBulk = "*"
	// NewLine 换行
	NewLine = "\r\n"
	// LineR 换行R
	LineR = '\r'
	// LineN 换行N
	LineN = '\n'
	// Space 分隔符
	Space = " "
)

// MarshalStr 字符串编码
func MarshalStr(cmd string) ([]byte, error) {
	return MarshalBytesArray(ParseCmdStr(cmd)...)
}

// ParseCmdStr 命令行字符串解析
func ParseCmdStr(cmd string) [][]byte {
	runes := []rune(cmd)
	args := make([][]byte, 0)
	arg := ""
	quotation := false
	disableSpaceFlag := false
	for _, v := range runes {
		s := string(v)
		if s == "\"" {
			if quotation == false {
				quotation = true
				disableSpaceFlag = true
			} else {
				quotation = false
				args = append(args, []byte(arg))
				arg = ""
			}
		} else {
			if !disableSpaceFlag && s == " " {
				args = append(args, []byte(arg))
				arg = ""
			} else {
				arg = arg + s
			}

		}
	}
	args = append(args, []byte(arg))
	return args
}

// MarshalBytes 字节数组编码
func MarshalBytes(cmd []byte) ([]byte, error) {
	r := bytes.Split(cmd, []byte(Space))
	if r == nil || len(r) == 0 {
		return nil, ErrProtocolEncodeEmpty
	}
	var buffer bytes.Buffer
	var reqBuffer bytes.Buffer
	argsLen := 0
	for _, v := range r {
		if len(v) > 0 {
			argsLen++
			argLen := len(v)
			buffer.WriteString(ReplyBulk + strconv.Itoa(argLen) + NewLine)
			buffer.Write(v)
			buffer.WriteString(NewLine)
		}
	}
	reqBuffer.WriteString(ReplyMultiBulk + strconv.Itoa(argsLen) + NewLine)
	reqBuffer.Write(buffer.Bytes())
	return reqBuffer.Bytes(), nil
}

// MarshalReplyBulkBytes 批量回复
func MarshalReplyBulkBytes(cmd []byte) ([]byte, error) {
	var respBuffer bytes.Buffer
	length := len(cmd)
	if cmd == nil || len(cmd) == 0 {
		length = -1
		respBuffer.WriteString(ReplyBulk + strconv.Itoa(length) + NewLine)
	} else {
		respBuffer.WriteString(ReplyBulk + strconv.Itoa(length) + NewLine)
		respBuffer.Write(cmd)
		respBuffer.WriteString(NewLine)
	}
	return respBuffer.Bytes(), nil
}

// MarshalBytesArray 多命令参数
func MarshalBytesArray(cmds ...[]byte) ([]byte, error) {
	var buffer bytes.Buffer
	var reqBuffer bytes.Buffer
	argsLen := 0
	for _, v := range cmds {
		if len(v) > 0 {
			argsLen++
			argLen := len(v)
			buffer.WriteString(ReplyBulk + strconv.Itoa(argLen) + NewLine)
			buffer.Write(v)
			buffer.WriteString(NewLine)
		}
	}
	reqBuffer.WriteString(ReplyMultiBulk + strconv.Itoa(argsLen) + NewLine)
	reqBuffer.Write(buffer.Bytes())
	return reqBuffer.Bytes(), nil
}

// UnMarshal 字节数组解码
func UnMarshal(reader *bufio.Reader) [][]byte {
	fieldsSize := 0
	readFiled := 0
	args := make([][]byte, 0)
	for {
		feild, _ := reader.ReadBytes(LineN)
		// fmt.Println(string(feild))
		if len(feild) > 2 && feild[len(feild)-1] == LineN && feild[len(feild)-2] == LineR {
			firstFlag := string(feild[0])
			switch firstFlag {
			case ReplyMultiBulk:
				fieldsSize = ReadRealLen(feild)
			case ReplyBulk:
				argLen := ReadRealLen(feild)
				if argLen > 0 {
					arg := ReadArgByReader(reader, argLen)
					args = append(args, arg)
					readFiled++
				}
			case ReplyStatus, ReplyError, ReplyInteger:
				arg := ReadSingleLineReply(feild)
				args = append(args, arg)
				break
			}
		}
		if readFiled >= fieldsSize {
			break
		}
	}
	return args
}

// UnMarshalBytes 字节数组解码
func UnMarshalBytes(cmd []byte) [][]byte {
	buff := bytes.NewBuffer(cmd)
	fieldsSize := 0
	readFiled := 0
	args := make([][]byte, 0)
	for {
		feild, _ := buff.ReadBytes(LineN)
		if len(feild) > 2 && feild[len(feild)-1] == LineN && feild[len(feild)-2] == LineR {
			firstFlag := string(feild[0])
			switch firstFlag {
			case ReplyMultiBulk:
				fieldsSize = ReadRealLen(feild)
			case ReplyBulk:
				argLen := ReadRealLen(feild)
				arg := ReadArg(buff, argLen)
				args = append(args, arg)
				readFiled++
			case ReplyStatus, ReplyError:
				arg := ReadSingleLineReply(feild)
				args = append(args, arg)
				break
			}
		}
		if readFiled >= fieldsSize {
			break
		}
	}
	return args
}

// ReadRealLen 读取真实数据长度
func ReadRealLen(feild []byte) int {
	v := feild[1 : len(feild)-2]
	s, _ := strconv.Atoi(string(v))
	return s
}

// ReadArg 读取指定长度字节数据
func ReadArg(buff *bytes.Buffer, size int) []byte {
	data := make([]byte, size)
	line := make([]byte, 2)
	buff.Read(data)
	buff.Read(line)
	return data
}

// ReadArgByReader reader读取指定字节
func ReadArgByReader(reader *bufio.Reader, size int) []byte {
	data := make([]byte, size)
	line := make([]byte, 2)
	reader.Read(data)
	reader.Read(line)
	return data
}

// ReadSingleLineReply 读取单行回复，包含状态回复和错误回复
func ReadSingleLineReply(data []byte) []byte {
	return data[1 : len(data)-2]
}

// MarshalIntergerReply 返回整型编码数据
func MarshalIntergerReply(length int) []byte {
	n := IntToBytes(length)
	buf := make([]byte, 0)
	buffer := bytes.NewBuffer(buf)
	buffer.WriteString(ReplyInteger)
	buffer.Write(n)
	buffer.WriteString(NewLine)
	return buffer.Bytes()
}

// IntToBytes int类型转bytes
func IntToBytes(n int) []byte {
	s := strconv.Itoa(n)
	return []byte(s)
}

// BytesToInt bytes转int类型
func BytesToInt(b []byte) (int, error) {
	s := string(b)
	return strconv.Atoi(s)
}

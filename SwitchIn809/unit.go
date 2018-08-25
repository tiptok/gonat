package SwitchIn809

import (
	"bytes"
	"errors"
)

func Byte809Descape(value []byte, startIndex int, length int) ([]byte, error) {
	ilength := len(value)
	if (startIndex + length) > ilength {
		return nil, errors.New("Byte809Descape 长度不足，下标越界")
	}
	buf := new(bytes.Buffer)
	/*去头去尾*/
	for i := startIndex + 1; i < ilength-1; i++ {
		if value[i] == 0x5a {
			if value[i+1] == 0x01 {
				buf.WriteByte(0x5b)
			} else if value[i+1] == 0x02 {
				buf.WriteByte(0x5a)
			}
			i++
		} else if value[i] == 0x5e {
			if value[i+1] == 0x01 {
				buf.WriteByte(0x5d)
			} else if value[i+1] == 0x02 {
				buf.WriteByte(0x5e)
			}
			i++
		} else {
			if i != ilength-1 {
				buf.WriteByte(value[i])
			}
			// if value[i] == 0x5d { //到达每个个包的最后一位 退出后面的字节不做追加
			// 	break
			// }
		}

	}
	return buf.Bytes(), nil
}

func Byte809Enscape(value []byte, startIndex int, length int) []byte {
	ilength := len(value)
	if (startIndex + length) > ilength {

	}
	buf := new(bytes.Buffer)
	buf.WriteByte(0x5b)
	for i := startIndex; i < ilength; i++ {
		if value[i] == 0x5b {
			buf.WriteByte(0x5a)
			buf.WriteByte(0x01)
		} else if value[i] == 0x5a {
			buf.WriteByte(0x5a)
			buf.WriteByte(0x02)
		} else if value[i] == 0x5d {
			buf.WriteByte(0x5e)
			buf.WriteByte(0x01)
		} else if value[i] == 0x5e {
			buf.WriteByte(0x5e)
			buf.WriteByte(0x02)
		} else {
			buf.WriteByte(value[i])
		}
	}
	buf.WriteByte(0x5d)
	return buf.Bytes()
}

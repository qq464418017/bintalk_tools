package btk

import (
	"bytes"
	"encoding/binary"
)

type ProtocolWriter struct {
	Bytes []byte
}

func (w *ProtocolWriter) Write(data []byte) {
	for _, v := range data {
		w.Bytes = append(w.Bytes, v)
	}
}

func (pw *ProtocolWriter) WriteData(a interface{}) {
	switch t := a.(type) {
	case string:
		pw.WriteString(t)
	case []byte:
		pw.WriteBinary(t)
	case bool, int8, uint8, int16, uint16, int32, uint32, float32, int64, uint64, float64:
		pw.WriteSingle(a)
	case []int8:
		//pw.WriteArray(t, len(t))
		pw.WriteDynSize(len(t))
		for i := 0; i < len(t); i++ {
			pw.WriteSingle(t[i])
		}
	case []int16:
		pw.WriteDynSize(len(t))
		for i := 0; i < len(t); i++ {
			pw.WriteSingle(t[i])
		}
	case []uint16:
		pw.WriteDynSize(len(t))
		for i := 0; i < len(t); i++ {
			pw.WriteSingle(t[i])
		}
	case []int32:
		pw.WriteDynSize(len(t))
		for i := 0; i < len(t); i++ {
			pw.WriteSingle(t[i])
		}
	case []uint32:
		pw.WriteDynSize(len(t))
		for i := 0; i < len(t); i++ {
			pw.WriteSingle(t[i])
		}
	case []int64:
		pw.WriteDynSize(len(t))
		for i := 0; i < len(t); i++ {
			pw.WriteSingle(t[i])
		}
	case []uint64:
		pw.WriteDynSize(len(t))
		for i := 0; i < len(t); i++ {
			pw.WriteSingle(t[i])
		}
	case []string:
		pw.WriteDynSize(len(t))
		for i := 0; i < len(t); i++ {
			pw.WriteString(t[i])
		}
	case []float32:
		pw.WriteDynSize(len(t))
		for i := 0; i < len(t); i++ {
			pw.WriteSingle(t[i])
		}
	case []float64:
		pw.WriteDynSize(len(t))
		for i := 0; i < len(t); i++ {
			pw.WriteSingle(t[i])
		}
	}
}

func (pw *ProtocolWriter) WriteSingle(a interface{}) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, a)
	if err == nil {
		pw.Write(buf.Bytes())
	}
}

//string
func (pw *ProtocolWriter) WriteString(v string) {
	if v == "" || len(v) == 0 {
		pw.WriteDynSize(0)
	} else {
		str := []byte(v)
		len := len(str)
		pw.WriteDynSize(len)
		if len > 0 {
			pw.Write(str)
		}
	}
}

//binary
func (pw *ProtocolWriter) WriteBinary(v []byte) {
	if v == nil || len(v) == 0 {
		pw.WriteDynSize(0)
	} else {
		s := len(v)
		pw.WriteDynSize(s)
		pw.Write(v)
	}
}

//dynamic
func (pw *ProtocolWriter) WriteDynSize(s int) {
	//bytesBuffer := new(bytes.Buffer)
	//err := binary.Write(bytesBuffer, binary.LittleEndian, int32(s))
	//if err == nil {
	//
	//}
	//b := bytesBuffer.Bytes()
	l := s
	//n := 0
	if s <= 0x3F {
		pw.WriteData((byte)(l & 0xFF))
	} else if s <= 0x3FFF {
		l |= 1 << 14
		pw.WriteData((byte)((l >> 8) & 0xFF))
		pw.WriteData((byte)(l & 0xFF))
	} else if s <= 0x3FFFFF {
		l |= 2 << 22
		pw.WriteData((byte)((l >> 16) & 0xFF))
		pw.WriteData((byte)((l >> 8) & 0xFF))
		pw.WriteData((byte)(l & 0xFF))
	} else if s <= 0x3FFFFFFF {
		l |= 3 << 30
		pw.WriteData((byte)((l >> 24) & 0xFF))
		pw.WriteData((byte)((l >> 16) & 0xFF))
		pw.WriteData((byte)((l >> 8) & 0xFF))
		pw.WriteData((byte)(l & 0xFF))
	}
	//b[n] |= byte(n << 6)
	//for i := n; i >= 0; i--{
	//	pw.WriteData(b[i])
	//}
}

func intDataSize(data interface{}) int {
	switch data := data.(type) {
	case bool, int8, uint8, *bool, *int8, *uint8:
		return 1
	case []int8:
		return len(data)
	case []uint8:
		return len(data)
	case int16, uint16, *int16, *uint16:
		return 2
	case []int16:
		return 2 * len(data)
	case []uint16:
		return 2 * len(data)
	case int32, uint32, *int32, *uint32, float32, *float32:
		return 4
	case []int32:
		return 4 * len(data)
	case []uint32:
		return 4 * len(data)
	case int64, uint64, *int64, *uint64, float64, *float64:
		return 8
	case []int64:
		return 8 * len(data)
	case []uint64:
		return 8 * len(data)
	}
	return 0
}

func Read(buff []byte, startid *int, a interface{}) bool {
	switch a.(type) {
	case *string:
		return ReadString(buff, startid, a)
	case string:
		return ReadString(buff, startid, a)
	case *[]byte:
		return ReadBytes(buff, startid, a)
	case bool, int8, uint8, *bool, *int8, *uint8, int16, uint16, *int16, *uint16,
		int32, uint32, *int32, *uint32, float32, *float32, int64, uint64,
		*int64, *uint64, float64, *float64:
		return ReadSignle(buff, startid, a)
	case *[]int8, *[]int16, *[]uint16, *[]int32, *[]uint32, *[]int64, *[]uint64:
		return ReadArray(buff, startid, a)
	case *[]string:
		return ReadArray(buff, startid, a)
	}

	return false
}

func ReadSignle(buff []byte, startid *int, a interface{}) bool {
	size := intDataSize(a)
	if size == 0 {
		if !ReadDynSize(buff, startid, &size) || size > len(buff) {
			return false
		}
	}
	end := *startid + size
	if end <= len(buff) {
		buf := bytes.NewBuffer(buff[*startid:end])
		binary.Read(buf, binary.LittleEndian, a)
		*startid = end
		return true
	}
	return false
}

func ReadArray(buff []byte, startid *int, a interface{}) bool {
	var s int
	if !ReadDynSize(buff, startid, &s) || s > len(buff) {
		return false
	}
	switch t := a.(type) {
	case *[]int8:
		v := make([]int8, s)
		for i := 0; i < s; i++ {
			var b int8
			ReadSignle(buff, startid, &b)
			v[i] = b
		}
		*t = v
		return true
	case *[]uint8:
		v := make([]uint8, s)
		for i := 0; i < s; i++ {
			var b uint8
			ReadSignle(buff, startid, &b)
			v[i] = b
		}
		*t = v
		return true
	case *[]int16:
		v := make([]int16, s)
		for i := 0; i < s; i++ {
			var b int16
			ReadSignle(buff, startid, &b)
			v[i] = b
		}
		*t = v
		return true
	case *[]uint16:
		v := make([]uint16, s)
		for i := 0; i < s; i++ {
			var b uint16
			ReadSignle(buff, startid, &b)
			v[i] = b
		}
		*t = v
		return true
	case *[]int32:
		v := make([]int32, s)
		for i := 0; i < s; i++ {
			var b int32
			ReadSignle(buff, startid, &b)
			v[i] = b
		}
		*t = v
		return true
	case *[]uint32:
		v := make([]uint32, s)
		for i := 0; i < s; i++ {
			var b uint32
			ReadSignle(buff, startid, &b)
			v[i] = b
		}
		*t = v
		return true
	case *[]int64:
		v := make([]int64, s)
		for i := 0; i < s; i++ {
			var b int64
			ReadSignle(buff, startid, &b)
			v[i] = b
		}
		a = v
		*t = v
	case *[]uint64:
		v := make([]uint64, s)
		for i := 0; i < s; i++ {
			var b uint64
			ReadSignle(buff, startid, &b)
			v[i] = b
		}
		*t = v
		return true
	case *[]string:
		v := make([]string, s)
		for i := 0; i < s; i++ {
			var b string
			ReadString(buff, startid, &b)
			v[i] = b
		}
		*t = v
		return true
	}
	return false
}

//string
func ReadString(buff []byte, startid *int, a interface{}) bool {
	var s int
	if !ReadDynSize(buff, startid, &s) || s > len(buff) {
		return false
	}
	end := *startid + s
	switch a := a.(type) {
	case *string:
		v := buff[*startid:end]
		*a = string(v)
		*startid = end
	default:
		return false
	}

	return true
}

//bytes
func ReadBytes(buff []byte, startid *int, a interface{}) bool {
	var s int
	if !ReadDynSize(buff, startid, &s) || s > len(buff) {
		return false
	}
	end := *startid + s
	switch a := a.(type) {
	case *[]byte:
		*a = buff[*startid:end]
		*startid = end
	default:
		return false
	}

	return true
}

//dynamic
func ReadDynSize(buff []byte, startid *int, s *int) bool {
	var b byte
	if !Read(buff, startid, &b) {
		return false
	}
	b = b & 0xFF
	n := b >> 6
	size := int(b & 0x3F)
	//for i := 0; i < int(n); i++ {
	//	if !Read(buff,startid,b) {
	//		return false
	//	}
	//	*s = int((*s << 8) | int(b))
	//}

	for n > 0 {
		var b byte
		if !Read(buff, startid, &b) {
			return false
		}
		b = b & 0xFF
		size = (size << 8) | int(b)
		n--
	}
	*s = size
	return true
}

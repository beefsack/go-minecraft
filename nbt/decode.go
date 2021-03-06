package nbt

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

func Decode(r io.Reader) (map[string]interface{}, error) {
	br := bufio.NewReader(r)
	_, name, tagPayload, err := DecodeTag(br)
	// Return a map with a single root tag
	return map[string]interface{}{
		name: tagPayload,
	}, err
}

func DecodeTag(br *bufio.Reader) (tagType int, name string, payload interface{},
	err error) {
	tagType, err = ReadTagType(br)
	if err != nil {
		return
	}
	if tagType != TAG_CLOSE {
		name, err = ReadTagName(br)
		if err != nil {
			return
		}
		payload, err = DecodePayload(tagType, br)
	}
	return
}

func ReadTagType(br *bufio.Reader) (int, error) {
	tagTypeByte, err := br.ReadByte()
	return int(tagTypeByte), err
}

func ReadTagName(br *bufio.Reader) (string, error) {
	// Read the tag length
	tagLength, err := DecodeShort(br)
	if err != nil {
		return "", err
	}
	// Read the tag name
	tagBytes := make([]byte, tagLength)
	_, err = br.Read(tagBytes)
	if err != nil {
		return "", err
	}
	return string(tagBytes), nil
}

func DecodePayload(tagType int, br *bufio.Reader) (payload interface{},
	err error) {
	switch tagType {
	case TAG_CLOSE:
	case TAG_BYTE:
		payload, err = DecodeByte(br)
	case TAG_SHORT:
		payload, err = DecodeShort(br)
	case TAG_INT:
		payload, err = DecodeInt(br)
	case TAG_LONG:
		payload, err = DecodeLong(br)
	case TAG_FLOAT:
		payload, err = DecodeFloat(br)
	case TAG_DOUBLE:
		payload, err = DecodeDouble(br)
	case TAG_BYTE_ARRAY:
		payload, err = DecodeByteArray(br)
	case TAG_STRING:
		payload, err = DecodeString(br)
	case TAG_LIST:
		payload, err = DecodeList(br)
	case TAG_COMPOUND:
		payload, err = DecodeCompound(br)
	case TAG_INT_ARRAY:
		payload, err = DecodeIntArray(br)
	default:
		err = errors.New(
			fmt.Sprintf("Unknown type: %d", tagType))
	}
	return
}

func DecodeByte(br *bufio.Reader) (b byte, err error) {
	err = binary.Read(br, binary.BigEndian, &b)
	return
}

func DecodeShort(br *bufio.Reader) (i int16, err error) {
	err = binary.Read(br, binary.BigEndian, &i)
	return
}

func DecodeInt(br *bufio.Reader) (i int32, err error) {
	err = binary.Read(br, binary.BigEndian, &i)
	return
}

func DecodeLong(br *bufio.Reader) (i int64, err error) {
	err = binary.Read(br, binary.BigEndian, &i)
	return
}

func DecodeFloat(br *bufio.Reader) (f float32, err error) {
	err = binary.Read(br, binary.BigEndian, &f)
	return
}

func DecodeDouble(br *bufio.Reader) (d float64, err error) {
	err = binary.Read(br, binary.BigEndian, &d)
	return
}

func DecodeByteArray(br *bufio.Reader) ([]byte, error) {
	length, err := DecodeInt(br)
	if err != nil {
		return nil, err
	}
	p := make([]byte, length)
	_, err = br.Read(p)
	return p, err
}

func DecodeString(br *bufio.Reader) (string, error) {
	length, err := DecodeShort(br)
	if err != nil {
		return "", err
	}
	strBytes := make([]byte, length)
	_, err = br.Read(strBytes)
	return string(strBytes), err
}

func DecodeList(br *bufio.Reader) (List, error) {
	tagType, err := ReadTagType(br)
	if err != nil {
		return List{}, err
	}
	length, err := DecodeInt(br)
	if err != nil {
		return List{}, err
	}
	list := List{
		ListType: tagType,
		Items:    make([]interface{}, length),
	}
	for i := int32(0); i < length; i++ {
		list.Items[i], err = DecodePayload(tagType, br)
		if err != nil {
			break
		}
	}
	return list, err
}

func DecodeCompound(br *bufio.Reader) (map[string]interface{}, error) {
	compound := map[string]interface{}{}
	// Keep reading tags until EOF or close
	for {
		tagType, name, payload, err := DecodeTag(br)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if tagType == TAG_CLOSE {
			break
		}
		compound[name] = payload
	}
	return compound, nil
}

func DecodeIntArray(br *bufio.Reader) ([]int32, error) {
	length, err := DecodeInt(br)
	if err != nil {
		return nil, err
	}
	intArray := make([]int32, length)
	var intVal int32
	for i := int32(0); i < length; i++ {
		err = binary.Read(br, binary.BigEndian, &intVal)
		if err != nil {
			break
		}
		intArray[i] = intVal
	}
	return intArray, err
}

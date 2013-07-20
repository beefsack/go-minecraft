package nbt

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

func Encode(data interface{}, w io.Writer) error {
	return EncodePayload(data, w)
}

func DetectPayloadType(payload interface{}) (tagType int, err error) {
	if _, ok := payload.(byte); ok {
		tagType = TAG_BYTE
	} else if _, ok := payload.(int16); ok {
		tagType = TAG_SHORT
	} else if _, ok := payload.(int32); ok {
		tagType = TAG_INT
	} else if _, ok := payload.(int64); ok {
		tagType = TAG_LONG
	} else if _, ok := payload.(float32); ok {
		tagType = TAG_FLOAT
	} else if _, ok := payload.(float64); ok {
		tagType = TAG_DOUBLE
	} else if _, ok := payload.([]byte); ok {
		tagType = TAG_BYTE_ARRAY
	} else if _, ok := payload.(string); ok {
		tagType = TAG_STRING
	} else if _, ok := payload.(*List); ok {
		tagType = TAG_LIST
	} else if _, ok := payload.(map[string]interface{}); ok {
		tagType = TAG_COMPOUND
	} else if _, ok := payload.([]int32); ok {
		tagType = TAG_INT_ARRAY
	} else {
		err = errors.New(fmt.Sprintf("Not a valid type: %#v", payload))
	}
	return
}

func EncodeTag(name string, payload interface{}, w io.Writer) error {
	tagType, err := DetectPayloadType(payload)
	if err != nil {
		return err
	}
	err = EncodeTagType(tagType, w)
	if err != nil {
		return err
	}
	err = EncodeTagName(name, w)
	if err != nil {
		return err
	}
	return EncodePayloadOfType(tagType, payload, w)
}

func EncodeTagType(tagType int, w io.Writer) error {
	_, err := w.Write([]byte{byte(tagType)})
	return err
}

func EncodeTagName(name string, w io.Writer) error {
	byteName := []byte(name)
	l := int16(len(byteName))
	err := binary.Write(w, binary.BigEndian, l)
	if err != nil {
		return err
	}
	_, err = w.Write(byteName)
	return err
}

func EncodePayload(payload interface{}, w io.Writer) error {
	tagType, err := DetectPayloadType(payload)
	if err != nil {
		return err
	}
	return EncodePayloadOfType(tagType, payload, w)
}

func EncodePayloadOfType(tagType int, payload interface{},
	w io.Writer) (err error) {
	switch tagType {
	case TAG_BYTE:
		err = EncodeByte(payload, w)
	case TAG_SHORT:
		err = EncodeShort(payload, w)
	case TAG_INT:
		err = EncodeInt(payload, w)
	case TAG_LONG:
		err = EncodeLong(payload, w)
	case TAG_FLOAT:
		err = EncodeFloat(payload, w)
	case TAG_DOUBLE:
		err = EncodeDouble(payload, w)
	case TAG_BYTE_ARRAY:
		err = EncodeByteArray(payload, w)
	case TAG_STRING:
		err = EncodeString(payload, w)
	case TAG_LIST:
		err = EncodeList(payload, w)
	case TAG_COMPOUND:
		err = EncodeCompound(payload, w)
	case TAG_INT_ARRAY:
		err = EncodeIntArray(payload, w)
	default:
		err = errors.New(fmt.Sprintf("Unsupported tag type: %d", tagType))
	}
	return
}

func EncodeByte(payload interface{}, w io.Writer) error {
	b, ok := payload.(byte)
	if !ok {
		return errors.New("Not a byte")
	}
	_, err := w.Write([]byte{b})
	return err
}

func EncodeShort(payload interface{}, w io.Writer) error {
	i, ok := payload.(int16)
	if !ok {
		return errors.New("Not a short")
	}
	return binary.Write(w, binary.BigEndian, i)
}

func EncodeInt(payload interface{}, w io.Writer) error {
	i, ok := payload.(int32)
	if !ok {
		return errors.New("Not an int")
	}
	return binary.Write(w, binary.BigEndian, i)
}

func EncodeLong(payload interface{}, w io.Writer) error {
	i, ok := payload.(int64)
	if !ok {
		return errors.New("Not a long")
	}
	return binary.Write(w, binary.BigEndian, i)
}

func EncodeFloat(payload interface{}, w io.Writer) error {
	f, ok := payload.(float32)
	if !ok {
		return errors.New("Not a float")
	}
	return binary.Write(w, binary.BigEndian, f)
}

func EncodeDouble(payload interface{}, w io.Writer) error {
	f, ok := payload.(float64)
	if !ok {
		return errors.New("Not a double")
	}
	return binary.Write(w, binary.BigEndian, f)
}

func EncodeByteArray(payload interface{}, w io.Writer) error {
	p, ok := payload.([]byte)
	if !ok {
		return errors.New("Not a byte array")
	}
	l := int32(len(p))
	err := binary.Write(w, binary.BigEndian, l)
	if err != nil {
		return err
	}
	_, err = w.Write(p)
	return err
}

func EncodeString(payload interface{}, w io.Writer) error {
	s, ok := payload.(string)
	if !ok {
		return errors.New("Not a string")
	}
	byteString := []byte(s)
	l := int16(len(byteString))
	err := binary.Write(w, binary.BigEndian, l)
	if err != nil {
		return err
	}
	_, err = w.Write(byteString)
	return err
}

func EncodeList(payload interface{}, w io.Writer) error {
	list, ok := payload.(*List)
	if !ok {
		return errors.New("Not a list")
	}
	err := EncodeTagType(list.ListType, w)
	if err != nil {
		return nil
	}
	l := int32(len(list.Items))
	err = binary.Write(w, binary.BigEndian, l)
	if err != nil {
		return err
	}
	for _, item := range list.Items {
		err = EncodePayloadOfType(list.ListType, item, w)
		if err != nil {
			break
		}
	}
	return err
}

func EncodeCompound(payload interface{}, w io.Writer) error {
	compound, ok := payload.(map[string]interface{})
	if !ok {
		return errors.New("Not a compound")
	}
	for name, p := range compound {
		err := EncodeTag(name, p, w)
		if err != nil {
			return err
		}
	}
	return EncodeTagType(TAG_CLOSE, w)
}

func EncodeIntArray(payload interface{}, w io.Writer) error {
	p, ok := payload.([]int32)
	if !ok {
		return errors.New("Not an int array")
	}
	l := int32(len(p))
	err := binary.Write(w, binary.BigEndian, l)
	if err != nil {
		return err
	}
	for _, i := range p {
		err := binary.Write(w, binary.BigEndian, i)
		if err != nil {
			break
		}
	}
	return err
}

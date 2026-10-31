package entities

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"slices"
)

type SimpleHighwayTokenizer struct {
	privateToken []byte
	algorithm    []byte
}

func NewSimpleHighwayTokenizer(privateToken string) (SimpleHighwayTokenizer, error) {
	tokenizer := SimpleHighwayTokenizer{
		privateToken: []byte(privateToken),
		algorithm:    []byte("simple"),
	}

	err := func() (err error) {
		defer func() {
			p := recover()
			if p != nil {
				err = fmt.Errorf("incorrect tokenizer work: panic: %v", p)
			}
		}()

		t, err := tokenizer.Validate(tokenizer.New(-1))
		if err != nil {
			return fmt.Errorf("work check: %w", err)
		}

		if t != -1 {
			return fmt.Errorf("work check: incorrect time: %d", t)
		}

		return nil
	}()

	if err != nil {
		return tokenizer, err
	}

	return tokenizer, nil
}

func (tokenizer SimpleHighwayTokenizer) New(validUntil int64) string {
	buff := bytes.Buffer{}
	buff.Write(tokenizer.algorithm)
	buff.Write(binary.LittleEndian.AppendUint64([]byte{}, uint64(validUntil)))

	tmp := [512]byte{}
	hash := md5.Sum(tmp[:tokenizer._2to1(tmp[:], buff.Bytes(), tokenizer.privateToken)])
	buff.Write(hash[:])

	return base64.RawStdEncoding.EncodeToString(buff.Bytes())
}

func (tokenizer SimpleHighwayTokenizer) Validate(token string) (int64, error) {
	targetLen := len(tokenizer.algorithm) + 8 + 16

	dataArr := [128]byte{}
	raw := dataArr[:]

	n, err := base64.RawStdEncoding.Decode(raw, []byte(token))
	if err != nil {
		return 0, fmt.Errorf("invalid format: %w", err)
	}

	if n != targetLen {
		return 0, fmt.Errorf("invalid len: %w", err)
	}

	raw = raw[:n]

	found := bytes.HasPrefix(raw, tokenizer.algorithm)
	if !found {
		return 0, fmt.Errorf("invalid algorithm format")
	}

	unix := int64(binary.LittleEndian.Uint64(raw[len(tokenizer.algorithm):]))

	tmp := [512]byte{}
	hash := md5.Sum(tmp[:tokenizer._2to1(tmp[:], raw[:len(tokenizer.algorithm)+8], tokenizer.privateToken)])

	ok := slices.Equal(hash[:], raw[len(tokenizer.algorithm)+8:])
	if !ok {
		return 0, fmt.Errorf("invalid sign")
	}

	return unix, nil
}

func (tokenizer SimpleHighwayTokenizer) _2to1(dst, src1, src2 []byte) int {
	_ = dst[len(src1)+len(src2)]

	copy(dst, src1)
	copy(dst[len(src1):], src2)

	return len(src1) + len(src2)
}

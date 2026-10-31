package entities

import (
	"testing"
	"time"
)

func TestSimpleHighway(t *testing.T) {
	timeUnix := time.Now().Add(time.Minute).Unix()

	tokenizer, err := NewSimpleHighwayTokenizer("abc")
	if err != nil {
		t.Fatal(err)
	}
	newTimeUnix, err := tokenizer.Validate(tokenizer.New(timeUnix))
	if err != nil {
		t.Fatal(err)
	}

	if timeUnix != newTimeUnix {
		t.Fatal("invalid payload")
	}

	tokenizer2, err := NewSimpleHighwayTokenizer("aac")
	if err != nil {
		t.Fatal(err)
	}

	_, err = tokenizer2.Validate(tokenizer.New(timeUnix))
	if err == nil {
		t.Fatal("sign not working")
	}
}

func BenchmarkSimpleHighwayValidate(b *testing.B) {
	timeUnix := time.Now().Add(time.Minute).Unix()
	tokenizer, _ := NewSimpleHighwayTokenizer("abc")
	token := tokenizer.New(timeUnix)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		newTimeUnix, err := tokenizer.Validate(token)
		if err != nil {
			b.Fatal(err)
		}

		if timeUnix != newTimeUnix {
			b.Fatal("invalid payload")
		}
	}
}

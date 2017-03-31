package main

import (
	"testing"

	"github.com/ikawaha/kagome/tokenizer"
)

func TestMorphString(t *testing.T) {
	tests := []struct {
		input Morph
		want  string
	}{
		{
			input: MorphBOS,
			want:  "\tBOS,,,,,,,,",
		},
		{
			input: Morph{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
			want:  "0\t1,2,3,4,5,6,7,8,9",
		},
	}

	for _, test := range tests {
		ans := test.input.String()
		if ans != test.want {
			t.Errorf("(*Morph).String(&%+v) = %q, want %q\n",
				test.input, ans, test.want)
		}
	}
}

func TestMorphsIsEqual(t *testing.T) {
	tests := []struct {
		input [2]Morphs
		want  bool
	}{
		{
			input: [2]Morphs{{}, {}},
			want:  true,
		},
		{
			input: [2]Morphs{
				{{Surface: "a"}, {Surface: "b"}},
				{{Surface: "a"}, {Surface: "b"}},
			},
			want: true,
		},
		{
			input: [2]Morphs{
				{{Surface: "a"}},
				{{Surface: "a"}, {Surface: "b"}},
			},
			want: false,
		},
		{
			input: [2]Morphs{
				{{Surface: "a"}, {Surface: "b"}},
				{{Surface: "a"}, {Surface: "c"}},
			},
			want: false,
		},
	}

	for _, test := range tests {
		if test.input[0].IsEqual(test.input[1]) != test.want {
			t.Errorf("(%v).IsEqual(%v) != %v\n",
				test.input[0], test.input[1], test.want)
		}
	}
}

func TestNewMorphs(t *testing.T) {
	tests := []struct {
		input string
		want  Morphs
	}{
		{
			input: "",
			want:  Morphs{},
		},
		{
			input: "愛",
			want: Morphs{
				&Morph{"愛", "名詞", "一般", "*", "*", "*", "*", "愛", "アイ", "アイ"},
			},
		},
	}

	kagome := tokenizer.New()

	for _, test := range tests {
		tokens := kagome.Tokenize(test.input)
		ms := NewMorphs(tokens)
		if !ms.IsEqual(test.want) {
			t.Errorf("NewMorphs(%q) = %v, want %v\n",
				test.input, ms, test.want)
		}
	}
}
package lexer

import (
	"errors"
)

type Input struct {
	source []rune
	lens   int
	pos    int
}

var inputEofErr = errors.New("文件结束")

func newInput(source string) *Input {
	sourceRunes := []rune(source)
	return &Input{sourceRunes, len(sourceRunes), -1}
}

func (input *Input) nextRune() (r rune, err error) {
	if err = input.advance(1); err != nil {
		return
	}

	r = input.source[input.pos]

	return
}

// 步进
func (input *Input) advance(num int) (err error) {
	if input.pos >= input.lens-num {
		err = inputEofErr
		return
	}

	input.pos = input.pos + num

	return
}

// 回退
func (input *Input) back(num int) (err error) {
	if input.pos-num < -1 {
		err = inputEofErr
		return
	}

	input.pos = input.pos - num

	return
}

// 往前查询num个字符，但不步进
func (input *Input) lookahead(num int) (runes []rune, err error) {
	if input.pos >= input.lens-num {
		err = inputEofErr
		return
	}

	for i := 0; i < num; i++ {
		runes = append(runes, input.source[input.pos+1+i])
	}

	return
}

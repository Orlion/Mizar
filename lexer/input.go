package lexer

import (
	"errors"
)

type Input struct {
	source            []rune
	lens              int
	pos               int
	lastLineColumnNum int // 上一行的列数
	LineNum           int // 当前所在行数
	ColumnNum         int // 当前所在列数
}

var inputEofErr = errors.New("文件结束")

func newInput(source string) *Input {
	sourceRunes := []rune(source)
	return &Input{sourceRunes, len(sourceRunes), -1, 1, 1, 1}
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

	// 检测步进过程中是否遇到了换行符，如果遇到换行符则将当前行数+1, 并修改当前列数
	for i := 1; i <= num; i++ {
		if input.source[input.pos+i] == 10 {
			input.LineNum++
			input.lastLineColumnNum = input.ColumnNum
			input.ColumnNum = 1
		} else {
			input.ColumnNum++
		}
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

	// 检测步进过程中是否遇到了换行符，如果遇到换行符则将当前行数+1, 并修改当前列数
	for i := 1; i < num; i++ {
		if input.pos-i != -1 && input.source[input.pos-i] == 10 {
			input.LineNum--
			input.ColumnNum = input.lastLineColumnNum
		} else {
			input.ColumnNum--
		}
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

// 判断输入流是否已结束
func (input *Input) isEof() bool {
	return input.pos >= input.lens-1
}

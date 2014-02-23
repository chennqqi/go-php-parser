package php

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type lexer struct {
	lastPos   int
	pos       int
	line      int
	start     int
	lastStart int
	width     int
	input     string
	file      string
	items     chan Item // channel of scanned items.
}

func newLexer(input string) *lexer {
	l := &lexer{
		line:  1,
		input: input,
		items: make(chan Item),
	}
	go l.run()
	return l
}

// stateFn represents the state of the scanner
// as a function that returns the next state.
type stateFn func(*lexer) stateFn

// run lexes the input by executing state functions until
// the state is nil.
func (l *lexer) run() {
	for state := lexHTML; state != nil; {
		state = state(l)
	}
	close(l.items) // No more tokens will be delivered.
}

func (l *lexer) emit(t ItemType) {
	i := Item{t, l.currentLocation(), l.input[l.start:l.pos]}
	l.incrementLines()
	l.items <- i
	l.start = l.pos
}

func (l *lexer) currentLocation() Location {
	return Location{Pos: l.start, Line: l.line, File: l.file}
}

// nextItem returns the next item from the input.
func (l *lexer) nextItem() Item {
	Item := <-l.items
	l.lastPos = Item.pos.Pos
	return Item
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	return r
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) skipSpace() {
	r := l.next()
	for isSpace(r) {
		r = l.next()
	}
	l.backup()
	l.ignore()
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	i := Item{itemError, l.currentLocation(), fmt.Sprintf(format, args...)}
	l.incrementLines()
	l.items <- i
	return nil
}

func (l *lexer) incrementLines() {
	l.line += strings.Count(l.input[l.lastStart:l.pos], "\n")
	l.lastStart = l.pos
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

func isKeyword(i ItemType) bool {
	is, ok := keywordMap[i]
	return is && ok
}

var keywordMap = map[ItemType]bool{
	itemFunction: true,

	itemReturn: true,
	itemEcho:   true,

	itemIf:      true,
	itemElse:    true,
	itemElseIf:  true,
	itemFor:     true,
	itemForeach: true,
	itemWhile:   true,
	itemDo:      true,

	itemTry:     true,
	itemCatch:   true,
	itemFinally: true,

	itemClass:       true,
	itemPrivate:     true,
	itemProtected:   true,
	itemPublic:      true,
	itemInterface:   true,
	itemImplements:  true,
	itemExtends:     true,
	itemNewOperator: true,

	itemInstanceofOperator: true,
	itemArray:              true,
}

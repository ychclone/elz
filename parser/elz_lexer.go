// Generated from Elz.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 35, 222,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3,
	28, 3, 28, 3, 29, 3, 29, 3, 30, 6, 30, 171, 10, 30, 13, 30, 14, 30, 172,
	3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 181, 10, 31, 12, 31, 14,
	31, 184, 11, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 7, 32, 192,
	10, 32, 12, 32, 14, 32, 195, 11, 32, 3, 33, 3, 33, 3, 34, 3, 34, 5, 34,
	201, 10, 34, 3, 35, 3, 35, 7, 35, 205, 10, 35, 12, 35, 14, 35, 208, 11,
	35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 7, 38, 216, 10, 38, 12, 38,
	14, 38, 219, 11, 38, 3, 38, 3, 38, 4, 182, 217, 2, 39, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63,
	33, 65, 2, 67, 2, 69, 34, 71, 2, 73, 2, 75, 35, 3, 2, 6, 5, 2, 11, 12,
	15, 15, 34, 34, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2, 48,
	48, 50, 59, 2, 223, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2,
	2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2,
	2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3,
	2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31,
	3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2,
	39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2,
	2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2,
	2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2,
	2, 2, 2, 63, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 3, 77, 3,
	2, 2, 2, 5, 84, 3, 2, 2, 2, 7, 86, 3, 2, 2, 2, 9, 88, 3, 2, 2, 2, 11, 95,
	3, 2, 2, 2, 13, 100, 3, 2, 2, 2, 15, 102, 3, 2, 2, 2, 17, 104, 3, 2, 2,
	2, 19, 110, 3, 2, 2, 2, 21, 113, 3, 2, 2, 2, 23, 115, 3, 2, 2, 2, 25, 117,
	3, 2, 2, 2, 27, 119, 3, 2, 2, 2, 29, 122, 3, 2, 2, 2, 31, 127, 3, 2, 2,
	2, 33, 129, 3, 2, 2, 2, 35, 131, 3, 2, 2, 2, 37, 135, 3, 2, 2, 2, 39, 139,
	3, 2, 2, 2, 41, 142, 3, 2, 2, 2, 43, 147, 3, 2, 2, 2, 45, 153, 3, 2, 2,
	2, 47, 155, 3, 2, 2, 2, 49, 157, 3, 2, 2, 2, 51, 159, 3, 2, 2, 2, 53, 161,
	3, 2, 2, 2, 55, 164, 3, 2, 2, 2, 57, 167, 3, 2, 2, 2, 59, 170, 3, 2, 2,
	2, 61, 176, 3, 2, 2, 2, 63, 189, 3, 2, 2, 2, 65, 196, 3, 2, 2, 2, 67, 200,
	3, 2, 2, 2, 69, 202, 3, 2, 2, 2, 71, 209, 3, 2, 2, 2, 73, 211, 3, 2, 2,
	2, 75, 213, 3, 2, 2, 2, 77, 78, 7, 107, 2, 2, 78, 79, 7, 111, 2, 2, 79,
	80, 7, 114, 2, 2, 80, 81, 7, 113, 2, 2, 81, 82, 7, 116, 2, 2, 82, 83, 7,
	118, 2, 2, 83, 4, 3, 2, 2, 2, 84, 85, 7, 42, 2, 2, 85, 6, 3, 2, 2, 2, 86,
	87, 7, 43, 2, 2, 87, 8, 3, 2, 2, 2, 88, 89, 7, 116, 2, 2, 89, 90, 7, 103,
	2, 2, 90, 91, 7, 118, 2, 2, 91, 92, 7, 119, 2, 2, 92, 93, 7, 116, 2, 2,
	93, 94, 7, 112, 2, 2, 94, 10, 3, 2, 2, 2, 95, 96, 7, 110, 2, 2, 96, 97,
	7, 113, 2, 2, 97, 98, 7, 113, 2, 2, 98, 99, 7, 114, 2, 2, 99, 12, 3, 2,
	2, 2, 100, 101, 7, 125, 2, 2, 101, 14, 3, 2, 2, 2, 102, 103, 7, 127, 2,
	2, 103, 16, 3, 2, 2, 2, 104, 105, 7, 111, 2, 2, 105, 106, 7, 99, 2, 2,
	106, 107, 7, 118, 2, 2, 107, 108, 7, 101, 2, 2, 108, 109, 7, 106, 2, 2,
	109, 18, 3, 2, 2, 2, 110, 111, 7, 63, 2, 2, 111, 112, 7, 64, 2, 2, 112,
	20, 3, 2, 2, 2, 113, 114, 7, 46, 2, 2, 114, 22, 3, 2, 2, 2, 115, 116, 7,
	63, 2, 2, 116, 24, 3, 2, 2, 2, 117, 118, 7, 66, 2, 2, 118, 26, 3, 2, 2,
	2, 119, 120, 7, 47, 2, 2, 120, 121, 7, 64, 2, 2, 121, 28, 3, 2, 2, 2, 122,
	123, 7, 107, 2, 2, 123, 124, 7, 111, 2, 2, 124, 125, 7, 114, 2, 2, 125,
	126, 7, 110, 2, 2, 126, 30, 3, 2, 2, 2, 127, 128, 7, 60, 2, 2, 128, 32,
	3, 2, 2, 2, 129, 130, 7, 45, 2, 2, 130, 34, 3, 2, 2, 2, 131, 132, 7, 110,
	2, 2, 132, 133, 7, 103, 2, 2, 133, 134, 7, 118, 2, 2, 134, 36, 3, 2, 2,
	2, 135, 136, 7, 111, 2, 2, 136, 137, 7, 119, 2, 2, 137, 138, 7, 118, 2,
	2, 138, 38, 3, 2, 2, 2, 139, 140, 7, 104, 2, 2, 140, 141, 7, 112, 2, 2,
	141, 40, 3, 2, 2, 2, 142, 143, 7, 118, 2, 2, 143, 144, 7, 123, 2, 2, 144,
	145, 7, 114, 2, 2, 145, 146, 7, 103, 2, 2, 146, 42, 3, 2, 2, 2, 147, 148,
	7, 118, 2, 2, 148, 149, 7, 116, 2, 2, 149, 150, 7, 99, 2, 2, 150, 151,
	7, 107, 2, 2, 151, 152, 7, 118, 2, 2, 152, 44, 3, 2, 2, 2, 153, 154, 7,
	96, 2, 2, 154, 46, 3, 2, 2, 2, 155, 156, 7, 44, 2, 2, 156, 48, 3, 2, 2,
	2, 157, 158, 7, 49, 2, 2, 158, 50, 3, 2, 2, 2, 159, 160, 7, 47, 2, 2, 160,
	52, 3, 2, 2, 2, 161, 162, 7, 35, 2, 2, 162, 163, 7, 63, 2, 2, 163, 54,
	3, 2, 2, 2, 164, 165, 7, 63, 2, 2, 165, 166, 7, 63, 2, 2, 166, 56, 3, 2,
	2, 2, 167, 168, 7, 65, 2, 2, 168, 58, 3, 2, 2, 2, 169, 171, 9, 2, 2, 2,
	170, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 172,
	173, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 175, 8, 30, 2, 2, 175, 60,
	3, 2, 2, 2, 176, 177, 7, 49, 2, 2, 177, 178, 7, 49, 2, 2, 178, 182, 3,
	2, 2, 2, 179, 181, 11, 2, 2, 2, 180, 179, 3, 2, 2, 2, 181, 184, 3, 2, 2,
	2, 182, 183, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 183, 185, 3, 2, 2, 2, 184,
	182, 3, 2, 2, 2, 185, 186, 7, 12, 2, 2, 186, 187, 3, 2, 2, 2, 187, 188,
	8, 31, 2, 2, 188, 62, 3, 2, 2, 2, 189, 193, 5, 65, 33, 2, 190, 192, 5,
	67, 34, 2, 191, 190, 3, 2, 2, 2, 192, 195, 3, 2, 2, 2, 193, 191, 3, 2,
	2, 2, 193, 194, 3, 2, 2, 2, 194, 64, 3, 2, 2, 2, 195, 193, 3, 2, 2, 2,
	196, 197, 9, 3, 2, 2, 197, 66, 3, 2, 2, 2, 198, 201, 9, 4, 2, 2, 199, 201,
	5, 65, 33, 2, 200, 198, 3, 2, 2, 2, 200, 199, 3, 2, 2, 2, 201, 68, 3, 2,
	2, 2, 202, 206, 5, 71, 36, 2, 203, 205, 5, 73, 37, 2, 204, 203, 3, 2, 2,
	2, 205, 208, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207,
	70, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 209, 210, 9, 5, 2, 2, 210, 72, 3,
	2, 2, 2, 211, 212, 9, 4, 2, 2, 212, 74, 3, 2, 2, 2, 213, 217, 7, 36, 2,
	2, 214, 216, 11, 2, 2, 2, 215, 214, 3, 2, 2, 2, 216, 219, 3, 2, 2, 2, 217,
	218, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 218, 220, 3, 2, 2, 2, 219, 217,
	3, 2, 2, 2, 220, 221, 7, 36, 2, 2, 221, 76, 3, 2, 2, 2, 9, 2, 172, 182,
	193, 200, 206, 217, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'import'", "'('", "')'", "'return'", "'loop'", "'{'", "'}'", "'match'",
	"'=>'", "','", "'='", "'@'", "'->'", "'impl'", "':'", "'+'", "'let'", "'mut'",
	"'fn'", "'type'", "'trait'", "'^'", "'*'", "'/'", "'-'", "'!='", "'=='",
	"'?'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "WS", "COMMENT", "ID", "NUM",
	"STRING",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "WS", "COMMENT", "ID", "StartLetter", "Letter",
	"NUM", "StartDigit", "Digit", "STRING",
}

type ElzLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewElzLexer(input antlr.CharStream) *ElzLexer {

	l := new(ElzLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Elz.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ElzLexer tokens.
const (
	ElzLexerT__0    = 1
	ElzLexerT__1    = 2
	ElzLexerT__2    = 3
	ElzLexerT__3    = 4
	ElzLexerT__4    = 5
	ElzLexerT__5    = 6
	ElzLexerT__6    = 7
	ElzLexerT__7    = 8
	ElzLexerT__8    = 9
	ElzLexerT__9    = 10
	ElzLexerT__10   = 11
	ElzLexerT__11   = 12
	ElzLexerT__12   = 13
	ElzLexerT__13   = 14
	ElzLexerT__14   = 15
	ElzLexerT__15   = 16
	ElzLexerT__16   = 17
	ElzLexerT__17   = 18
	ElzLexerT__18   = 19
	ElzLexerT__19   = 20
	ElzLexerT__20   = 21
	ElzLexerT__21   = 22
	ElzLexerT__22   = 23
	ElzLexerT__23   = 24
	ElzLexerT__24   = 25
	ElzLexerT__25   = 26
	ElzLexerT__26   = 27
	ElzLexerT__27   = 28
	ElzLexerWS      = 29
	ElzLexerCOMMENT = 30
	ElzLexerID      = 31
	ElzLexerNUM     = 32
	ElzLexerSTRING  = 33
)

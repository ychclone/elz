// Code generated from Elz.g4 by ANTLR 4.7.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 42, 260,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3,
	24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29,
	3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 33, 3,
	33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37,
	6, 37, 208, 10, 37, 13, 37, 14, 37, 209, 3, 37, 3, 37, 3, 38, 3, 38, 3,
	38, 3, 38, 7, 38, 218, 10, 38, 12, 38, 14, 38, 221, 11, 38, 3, 38, 3, 38,
	3, 38, 3, 38, 3, 39, 3, 39, 7, 39, 229, 10, 39, 12, 39, 14, 39, 232, 11,
	39, 3, 40, 5, 40, 235, 10, 40, 3, 41, 3, 41, 5, 41, 239, 10, 41, 3, 42,
	3, 42, 7, 42, 243, 10, 42, 12, 42, 14, 42, 246, 11, 42, 3, 43, 3, 43, 3,
	44, 3, 44, 3, 45, 3, 45, 7, 45, 254, 10, 45, 12, 45, 14, 45, 257, 11, 45,
	3, 45, 3, 45, 4, 219, 255, 2, 46, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8,
	15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35,
	69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 2, 81, 2, 83, 41, 85, 2, 87,
	2, 89, 42, 3, 2, 7, 5, 2, 11, 12, 15, 15, 34, 34, 16, 2, 67, 92, 97, 97,
	99, 124, 194, 216, 218, 248, 250, 769, 882, 895, 897, 8193, 8206, 8207,
	8306, 8593, 11266, 12273, 12291, 55297, 63746, 64977, 65010, 65535, 6,
	2, 50, 59, 185, 185, 770, 881, 8257, 8258, 4, 2, 48, 48, 50, 59, 3, 2,
	50, 59, 2, 261, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2,
	2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3,
	2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47,
	3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2,
	55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2,
	2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2,
	2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2,
	2, 2, 2, 83, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 3, 91, 3, 2, 2, 2, 5, 94, 3,
	2, 2, 2, 7, 101, 3, 2, 2, 2, 9, 103, 3, 2, 2, 2, 11, 105, 3, 2, 2, 2, 13,
	112, 3, 2, 2, 2, 15, 117, 3, 2, 2, 2, 17, 119, 3, 2, 2, 2, 19, 121, 3,
	2, 2, 2, 21, 127, 3, 2, 2, 2, 23, 130, 3, 2, 2, 2, 25, 132, 3, 2, 2, 2,
	27, 134, 3, 2, 2, 2, 29, 136, 3, 2, 2, 2, 31, 139, 3, 2, 2, 2, 33, 144,
	3, 2, 2, 2, 35, 146, 3, 2, 2, 2, 37, 148, 3, 2, 2, 2, 39, 152, 3, 2, 2,
	2, 41, 156, 3, 2, 2, 2, 43, 159, 3, 2, 2, 2, 45, 168, 3, 2, 2, 2, 47, 174,
	3, 2, 2, 2, 49, 176, 3, 2, 2, 2, 51, 178, 3, 2, 2, 2, 53, 180, 3, 2, 2,
	2, 55, 182, 3, 2, 2, 2, 57, 184, 3, 2, 2, 2, 59, 186, 3, 2, 2, 2, 61, 189,
	3, 2, 2, 2, 63, 192, 3, 2, 2, 2, 65, 195, 3, 2, 2, 2, 67, 198, 3, 2, 2,
	2, 69, 201, 3, 2, 2, 2, 71, 204, 3, 2, 2, 2, 73, 207, 3, 2, 2, 2, 75, 213,
	3, 2, 2, 2, 77, 226, 3, 2, 2, 2, 79, 234, 3, 2, 2, 2, 81, 238, 3, 2, 2,
	2, 83, 240, 3, 2, 2, 2, 85, 247, 3, 2, 2, 2, 87, 249, 3, 2, 2, 2, 89, 251,
	3, 2, 2, 2, 91, 92, 7, 60, 2, 2, 92, 93, 7, 60, 2, 2, 93, 4, 3, 2, 2, 2,
	94, 95, 7, 107, 2, 2, 95, 96, 7, 111, 2, 2, 96, 97, 7, 114, 2, 2, 97, 98,
	7, 113, 2, 2, 98, 99, 7, 116, 2, 2, 99, 100, 7, 118, 2, 2, 100, 6, 3, 2,
	2, 2, 101, 102, 7, 42, 2, 2, 102, 8, 3, 2, 2, 2, 103, 104, 7, 43, 2, 2,
	104, 10, 3, 2, 2, 2, 105, 106, 7, 116, 2, 2, 106, 107, 7, 103, 2, 2, 107,
	108, 7, 118, 2, 2, 108, 109, 7, 119, 2, 2, 109, 110, 7, 116, 2, 2, 110,
	111, 7, 112, 2, 2, 111, 12, 3, 2, 2, 2, 112, 113, 7, 110, 2, 2, 113, 114,
	7, 113, 2, 2, 114, 115, 7, 113, 2, 2, 115, 116, 7, 114, 2, 2, 116, 14,
	3, 2, 2, 2, 117, 118, 7, 125, 2, 2, 118, 16, 3, 2, 2, 2, 119, 120, 7, 127,
	2, 2, 120, 18, 3, 2, 2, 2, 121, 122, 7, 111, 2, 2, 122, 123, 7, 99, 2,
	2, 123, 124, 7, 118, 2, 2, 124, 125, 7, 101, 2, 2, 125, 126, 7, 106, 2,
	2, 126, 20, 3, 2, 2, 2, 127, 128, 7, 63, 2, 2, 128, 129, 7, 64, 2, 2, 129,
	22, 3, 2, 2, 2, 130, 131, 7, 46, 2, 2, 131, 24, 3, 2, 2, 2, 132, 133, 7,
	63, 2, 2, 133, 26, 3, 2, 2, 2, 134, 135, 7, 66, 2, 2, 135, 28, 3, 2, 2,
	2, 136, 137, 7, 47, 2, 2, 137, 138, 7, 64, 2, 2, 138, 30, 3, 2, 2, 2, 139,
	140, 7, 107, 2, 2, 140, 141, 7, 111, 2, 2, 141, 142, 7, 114, 2, 2, 142,
	143, 7, 110, 2, 2, 143, 32, 3, 2, 2, 2, 144, 145, 7, 60, 2, 2, 145, 34,
	3, 2, 2, 2, 146, 147, 7, 45, 2, 2, 147, 36, 3, 2, 2, 2, 148, 149, 7, 110,
	2, 2, 149, 150, 7, 103, 2, 2, 150, 151, 7, 118, 2, 2, 151, 38, 3, 2, 2,
	2, 152, 153, 7, 111, 2, 2, 153, 154, 7, 119, 2, 2, 154, 155, 7, 118, 2,
	2, 155, 40, 3, 2, 2, 2, 156, 157, 7, 104, 2, 2, 157, 158, 7, 112, 2, 2,
	158, 42, 3, 2, 2, 2, 159, 160, 7, 118, 2, 2, 160, 161, 7, 123, 2, 2, 161,
	162, 7, 114, 2, 2, 162, 163, 7, 103, 2, 2, 163, 164, 7, 72, 2, 2, 164,
	165, 7, 113, 2, 2, 165, 166, 7, 116, 2, 2, 166, 167, 7, 111, 2, 2, 167,
	44, 3, 2, 2, 2, 168, 169, 7, 118, 2, 2, 169, 170, 7, 116, 2, 2, 170, 171,
	7, 99, 2, 2, 171, 172, 7, 107, 2, 2, 172, 173, 7, 118, 2, 2, 173, 46, 3,
	2, 2, 2, 174, 175, 7, 96, 2, 2, 175, 48, 3, 2, 2, 2, 176, 177, 7, 44, 2,
	2, 177, 50, 3, 2, 2, 2, 178, 179, 7, 49, 2, 2, 179, 52, 3, 2, 2, 2, 180,
	181, 7, 47, 2, 2, 181, 54, 3, 2, 2, 2, 182, 183, 7, 62, 2, 2, 183, 56,
	3, 2, 2, 2, 184, 185, 7, 64, 2, 2, 185, 58, 3, 2, 2, 2, 186, 187, 7, 62,
	2, 2, 187, 188, 7, 63, 2, 2, 188, 60, 3, 2, 2, 2, 189, 190, 7, 64, 2, 2,
	190, 191, 7, 63, 2, 2, 191, 62, 3, 2, 2, 2, 192, 193, 7, 35, 2, 2, 193,
	194, 7, 63, 2, 2, 194, 64, 3, 2, 2, 2, 195, 196, 7, 63, 2, 2, 196, 197,
	7, 63, 2, 2, 197, 66, 3, 2, 2, 2, 198, 199, 7, 40, 2, 2, 199, 200, 7, 40,
	2, 2, 200, 68, 3, 2, 2, 2, 201, 202, 7, 126, 2, 2, 202, 203, 7, 126, 2,
	2, 203, 70, 3, 2, 2, 2, 204, 205, 7, 65, 2, 2, 205, 72, 3, 2, 2, 2, 206,
	208, 9, 2, 2, 2, 207, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 207,
	3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 212, 8, 37,
	2, 2, 212, 74, 3, 2, 2, 2, 213, 214, 7, 49, 2, 2, 214, 215, 7, 49, 2, 2,
	215, 219, 3, 2, 2, 2, 216, 218, 11, 2, 2, 2, 217, 216, 3, 2, 2, 2, 218,
	221, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 219, 217, 3, 2, 2, 2, 220, 222,
	3, 2, 2, 2, 221, 219, 3, 2, 2, 2, 222, 223, 7, 12, 2, 2, 223, 224, 3, 2,
	2, 2, 224, 225, 8, 38, 2, 2, 225, 76, 3, 2, 2, 2, 226, 230, 5, 79, 40,
	2, 227, 229, 5, 81, 41, 2, 228, 227, 3, 2, 2, 2, 229, 232, 3, 2, 2, 2,
	230, 228, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 78, 3, 2, 2, 2, 232, 230,
	3, 2, 2, 2, 233, 235, 9, 3, 2, 2, 234, 233, 3, 2, 2, 2, 235, 80, 3, 2,
	2, 2, 236, 239, 5, 79, 40, 2, 237, 239, 9, 4, 2, 2, 238, 236, 3, 2, 2,
	2, 238, 237, 3, 2, 2, 2, 239, 82, 3, 2, 2, 2, 240, 244, 5, 85, 43, 2, 241,
	243, 5, 87, 44, 2, 242, 241, 3, 2, 2, 2, 243, 246, 3, 2, 2, 2, 244, 242,
	3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 84, 3, 2, 2, 2, 246, 244, 3, 2,
	2, 2, 247, 248, 9, 5, 2, 2, 248, 86, 3, 2, 2, 2, 249, 250, 9, 6, 2, 2,
	250, 88, 3, 2, 2, 2, 251, 255, 7, 36, 2, 2, 252, 254, 11, 2, 2, 2, 253,
	252, 3, 2, 2, 2, 254, 257, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 255, 253,
	3, 2, 2, 2, 256, 258, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 258, 259, 7, 36,
	2, 2, 259, 90, 3, 2, 2, 2, 10, 2, 209, 219, 230, 234, 238, 244, 255, 3,
	2, 3, 2,
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
	"", "'::'", "'import'", "'('", "')'", "'return'", "'loop'", "'{'", "'}'",
	"'match'", "'=>'", "','", "'='", "'@'", "'->'", "'impl'", "':'", "'+'",
	"'let'", "'mut'", "'fn'", "'typeForm'", "'trait'", "'^'", "'*'", "'/'",
	"'-'", "'<'", "'>'", "'<='", "'>='", "'!='", "'=='", "'&&'", "'||'", "'?'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"WS", "COMMENT", "ID", "NUM", "STRING",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "WS", "COMMENT", "ID", "StartLetter", "Letter", "NUM",
	"StartDigit", "Digit", "STRING",
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
	ElzLexerT__28   = 29
	ElzLexerT__29   = 30
	ElzLexerT__30   = 31
	ElzLexerT__31   = 32
	ElzLexerT__32   = 33
	ElzLexerT__33   = 34
	ElzLexerT__34   = 35
	ElzLexerWS      = 36
	ElzLexerCOMMENT = 37
	ElzLexerID      = 38
	ElzLexerNUM     = 39
	ElzLexerSTRING  = 40
)

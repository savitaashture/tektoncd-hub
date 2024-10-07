// Code generated by goyacc -o parser.go parser.go.y. DO NOT EDIT.

//line parser.go.y:2
package gojq

import __yyfmt__ "fmt"

//line parser.go.y:2

func reverseFuncDef(xs []*FuncDef) []*FuncDef {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}
	return xs
}

func prependFuncDef(xs []*FuncDef, x *FuncDef) []*FuncDef {
	xs = append(xs, nil)
	copy(xs[1:], xs)
	xs[0] = x
	return xs
}

//line parser.go.y:19
type yySymType struct {
	yys      int
	value    any
	token    string
	operator Operator
}

const tokAltOp = 57346
const tokUpdateOp = 57347
const tokDestAltOp = 57348
const tokCompareOp = 57349
const tokOrOp = 57350
const tokAndOp = 57351
const tokModule = 57352
const tokImport = 57353
const tokInclude = 57354
const tokDef = 57355
const tokAs = 57356
const tokLabel = 57357
const tokBreak = 57358
const tokNull = 57359
const tokTrue = 57360
const tokFalse = 57361
const tokIf = 57362
const tokThen = 57363
const tokElif = 57364
const tokElse = 57365
const tokEnd = 57366
const tokTry = 57367
const tokCatch = 57368
const tokReduce = 57369
const tokForeach = 57370
const tokIdent = 57371
const tokVariable = 57372
const tokModuleIdent = 57373
const tokModuleVariable = 57374
const tokRecurse = 57375
const tokIndex = 57376
const tokNumber = 57377
const tokFormat = 57378
const tokString = 57379
const tokStringStart = 57380
const tokStringQuery = 57381
const tokStringEnd = 57382
const tokInvalid = 57383
const tokInvalidEscapeSequence = 57384
const tokUnterminatedString = 57385
const tokFuncDefQuery = 57386
const tokExpr = 57387
const tokTerm = 57388
const tokEmptyCatch = 57389

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"tokAltOp",
	"tokUpdateOp",
	"tokDestAltOp",
	"tokCompareOp",
	"tokOrOp",
	"tokAndOp",
	"tokModule",
	"tokImport",
	"tokInclude",
	"tokDef",
	"tokAs",
	"tokLabel",
	"tokBreak",
	"tokNull",
	"tokTrue",
	"tokFalse",
	"tokIf",
	"tokThen",
	"tokElif",
	"tokElse",
	"tokEnd",
	"tokTry",
	"tokCatch",
	"tokReduce",
	"tokForeach",
	"tokIdent",
	"tokVariable",
	"tokModuleIdent",
	"tokModuleVariable",
	"tokRecurse",
	"tokIndex",
	"tokNumber",
	"tokFormat",
	"tokString",
	"tokStringStart",
	"tokStringQuery",
	"tokStringEnd",
	"tokInvalid",
	"tokInvalidEscapeSequence",
	"tokUnterminatedString",
	"tokFuncDefQuery",
	"tokExpr",
	"tokTerm",
	"'|'",
	"','",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"'?'",
	"tokEmptyCatch",
	"'['",
	"';'",
	"':'",
	"'('",
	"')'",
	"']'",
	"'{'",
	"'}'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:671

//line yacctab:1
var yyExca = [...]int16{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 145,
	5, 0,
	-2, 27,
	-1, 148,
	7, 0,
	-2, 30,
	-1, 199,
	59, 114,
	-2, 49,
}

const yyPrivate = 57344

const yyLast = 782

var yyAct = [...]int16{
	78, 134, 186, 102, 103, 10, 175, 195, 32, 211,
	48, 108, 81, 176, 131, 6, 229, 5, 50, 73,
	74, 159, 14, 180, 181, 182, 124, 98, 110, 135,
	280, 97, 228, 279, 115, 104, 16, 158, 265, 121,
	114, 178, 123, 179, 244, 73, 74, 180, 181, 182,
	73, 74, 112, 113, 154, 155, 136, 117, 117, 117,
	254, 243, 137, 183, 282, 178, 255, 179, 220, 6,
	247, 116, 118, 119, 128, 129, 73, 74, 99, 73,
	74, 227, 73, 74, 246, 141, 238, 183, 201, 237,
	132, 200, 139, 6, 235, 226, 138, 163, 208, 80,
	157, 207, 241, 231, 230, 161, 162, 73, 74, 117,
	117, 117, 117, 117, 117, 117, 117, 117, 117, 83,
	82, 278, 84, 144, 145, 146, 147, 148, 149, 150,
	151, 152, 153, 184, 185, 174, 50, 160, 193, 73,
	74, 127, 196, 202, 203, 126, 197, 73, 74, 125,
	73, 74, 248, 253, 189, 204, 45, 240, 206, 73,
	74, 245, 143, 210, 214, 215, 73, 74, 104, 217,
	218, 213, 79, 219, 86, 87, 76, 90, 88, 89,
	169, 43, 44, 117, 117, 73, 74, 75, 166, 117,
	222, 224, 80, 225, 73, 74, 273, 212, 212, 232,
	132, 223, 234, 216, 120, 271, 73, 74, 191, 239,
	43, 44, 83, 82, 85, 84, 274, 270, 96, 91,
	92, 93, 94, 95, 73, 74, 93, 94, 95, 249,
	84, 164, 251, 252, 196, 236, 267, 250, 197, 130,
	25, 256, 73, 74, 262, 263, 187, 188, 3, 190,
	257, 258, 260, 261, 264, 24, 266, 73, 74, 9,
	221, 268, 269, 117, 117, 111, 171, 272, 172, 170,
	13, 275, 276, 77, 90, 277, 89, 212, 212, 13,
	177, 281, 52, 53, 54, 55, 56, 57, 58, 59,
	60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
	70, 71, 72, 106, 107, 91, 92, 93, 94, 95,
	47, 43, 44, 101, 165, 259, 91, 92, 93, 94,
	95, 242, 156, 122, 194, 17, 192, 15, 37, 21,
	22, 23, 33, 133, 105, 205, 7, 34, 209, 35,
	36, 39, 41, 40, 42, 19, 20, 28, 31, 43,
	44, 8, 4, 2, 86, 87, 1, 90, 88, 89,
	0, 29, 30, 0, 168, 90, 18, 0, 0, 27,
	0, 142, 38, 0, 140, 26, 52, 53, 54, 55,
	56, 57, 58, 59, 60, 61, 62, 63, 64, 65,
	66, 67, 68, 69, 70, 71, 72, 106, 107, 91,
	92, 93, 94, 95, 0, 43, 44, 91, 92, 93,
	94, 95, 0, 0, 0, 0, 0, 11, 12, 17,
	0, 15, 37, 21, 22, 23, 33, 0, 105, 0,
	0, 34, 100, 35, 36, 39, 41, 40, 42, 19,
	20, 28, 31, 43, 44, 0, 0, 0, 0, 86,
	87, 0, 90, 88, 89, 29, 30, 0, 0, 167,
	18, 0, 0, 27, 0, 0, 38, 0, 17, 26,
	15, 37, 21, 22, 23, 33, 0, 0, 0, 0,
	34, 0, 35, 36, 39, 41, 40, 42, 19, 20,
	28, 31, 43, 44, 91, 92, 93, 94, 95, 0,
	0, 0, 0, 0, 29, 30, 90, 88, 89, 18,
	0, 0, 27, 0, 0, 38, 0, 233, 26, 17,
	0, 15, 37, 21, 22, 23, 33, 0, 0, 0,
	0, 34, 0, 35, 36, 39, 41, 40, 42, 19,
	20, 28, 31, 43, 44, 0, 0, 0, 91, 92,
	93, 94, 95, 0, 0, 29, 30, 0, 0, 0,
	18, 0, 0, 27, 0, 0, 38, 0, 109, 26,
	17, 0, 15, 37, 21, 22, 23, 33, 0, 0,
	0, 0, 34, 0, 35, 36, 39, 41, 40, 42,
	19, 20, 28, 31, 43, 44, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 29, 30, 0, 0,
	0, 18, 0, 0, 27, 0, 0, 38, 0, 0,
	26, 52, 53, 54, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 72, 49, 0, 0, 0, 0, 0, 0, 0,
	51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 72, 49, 0, 0, 0, 0, 173, 0, 0,
	51, 37, 21, 22, 23, 33, 0, 0, 0, 0,
	34, 0, 35, 36, 39, 41, 40, 42, 19, 20,
	28, 31, 43, 44, 0, 0, 0, 46, 0, 0,
	0, 0, 0, 0, 29, 30, 0, 0, 0, 18,
	0, 0, 27, 0, 0, 38, 0, 0, 26, 52,
	53, 54, 55, 56, 57, 58, 59, 60, 61, 62,
	63, 64, 65, 66, 67, 68, 69, 70, 71, 72,
	106, 199, 0, 0, 0, 0, 0, 0, 43, 44,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 198,
}

var yyPact = [...]int16{
	238, -1000, -1000, -48, 406, 98, 643, -1000, -1000, -1000,
	112, 150, 139, 557, 158, 184, 170, 189, 173, -1000,
	-1000, -1000, -1000, -1000, 18, -1000, 368, 506, -1000, 665,
	665, 144, -1000, 557, 665, 665, 665, 174, 557, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -22, -1000, 90,
	86, 82, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 557, 557, 225, -48, -1000, 112, -1,
	-1000, -1000, -1000, 173, 312, 115, 665, 665, 665, 665,
	665, 665, 665, 665, 665, 665, -5, -1000, -1000, 557,
	-1000, -27, -1000, 78, 46, 557, -1000, -1000, -1000, -1000,
	35, 557, 65, 65, -1000, 210, 162, 65, 445, 350,
	-1000, 119, 229, -1000, 613, 30, 30, 30, 112, -1000,
	217, 96, -1000, 202, -1000, -1000, -1, 721, -1000, -1000,
	-1000, 29, 557, 557, 170, 499, 267, 358, 256, 175,
	175, -1000, -1000, -1000, 557, 217, 40, 112, -1000, 274,
	665, 665, 103, -1000, 557, -1000, 665, -1, -1, -1000,
	-1000, -1000, 557, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 6, -1000, -1000, -48, -1000, -1000, -1000,
	557, -1, 33, -1000, -32, -1000, 45, 44, 557, -1000,
	-1000, 455, 32, 112, 177, 28, -1000, -1000, 557, -1000,
	-1000, 110, 170, 110, 43, 112, -1000, 1, -16, 100,
	-1000, 22, -1000, 94, 112, -1000, -1000, -1, -1000, 721,
	-1, -1, 92, -1000, -2, -1000, -1000, 7, 217, 112,
	665, 665, 230, 557, 557, -1000, -1000, 30, -1000, -1000,
	-1000, -1000, -1000, -21, -1000, 557, -1000, 110, 110, 212,
	557, 557, 159, 147, -1000, -1, 138, -1000, 195, 112,
	557, 557, -1000, -1000, 557, 60, -28, 112, -1000, -1000,
	557, 3, -1000,
}

var yyPgo = [...]int16{
	0, 356, 353, 352, 351, 14, 336, 259, 265, 335,
	0, 333, 1, 326, 324, 7, 36, 22, 8, 323,
	12, 322, 321, 315, 314, 313, 3, 9, 6, 13,
	310, 10, 280, 260, 2, 255, 240, 11, 4,
}

var yyR1 = [...]int8{
	0, 1, 2, 2, 3, 3, 4, 4, 5, 5,
	6, 6, 7, 7, 8, 8, 9, 9, 34, 34,
	10, 10, 10, 10, 10, 10, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 11, 11, 12,
	12, 12, 13, 13, 14, 14, 15, 15, 15, 15,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 18, 18, 19, 19, 19, 35,
	35, 36, 36, 20, 20, 20, 20, 20, 21, 21,
	22, 22, 23, 23, 24, 24, 25, 25, 26, 26,
	26, 26, 26, 38, 38, 38, 27, 27, 28, 28,
	28, 28, 28, 28, 28, 29, 29, 29, 30, 30,
	31, 31, 31, 32, 32, 33, 33, 37, 37, 37,
	37, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 37, 37, 37, 37, 37, 37,
}

var yyR2 = [...]int8{
	0, 3, 0, 3, 0, 2, 6, 4, 0, 1,
	1, 1, 0, 2, 5, 8, 1, 3, 1, 1,
	2, 3, 5, 4, 3, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 1, 1, 3, 1,
	3, 3, 1, 3, 1, 3, 3, 3, 5, 1,
	1, 1, 1, 2, 2, 1, 1, 1, 1, 4,
	1, 2, 3, 4, 2, 3, 1, 2, 2, 1,
	2, 1, 7, 3, 9, 9, 11, 2, 3, 2,
	2, 2, 3, 3, 1, 3, 0, 2, 4, 1,
	1, 1, 1, 2, 3, 4, 4, 5, 1, 3,
	0, 5, 0, 2, 0, 2, 1, 3, 3, 3,
	5, 1, 1, 1, 1, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 3, 4, 1, 3,
	3, 3, 3, 2, 3, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -1, -2, 10, -3, -29, 63, -6, -4, -7,
	-10, 11, 12, -8, -17, 15, -16, 13, 54, 33,
	34, 17, 18, 19, -35, -36, 63, 57, 35, 49,
	50, 36, -18, 20, 25, 27, 28, 16, 60, 29,
	31, 30, 32, 37, 38, 58, 64, -30, -31, 29,
	-37, 37, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 47, 48, 37, 37, -7, -10, 14,
	34, -20, 55, 54, 57, 30, 4, 5, 8, 9,
	7, 49, 50, 51, 52, 53, 29, -20, -18, 60,
	64, -25, -26, -38, -18, 60, 29, 30, -37, 62,
	-10, -8, -17, -17, -18, -10, -16, -17, -16, -16,
	30, -10, -19, 64, 48, 59, 59, 59, -10, -10,
	14, -5, -29, -11, -12, 30, 57, 63, -20, -18,
	62, -10, 59, 47, -16, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, 59, 60, -21, -10, 64, 48,
	59, 59, -10, 62, 21, -24, 26, 14, 14, 61,
	40, 37, 39, 64, -31, -28, -29, -32, 35, 37,
	17, 18, 19, 57, -28, -28, -34, 29, 30, 58,
	47, 6, -13, -12, -14, -15, -38, -18, 60, 30,
	62, 59, -10, -10, -10, -9, -34, 61, 58, 64,
	-26, -27, -16, -27, 61, -10, -16, -12, -12, -10,
	62, -33, -28, -5, -10, -12, 62, 48, 64, 48,
	59, 59, -10, 62, -10, 62, 58, 61, 58, -10,
	47, 59, -22, 60, 60, 61, 62, 48, 58, -12,
	-15, -12, -12, 61, 62, 59, -34, -27, -27, -23,
	22, 23, -10, -10, -28, 59, -10, 24, -10, -10,
	58, 58, -12, 58, 21, -10, -10, -10, 61, 61,
	58, -10, 61,
}

var yyDef = [...]int16{
	2, -2, 4, 0, 12, 0, 0, 1, 5, 10,
	11, 0, 0, 12, 36, 0, 25, 0, 50, 51,
	52, 55, 56, 57, 58, 60, 0, 0, 66, 0,
	0, 69, 71, 0, 0, 0, 0, 0, 0, 89,
	90, 91, 92, 84, 86, 3, 125, 0, 128, 0,
	0, 0, 137, 138, 139, 140, 141, 142, 143, 144,
	145, 146, 147, 148, 149, 150, 151, 152, 153, 154,
	155, 156, 157, 0, 0, 0, 8, 13, 20, 0,
	79, 80, 81, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 53, 54, 0,
	61, 0, 106, 111, 112, 0, 113, 114, 115, 64,
	0, 0, 67, 68, 70, 0, 104, 36, 0, 0,
	77, 0, 0, 126, 0, 0, 0, 0, 21, 24,
	0, 0, 9, 0, 37, 39, 0, 0, 82, 83,
	93, 0, 0, 0, 26, -2, 28, 29, -2, 31,
	32, 33, 34, 35, 0, 0, 0, 98, 62, 0,
	0, 0, 0, 65, 0, 73, 0, 0, 0, 78,
	85, 87, 0, 127, 129, 130, 118, 119, 120, 121,
	122, 123, 124, 0, 131, 132, 8, 18, 19, 7,
	0, 0, 0, 42, 0, 44, 0, 0, 0, -2,
	94, 0, 0, 23, 0, 0, 16, 59, 0, 63,
	107, 108, 117, 109, 0, 100, 105, 0, 0, 0,
	133, 0, 135, 0, 22, 38, 40, 0, 41, 0,
	0, 0, 0, 95, 0, 96, 14, 0, 0, 99,
	0, 0, 102, 0, 0, 88, 134, 0, 6, 43,
	45, 46, 47, 0, 97, 0, 17, 116, 110, 0,
	0, 0, 0, 0, 136, 0, 0, 72, 0, 103,
	0, 0, 48, 15, 0, 0, 0, 101, 74, 75,
	0, 0, 76,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 53, 3, 3,
	60, 61, 51, 49, 48, 50, 54, 52, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 59, 58,
	3, 3, 3, 55, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 57, 3, 62, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 63, 47, 64,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 56,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:58
		{
			query := yyDollar[3].value.(*Query)
			query.Meta = yyDollar[1].value.(*ConstObject)
			query.Imports = yyDollar[2].value.([]*Import)
			yylex.(*lexer).result = query
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:67
		{
			yyVAL.value = (*ConstObject)(nil)
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:71
		{
			yyVAL.value = yyDollar[2].value
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:77
		{
			yyVAL.value = []*Import(nil)
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:81
		{
			yyVAL.value = append(yyDollar[1].value.([]*Import), yyDollar[2].value.(*Import))
		}
	case 6:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:87
		{
			yyVAL.value = &Import{ImportPath: yyDollar[2].token, ImportAlias: yyDollar[4].token, Meta: yyDollar[5].value.(*ConstObject)}
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:91
		{
			yyVAL.value = &Import{IncludePath: yyDollar[2].token, Meta: yyDollar[3].value.(*ConstObject)}
		}
	case 8:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:97
		{
			yyVAL.value = (*ConstObject)(nil)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:104
		{
			yyVAL.value = &Query{FuncDefs: reverseFuncDef(yyDollar[1].value.([]*FuncDef))}
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:111
		{
			yyVAL.value = []*FuncDef(nil)
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:115
		{
			yyVAL.value = append(yyDollar[2].value.([]*FuncDef), yyDollar[1].value.(*FuncDef))
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:121
		{
			yyVAL.value = &FuncDef{Name: yyDollar[2].token, Body: yyDollar[4].value.(*Query)}
		}
	case 15:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.go.y:125
		{
			yyVAL.value = &FuncDef{yyDollar[2].token, yyDollar[4].value.([]string), yyDollar[7].value.(*Query)}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:131
		{
			yyVAL.value = []string{yyDollar[1].token}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:135
		{
			yyVAL.value = append(yyDollar[1].value.([]string), yyDollar[3].token)
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:145
		{
			query := yyDollar[2].value.(*Query)
			query.FuncDefs = prependFuncDef(query.FuncDefs, yyDollar[1].value.(*FuncDef))
			yyVAL.value = query
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:151
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpPipe, Right: yyDollar[3].value.(*Query)}
		}
	case 22:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:155
		{
			term := yyDollar[1].value.(*Term)
			term.SuffixList = append(term.SuffixList, &Suffix{Bind: &Bind{yyDollar[3].value.([]*Pattern), yyDollar[5].value.(*Query)}})
			yyVAL.value = &Query{Term: term}
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:161
		{
			yyVAL.value = &Query{Term: &Term{Type: TermTypeLabel, Label: &Label{yyDollar[2].token, yyDollar[4].value.(*Query)}}}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:165
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpComma, Right: yyDollar[3].value.(*Query)}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:172
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: yyDollar[2].operator, Right: yyDollar[3].value.(*Query)}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:176
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: yyDollar[2].operator, Right: yyDollar[3].value.(*Query)}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:180
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpOr, Right: yyDollar[3].value.(*Query)}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:184
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpAnd, Right: yyDollar[3].value.(*Query)}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:188
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: yyDollar[2].operator, Right: yyDollar[3].value.(*Query)}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:192
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpAdd, Right: yyDollar[3].value.(*Query)}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:196
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpSub, Right: yyDollar[3].value.(*Query)}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:200
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpMul, Right: yyDollar[3].value.(*Query)}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:204
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpDiv, Right: yyDollar[3].value.(*Query)}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:208
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpMod, Right: yyDollar[3].value.(*Query)}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:212
		{
			yyVAL.value = &Query{Term: yyDollar[1].value.(*Term)}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:218
		{
			yyVAL.value = []*Pattern{yyDollar[1].value.(*Pattern)}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:222
		{
			yyVAL.value = append(yyDollar[1].value.([]*Pattern), yyDollar[3].value.(*Pattern))
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:228
		{
			yyVAL.value = &Pattern{Name: yyDollar[1].token}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:232
		{
			yyVAL.value = &Pattern{Array: yyDollar[2].value.([]*Pattern)}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:236
		{
			yyVAL.value = &Pattern{Object: yyDollar[2].value.([]*PatternObject)}
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:242
		{
			yyVAL.value = []*Pattern{yyDollar[1].value.(*Pattern)}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:246
		{
			yyVAL.value = append(yyDollar[1].value.([]*Pattern), yyDollar[3].value.(*Pattern))
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:252
		{
			yyVAL.value = []*PatternObject{yyDollar[1].value.(*PatternObject)}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:256
		{
			yyVAL.value = append(yyDollar[1].value.([]*PatternObject), yyDollar[3].value.(*PatternObject))
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:262
		{
			yyVAL.value = &PatternObject{Key: yyDollar[1].token, Val: yyDollar[3].value.(*Pattern)}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:266
		{
			yyVAL.value = &PatternObject{KeyString: yyDollar[1].value.(*String), Val: yyDollar[3].value.(*Pattern)}
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:270
		{
			yyVAL.value = &PatternObject{KeyQuery: yyDollar[2].value.(*Query), Val: yyDollar[5].value.(*Pattern)}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:274
		{
			yyVAL.value = &PatternObject{Key: yyDollar[1].token}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:280
		{
			yyVAL.value = &Term{Type: TermTypeIdentity}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:284
		{
			yyVAL.value = &Term{Type: TermTypeRecurse}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:288
		{
			yyVAL.value = &Term{Type: TermTypeIndex, Index: &Index{Name: yyDollar[1].token}}
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:292
		{
			suffix := yyDollar[2].value.(*Suffix)
			if suffix.Iter {
				yyVAL.value = &Term{Type: TermTypeIdentity, SuffixList: []*Suffix{suffix}}
			} else {
				yyVAL.value = &Term{Type: TermTypeIndex, Index: suffix.Index}
			}
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:301
		{
			yyVAL.value = &Term{Type: TermTypeIndex, Index: &Index{Str: yyDollar[2].value.(*String)}}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:305
		{
			yyVAL.value = &Term{Type: TermTypeNull}
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:309
		{
			yyVAL.value = &Term{Type: TermTypeTrue}
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:313
		{
			yyVAL.value = &Term{Type: TermTypeFalse}
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:317
		{
			yyVAL.value = &Term{Type: TermTypeFunc, Func: &Func{Name: yyDollar[1].token}}
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:321
		{
			yyVAL.value = &Term{Type: TermTypeFunc, Func: &Func{Name: yyDollar[1].token, Args: yyDollar[3].value.([]*Query)}}
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:325
		{
			yyVAL.value = &Term{Type: TermTypeFunc, Func: &Func{Name: yyDollar[1].token}}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:329
		{
			yyVAL.value = &Term{Type: TermTypeObject, Object: &Object{}}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:333
		{
			yyVAL.value = &Term{Type: TermTypeObject, Object: &Object{yyDollar[2].value.([]*ObjectKeyVal)}}
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:337
		{
			yyVAL.value = &Term{Type: TermTypeObject, Object: &Object{yyDollar[2].value.([]*ObjectKeyVal)}}
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:341
		{
			yyVAL.value = &Term{Type: TermTypeArray, Array: &Array{}}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:345
		{
			yyVAL.value = &Term{Type: TermTypeArray, Array: &Array{yyDollar[2].value.(*Query)}}
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:349
		{
			yyVAL.value = &Term{Type: TermTypeNumber, Number: yyDollar[1].token}
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:353
		{
			yyVAL.value = &Term{Type: TermTypeUnary, Unary: &Unary{OpAdd, yyDollar[2].value.(*Term)}}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:357
		{
			yyVAL.value = &Term{Type: TermTypeUnary, Unary: &Unary{OpSub, yyDollar[2].value.(*Term)}}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:361
		{
			yyVAL.value = &Term{Type: TermTypeFormat, Format: yyDollar[1].token}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:365
		{
			yyVAL.value = &Term{Type: TermTypeFormat, Format: yyDollar[1].token, Str: yyDollar[2].value.(*String)}
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:369
		{
			yyVAL.value = &Term{Type: TermTypeString, Str: yyDollar[1].value.(*String)}
		}
	case 72:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:373
		{
			yyVAL.value = &Term{Type: TermTypeIf, If: &If{yyDollar[2].value.(*Query), yyDollar[4].value.(*Query), yyDollar[5].value.([]*IfElif), yyDollar[6].value.(*Query)}}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:377
		{
			yyVAL.value = &Term{Type: TermTypeTry, Try: &Try{yyDollar[2].value.(*Query), yyDollar[3].value.(*Query)}}
		}
	case 74:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:381
		{
			yyVAL.value = &Term{Type: TermTypeReduce, Reduce: &Reduce{yyDollar[2].value.(*Query), yyDollar[4].value.(*Pattern), yyDollar[6].value.(*Query), yyDollar[8].value.(*Query)}}
		}
	case 75:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:385
		{
			yyVAL.value = &Term{Type: TermTypeForeach, Foreach: &Foreach{yyDollar[2].value.(*Query), yyDollar[4].value.(*Pattern), yyDollar[6].value.(*Query), yyDollar[8].value.(*Query), nil}}
		}
	case 76:
		yyDollar = yyS[yypt-11 : yypt+1]
//line parser.go.y:389
		{
			yyVAL.value = &Term{Type: TermTypeForeach, Foreach: &Foreach{yyDollar[2].value.(*Query), yyDollar[4].value.(*Pattern), yyDollar[6].value.(*Query), yyDollar[8].value.(*Query), yyDollar[10].value.(*Query)}}
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:393
		{
			yyVAL.value = &Term{Type: TermTypeBreak, Break: yyDollar[2].token}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:397
		{
			yyVAL.value = &Term{Type: TermTypeQuery, Query: yyDollar[2].value.(*Query)}
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:401
		{
			yyDollar[1].value.(*Term).SuffixList = append(yyDollar[1].value.(*Term).SuffixList, &Suffix{Index: &Index{Name: yyDollar[2].token}})
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:405
		{
			yyDollar[1].value.(*Term).SuffixList = append(yyDollar[1].value.(*Term).SuffixList, yyDollar[2].value.(*Suffix))
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:409
		{
			yyDollar[1].value.(*Term).SuffixList = append(yyDollar[1].value.(*Term).SuffixList, &Suffix{Optional: true})
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:413
		{
			yyDollar[1].value.(*Term).SuffixList = append(yyDollar[1].value.(*Term).SuffixList, yyDollar[3].value.(*Suffix))
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:417
		{
			yyDollar[1].value.(*Term).SuffixList = append(yyDollar[1].value.(*Term).SuffixList, &Suffix{Index: &Index{Str: yyDollar[3].value.(*String)}})
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:423
		{
			yyVAL.value = &String{Str: yyDollar[1].token}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:427
		{
			yyVAL.value = &String{Queries: yyDollar[2].value.([]*Query)}
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:433
		{
			yyVAL.value = []*Query{}
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:437
		{
			yyVAL.value = append(yyDollar[1].value.([]*Query), &Query{Term: &Term{Type: TermTypeString, Str: &String{Str: yyDollar[2].token}}})
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:441
		{
			yylex.(*lexer).inString = true
			yyVAL.value = append(yyDollar[1].value.([]*Query), &Query{Term: &Term{Type: TermTypeQuery, Query: yyDollar[3].value.(*Query)}})
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:456
		{
			yyVAL.value = &Suffix{Iter: true}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:460
		{
			yyVAL.value = &Suffix{Index: &Index{Start: yyDollar[2].value.(*Query)}}
		}
	case 95:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:464
		{
			yyVAL.value = &Suffix{Index: &Index{Start: yyDollar[2].value.(*Query), IsSlice: true}}
		}
	case 96:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:468
		{
			yyVAL.value = &Suffix{Index: &Index{End: yyDollar[3].value.(*Query), IsSlice: true}}
		}
	case 97:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:472
		{
			yyVAL.value = &Suffix{Index: &Index{Start: yyDollar[2].value.(*Query), End: yyDollar[4].value.(*Query), IsSlice: true}}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:478
		{
			yyVAL.value = []*Query{yyDollar[1].value.(*Query)}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:482
		{
			yyVAL.value = append(yyDollar[1].value.([]*Query), yyDollar[3].value.(*Query))
		}
	case 100:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:488
		{
			yyVAL.value = []*IfElif(nil)
		}
	case 101:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:492
		{
			yyVAL.value = append(yyDollar[1].value.([]*IfElif), &IfElif{yyDollar[3].value.(*Query), yyDollar[5].value.(*Query)})
		}
	case 102:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:498
		{
			yyVAL.value = (*Query)(nil)
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:502
		{
			yyVAL.value = yyDollar[2].value
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:508
		{
			yyVAL.value = (*Query)(nil)
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:512
		{
			yyVAL.value = yyDollar[2].value
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:518
		{
			yyVAL.value = []*ObjectKeyVal{yyDollar[1].value.(*ObjectKeyVal)}
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:522
		{
			yyVAL.value = append(yyDollar[1].value.([]*ObjectKeyVal), yyDollar[3].value.(*ObjectKeyVal))
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:528
		{
			yyVAL.value = &ObjectKeyVal{Key: yyDollar[1].token, Val: yyDollar[3].value.(*Query)}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:532
		{
			yyVAL.value = &ObjectKeyVal{KeyString: yyDollar[1].value.(*String), Val: yyDollar[3].value.(*Query)}
		}
	case 110:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:536
		{
			yyVAL.value = &ObjectKeyVal{KeyQuery: yyDollar[2].value.(*Query), Val: yyDollar[5].value.(*Query)}
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:540
		{
			yyVAL.value = &ObjectKeyVal{Key: yyDollar[1].token}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:544
		{
			yyVAL.value = &ObjectKeyVal{KeyString: yyDollar[1].value.(*String)}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:555
		{
			yyVAL.value = &Query{Left: yyDollar[1].value.(*Query), Op: OpPipe, Right: yyDollar[3].value.(*Query)}
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:562
		{
			yyVAL.value = &ConstTerm{Object: yyDollar[1].value.(*ConstObject)}
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:566
		{
			yyVAL.value = &ConstTerm{Array: yyDollar[1].value.(*ConstArray)}
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:570
		{
			yyVAL.value = &ConstTerm{Number: yyDollar[1].token}
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:574
		{
			yyVAL.value = &ConstTerm{Str: yyDollar[1].token}
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:578
		{
			yyVAL.value = &ConstTerm{Null: true}
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:582
		{
			yyVAL.value = &ConstTerm{True: true}
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:586
		{
			yyVAL.value = &ConstTerm{False: true}
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:592
		{
			yyVAL.value = &ConstObject{}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:596
		{
			yyVAL.value = &ConstObject{yyDollar[2].value.([]*ConstObjectKeyVal)}
		}
	case 127:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:600
		{
			yyVAL.value = &ConstObject{yyDollar[2].value.([]*ConstObjectKeyVal)}
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:606
		{
			yyVAL.value = []*ConstObjectKeyVal{yyDollar[1].value.(*ConstObjectKeyVal)}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:610
		{
			yyVAL.value = append(yyDollar[1].value.([]*ConstObjectKeyVal), yyDollar[3].value.(*ConstObjectKeyVal))
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:616
		{
			yyVAL.value = &ConstObjectKeyVal{Key: yyDollar[1].token, Val: yyDollar[3].value.(*ConstTerm)}
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:620
		{
			yyVAL.value = &ConstObjectKeyVal{Key: yyDollar[1].token, Val: yyDollar[3].value.(*ConstTerm)}
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:624
		{
			yyVAL.value = &ConstObjectKeyVal{KeyString: yyDollar[1].token, Val: yyDollar[3].value.(*ConstTerm)}
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:630
		{
			yyVAL.value = &ConstArray{}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:634
		{
			yyVAL.value = &ConstArray{yyDollar[2].value.([]*ConstTerm)}
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:640
		{
			yyVAL.value = []*ConstTerm{yyDollar[1].value.(*ConstTerm)}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:644
		{
			yyVAL.value = append(yyDollar[1].value.([]*ConstTerm), yyDollar[3].value.(*ConstTerm))
		}
	}
	goto yystack /* stack new state and value */
}
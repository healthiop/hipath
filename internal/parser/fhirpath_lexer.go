// Code generated from fhirpath.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 66, 525,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22,
	3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3,
	29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3,
	39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3,
	47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48,
	3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3,
	49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51,
	3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3,
	53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54,
	3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3,
	55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57,
	3, 57, 3, 57, 3, 57, 5, 57, 388, 10, 57, 5, 57, 390, 10, 57, 3, 58, 3,
	58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59,
	3, 59, 3, 59, 5, 59, 406, 10, 59, 5, 59, 408, 10, 59, 3, 60, 3, 60, 3,
	60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 6, 60, 420, 10, 60,
	13, 60, 14, 60, 421, 5, 60, 424, 10, 60, 5, 60, 426, 10, 60, 5, 60, 428,
	10, 60, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 5, 61, 437, 10,
	61, 3, 62, 5, 62, 440, 10, 62, 3, 62, 7, 62, 443, 10, 62, 12, 62, 14, 62,
	446, 11, 62, 3, 63, 3, 63, 3, 63, 7, 63, 451, 10, 63, 12, 63, 14, 63, 454,
	11, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64, 7, 64, 461, 10, 64, 12, 64,
	14, 64, 464, 11, 64, 3, 64, 3, 64, 3, 65, 6, 65, 469, 10, 65, 13, 65, 14,
	65, 470, 3, 65, 3, 65, 6, 65, 475, 10, 65, 13, 65, 14, 65, 476, 5, 65,
	479, 10, 65, 3, 66, 6, 66, 482, 10, 66, 13, 66, 14, 66, 483, 3, 66, 3,
	66, 3, 67, 3, 67, 3, 67, 3, 67, 7, 67, 492, 10, 67, 12, 67, 14, 67, 495,
	11, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 68, 3, 68, 3, 68, 3, 68,
	7, 68, 506, 10, 68, 12, 68, 14, 68, 509, 11, 68, 3, 68, 3, 68, 3, 69, 3,
	69, 3, 69, 5, 69, 516, 10, 69, 3, 70, 3, 70, 3, 70, 3, 70, 3, 70, 3, 70,
	3, 71, 3, 71, 5, 452, 462, 493, 2, 72, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35,
	69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44,
	87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53,
	105, 54, 107, 55, 109, 56, 111, 57, 113, 58, 115, 59, 117, 2, 119, 2, 121,
	2, 123, 60, 125, 61, 127, 62, 129, 63, 131, 64, 133, 65, 135, 66, 137,
	2, 139, 2, 141, 2, 3, 2, 10, 3, 2, 50, 59, 4, 2, 45, 45, 47, 47, 5, 2,
	67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 5, 2, 11,
	12, 15, 15, 34, 34, 4, 2, 12, 12, 15, 15, 10, 2, 41, 41, 49, 49, 94, 94,
	98, 98, 104, 104, 112, 112, 116, 116, 118, 118, 5, 2, 50, 59, 67, 72, 99,
	104, 2, 539, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2,
	2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2,
	2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3,
	2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55,
	3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2,
	63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2,
	2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2,
	2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2,
	2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3,
	2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101,
	3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2,
	2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3,
	2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 2,
	129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3, 2,
	2, 2, 3, 143, 3, 2, 2, 2, 5, 145, 3, 2, 2, 2, 7, 147, 3, 2, 2, 2, 9, 149,
	3, 2, 2, 2, 11, 151, 3, 2, 2, 2, 13, 153, 3, 2, 2, 2, 15, 155, 3, 2, 2,
	2, 17, 157, 3, 2, 2, 2, 19, 161, 3, 2, 2, 2, 21, 165, 3, 2, 2, 2, 23, 167,
	3, 2, 2, 2, 25, 170, 3, 2, 2, 2, 27, 173, 3, 2, 2, 2, 29, 175, 3, 2, 2,
	2, 31, 178, 3, 2, 2, 2, 33, 180, 3, 2, 2, 2, 35, 182, 3, 2, 2, 2, 37, 185,
	3, 2, 2, 2, 39, 187, 3, 2, 2, 2, 41, 189, 3, 2, 2, 2, 43, 192, 3, 2, 2,
	2, 45, 195, 3, 2, 2, 2, 47, 198, 3, 2, 2, 2, 49, 207, 3, 2, 2, 2, 51, 211,
	3, 2, 2, 2, 53, 214, 3, 2, 2, 2, 55, 218, 3, 2, 2, 2, 57, 226, 3, 2, 2,
	2, 59, 228, 3, 2, 2, 2, 61, 230, 3, 2, 2, 2, 63, 232, 3, 2, 2, 2, 65, 234,
	3, 2, 2, 2, 67, 239, 3, 2, 2, 2, 69, 245, 3, 2, 2, 2, 71, 247, 3, 2, 2,
	2, 73, 253, 3, 2, 2, 2, 75, 260, 3, 2, 2, 2, 77, 267, 3, 2, 2, 2, 79, 269,
	3, 2, 2, 2, 81, 274, 3, 2, 2, 2, 83, 280, 3, 2, 2, 2, 85, 285, 3, 2, 2,
	2, 87, 289, 3, 2, 2, 2, 89, 294, 3, 2, 2, 2, 91, 301, 3, 2, 2, 2, 93, 308,
	3, 2, 2, 2, 95, 320, 3, 2, 2, 2, 97, 326, 3, 2, 2, 2, 99, 333, 3, 2, 2,
	2, 101, 339, 3, 2, 2, 2, 103, 344, 3, 2, 2, 2, 105, 350, 3, 2, 2, 2, 107,
	358, 3, 2, 2, 2, 109, 366, 3, 2, 2, 2, 111, 379, 3, 2, 2, 2, 113, 382,
	3, 2, 2, 2, 115, 391, 3, 2, 2, 2, 117, 395, 3, 2, 2, 2, 119, 409, 3, 2,
	2, 2, 121, 436, 3, 2, 2, 2, 123, 439, 3, 2, 2, 2, 125, 447, 3, 2, 2, 2,
	127, 457, 3, 2, 2, 2, 129, 468, 3, 2, 2, 2, 131, 481, 3, 2, 2, 2, 133,
	487, 3, 2, 2, 2, 135, 501, 3, 2, 2, 2, 137, 512, 3, 2, 2, 2, 139, 517,
	3, 2, 2, 2, 141, 523, 3, 2, 2, 2, 143, 144, 7, 48, 2, 2, 144, 4, 3, 2,
	2, 2, 145, 146, 7, 93, 2, 2, 146, 6, 3, 2, 2, 2, 147, 148, 7, 95, 2, 2,
	148, 8, 3, 2, 2, 2, 149, 150, 7, 45, 2, 2, 150, 10, 3, 2, 2, 2, 151, 152,
	7, 47, 2, 2, 152, 12, 3, 2, 2, 2, 153, 154, 7, 44, 2, 2, 154, 14, 3, 2,
	2, 2, 155, 156, 7, 49, 2, 2, 156, 16, 3, 2, 2, 2, 157, 158, 7, 102, 2,
	2, 158, 159, 7, 107, 2, 2, 159, 160, 7, 120, 2, 2, 160, 18, 3, 2, 2, 2,
	161, 162, 7, 111, 2, 2, 162, 163, 7, 113, 2, 2, 163, 164, 7, 102, 2, 2,
	164, 20, 3, 2, 2, 2, 165, 166, 7, 40, 2, 2, 166, 22, 3, 2, 2, 2, 167, 168,
	7, 107, 2, 2, 168, 169, 7, 117, 2, 2, 169, 24, 3, 2, 2, 2, 170, 171, 7,
	99, 2, 2, 171, 172, 7, 117, 2, 2, 172, 26, 3, 2, 2, 2, 173, 174, 7, 126,
	2, 2, 174, 28, 3, 2, 2, 2, 175, 176, 7, 62, 2, 2, 176, 177, 7, 63, 2, 2,
	177, 30, 3, 2, 2, 2, 178, 179, 7, 62, 2, 2, 179, 32, 3, 2, 2, 2, 180, 181,
	7, 64, 2, 2, 181, 34, 3, 2, 2, 2, 182, 183, 7, 64, 2, 2, 183, 184, 7, 63,
	2, 2, 184, 36, 3, 2, 2, 2, 185, 186, 7, 63, 2, 2, 186, 38, 3, 2, 2, 2,
	187, 188, 7, 128, 2, 2, 188, 40, 3, 2, 2, 2, 189, 190, 7, 35, 2, 2, 190,
	191, 7, 63, 2, 2, 191, 42, 3, 2, 2, 2, 192, 193, 7, 35, 2, 2, 193, 194,
	7, 128, 2, 2, 194, 44, 3, 2, 2, 2, 195, 196, 7, 107, 2, 2, 196, 197, 7,
	112, 2, 2, 197, 46, 3, 2, 2, 2, 198, 199, 7, 101, 2, 2, 199, 200, 7, 113,
	2, 2, 200, 201, 7, 112, 2, 2, 201, 202, 7, 118, 2, 2, 202, 203, 7, 99,
	2, 2, 203, 204, 7, 107, 2, 2, 204, 205, 7, 112, 2, 2, 205, 206, 7, 117,
	2, 2, 206, 48, 3, 2, 2, 2, 207, 208, 7, 99, 2, 2, 208, 209, 7, 112, 2,
	2, 209, 210, 7, 102, 2, 2, 210, 50, 3, 2, 2, 2, 211, 212, 7, 113, 2, 2,
	212, 213, 7, 116, 2, 2, 213, 52, 3, 2, 2, 2, 214, 215, 7, 122, 2, 2, 215,
	216, 7, 113, 2, 2, 216, 217, 7, 116, 2, 2, 217, 54, 3, 2, 2, 2, 218, 219,
	7, 107, 2, 2, 219, 220, 7, 111, 2, 2, 220, 221, 7, 114, 2, 2, 221, 222,
	7, 110, 2, 2, 222, 223, 7, 107, 2, 2, 223, 224, 7, 103, 2, 2, 224, 225,
	7, 117, 2, 2, 225, 56, 3, 2, 2, 2, 226, 227, 7, 42, 2, 2, 227, 58, 3, 2,
	2, 2, 228, 229, 7, 43, 2, 2, 229, 60, 3, 2, 2, 2, 230, 231, 7, 125, 2,
	2, 231, 62, 3, 2, 2, 2, 232, 233, 7, 127, 2, 2, 233, 64, 3, 2, 2, 2, 234,
	235, 7, 118, 2, 2, 235, 236, 7, 116, 2, 2, 236, 237, 7, 119, 2, 2, 237,
	238, 7, 103, 2, 2, 238, 66, 3, 2, 2, 2, 239, 240, 7, 104, 2, 2, 240, 241,
	7, 99, 2, 2, 241, 242, 7, 110, 2, 2, 242, 243, 7, 117, 2, 2, 243, 244,
	7, 103, 2, 2, 244, 68, 3, 2, 2, 2, 245, 246, 7, 39, 2, 2, 246, 70, 3, 2,
	2, 2, 247, 248, 7, 38, 2, 2, 248, 249, 7, 118, 2, 2, 249, 250, 7, 106,
	2, 2, 250, 251, 7, 107, 2, 2, 251, 252, 7, 117, 2, 2, 252, 72, 3, 2, 2,
	2, 253, 254, 7, 38, 2, 2, 254, 255, 7, 107, 2, 2, 255, 256, 7, 112, 2,
	2, 256, 257, 7, 102, 2, 2, 257, 258, 7, 103, 2, 2, 258, 259, 7, 122, 2,
	2, 259, 74, 3, 2, 2, 2, 260, 261, 7, 38, 2, 2, 261, 262, 7, 118, 2, 2,
	262, 263, 7, 113, 2, 2, 263, 264, 7, 118, 2, 2, 264, 265, 7, 99, 2, 2,
	265, 266, 7, 110, 2, 2, 266, 76, 3, 2, 2, 2, 267, 268, 7, 46, 2, 2, 268,
	78, 3, 2, 2, 2, 269, 270, 7, 123, 2, 2, 270, 271, 7, 103, 2, 2, 271, 272,
	7, 99, 2, 2, 272, 273, 7, 116, 2, 2, 273, 80, 3, 2, 2, 2, 274, 275, 7,
	111, 2, 2, 275, 276, 7, 113, 2, 2, 276, 277, 7, 112, 2, 2, 277, 278, 7,
	118, 2, 2, 278, 279, 7, 106, 2, 2, 279, 82, 3, 2, 2, 2, 280, 281, 7, 121,
	2, 2, 281, 282, 7, 103, 2, 2, 282, 283, 7, 103, 2, 2, 283, 284, 7, 109,
	2, 2, 284, 84, 3, 2, 2, 2, 285, 286, 7, 102, 2, 2, 286, 287, 7, 99, 2,
	2, 287, 288, 7, 123, 2, 2, 288, 86, 3, 2, 2, 2, 289, 290, 7, 106, 2, 2,
	290, 291, 7, 113, 2, 2, 291, 292, 7, 119, 2, 2, 292, 293, 7, 116, 2, 2,
	293, 88, 3, 2, 2, 2, 294, 295, 7, 111, 2, 2, 295, 296, 7, 107, 2, 2, 296,
	297, 7, 112, 2, 2, 297, 298, 7, 119, 2, 2, 298, 299, 7, 118, 2, 2, 299,
	300, 7, 103, 2, 2, 300, 90, 3, 2, 2, 2, 301, 302, 7, 117, 2, 2, 302, 303,
	7, 103, 2, 2, 303, 304, 7, 101, 2, 2, 304, 305, 7, 113, 2, 2, 305, 306,
	7, 112, 2, 2, 306, 307, 7, 102, 2, 2, 307, 92, 3, 2, 2, 2, 308, 309, 7,
	111, 2, 2, 309, 310, 7, 107, 2, 2, 310, 311, 7, 110, 2, 2, 311, 312, 7,
	110, 2, 2, 312, 313, 7, 107, 2, 2, 313, 314, 7, 117, 2, 2, 314, 315, 7,
	103, 2, 2, 315, 316, 7, 101, 2, 2, 316, 317, 7, 113, 2, 2, 317, 318, 7,
	112, 2, 2, 318, 319, 7, 102, 2, 2, 319, 94, 3, 2, 2, 2, 320, 321, 7, 123,
	2, 2, 321, 322, 7, 103, 2, 2, 322, 323, 7, 99, 2, 2, 323, 324, 7, 116,
	2, 2, 324, 325, 7, 117, 2, 2, 325, 96, 3, 2, 2, 2, 326, 327, 7, 111, 2,
	2, 327, 328, 7, 113, 2, 2, 328, 329, 7, 112, 2, 2, 329, 330, 7, 118, 2,
	2, 330, 331, 7, 106, 2, 2, 331, 332, 7, 117, 2, 2, 332, 98, 3, 2, 2, 2,
	333, 334, 7, 121, 2, 2, 334, 335, 7, 103, 2, 2, 335, 336, 7, 103, 2, 2,
	336, 337, 7, 109, 2, 2, 337, 338, 7, 117, 2, 2, 338, 100, 3, 2, 2, 2, 339,
	340, 7, 102, 2, 2, 340, 341, 7, 99, 2, 2, 341, 342, 7, 123, 2, 2, 342,
	343, 7, 117, 2, 2, 343, 102, 3, 2, 2, 2, 344, 345, 7, 106, 2, 2, 345, 346,
	7, 113, 2, 2, 346, 347, 7, 119, 2, 2, 347, 348, 7, 116, 2, 2, 348, 349,
	7, 117, 2, 2, 349, 104, 3, 2, 2, 2, 350, 351, 7, 111, 2, 2, 351, 352, 7,
	107, 2, 2, 352, 353, 7, 112, 2, 2, 353, 354, 7, 119, 2, 2, 354, 355, 7,
	118, 2, 2, 355, 356, 7, 103, 2, 2, 356, 357, 7, 117, 2, 2, 357, 106, 3,
	2, 2, 2, 358, 359, 7, 117, 2, 2, 359, 360, 7, 103, 2, 2, 360, 361, 7, 101,
	2, 2, 361, 362, 7, 113, 2, 2, 362, 363, 7, 112, 2, 2, 363, 364, 7, 102,
	2, 2, 364, 365, 7, 117, 2, 2, 365, 108, 3, 2, 2, 2, 366, 367, 7, 111, 2,
	2, 367, 368, 7, 107, 2, 2, 368, 369, 7, 110, 2, 2, 369, 370, 7, 110, 2,
	2, 370, 371, 7, 107, 2, 2, 371, 372, 7, 117, 2, 2, 372, 373, 7, 103, 2,
	2, 373, 374, 7, 101, 2, 2, 374, 375, 7, 113, 2, 2, 375, 376, 7, 112, 2,
	2, 376, 377, 7, 102, 2, 2, 377, 378, 7, 117, 2, 2, 378, 110, 3, 2, 2, 2,
	379, 380, 7, 66, 2, 2, 380, 381, 5, 117, 59, 2, 381, 112, 3, 2, 2, 2, 382,
	383, 7, 66, 2, 2, 383, 384, 5, 117, 59, 2, 384, 389, 7, 86, 2, 2, 385,
	387, 5, 119, 60, 2, 386, 388, 5, 121, 61, 2, 387, 386, 3, 2, 2, 2, 387,
	388, 3, 2, 2, 2, 388, 390, 3, 2, 2, 2, 389, 385, 3, 2, 2, 2, 389, 390,
	3, 2, 2, 2, 390, 114, 3, 2, 2, 2, 391, 392, 7, 66, 2, 2, 392, 393, 7, 86,
	2, 2, 393, 394, 5, 119, 60, 2, 394, 116, 3, 2, 2, 2, 395, 396, 9, 2, 2,
	2, 396, 397, 9, 2, 2, 2, 397, 398, 9, 2, 2, 2, 398, 407, 9, 2, 2, 2, 399,
	400, 7, 47, 2, 2, 400, 401, 9, 2, 2, 2, 401, 405, 9, 2, 2, 2, 402, 403,
	7, 47, 2, 2, 403, 404, 9, 2, 2, 2, 404, 406, 9, 2, 2, 2, 405, 402, 3, 2,
	2, 2, 405, 406, 3, 2, 2, 2, 406, 408, 3, 2, 2, 2, 407, 399, 3, 2, 2, 2,
	407, 408, 3, 2, 2, 2, 408, 118, 3, 2, 2, 2, 409, 410, 9, 2, 2, 2, 410,
	427, 9, 2, 2, 2, 411, 412, 7, 60, 2, 2, 412, 413, 9, 2, 2, 2, 413, 425,
	9, 2, 2, 2, 414, 415, 7, 60, 2, 2, 415, 416, 9, 2, 2, 2, 416, 423, 9, 2,
	2, 2, 417, 419, 7, 48, 2, 2, 418, 420, 9, 2, 2, 2, 419, 418, 3, 2, 2, 2,
	420, 421, 3, 2, 2, 2, 421, 419, 3, 2, 2, 2, 421, 422, 3, 2, 2, 2, 422,
	424, 3, 2, 2, 2, 423, 417, 3, 2, 2, 2, 423, 424, 3, 2, 2, 2, 424, 426,
	3, 2, 2, 2, 425, 414, 3, 2, 2, 2, 425, 426, 3, 2, 2, 2, 426, 428, 3, 2,
	2, 2, 427, 411, 3, 2, 2, 2, 427, 428, 3, 2, 2, 2, 428, 120, 3, 2, 2, 2,
	429, 437, 7, 92, 2, 2, 430, 431, 9, 3, 2, 2, 431, 432, 9, 2, 2, 2, 432,
	433, 9, 2, 2, 2, 433, 434, 7, 60, 2, 2, 434, 435, 9, 2, 2, 2, 435, 437,
	9, 2, 2, 2, 436, 429, 3, 2, 2, 2, 436, 430, 3, 2, 2, 2, 437, 122, 3, 2,
	2, 2, 438, 440, 9, 4, 2, 2, 439, 438, 3, 2, 2, 2, 440, 444, 3, 2, 2, 2,
	441, 443, 9, 5, 2, 2, 442, 441, 3, 2, 2, 2, 443, 446, 3, 2, 2, 2, 444,
	442, 3, 2, 2, 2, 444, 445, 3, 2, 2, 2, 445, 124, 3, 2, 2, 2, 446, 444,
	3, 2, 2, 2, 447, 452, 7, 98, 2, 2, 448, 451, 5, 137, 69, 2, 449, 451, 11,
	2, 2, 2, 450, 448, 3, 2, 2, 2, 450, 449, 3, 2, 2, 2, 451, 454, 3, 2, 2,
	2, 452, 453, 3, 2, 2, 2, 452, 450, 3, 2, 2, 2, 453, 455, 3, 2, 2, 2, 454,
	452, 3, 2, 2, 2, 455, 456, 7, 98, 2, 2, 456, 126, 3, 2, 2, 2, 457, 462,
	7, 41, 2, 2, 458, 461, 5, 137, 69, 2, 459, 461, 11, 2, 2, 2, 460, 458,
	3, 2, 2, 2, 460, 459, 3, 2, 2, 2, 461, 464, 3, 2, 2, 2, 462, 463, 3, 2,
	2, 2, 462, 460, 3, 2, 2, 2, 463, 465, 3, 2, 2, 2, 464, 462, 3, 2, 2, 2,
	465, 466, 7, 41, 2, 2, 466, 128, 3, 2, 2, 2, 467, 469, 9, 2, 2, 2, 468,
	467, 3, 2, 2, 2, 469, 470, 3, 2, 2, 2, 470, 468, 3, 2, 2, 2, 470, 471,
	3, 2, 2, 2, 471, 478, 3, 2, 2, 2, 472, 474, 7, 48, 2, 2, 473, 475, 9, 2,
	2, 2, 474, 473, 3, 2, 2, 2, 475, 476, 3, 2, 2, 2, 476, 474, 3, 2, 2, 2,
	476, 477, 3, 2, 2, 2, 477, 479, 3, 2, 2, 2, 478, 472, 3, 2, 2, 2, 478,
	479, 3, 2, 2, 2, 479, 130, 3, 2, 2, 2, 480, 482, 9, 6, 2, 2, 481, 480,
	3, 2, 2, 2, 482, 483, 3, 2, 2, 2, 483, 481, 3, 2, 2, 2, 483, 484, 3, 2,
	2, 2, 484, 485, 3, 2, 2, 2, 485, 486, 8, 66, 2, 2, 486, 132, 3, 2, 2, 2,
	487, 488, 7, 49, 2, 2, 488, 489, 7, 44, 2, 2, 489, 493, 3, 2, 2, 2, 490,
	492, 11, 2, 2, 2, 491, 490, 3, 2, 2, 2, 492, 495, 3, 2, 2, 2, 493, 494,
	3, 2, 2, 2, 493, 491, 3, 2, 2, 2, 494, 496, 3, 2, 2, 2, 495, 493, 3, 2,
	2, 2, 496, 497, 7, 44, 2, 2, 497, 498, 7, 49, 2, 2, 498, 499, 3, 2, 2,
	2, 499, 500, 8, 67, 2, 2, 500, 134, 3, 2, 2, 2, 501, 502, 7, 49, 2, 2,
	502, 503, 7, 49, 2, 2, 503, 507, 3, 2, 2, 2, 504, 506, 10, 7, 2, 2, 505,
	504, 3, 2, 2, 2, 506, 509, 3, 2, 2, 2, 507, 505, 3, 2, 2, 2, 507, 508,
	3, 2, 2, 2, 508, 510, 3, 2, 2, 2, 509, 507, 3, 2, 2, 2, 510, 511, 8, 68,
	2, 2, 511, 136, 3, 2, 2, 2, 512, 515, 7, 94, 2, 2, 513, 516, 9, 8, 2, 2,
	514, 516, 5, 139, 70, 2, 515, 513, 3, 2, 2, 2, 515, 514, 3, 2, 2, 2, 516,
	138, 3, 2, 2, 2, 517, 518, 7, 119, 2, 2, 518, 519, 5, 141, 71, 2, 519,
	520, 5, 141, 71, 2, 520, 521, 5, 141, 71, 2, 521, 522, 5, 141, 71, 2, 522,
	140, 3, 2, 2, 2, 523, 524, 9, 9, 2, 2, 524, 142, 3, 2, 2, 2, 26, 2, 387,
	389, 405, 407, 421, 423, 425, 427, 436, 439, 442, 444, 450, 452, 460, 462,
	470, 476, 478, 483, 493, 507, 515, 3, 2, 3, 2,
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
	"", "'.'", "'['", "']'", "'+'", "'-'", "'*'", "'/'", "'div'", "'mod'",
	"'&'", "'is'", "'as'", "'|'", "'<='", "'<'", "'>'", "'>='", "'='", "'~'",
	"'!='", "'!~'", "'in'", "'contains'", "'and'", "'or'", "'xor'", "'implies'",
	"'('", "')'", "'{'", "'}'", "'true'", "'false'", "'%'", "'$this'", "'$index'",
	"'$total'", "','", "'year'", "'month'", "'week'", "'day'", "'hour'", "'minute'",
	"'second'", "'millisecond'", "'years'", "'months'", "'weeks'", "'days'",
	"'hours'", "'minutes'", "'seconds'", "'milliseconds'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "DATE", "DATETIME", "TIME", "IDENTIFIER", "DELIMITEDIDENTIFIER", "STRING",
	"NUMBER", "WS", "COMMENT", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
	"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48",
	"T__49", "T__50", "T__51", "T__52", "T__53", "DATE", "DATETIME", "TIME",
	"DATEFORMAT", "TIMEFORMAT", "TIMEZONEOFFSETFORMAT", "IDENTIFIER", "DELIMITEDIDENTIFIER",
	"STRING", "NUMBER", "WS", "COMMENT", "LINE_COMMENT", "ESC", "UNICODE",
	"HEX",
}

type fhirpathLexer struct {
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

func NewfhirpathLexer(input antlr.CharStream) *fhirpathLexer {

	l := new(fhirpathLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "fhirpath.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// fhirpathLexer tokens.
const (
	fhirpathLexerT__0                = 1
	fhirpathLexerT__1                = 2
	fhirpathLexerT__2                = 3
	fhirpathLexerT__3                = 4
	fhirpathLexerT__4                = 5
	fhirpathLexerT__5                = 6
	fhirpathLexerT__6                = 7
	fhirpathLexerT__7                = 8
	fhirpathLexerT__8                = 9
	fhirpathLexerT__9                = 10
	fhirpathLexerT__10               = 11
	fhirpathLexerT__11               = 12
	fhirpathLexerT__12               = 13
	fhirpathLexerT__13               = 14
	fhirpathLexerT__14               = 15
	fhirpathLexerT__15               = 16
	fhirpathLexerT__16               = 17
	fhirpathLexerT__17               = 18
	fhirpathLexerT__18               = 19
	fhirpathLexerT__19               = 20
	fhirpathLexerT__20               = 21
	fhirpathLexerT__21               = 22
	fhirpathLexerT__22               = 23
	fhirpathLexerT__23               = 24
	fhirpathLexerT__24               = 25
	fhirpathLexerT__25               = 26
	fhirpathLexerT__26               = 27
	fhirpathLexerT__27               = 28
	fhirpathLexerT__28               = 29
	fhirpathLexerT__29               = 30
	fhirpathLexerT__30               = 31
	fhirpathLexerT__31               = 32
	fhirpathLexerT__32               = 33
	fhirpathLexerT__33               = 34
	fhirpathLexerT__34               = 35
	fhirpathLexerT__35               = 36
	fhirpathLexerT__36               = 37
	fhirpathLexerT__37               = 38
	fhirpathLexerT__38               = 39
	fhirpathLexerT__39               = 40
	fhirpathLexerT__40               = 41
	fhirpathLexerT__41               = 42
	fhirpathLexerT__42               = 43
	fhirpathLexerT__43               = 44
	fhirpathLexerT__44               = 45
	fhirpathLexerT__45               = 46
	fhirpathLexerT__46               = 47
	fhirpathLexerT__47               = 48
	fhirpathLexerT__48               = 49
	fhirpathLexerT__49               = 50
	fhirpathLexerT__50               = 51
	fhirpathLexerT__51               = 52
	fhirpathLexerT__52               = 53
	fhirpathLexerT__53               = 54
	fhirpathLexerDATE                = 55
	fhirpathLexerDATETIME            = 56
	fhirpathLexerTIME                = 57
	fhirpathLexerIDENTIFIER          = 58
	fhirpathLexerDELIMITEDIDENTIFIER = 59
	fhirpathLexerSTRING              = 60
	fhirpathLexerNUMBER              = 61
	fhirpathLexerWS                  = 62
	fhirpathLexerCOMMENT             = 63
	fhirpathLexerLINE_COMMENT        = 64
)

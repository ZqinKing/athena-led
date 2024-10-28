package athenaLed

// 字符字典
// 从左到右，每个字节代表一列led
// 只使用字节的低5位，低位到高位是自顶向下
var charDict = map[rune][]byte{
	'A': {0b00011110, 0b00001001, 0b00001001, 0b00011110, 0b00000000},
	'B': {0b00011111, 0b00010101, 0b00010101, 0b00001110, 0b00000000},
	'C': {0b00001110, 0b00010001, 0b00010001, 0b00010001, 0b00000000},
	'D': {0b00011111, 0b00010001, 0b00010001, 0b00001110, 0b00000000},
	'E': {0b00011111, 0b00010101, 0b00010101, 0b00010101, 0b00000000},
	'F': {0b00011111, 0b00000101, 0b00000101, 0b00000101, 0b00000000},
	'G': {0b00011111, 0b00010001, 0b00010101, 0b00011101, 0b00000000},
	'H': {0b00011111, 0b00000100, 0b00000100, 0b00011111, 0b00000000},
	'I': {0b00000000, 0b00010001, 0b00011111, 0b00010001, 0b00000000},
	'J': {0b00000000, 0b00010000, 0b00010000, 0b00001111, 0b00000000},
	'K': {0b00011111, 0b00000100, 0b00001010, 0b00010001, 0b00000000},
	'L': {0b00011111, 0b00010000, 0b00010000, 0b00010000, 0b00000000},
	'M': {0b00011111, 0b00000001, 0b00011110, 0b00000001, 0b00011111, 0b00000000},
	'N': {0b00011111, 0b00000010, 0b00000100, 0b00001000, 0b00011111, 0b00000000},
	'O': {0b00011111, 0b00010001, 0b00010001, 0b00011111, 0b00000000},
	'P': {0b00011111, 0b00000101, 0b00000101, 0b00000111, 0b00000000},
	'Q': {0b00001111, 0b00001001, 0b00001001, 0b00011111, 0b00000000},
	'R': {0b00011111, 0b00000101, 0b00001101, 0b00010010, 0b00000000},
	'S': {0b00010010, 0b00010101, 0b00010101, 0b00001001, 0b00000000},
	'T': {0b00000001, 0b00000001, 0b00011111, 0b00000001, 0b00000001, 0b00000000},
	'U': {0b00011111, 0b00010000, 0b00010000, 0b00011111, 0b00000000},
	'V': {0b00001111, 0b00010000, 0b00010000, 0b00001111, 0b00000000},
	'W': {0b00000111, 0b00011000, 0b00000111, 0b00011000, 0b00000111, 0b00000000},
	'X': {0b00010001, 0b00001010, 0b00000100, 0b00001010, 0b00010001, 0b00000000},
	'Y': {0b00000001, 0b00000010, 0b00011100, 0b00000010, 0b00000001, 0b00000000},
	'Z': {0b00010001, 0b00011001, 0b00010101, 0b00010011, 0b00010001, 0b00000000},
	' ': {0b00000000},
	'.': {0b00010000, 0b00000000},
	'℃': {0b00000111, 0b00000101, 0b00000111, 0b00000000, 0b00001110, 0b00010001, 0b00010001, 0b00010001, 0b00000000},
	'-': {0b00000100, 0b00000100, 0b00000000},
	'+': {0b00000100, 0b00001110, 0b00000100, 0b00000000},
	'*': {0b00010101, 0b00001110, 0b00011111, 0b00001110, 0b00010101, 0b00000000},
	'/': {0b00010000, 0b00001000, 0b00000100, 0b00000010, 0b00000001, 0b00000000},
	'=': {0b00001010, 0b00001010, 0b00001010, 0b00001010, 0b00001010, 0b00000000},
	':': {0b00001010, 0b00000000},
	'：': {0b00001010, 0b00000000},
	'1': {0b00010010, 0b00011111, 0b00010000, 0b00000000},
	'2': {0b00011101, 0b00010101, 0b00010111, 0b00000000},
	'3': {0b00010101, 0b00010101, 0b00011111, 0b00000000},
	'4': {0b00000111, 0b00000100, 0b00011111, 0b00000000},
	'5': {0b00010111, 0b00010101, 0b00011101, 0b00000000},
	'6': {0b00011111, 0b00010101, 0b00011101, 0b00000000},
	'7': {0b00000001, 0b00000001, 0b00011111, 0b00000000},
	'8': {0b00011111, 0b00010101, 0b00011111, 0b00000000},
	'9': {0b00010111, 0b00010101, 0b00011111, 0b00000000},
	'0': {0b00011111, 0b00010001, 0b00011111, 0b00000000},
}
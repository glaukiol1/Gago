package parser

type Token struct {
	pos   int // position of the token (inside the line)
	line  int
	value string // the value of the token its-self
	code  int    // the token code (specified in token_codes.go)
}

func NewToken(token string, pos int, line int) *Token {
	return &Token{pos, line, token, getTokenCode(token)}
}

func getTokenCode(token string) int {
	switch token {
	case "a":
		return CHAR_A_LOWER
	case "b":
		return CHAR_B_LOWER
	case "c":
		return CHAR_C_LOWER
	case "d":
		return CHAR_D_LOWER
	case "e":
		return CHAR_E_LOWER
	case "f":
		return CHAR_F_LOWER
	case "g":
		return CHAR_G_LOWER
	case "h":
		return CHAR_H_LOWER
	case "i":
		return CHAR_I_LOWER
	case "j":
		return CHAR_J_LOWER
	case "k":
		return CHAR_K_LOWER
	case "l":
		return CHAR_L_LOWER
	case "m":
		return CHAR_M_LOWER
	case "n":
		return CHAR_N_LOWER
	case "o":
		return CHAR_O_LOWER
	case "p":
		return CHAR_P_LOWER
	case "q":
		return CHAR_Q_LOWER
	case "r":
		return CHAR_R_LOWER
	case "s":
		return CHAR_S_LOWER
	case "t":
		return CHAR_T_LOWER
	case "u":
		return CHAR_U_LOWER
	case "v":
		return CHAR_V_LOWER
	case "w":
		return CHAR_W_LOWER
	case "x":
		return CHAR_X_LOWER
	case "y":
		return CHAR_Y_LOWER
	case "z":
		return CHAR_Z_LOWER
	case "A":
		return CHAR_A_UPPER
	case "B":
		return CHAR_B_UPPER
	case "C":
		return CHAR_C_UPPER
	case "D":
		return CHAR_D_UPPER
	case "E":
		return CHAR_E_UPPER
	case "F":
		return CHAR_F_UPPER
	case "G":
		return CHAR_G_UPPER
	case "H":
		return CHAR_H_UPPER
	case "I":
		return CHAR_I_UPPER
	case "J":
		return CHAR_J_UPPER
	case "K":
		return CHAR_K_UPPER
	case "L":
		return CHAR_L_UPPER
	case "M":
		return CHAR_M_UPPER
	case "N":
		return CHAR_N_UPPER
	case "O":
		return CHAR_O_UPPER
	case "P":
		return CHAR_P_UPPER
	case "Q":
		return CHAR_Q_UPPER
	case "R":
		return CHAR_R_UPPER
	case "S":
		return CHAR_S_UPPER
	case "T":
		return CHAR_T_UPPER
	case "U":
		return CHAR_U_UPPER
	case "V":
		return CHAR_V_UPPER
	case "W":
		return CHAR_W_UPPER
	case "X":
		return CHAR_X_UPPER
	case "Y":
		return CHAR_Y_UPPER
	case "Z":
		return CHAR_Z_UPPER
	case "0":
		return NUM_0
	case "1":
		return NUM_1
	case "2":
		return NUM_2
	case "3":
		return NUM_3
	case "4":
		return NUM_4
	case "5":
		return NUM_5
	case "6":
		return NUM_6
	case "7":
		return NUM_7
	case "8":
		return NUM_8
	case "9":
		return NUM_9
	}
	return -1
}
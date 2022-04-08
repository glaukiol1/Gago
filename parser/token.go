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
	}
	return -1
}

package lexer

// int representation for tokens

/* alphabet a-z A-Z start codes */
const CHAR_A_LOWER = 0
const CHAR_B_LOWER = 1
const CHAR_C_LOWER = 2
const CHAR_D_LOWER = 3
const CHAR_E_LOWER = 4
const CHAR_F_LOWER = 5
const CHAR_G_LOWER = 6
const CHAR_H_LOWER = 7
const CHAR_I_LOWER = 8
const CHAR_J_LOWER = 9
const CHAR_K_LOWER = 10
const CHAR_L_LOWER = 11
const CHAR_M_LOWER = 12
const CHAR_N_LOWER = 13
const CHAR_O_LOWER = 14
const CHAR_P_LOWER = 15
const CHAR_Q_LOWER = 16
const CHAR_R_LOWER = 17
const CHAR_S_LOWER = 18
const CHAR_T_LOWER = 19
const CHAR_U_LOWER = 20
const CHAR_V_LOWER = 21
const CHAR_W_LOWER = 22
const CHAR_X_LOWER = 23
const CHAR_Y_LOWER = 24
const CHAR_Z_LOWER = 25

const CHAR_A_UPPER = 26
const CHAR_B_UPPER = 27
const CHAR_C_UPPER = 28
const CHAR_D_UPPER = 29
const CHAR_E_UPPER = 30
const CHAR_F_UPPER = 31
const CHAR_G_UPPER = 32
const CHAR_H_UPPER = 33
const CHAR_I_UPPER = 34
const CHAR_J_UPPER = 35
const CHAR_K_UPPER = 36
const CHAR_L_UPPER = 37
const CHAR_M_UPPER = 38
const CHAR_N_UPPER = 39
const CHAR_O_UPPER = 40
const CHAR_P_UPPER = 41
const CHAR_Q_UPPER = 42
const CHAR_R_UPPER = 43
const CHAR_S_UPPER = 44
const CHAR_T_UPPER = 45
const CHAR_U_UPPER = 46
const CHAR_V_UPPER = 47
const CHAR_W_UPPER = 48
const CHAR_X_UPPER = 49
const CHAR_Y_UPPER = 50
const CHAR_Z_UPPER = 51

/* alphabet a-z end codes */

/* numbers 0-9 start codes */

const NUM_0 = 52
const NUM_1 = 53
const NUM_2 = 54
const NUM_3 = 55
const NUM_4 = 56
const NUM_5 = 57
const NUM_6 = 58
const NUM_7 = 59
const NUM_8 = 60
const NUM_9 = 61

/* numbers 0-9 end codes */

/* other chars start codes */

const OPEN_SQUARE_BRACKET = 62
const CLOSE_SQUARE_BRACKET = 63
const OPEN_CURLY_BRACKET = 64
const CLOSE_CURLY_BRACKET = 65
const SEMICOLON = 66
const COLON = 67
const FORWARD_SLASH = 68
const BACK_SLASH = 69
const DOUBLE_QUOTE = 70
const SINGLE_QUOTE = 71
const DOT = 72
const COMMA = 73
const LESS_THAN = 74
const GREATER_THAN = 75
const QUESTION_MARK = 76
const VERTICAL_BAR = 77
const EQUALS = 78
const PLUS = 79
const HYPHEN = 80
const UNDERSCORE = 81
const EXCLAMATION_MARK = 82
const AT_SIGN = 83
const NUMBER_SIGN = 84
const DOLLAR_SIGN = 85
const PERCENTAGE_SIGN = 86
const CARET = 87    // ^
const AND_SIGN = 88 // &
const ASTERIK = 89  // *
const TILDE = 90    // ~
const BACKTICK = 91 // `
const OPEN_PAREN = 92
const CLOSE_PAREN = 93
const SPACE = 94
const NEWLINE = 95
const EOF = 96

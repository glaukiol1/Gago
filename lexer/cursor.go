package lexer

import "fmt"

type Cursor struct {
	line  *Line
	Token *Token // the token of this Cursor
	Pos   int    // position in line
}

func NewCursor(line *Line, token string, pos int) (*Cursor, error) {
	cursor := &Cursor{line, nil, pos}
	err := cursor.tokenize(token)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

func (cursor *Cursor) tokenize(token string) error {
	cursor.Token = NewToken(token, cursor.Pos, cursor.line.linepos)
	if cursor.Token.code == -1 {
		cursor.line.error = true
		cursor.line.errorString = "TokenError: Unrecognized token at " + cursor.line.lexer.filename + ":" + fmt.Sprint(cursor.line.linepos) + ":" + fmt.Sprint(cursor.Pos)
		return fmt.Errorf("")
	}
	if cursor.Token.code == SPACE {
		cursor.line.indents += 1
	}
	if cursor.Token.code == OPEN_CURLY_BRACKET || cursor.Token.code == CLOSE_CURLY_BRACKET {
		cursor.line.curly_braces += 1
	}
	if cursor.Token.code == OPEN_PAREN || cursor.Token.code == CLOSE_PAREN {
		cursor.line.parenthesis += 1
	}
	if cursor.Token.code == OPEN_SQUARE_BRACKET || cursor.Token.code == CLOSE_SQUARE_BRACKET {
		cursor.line.bracket += 1
	}
	return nil
}

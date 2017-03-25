package main

import (
"strings"
"fmt"
)

type Express interface {
	interpret(context string) bool
}

type TerminalExpress struct {
	data string
}

func (t *TerminalExpress)interpret(context string) bool {
	if (strings.Contains(context, t.data)) {
		return true;
	}
	return false;
}

type AndExpress struct {
	exp1 *TerminalExpress
	exp2 *TerminalExpress
}

func (a *AndExpress)interpret(context string) bool {
	return a.exp1.interpret(context) && a.exp2.interpret(context);
}

type OrExpress struct {
	exp1 *TerminalExpress
	exp2 *TerminalExpress
}

func (a *OrExpress)interpret(context string) bool {
	return a.exp1.interpret(context) || a.exp2.interpret(context);
}

func isMale() Express {
	robert := &TerminalExpress{"Robert"}
	join := &TerminalExpress{"Join"}
	return &OrExpress{robert, join}
}

func isMarriedWoman() Express {
	married := &TerminalExpress{"Married"}
	julie := &TerminalExpress{"Julie"}
	return &AndExpress{married, julie}
}

func main() {
	isMale := isMale()
	isMarriedWoman := isMarriedWoman()
	fmt.Println("John is male? ", isMale.interpret("Join"));
	fmt.Println("Juile is a married women?", isMarriedWoman.interpret("Married Julie"));
}

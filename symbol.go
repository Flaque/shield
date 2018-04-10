package main

type Symbol int

const (
	FUNC Symbol = iota
	LABEL
	COLON
	L_PAREN
	R_PAREN
	END
)

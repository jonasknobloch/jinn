package corenlp

type Token struct {
	Index                int    `json:"index"`
	Word                 string `json:"word"`
	OriginalText         string `json:"originalText"`
	CharacterOffsetBegin int    `json:"characterOffsetBegin"`
	CharacterOffsetEnd   int    `json:"characterOffsetEnd"`
	Pos                  string `json:"pos"`
	Before               string `json:"before"`
	After                string `json:"after"`
}

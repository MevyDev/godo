package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type task struct {
	Text		string	`json:"text"`
	Status		string	`json:"status"`
	Difficulty	string	`json:"difficulty"`
}

package main

import (
	"net/http"
	"strconv"
)

type LevelState uint16

func ParseStateFromCookie(c *http.Cookie) (LevelState, error) {
  if c == nil {
		return LevelState(0), nil
	}
	value, err := strconv.ParseUint(c.Value, 10, 8)
	return LevelState(value), err
}

func (ls *LevelState) String() string {
	return strconv.FormatUint(uint64(*ls), 10)
}

func (ls *LevelState) Num() uint16 {
	return uint16(*ls)
}

type GameLevel interface {
	ValidRequest(r *http.Request, ls *LevelState) error
	IsComplete(r *http.Request, ls *LevelState) (bool, error)
	Render(w http.ResponseWriter, r *http.Request, ls *LevelState)
}

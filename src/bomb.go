package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

const LevelCookieId = "state"

func serveIcon(w http.ResponseWriter, r *http.Request) {
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	all_levels := [...]GameLevel{new(CookieLevel)}

	level_cookie, err := r.Cookie(LevelCookieId)
	level_state, err := ParseStateFromCookie(level_cookie)
	l := all_levels[level_state.Num()]
	err = l.ValidRequest(r, &level_state)
	if err != nil {
		explode(w, r, err)
		return
	}
	done, err := l.IsComplete(r, &level_state)
	if err != nil {
		explode(w, r, err)
		return
	}
	if done {
		fmt.Fprintln(w, "onwards")
	} else {
		l.Render(w, r, &level_state)
	}
}

func explode(w http.ResponseWriter, r *http.Request, err error) {
	for _, cookie := range r.Cookies() {
		cookie.MaxAge = -1
		cookie.Value = ""
		http.SetCookie(w, cookie)
	}
	fmt.Fprintf(w, "bang %s", err)
}

func main() {
  http.HandleFunc("/favicon.ico", serveIcon)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

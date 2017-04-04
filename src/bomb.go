package main

import (
	"fmt"
	"net/http"
)

const LevelCookieId = "state"

func serveIcon(w http.ResponseWriter, r *http.Request) {
	return
}

func handleLevel(w http.ResponseWriter, r *http.Request, l GameLevel,
	level_state LevelState) {
	err := l.ValidRequest(r, &level_state)
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

func handler(w http.ResponseWriter, r *http.Request) {
	all_levels := [...]GameLevel{new(CookieLevel)}

	level_cookie, _ := r.Cookie(LevelCookieId)
	level_state, _ := ParseStateFromCookie(level_cookie)
	handleLevel(w, r, all_levels[level_state.Num()], level_state)
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

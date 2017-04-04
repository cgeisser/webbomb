package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

type CookieLevel bool

func (cl *CookieLevel) ValidRequest(r *http.Request, ls *LevelState) error {
	if r.URL.Path == "/" ||
		r.URL.Path == "/eat_cookie" {
		return nil
	}
	return errors.New("you shouldn't be here")
}

func (cl *CookieLevel) IsComplete(r *http.Request, ls *LevelState) (bool, error) {
	if r.URL.Path == "/eat_cookie" {
		c, err := r.Cookie("sekret")
		fmt.Println(c)
		fmt.Println(err)
		fmt.Println(r.FormValue("cookie"))
		if err == nil &&
			c.Value == r.FormValue("cookie") {
			return true, nil
		}
		return false, errors.New("yuck")
	}
	return false, nil
}

func (cl *CookieLevel) Render(w http.ResponseWriter, r *http.Request, ls *LevelState) {
	t, err := template.ParseFiles("level1.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "sekret",
		Value: "cookie",
	})
	t.Execute(w, nil)
}

package handlers

import (
    "github.com/kollekcioner47/finance-app/internal/service"
    "github.com/kollekcioner47/finance-app/internal/session"
    "net/http"
)

type UserHandler struct {
    userService *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
    return &UserHandler{userService: us}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        renderTemplate(w, "auth/register.html", nil)
        return
    }
    // POST
    email := r.FormValue("email")
    password := r.FormValue("password")
    _, err := h.userService.Register(email, password)
    if err != nil {
        renderTemplate(w, "auth/register.html", map[string]interface{}{"Error": err.Error()})
        return
    }
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        renderTemplate(w, "auth/login.html", nil)
        return
    }
    email := r.FormValue("email")
    password := r.FormValue("password")
    user, err := h.userService.Login(email, password)
    if err != nil {
        renderTemplate(w, "auth/login.html", map[string]interface{}{"Error": "Invalid credentials"})
        return
    }
    sess, _ := session.Store.Get(r, "finance-session")
    sess.Values["authenticated"] = true
    sess.Values["userID"] = user.ID
    sess.Save(r, w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
    sess, _ := session.Store.Get(r, "finance-session")
    sess.Values["authenticated"] = false
    delete(sess.Values, "userID")
    sess.Save(r, w)
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}

package middleware

import (
	"context"
	"net/http"

	"github.com/kollekcioner47/finance-app/internal/session"
)

func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, _ := session.Store.Get(r, "finance-session")
		if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		// Добавляем userID в контекст для дальнейшего использования
		ctx := context.WithValue(r.Context(), "userID", sess.Values["userID"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

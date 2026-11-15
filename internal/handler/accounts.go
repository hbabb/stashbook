// Package handler holds the HTTP handlers for Stashbook's API routes.
package handler

import (
	"log/slog"
	"net/http"

	"github.com/hbabb/stashbook/internal/repository"
)

// Accounts handles the /api/accounts routes.
type Accounts struct {
	queries *repository.Queries
}

// NewAccounts creates the accounts handler.
func NewAccounts(queries *repository.Queries) *Accounts {
	return &Accounts{queries: queries}
}

// List returns every account as JSON.
func (h *Accounts) List(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.queries.ListAccounts(r.Context())
	if err != nil {
		slog.Error("list accounts", "err", err)
		writeError(w, http.StatusInternalServerError, "could not load accounts")
		return
	}

	if accounts == nil {
		accounts = []repository.Account{}
	}

	writeJSON(w, http.StatusOK, accounts)
}

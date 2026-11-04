// Package handler holds the HTTP handlers for Stashbook's API routes.
package handler

import (
	"encoding/json"
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
		http.Error(w, "could not load accounts", http.StatusInternalServerError)
		return
	}

	if accounts == nil {
		accounts = []repository.Account{}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		slog.Error("encode accounts", "err", err)
	}
}

// Package handler holds the HTTP handlers for Stashbook's API routes.
package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

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

type (
	// accountRequest is the JSON request body for creating or updating an account.
	accountRequest struct {
		Name                string                  `json:"name"`
		Type                repository.AccountType  `json:"type"`
		Owner               repository.AccountOwner `json:"owner"`
		OpeningBalanceCents int64                   `json:"opening_balance_cents"`
	}
)

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

// Get returns one account by its ID as JSON.
func (h *Accounts) Get(w http.ResponseWriter, r *http.Request) {
	id, err := accountID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid account id")
		return
	}

	account, err := h.queries.GetAccount(r.Context(), id)
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, http.StatusNotFound, "account not found")
		return
	}
	if err != nil {
		slog.Error("get account", "id", id, "err", err)
		writeError(w, http.StatusInternalServerError, "could not load account")
		return
	}

	writeJSON(w, http.StatusOK, account)
}

// Create adds a new account from the JSON body.
func (h *Accounts) Create(w http.ResponseWriter, r *http.Request) {
	req, err := readAccountRequest(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	account, err := h.queries.CreateAccount(r.Context(), repository.CreateAccountParams{
		Name:                req.Name,
		Type:                req.Type,
		Owner:               req.Owner,
		OpeningBalanceCents: req.OpeningBalanceCents,
	})
	if err != nil {
		slog.Error("create account", "err", err)
		writeError(w, http.StatusInternalServerError, "could not create account")
		return
	}

	writeJSON(w, http.StatusCreated, account)
}

// Update replaces an account's fields with the JSON body.
func (h *Accounts) Update(w http.ResponseWriter, r *http.Request) {
	id, err := accountID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid account id")
		return
	}

	req, err := readAccountRequest(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	account, err := h.queries.UpdateAccount(r.Context(), repository.UpdateAccountParams{
		ID:                  id,
		Name:                req.Name,
		Type:                req.Type,
		Owner:               req.Owner,
		OpeningBalanceCents: req.OpeningBalanceCents,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, http.StatusNotFound, "account not found")
		return
	}
	if err != nil {
		slog.Error("update account", "id", id, "err", err)
		writeError(w, http.StatusInternalServerError, "could not update account")
		return
	}

	writeJSON(w, http.StatusOK, account)
}

// Delete removes an account by its ID.
func (h *Accounts) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := accountID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid account id")
		return
	}

	rows, err := h.queries.DeleteAccount(r.Context(), id)
	if err != nil {
		slog.Error("delete account", "id", id, "err", err)
		writeError(w, http.StatusInternalServerError, "could not delete account")
		return
	}
	if rows == 0 {
		writeError(w, http.StatusNotFound, "account not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// accountID reads the {id} URL parameter as an int32.
func accountID(r *http.Request) (int32, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(id), nil
}

// readAccountRequest decodes and checks the JSON body for Create and Update.
// An empty owner becomes "personal", the same default as the database.
func readAccountRequest(r *http.Request) (accountRequest, error) {
	var req accountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.New("invalid request body")
	}

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return req, errors.New("name is required")
	}

	switch req.Type {
	case repository.AccountTypeChecking,
		repository.AccountTypeSavings,
		repository.AccountTypeLoan,
		repository.AccountTypeCreditCard,
		repository.AccountTypeCash,
		repository.AccountTypeBrokerage,
		repository.AccountTypeRetirement:
	default:
		return req, errors.New("invalid account type")
	}

	if req.Owner == "" {
		req.Owner = repository.AccountOwnerPersonal
	}
	switch req.Owner {
	case repository.AccountOwnerPersonal, repository.AccountOwnerBusiness:
	default:
		return req, errors.New("invalid account owner")
	}

	return req, nil
}

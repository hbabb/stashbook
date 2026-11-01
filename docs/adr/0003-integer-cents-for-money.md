# 3. Integer cents for money

Date: 2026-11-01

## Status

Accepted

## Context

Floating-point numbers can't represent most decimal amounts exactly, so totals and balances can drift by fractions of a cent.

Stashbook is a ledger. Every amount, including each line of a split entry, is entered by hand. The app only adds and subtracts. It never divides amounts.

## Decision

Store and handle all money as whole numbers of cents. `$45.01` is stored as `4501`.

- **Database:** `bigint` columns ending in `_cents`, like `debit_cents` and `amount_cents`.
- **Go:** `int64`.
- **API:** plain JSON integers.
- **Frontend:** converts cents to dollars only for display, and dollars to cents on input.

## Consequences

- Adding and subtracting whole numbers is always exact, so totals, running balances, and reconciliation never drift.
- No rounding is needed anywhere, because the app never divides amounts.
- JavaScript represents whole numbers exactly up to about 9 quadrillion, roughly $90 trillion in cents, so the frontend can handle any amount safely.
- Every place that shows or accepts an amount must convert between dollars and cents.

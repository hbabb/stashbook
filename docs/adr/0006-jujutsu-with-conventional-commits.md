# 6. Jujutsu with Conventional Commits

Date: 2026-11-01

## Status

Accepted

## Context

The project needs version control that pushes to GitHub, and a commit history that is easy to read and scan by type of change.

## Decision

Use Jujutsu (`jj`) as the version control tool, backed by Git so it pushes to GitHub.

Write every commit message in the Conventional Commits 1.0.0 format:

```
<type>(<optional scope>): <description>
```

For example, `jj describe -m "docs(adr): add architecture decision records"`.

Types used:

- `feat`: a new feature
- `fix`: a bug fix
- `docs`: documentation only
- `refactor`: a code change that neither fixes a bug nor adds a feature
- `test`: adding or changing tests
- `build`: the build system or dependencies
- `ci`: CI workflows
- `chore`: other changes that don't touch source or tests

A breaking change adds `!` after the type, like `feat!:`, or a `BREAKING CHANGE:` footer.

Work happens on a bookmark, is pushed to GitHub, and is merged into `main` through a pull request. `main` is protected.

## Consequences

- The history shows what kind of change each commit is at a glance.
- Tools can generate a changelog and pick version numbers from the commit types: `fix` is a patch, `feat` is a minor version, and a breaking change is a major version.
- Commits can't go straight to `main`. Every change goes through a pull request.

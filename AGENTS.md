
# AGENTS.md

> **Purpose**  
> This file provides *machine‑readable* guidelines for OpenAI **Codex** (and other autonomous coding agents) working on the **Market‑Making + Hedging Simulation** written in **Go 1.22**.  
> Follow every instruction unless a more specific inline comment overrides it.  
> The original functional spec is in `Quant_Developer_Assessment_Market_Making_Hedging_Simulation.md`.

---

## 1. Project Overview

A discrete‑event **simulation** that:

* Quotes bid/ask prices on **Exchange 1**.  
* Immediately hedges filled quotes on **Exchange 2** to stay *delta‑neutral*.  
* Tracks inventory, PnL, and risk in real‑time.  
* Exposes a minimal **REST API** (`/quotes`, `/positions`, `/pnl`, etc.).

Target run‑time: local single‑process app (no external services).The code **must compile**, **pass tests**, and **race‑free** (`go test -race ./...`).

---

## 2. Directory & Package Layout

| Path | Package | Responsibility |
|------|---------|----------------|
| `cmd/sim/main.go` | **main** | CLI entrypoint. Parses flags, wires deps, starts REST server & event loop. |
| `internal/exchange/` | `exchange` | Order‑book, matching engine, price feed generator. |
| `internal/strategy/` | `strategy` | Market‑making quote logic (spread, inventory‑skew, volatility). |
| `internal/hedge/` | `hedge` | Executes hedge orders on Exchange 2; latency/slippage model. |
| `internal/risk/` | `risk` | Delta, inventory limits, kill‑switch logic. |
| `internal/api/` | `api` | HTTP handlers (mux: Chi). Returns JSON. |
| `internal/sim/` | `sim` | Discrete‑event scheduler, ticks loop, orchestrates interactions. |
| `pkg/types/` | `types` | Shared domain objects (`Order`,`Trade`,`Quote`,`Position`). |

> **Codex**: *create any missing dirs/files; avoid circular imports.*

---

## 3. Coding Conventions

* Go module path: `github.com/yourorg/marketmaker` (adjust if repo differs).  
* **Formatting**: run `gofmt`, `goimports`, and `go vet` on every change.  
* **Error handling**: return typed errors (`var ErrXYZ = errors.New("…")`) rather than panicking.  
* **Concurrency**: use goroutines + channels; protect shared state with `sync.Mutex` or design lock‑free queues.  
* **Logging**: use stdlib `log` for now; structured logging optional.  
* **Config**: use a simple `config.Config` struct loaded from flags/env (no Viper needed).

---

## 4. Testing Standards

* Unit tests live next to code (`foo_test.go`).  
* Mock time via an injected `clock.Clock` interface (see `github.com/benbjohnson/clock`).  
* Minimum **80 %** coverage (`go test -cover ./...`).  
* Include one integration test in `internal/sim` that runs a 1 000‑tick simulation and asserts:  
  * Net delta ≈ 0 after each hedge.  
  * Positive expected PnL with random‑walk prices.

---

## 5. Tasks for Codex

The implementation may be delivered incrementally as pull requests that satisfy the checklist below.  
Codex should **open a new PR per logical task** and mark items complete.

- [ ] **types**: Define core structs & enums in `pkg/types`.  
- [ ] **exchange**: Matching engine (`Submit()`, `Cancel()`, `Match()`), price feed (`Step()` random walk).  
- [ ] **strategy**: Quote generator (`GenerateQuotes()`), inventory‑aware spread.  
- [ ] **hedge**: `Hedger` that mirrors fills from Exchange 1 onto Exchange 2 with latency.  
- [ ] **risk**: Inventory & PnL monitors, kill‑switch triggers.  
- [ ] **sim**: Event loop driving time steps and orchestrating modules.  
- [ ] **api**: REST handlers using `chi` router; JSON marshaling via `encoding/json`.  
- [ ] **cmd/sim**: Wire everything together, parse CLI flags (`--ticks`, `--log-level`).  
- [ ] **tests**: Unit + integration tests meeting coverage threshold.  
- [ ] **Makefile**: targets `run`, `test`, `lint`, `cover`.  
- [ ] **Docs**: Update README with build/run instructions.

---

## 6. Pull‑Request Template

```
### Summary
Short description of change.

### Changes
- high‑level bullet
- bullet 2

### Tests
```
make test
```

### Checklist
- [ ] go vet / lint passes
- [ ] tests pass
- [ ] docs updated
```

Codex should populate the template automatically.

---

## 7. Out‑of‑Scope for v1

* Persistent storage (in‑memory only).  
* External market data APIs.  
* GUI dashboards (CLI/REST sufficient).

---

## 8. Helpful Resources

* Original spec: `Quant_Developer_Assessment_Market_Making_Hedging_Simulation.md`  
* Go Concurrency Patterns: <https://go.dev/doc/effective_go>  
* Chi router: <https://github.com/go-chi/chi>

---

_Last updated: 2025-06-09T22:28:18Z_

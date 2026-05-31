<![CDATA[<div align="center">

# 🚀 My Backend Journey

### Becoming a Top 1% Backend Engineer — One Phase at a Time

[![Progress](https://img.shields.io/badge/Phase-01%20of%2008-blue?style=for-the-badge)]()
[![Duration](https://img.shields.io/badge/Timeline-24%20Months-purple?style=for-the-badge)]()
[![Status](https://img.shields.io/badge/Status-In%20Progress-orange?style=for-the-badge)]()

*A 24-month, topic-by-topic execution plan with curated resources,
exercises, and precise timelines — following a structured roadmap from foundational CS to production-grade distributed systems.*

---

</div>

## 📖 About This Repo

This repository documents my journey to becoming a top-tier backend engineer. I'm following a comprehensive **8-phase mastery roadmap** that covers everything from CPU architecture and memory hierarchies to distributed consensus algorithms and chaos engineering.

Every concept, exercise, project, and insight along the way will be committed here — code, notes, and all.

> **Philosophy:** *Do not move to the next topic until you can explain the current one from first principles without referencing notes.*

---

## 🗺️ The Roadmap

| Phase | Title | Duration | Key Outcome |
|:-----:|-------|----------|-------------|
| **01** | 🧱 The Unshakeable Foundation | Months 1–4 | Own the computer from CPU to network |
| **02** | 🗄️ Databases: Beyond CRUD | Months 3–8 | Read `EXPLAIN ANALYZE` fluently, design schemas that scale |
| **03** | 🌐 Distributed Systems Engineering | Months 6–14 | Design and debug distributed systems under failure |
| **04** | ⚙️ Production Engineering | Months 10–18 | Run systems reliably at scale, on-call ready |
| **05** | 🏗️ System Design as Craft | Months 12–20 | Lead architecture discussions, write ADRs |
| **06** | 🔤 Language Mastery | Months 1–24 | Own one language at the runtime level |
| **07** | 🔒 Security & Reliability Patterns | Months 16–22 | Threat model, design for zero-trust |
| **08** | 🏆 The Elite Habits | Months 1–24 | Source reading, writing, teaching, estimation |

---

## 📂 Repository Structure

```
my-backend-journey/
│
├── 📄 README.md                  ← You are here
├── 📁 roadmap/                   ← The original detailed roadmap document
│
├── 📁 phase-01-foundations/      ← CS Fundamentals (CPU, OS, Networking, Concurrency, Data Structures)
├── 📁 phase-02-databases/        ← Query internals, indexing, MVCC, replication, Redis, CDC
├── 📁 phase-03-distributed/      ← Failure models, consensus, Kafka, sagas, idempotency
├── 📁 phase-04-production/       ← Observability, profiling, SLOs, chaos engineering
├── 📁 phase-05-system-design/    ← API design, rate limiting, estimation
├── 📁 phase-06-language/         ← Go runtime internals, Rust ownership & async
├── 📁 phase-07-security/         ← Auth, JWT, OAuth, injection attacks, secrets management
├── 📁 phase-08-elite-habits/     ← Source code reading, technical writing, engineering blogs
│
└── 📁 projects/                  ← Standalone projects built along the way
```

> **Note:** Folders will be created as I progress through each phase.

---

## 🧱 Phase 01 — The Unshakeable Foundation

**Duration:** Months 1–4 (16 weeks) · **Anchor Book:** *CS:APP (Bryant & O'Hallaron)*

CS fundamentals that never expire — the same in 2025 as they were in 1985.

| # | Topic | Timeline | Status |
|:-:|-------|----------|:------:|
| 1.1 | How Computers Actually Work — CPU, Memory & Execution | Weeks 1–2 | ⬜ |
| 1.2 | OS Internals — Processes, Threads & the Kernel | Weeks 2–4 | ⬜ |
| 1.3 | Networking from First Principles — TCP/IP, TLS, HTTP | Weeks 3–5 | ⬜ |
| 1.4 | Concurrency Theory — The Hardest Fundamental | Weeks 4–7 | ⬜ |
| 1.5 | Data Structures: The Right Depth | Weeks 6–8 | ⬜ |

<details>
<summary><b>📚 Key Resources</b></summary>

- **CS:APP** — Chapters 1, 5, 6, 8, 9, 10, 12
- **OSTEP (Operating Systems: Three Easy Pieces)** — [ostep.org](https://ostep.org)
- **High Performance Browser Networking** — [hpbn.co](https://hpbn.co)
- **The Art of Multiprocessor Programming** — Herlihy & Shavit
- **Latency Numbers Every Programmer Should Know** — [napkin-math](https://github.com/sirupsen/napkin-math)

</details>

---

## 🗄️ Phase 02 — Databases: Beyond CRUD

**Duration:** Months 3–8 (20 weeks) · **Anchor Books:** *DDIA (Kleppmann) + Database Internals (Petrov)*

Go 10x deeper than anyone around you. The database is where every serious backend problem lives.

| # | Topic | Timeline | Status |
|:-:|-------|----------|:------:|
| 2.1 | Query Execution Internals | Weeks 1–3 | ⬜ |
| 2.2 | Indexing Mastery | Weeks 2–5 | ⬜ |
| 2.3 | Transactions, MVCC & Isolation Levels | Weeks 4–7 | ⬜ |
| 2.4 | Distributed Databases, Redis Internals & CDC | Weeks 8–14 | ⬜ |

<details>
<summary><b>📚 Key Resources</b></summary>

- **Designing Data-Intensive Applications** — Chapters 5, 6, 7
- **Use The Index, Luke!** — [use-the-index-luke.com](https://use-the-index-luke.com)
- **PostgreSQL Documentation** — [postgresql.org/docs](https://postgresql.org/docs)
- **Redis Documentation** — [redis.io/docs](https://redis.io/docs)
- **Debezium Documentation** — [debezium.io](https://debezium.io/documentation)

</details>

---

## 🌐 Phase 03 — Distributed Systems Engineering

**Duration:** Months 6–14 · **Anchor Book:** *DDIA (Kleppmann) — second read*

Where backend engineering becomes truly hard. Most engineers never get here.

| # | Topic | Timeline | Status |
|:-:|-------|----------|:------:|
| 3.1 | Distributed Systems Fundamentals — Failure Models | Weeks 1–3 | ⬜ |
| 3.2 | Consensus Algorithms — Raft & Paxos | Weeks 3–6 | ⬜ |
| 3.3 | Kafka Deeply — Beyond Produce & Consume | Weeks 5–9 | ⬜ |
| 3.4 | Idempotency, Sagas & Distributed Correctness | Weeks 9–14 | ⬜ |

<details>
<summary><b>📚 Key Resources</b></summary>

- **DDIA** — Chapters 8, 9, 12
- **The Raft Paper** — [raft.github.io](https://raft.github.io/raft.pdf)
- **MIT 6.824 Distributed Systems Labs** — [pdos.csail.mit.edu/6.824](https://pdos.csail.mit.edu/6.824)
- **Kafka: The Definitive Guide** — [kafka.apache.org](https://kafka.apache.org/documentation)
- **Stripe Blog — Idempotency Keys** — [stripe.com/blog/idempotency](https://stripe.com/blog/idempotency)

</details>

---

## ⚙️ Phase 04 — Production Engineering

**Duration:** Months 10–18 · **Anchor Book:** *Google SRE Book (free online)*

Code that works on your machine is the beginning. Code that works at 3am is the craft.

| # | Topic | Timeline | Status |
|:-:|-------|----------|:------:|
| 4.1 | Observability — Philosophy & Practice | Weeks 1–4 | ⬜ |
| 4.2 | Performance Profiling & Debugging | Weeks 4–8 | ⬜ |
| 4.3 | SLOs, Error Budgets & Reliability Engineering | Weeks 6–10 | ⬜ |
| 4.4 | Chaos Engineering & Zero-Downtime Deployments | Weeks 10–18 | ⬜ |

<details>
<summary><b>📚 Key Resources</b></summary>

- **Google SRE Book** — [sre.google](https://sre.google/sre-book/table-of-contents)
- **OpenTelemetry Docs** — [opentelemetry.io](https://opentelemetry.io/docs)
- **Systems Performance (Brendan Gregg)** — Flame graphs, profiling methodology
- **Toxiproxy** — [github.com/Shopify/toxiproxy](https://github.com/Shopify/toxiproxy)

</details>

---

## 🏗️ Phase 05 — System Design as Craft

**Duration:** Months 12–20 · **Anchor Book:** *Software Architecture: The Hard Parts*

The synthesis of everything. Where all deep knowledge creates real leverage.

| # | Topic | Timeline | Status |
|:-:|-------|----------|:------:|
| 5.1 | API Design as a Contract | Weeks 1–4 | ⬜ |
| 5.2 | Rate Limiting, Backpressure & Flow Control | Weeks 3–6 | ⬜ |
| 5.3 | Back-of-Envelope Estimation | Weeks 5–8 | ⬜ |

<details>
<summary><b>📚 Key Resources</b></summary>

- **Stripe API Reference** — [stripe.com/docs/api](https://stripe.com/docs/api)
- **gRPC Documentation** — [grpc.io/docs](https://grpc.io/docs)
- **System Design Interview (Alex Xu)** — Volumes 1 & 2
- **Napkin Math** — [github.com/sirupsen/napkin-math](https://github.com/sirupsen/napkin-math)

</details>

---

## 🔤 Phase 06 — Language Mastery

**Duration:** Months 1–24 (parallel track) · **Primary Language:** Go · **Secondary:** Rust

Go deep in one language. Know its runtime, memory model, and failure modes completely.

| # | Topic | Timeline | Status |
|:-:|-------|----------|:------:|
| 6.1 | Go Runtime Internals (GMP Scheduler, GC, Channels) | Months 1–12 | ⬜ |
| 6.2 | Rust — Memory Safety Without GC | Months 6–18 | ⬜ |

<details>
<summary><b>📚 Key Resources</b></summary>

- **The Go Programming Language** — Donovan & Kernighan
- **100 Go Mistakes** — Teiva Harsanyi
- **Go Runtime Source** — [github.com/golang/go](https://github.com/golang/go/tree/master/src/runtime)
- **The Rust Programming Language** — [rust-book.cs.brown.edu](https://rust-book.cs.brown.edu)
- **Tokio Tutorial** — [tokio.rs](https://tokio.rs/tokio/tutorial)

</details>

---

## 🔒 Phase 07 — Security & Reliability Patterns

**Duration:** Months 16–22 · **Anchor Resource:** *OWASP Top 10 + Web Application Security (O'Reilly)*

Threat model first. Build security into the architecture, not as an afterthought.

| # | Topic | Timeline | Status |
|:-:|-------|----------|:------:|
| 7.1 | Authentication & Authorization Deeply | Weeks 1–6 | ⬜ |
| 7.2 | Injection Attacks, Secrets Management & Defense in Depth | Weeks 4–10 | ⬜ |

<details>
<summary><b>📚 Key Resources</b></summary>

- **OWASP Top 10** — [owasp.org](https://owasp.org/www-project-top-ten)
- **PortSwigger Web Security Academy** — [portswigger.net/web-security](https://portswigger.net/web-security)
- **HashiCorp Vault** — [developer.hashicorp.com/vault](https://developer.hashicorp.com/vault)
- **Google Zanzibar Paper** — [research.google](https://research.google/pubs/pub48190)

</details>

---

## 🏆 Phase 08 — The Elite Habits

**Duration:** Months 1–24 (woven into every phase)

What separates the top 1% is not just what they know — it's how they operate every day.

| # | Topic | Cadence | Status |
|:-:|-------|---------|:------:|
| 8.1 | Source Code Reading as a Discipline | 1 hr/week minimum | ⬜ |
| 8.2 | Writing to Think and to Lead | 2 posts/month minimum | ⬜ |
| 8.3 | Study Real Systems — Engineering Blogs | 2 hrs/week | ⬜ |

<details>
<summary><b>📚 Key Resources & Blogs</b></summary>

**Source Code to Read:**
- PostgreSQL — `src/backend/storage/buffer`, `src/backend/access/heap`
- Redis — `src/ae.c`, `src/t_zset.c`, `src/rdb.c`
- Go Runtime — `runtime/proc.go`, `runtime/chan.go`, `sync/mutex.go`

**Engineering Blogs to Follow:**
- [Cloudflare Blog](https://blog.cloudflare.com)
- [Netflix Tech Blog](https://netflixtechblog.com)
- [Stripe Engineering](https://stripe.com/blog/engineering)
- [Discord Engineering](https://discord.com/blog/engineering)
- [Martin Kleppmann's Blog](https://martin.kleppmann.com)

</details>

---

## 📊 Progress Tracker

```
Phase 01 ██░░░░░░░░░░░░░░░░░░  0%   ← Current
Phase 02 ░░░░░░░░░░░░░░░░░░░░  0%
Phase 03 ░░░░░░░░░░░░░░░░░░░░  0%
Phase 04 ░░░░░░░░░░░░░░░░░░░░  0%
Phase 05 ░░░░░░░░░░░░░░░░░░░░  0%
Phase 06 ░░░░░░░░░░░░░░░░░░░░  0%   (parallel)
Phase 07 ░░░░░░░░░░░░░░░░░░░░  0%
Phase 08 ░░░░░░░░░░░░░░░░░░░░  0%   (parallel)
```

---

## 📝 How I'm Using This Repo

1. **Notes & Summaries** — Written explanations of every concept in my own words
2. **Code Exercises** — Implementations of every exercise from the roadmap
3. **Projects** — Standalone projects that integrate knowledge across topics
4. **Blog Posts** — Technical writing published from this repo (Phase 08)
5. **Reading Notes** — Annotations and key takeaways from books and papers

---

## 🔑 Core Books

| Book | Author | Phases |
|------|--------|--------|
| Computer Systems: A Programmer's Perspective (CS:APP) | Bryant & O'Hallaron | 01 |
| Operating Systems: Three Easy Pieces (OSTEP) | Arpaci-Dusseau | 01 |
| Designing Data-Intensive Applications (DDIA) | Martin Kleppmann | 02, 03, 05 |
| Database Internals | Alex Petrov | 02 |
| The Google SRE Book | Beyer, Jones et al. | 04 |
| Software Architecture: The Hard Parts | Ford, Richards et al. | 05 |
| The Go Programming Language | Donovan & Kernighan | 06 |
| The Rust Programming Language | Klabnik & Nichols | 06 |

---

<div align="center">

### 💡 The Final Word

*There is no shortcut to the top 1%. The engineers who reach the top are not the ones who read the most tutorials — they are the ones who understand the fewest things at the deepest possible level, and they never stop going deeper.*

---

**Started:** June 2025 · **Target Completion:** 2027

*⭐ Star this repo if you find the roadmap useful!*

</div>
]]>

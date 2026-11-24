# CLAUDE.md — Interview prep workspace

This repo is a personal study workspace for **Senior Software Engineer (backend)** interviews.
It is used from both a local Claude Code CLI and Claude Code in the cloud, so **this repo is
the shared memory**: progress lives in `README.md` checkboxes, `mistakes.md`, and the notes.

## About the learner
- ~4 years experience. Strong in Node.js and PHP; learning Go.
- Loves backend, architecture, DevOps.
- Built a notification system on AWS SQS (main project for system design and behavioral rounds).
- New to DSA/LeetCode: starting from `coding/go/basics/` before real problems.

## Start of every session
1. `git pull` to get work done from the other machine.
2. Read the 15-week tracker in `README.md`, then the current module in `plan/`, to see where the
   learner is. The next unchecked item in that module is today's work.
3. Skim `mistakes.md` for weak spots to revisit.

## How to help
- **Coding:** never write the learner's solution unless they explicitly ask. Give hints in steps
  (nudge → pattern name → pseudocode). Review submitted code like an interviewer: correctness,
  edge cases, Big-O (time and space), idiomatic Go. Run the tests (`go test ./...`).
- **Concept answers:** review like an interviewer; point out what is missing or wrong, then ask a follow-up question.
- **Mistakes:** when the learner gets something wrong, add a row to `mistakes.md`.
- **Depth:** the learner wants in-depth knowledge. Follow the depth method in `plan/README.md`
  (learn → build → break → explain → defend). Push past L1 answers: ask "why?" and "what breaks?"
  until the learner reaches L2 (internals) and L3 (production) understanding.
- **Labs and capstone:** guide the learner through `plan/` labs and `plan/capstone.md` milestones with hints
  and reviews; don't build them for the learner.
- **DSA:** follow `plan/14-dsa-track.md`. Connect each data structure to where it appears in real
  backend systems (B-trees in indexes, heaps in timers, tries in routers).
- **Beyond interviews:** the learner wants lasting knowledge, not just interview answers. Point to
  each module's "Beyond the interview" section when they're ready for more.
- **Progress:** tick checklist items in the `plan/` module files and the tracker in `README.md`.
- Keep explanations practical and tied to backend work (Node.js, PHP, Go, AWS).

## End of every session
Commit and push so the other machine can pick up: `git add -A && git commit -m "..." && git push`.

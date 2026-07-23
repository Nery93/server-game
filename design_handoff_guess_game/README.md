# Handoff: Number Guess Game Frontend ("Neon Arcade")

## Overview
Frontend for a Mastermind/Wordle-style number guessing game backed by an existing Go game server. Player creates a match, guesses a number 0–100, gets "higher"/"lower" hints, has 10 attempts, and wins or loses.

## About the Design Files
The file in this bundle (`neon-arcade-mockup.html`) is a **static HTML/CSS design reference** — a high-fidelity mockup of look and feel, not production code to copy directly. It has no state, no fetch calls, no real interactivity. The task is to **recreate this design in a real frontend app** (React + Vite recommended, or whatever this repo already uses) that talks to the Go backend over fetch/HTTP.

## Fidelity
**High-fidelity.** Colors, typography, spacing, and copy below are final — recreate pixel-close using the target framework's own components/styling approach (inline styles, CSS modules, Tailwind, etc. — whichever fits the codebase).

## Visual Direction: "Neon Arcade"
Dark retro-CRT arcade aesthetic. Cyan + magenta neon accents on near-black background, monospace/pixel type, competitive framing (wins/streak counters, "insert coin" language).

### Design Tokens
- Background: `#0a0a14` (screens), `#14060a` (lose screen variant)
- Panel/card background: `#12122a`, `#12121e`
- Border/hairline: `#2a2a3e`
- Cyan (primary/info): `#5ff` — glow: `text-shadow: 0 0 18px rgba(85,255,255,0.6)`
- Magenta (accent/CTA): `#f5f` — glow: `0 0 24px rgba(255,85,255,0.55)`
- Yellow (warning/"higher" hint): `#ff5`
- Green (win): `#5f5` — glow `0 0 26px rgba(85,255,85,0.7)`
- Red (lose/empty life): `#f55`, dim empty heart `#2e1a1e` / `#2a2a3e`
- Muted text: `#8888a8`, dim label: `#44445c`
- Fonts: `'Press Start 2P'` (headlines, pixel display, use sparingly — large sizes only) and `'JetBrains Mono'` (body/data/log). Both via Google Fonts.
- Border radius: 4–6px (chunky pixel feel, not rounded)
- Corner style: sharp-ish, small radius only

## Screens

### 1. Start Screen
- Purpose: create a new match.
- Layout: centered column, full-bleed dark background, top-right HUD row (wins/streak).
- Components:
  - HUD top-right: `WINS 12` / `STREAK 4`, JetBrains Mono 12px, cyan/magenta, letter-spacing 1px.
  - Title: "HI-LO / ARCADE", Press Start 2P ~34px, cyan with glow, centered, 2 lines.
  - Subtitle: "Guess the secret number 0–100. / 10 lives. No mercy." JetBrains Mono 14px, muted.
  - CTA button: "▶ INSERT COIN" — Press Start 2P 15px, dark text on magenta `#f5f` bg, glow shadow, padding 18px/34px, radius 4px. Click → calls create-match endpoint, navigates to Gameplay.
  - Footer hint: "press ENTER to start", JetBrains Mono 11px, dim.

### 2. Gameplay Screen
- Purpose: submit guesses, see hint + history.
- Layout: vertical stack, padding 28px/36px.
  - Header row: `ROUND 6/10` (Press Start 2P 13px cyan) left, 10-heart life row right (♥, filled magenta with glow for remaining lives, dim `#2a2a3e` for spent).
  - Hint block, centered: big directional call — "▲ GO HIGHER" or "▼ GO LOWER", Press Start 2P 26px, yellow (higher) or magenta (lower), glow; sub-line "the secret number is HIGHER/LOWER than {last guess}" JetBrains Mono 13px muted.
  - Input row, centered: number input styled as pixel display (JetBrains Mono/Press Start 2P 28px, cyan text, dark bg `#12122a`, cyan 2px border, inset glow) + "FIRE" button (Press Start 2P 13px, dark text on cyan bg, glow).
  - History/log section, flex-1, label "— LOG —" (dim, letter-spacing 2px), then a list newest-first: each row is a flex space-between pill (bg `#12121e`, radius 4px, padding 10px/16px) showing `#{attempt} {guess}` (mono, muted number + white guess) and hint text `TOO LOW ▲` (yellow) or `TOO HIGH ▼` (magenta).
- Behavior: submit guess → POST to guess endpoint → append to history, decrement lives, update hint. Enter key submits. Disable input at 0 lives or on win. Attempts counter and life-hearts must derive from the same server-reported state (don't track separately client-side).

### 3. Win Screen
- Purpose: celebrate a correct guess.
- Layout: centered column.
- Components:
  - "YOU WIN" — Press Start 2P 40px, green with strong glow.
  - Result line: "The number was **63** — cracked in **6 guesses**" JetBrains Mono 15px, number/count bolded green.
  - Stats row: `WINS 13` (cyan) / `STREAK 5` (magenta) / `BEST 3` (yellow), JetBrains Mono 12px muted labels.
  - CTA: "↻ PLAY AGAIN" — Press Start 2P 13px, dark text on green bg, glow, radius 4px. Click → new match, back to Start or straight into new Gameplay.

### 4. Lose Screen
- Purpose: show loss after 10 failed attempts.
- Layout: centered column, background shifts to `#14060a` (slightly red-tinted black), border `#3e2a2e`.
- Components:
  - 10-heart row, all dim/spent (`#2e1a1e`).
  - "GAME OVER" — Press Start 2P 40px, red with glow.
  - Result line: "The number was **63**. Streak lost." JetBrains Mono 15px.
  - CTA: "↻ CONTINUE?" — Press Start 2P 13px, red text, transparent bg, red 2px border, soft red glow. Click → new match.

## Interactions & Behavior
- Start → New Game creates a match server-side and returns a match id (kept internal, not shown in UI) plus initial state (10 lives/attempts remaining).
- Each guess: client sends the guessed integer; server returns hint (`higher`/`lower`/`correct`) and remaining attempts.
- Win: triggered when server reports `correct`.
- Lose: triggered when attempts remaining hits 0 without a correct guess.
- History list is append-only per match, newest guess on top.
- Hearts: total 10, filled count = attempts remaining, decrementing after each wrong guess (small "pop"/fade transition on the heart that just went dim is a nice touch, not required).
- No page reload between screens — single-page flow (Start → Gameplay → Win/Lose → back to Start).
- Keyboard: Enter submits guess in Gameplay; Enter also triggers "Insert Coin" on Start (per footer hint).

## State Management (suggested)
- `matchId` (string, internal only)
- `attemptsRemaining` (number, starts 10)
- `guesses`: array of `{ value: number, hint: 'higher'|'lower' }`, newest first
- `status`: `'idle' | 'playing' | 'won' | 'lost'`
- `secretNumber` (revealed only on win/lose, from server response)
- `wins`, `streak`, `bestAttempts`: persisted client-side (e.g. localStorage) across matches, since this is presentation-layer stat tracking not required to live on the server unless the API already provides it — confirm against actual API shape.

## Assets
No images/icons beyond Unicode glyphs (▲ ▼ ↻ ▶ ♥) and two Google Fonts: `Press Start 2P`, `JetBrains Mono` (load via `fonts.googleapis.com`, weights default/700 for JetBrains Mono).

## API Integration Notes
This README does not know the exact Go server's endpoint paths/payload shapes — inspect the backend's routes/handlers (or its OpenAPI/docs if any) and wire:
- create match → Start screen CTA
- submit guess → Gameplay submit
- read current match state → to resume/hydrate lives, history, status
Map server hint values (whatever strings/enums the API returns for "secret is bigger/smaller") to the "higher"/"lower" display strings above.

## Files
- `neon-arcade-mockup.html` — static design reference (Start, Gameplay, Win, Lose), all four screens visible at once for reference. Open directly in a browser.

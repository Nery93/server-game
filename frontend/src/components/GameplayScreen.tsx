import { useState, type FormEvent } from 'react';
import { Hearts } from './Hearts';
import type { GuessEntry } from '../types';
import type { HintDirection } from '../api/types';
import styles from './GameplayScreen.module.css';

interface GameplayScreenProps {
  round: number;
  maxAttempts: number;
  livesRemaining: number;
  lastHint: HintDirection | null;
  lastGuessValue: number | null;
  guesses: GuessEntry[];
  loading: boolean;
  error: string | null;
  onSubmit: (value: number) => void;
}

export function GameplayScreen({
  round,
  maxAttempts,
  livesRemaining,
  lastHint,
  lastGuessValue,
  guesses,
  loading,
  error,
  onSubmit,
}: GameplayScreenProps) {
  const [draft, setDraft] = useState('');

  function handleSubmit(e: FormEvent) {
    e.preventDefault();
    const value = Number(draft);
    if (draft === '' || Number.isNaN(value) || value < 0 || value > 100) return;
    onSubmit(value);
    setDraft('');
  }

  return (
    <div className={styles.screen}>
      <div className={styles.header}>
        <div className={styles.round}>
          ROUND {round}/{maxAttempts}
        </div>
        <Hearts total={maxAttempts} filled={livesRemaining} />
      </div>

      <div className={styles.hintBlock}>
        {lastHint && (
          <>
            <p className={styles.hintLine} data-direction={lastHint}>
              {lastHint === 'higher' ? '▲ GO HIGHER' : '▼ GO LOWER'}
            </p>
            <p className={styles.hintSub}>
              the secret number is {lastHint === 'higher' ? 'HIGHER' : 'LOWER'} than {lastGuessValue}
            </p>
          </>
        )}
      </div>

      <form className={styles.inputRow} onSubmit={handleSubmit}>
        <input
          className={styles.input}
          type="number"
          inputMode="numeric"
          min={0}
          max={100}
          value={draft}
          onChange={(e) => setDraft(e.target.value)}
          disabled={loading}
          autoFocus
        />
        <button className={styles.fireButton} type="submit" disabled={loading}>
          FIRE
        </button>
      </form>

      {error && <p className={styles.error}>{error}</p>}

      <div className={styles.logSection}>
        <div className={styles.logLabel}>— LOG —</div>
        {guesses.map((g) => (
          <div className={styles.logRow} key={g.attempt}>
            <span>
              <span className={styles.logAttempt}>#{g.attempt}&nbsp;&nbsp;</span>
              <span className={styles.logGuess}>{g.value}</span>
            </span>
            <span className={styles.logHint} data-direction={g.hint}>
              {g.hint === 'higher' ? 'TOO LOW ▲' : 'TOO HIGH ▼'}
            </span>
          </div>
        ))}
      </div>
    </div>
  );
}

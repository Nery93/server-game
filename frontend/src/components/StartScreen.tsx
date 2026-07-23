import { useEffect } from 'react';
import styles from './StartScreen.module.css';

interface StartScreenProps {
  wins: number;
  streak: number;
  loading: boolean;
  error: string | null;
  onStart: () => void;
}

export function StartScreen({ wins, streak, loading, error, onStart }: StartScreenProps) {
  useEffect(() => {
    function handleKeyDown(e: KeyboardEvent) {
      if (e.key === 'Enter') onStart();
    }
    window.addEventListener('keydown', handleKeyDown);
    return () => window.removeEventListener('keydown', handleKeyDown);
  }, [onStart]);

  return (
    <div className={styles.screen}>
      <div className={styles.hud}>
        <span>WINS {wins}</span>
        <span>STREAK {streak}</span>
      </div>

      <h1 className={styles.title}>
        HI-LO
        <br />
        ARCADE
      </h1>

      <p className={styles.subtitle}>
        Guess the secret number 0–100.
        <br />
        10 lives. No mercy.
      </p>

      <button className={styles.cta} onClick={onStart} disabled={loading}>
        ▶ INSERT COIN
      </button>

      {error && <p className={styles.error}>{error}</p>}

      <p className={styles.footerHint}>press ENTER to start</p>
    </div>
  );
}

import { Hearts } from './Hearts';
import styles from './LoseScreen.module.css';

interface LoseScreenProps {
  secretNumber: number;
  maxAttempts: number;
  onContinue: () => void;
}

export function LoseScreen({ secretNumber, maxAttempts, onContinue }: LoseScreenProps) {
  return (
    <div className={styles.screen}>
      <Hearts total={maxAttempts} filled={0} emptyColor="var(--heart-empty-lose)" />

      <h1 className={styles.title}>GAME OVER</h1>

      <p className={styles.result}>
        The number was <b>{secretNumber}</b>. Streak lost.
      </p>

      <button className={styles.cta} onClick={onContinue}>
        ↻ CONTINUE?
      </button>
    </div>
  );
}

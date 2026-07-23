import styles from './WinScreen.module.css';

interface WinScreenProps {
  secretNumber: number;
  attemptsUsed: number;
  wins: number;
  streak: number;
  bestAttempts: number | null;
  onPlayAgain: () => void;
}

export function WinScreen({
  secretNumber,
  attemptsUsed,
  wins,
  streak,
  bestAttempts,
  onPlayAgain,
}: WinScreenProps) {
  return (
    <div className={styles.screen}>
      <h1 className={styles.title}>YOU WIN</h1>

      <p className={styles.result}>
        The number was <b>{secretNumber}</b> — cracked in <b>{attemptsUsed} guesses</b>
      </p>

      <div className={styles.stats}>
        <span>
          WINS <b>{wins}</b>
        </span>
        <span>
          STREAK <b>{streak}</b>
        </span>
        <span>
          BEST <b>{bestAttempts ?? attemptsUsed}</b>
        </span>
      </div>

      <button className={styles.cta} onClick={onPlayAgain}>
        ↻ PLAY AGAIN
      </button>
    </div>
  );
}

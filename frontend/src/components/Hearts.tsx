import styles from './Hearts.module.css';

interface HeartsProps {
  total: number;
  filled: number;
  /** Overrides the empty-heart color (the Lose screen renders every heart dim). */
  emptyColor?: string;
}

export function Hearts({ total, filled, emptyColor }: HeartsProps) {
  return (
    <div className={styles.row}>
      {Array.from({ length: total }, (_, i) => {
        const isFilled = i < filled;
        return (
          <span
            key={i}
            className={isFilled ? styles.filled : styles.empty}
            style={!isFilled && emptyColor ? { color: emptyColor } : undefined}
          >
            ♥
          </span>
        );
      })}
    </div>
  );
}

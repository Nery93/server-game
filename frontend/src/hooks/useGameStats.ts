import { useCallback, useState } from 'react';

interface GameStats {
  wins: number;
  streak: number;
  bestAttempts: number | null;
}

const STORAGE_KEY = 'hilo-arcade-stats';
const EMPTY_STATS: GameStats = { wins: 0, streak: 0, bestAttempts: null };

function loadStats(): GameStats {
  try {
    const raw = localStorage.getItem(STORAGE_KEY);
    return raw ? (JSON.parse(raw) as GameStats) : EMPTY_STATS;
  } catch {
    return EMPTY_STATS;
  }
}

// Wins/streak/best are presentation-layer stats the API doesn't track, so
// they live in localStorage and survive across matches on this browser only.
export function useGameStats() {
  const [stats, setStats] = useState<GameStats>(loadStats);

  const recordWin = useCallback((attemptsUsed: number) => {
    setStats((prev) => {
      const next: GameStats = {
        wins: prev.wins + 1,
        streak: prev.streak + 1,
        bestAttempts:
          prev.bestAttempts === null ? attemptsUsed : Math.min(prev.bestAttempts, attemptsUsed),
      };
      localStorage.setItem(STORAGE_KEY, JSON.stringify(next));
      return next;
    });
  }, []);

  const recordLoss = useCallback(() => {
    setStats((prev) => {
      const next: GameStats = { ...prev, streak: 0 };
      localStorage.setItem(STORAGE_KEY, JSON.stringify(next));
      return next;
    });
  }, []);

  return { stats, recordWin, recordLoss };
}

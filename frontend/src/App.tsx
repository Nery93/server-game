import { useCallback, useState } from 'react';
import { createGame, submitGuess, parseHintDirection } from './api/client';
import { MAX_ATTEMPTS } from './constants';
import { useGameStats } from './hooks/useGameStats';
import { StartScreen } from './components/StartScreen';
import { GameplayScreen } from './components/GameplayScreen';
import { WinScreen } from './components/WinScreen';
import { LoseScreen } from './components/LoseScreen';
import type { GuessEntry, Status } from './types';

function App() {
  const [status, setStatus] = useState<Status>('idle');
  const [gameId, setGameId] = useState<string | null>(null);
  const [guesses, setGuesses] = useState<GuessEntry[]>([]);
  const [attemptsMade, setAttemptsMade] = useState(0);
  const [secretNumber, setSecretNumber] = useState<number | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const { stats, recordWin, recordLoss } = useGameStats();

  const handleStart = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      const res = await createGame();
      setGameId(res.id);
      setGuesses([]);
      setAttemptsMade(0);
      setSecretNumber(null);
      setStatus('playing');
    } catch (e) {
      setError(e instanceof Error ? e.message : 'Failed to start a new game');
    } finally {
      setLoading(false);
    }
  }, []);

  const handleGuess = useCallback(
    async (value: number) => {
      if (!gameId) return;
      setLoading(true);
      setError(null);
      try {
        const res = await submitGuess(gameId, value);
        setAttemptsMade(res.attempts);

        if (res.correct) {
          setSecretNumber(res.secret_number ?? value);
          setStatus('won');
          recordWin(res.attempts);
          return;
        }

        if (res.attempts >= MAX_ATTEMPTS) {
          setSecretNumber(res.secret_number ?? null);
          setStatus('lost');
          recordLoss();
          return;
        }

        const direction = parseHintDirection(res.hint ?? '');
        setGuesses((prev) => [{ attempt: res.attempts, value, hint: direction }, ...prev]);
      } catch (e) {
        setError(e instanceof Error ? e.message : 'Failed to submit guess');
      } finally {
        setLoading(false);
      }
    },
    [gameId, recordWin, recordLoss],
  );

  if (status === 'idle') {
    return (
      <StartScreen
        wins={stats.wins}
        streak={stats.streak}
        loading={loading}
        error={error}
        onStart={handleStart}
      />
    );
  }

  if (status === 'won') {
    return (
      <WinScreen
        secretNumber={secretNumber ?? 0}
        attemptsUsed={attemptsMade}
        wins={stats.wins}
        streak={stats.streak}
        bestAttempts={stats.bestAttempts}
        onPlayAgain={handleStart}
      />
    );
  }

  if (status === 'lost') {
    return (
      <LoseScreen secretNumber={secretNumber ?? 0} maxAttempts={MAX_ATTEMPTS} onContinue={handleStart} />
    );
  }

  const lastGuess = guesses[0] ?? null;

  return (
    <GameplayScreen
      round={attemptsMade + 1}
      maxAttempts={MAX_ATTEMPTS}
      livesRemaining={MAX_ATTEMPTS - attemptsMade}
      lastHint={lastGuess?.hint ?? null}
      lastGuessValue={lastGuess?.value ?? null}
      guesses={guesses}
      loading={loading}
      error={error}
      onSubmit={handleGuess}
    />
  );
}

export default App;

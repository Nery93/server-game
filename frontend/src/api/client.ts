import type { CreateGameResponse, GuessResponse, HintDirection } from './types';

async function parseErrorMessage(res: Response): Promise<string> {
  const text = await res.text();
  return text || `Request failed with status ${res.status}`;
}

export async function createGame(): Promise<CreateGameResponse> {
  const res = await fetch('/game', { method: 'POST' });
  if (!res.ok) throw new Error(await parseErrorMessage(res));
  return res.json();
}

export async function submitGuess(gameId: string, guess: number): Promise<GuessResponse> {
  const res = await fetch(`/game/${gameId}/guess`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ number: guess }),
  });
  if (!res.ok) throw new Error(await parseErrorMessage(res));
  return res.json();
}

// The backend returns full sentences ("No, the secret number is higher."),
// not an enum, so we detect direction from the wording instead of an exact match.
export function parseHintDirection(hint: string): HintDirection {
  return hint.toLowerCase().includes('lower') ? 'lower' : 'higher';
}

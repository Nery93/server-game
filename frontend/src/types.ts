import type { HintDirection } from './api/types';

export type Status = 'idle' | 'playing' | 'won' | 'lost';

export interface GuessEntry {
  attempt: number;
  value: number;
  hint: HintDirection;
}

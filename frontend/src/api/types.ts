export interface CreateGameResponse {
  id: string;
}

// Mirrors internal/handler/handler.go's GuessResponse. hint, attempts_left and
// secret_number are only present in some states (see handler.go for when).
export interface GuessResponse {
  correct: boolean;
  attempts: number;
  hint?: string;
  attempts_left?: string;
  secret_number?: number;
}

export type HintDirection = 'higher' | 'lower';

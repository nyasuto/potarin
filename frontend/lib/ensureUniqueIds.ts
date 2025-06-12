import { randomUUID } from 'crypto';
import { Suggestion } from 'potarin-shared/types';

export function ensureUniqueIds(suggestions: Suggestion[]): Suggestion[] {
  const used = new Set<string>();
  return suggestions.map((s) => {
    let id = s.id;
    if (!id || used.has(id)) {
      id = randomUUID();
    }
    used.add(id);
    return { ...s, id };
  });
}

# Contributor Guide
## Rule
- Always generate branch names in English using only lowercase letters, numbers, and hyphens. Translate task descriptions to English. No Japanese or non-ASCII characters.

## Dev Environment Tips
- Use golang for backend
- Use TypeScript for frontend
- When adding new npm package, run `bun add <npm package>` instead

## Testing Instructions
Do this in backend folder
- go build ./...
- go test ./...

Do this in frontend folder
- bun install
- bun run build

## PR instructions
- Title format: [<Potarin>] <Title>
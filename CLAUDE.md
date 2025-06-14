# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Potarin is an AI-powered cycling course recommendation application. Users input their preferences and the AI generates personalized cycling routes with map visualization. The application features a Go backend with Fiber framework and a Next.js frontend.

## Development Commands

### Backend (Go + Fiber)
```bash
cd backend
go install github.com/cosmtrek/air@latest  # Install hot reload tool (one-time)
air  # Start backend with hot reload on :8080
```

### Frontend (Next.js + Bun)
```bash
cd frontend
bun install
bun run dev  # Start frontend dev server
```

### Testing
Backend:
```bash
cd backend
go build ./...
go test ./...
```

Frontend:
```bash
cd frontend
bun install
bun run build
bun run lint
```

## Architecture

### Monorepo Structure
- `/backend` - Go API server using Fiber framework
- `/frontend` - Next.js 15 app with TypeScript and React 19
- `/shared` - Type definitions and JSON schemas shared between backend and frontend

### Key Components

**Backend (`/backend/main.go`)**
- Fiber web server with CORS enabled
- Two main endpoints:
  - `GET /api/v1/suggestions` - Returns AI-generated cycling course suggestions
  - `POST /api/v1/details` - Returns detailed route information for a specific suggestion
- OpenAI integration with structured JSON responses using response_format with json_schema
- Type-safe Go structs that mirror the shared types

**Frontend Architecture**
- Next.js 15 with App Router
- React-Leaflet for map visualization (requires `{ ssr: false }` for MapClient components)
- Tailwind CSS for styling
- TypeScript with strict type checking using shared types from `/shared/types.ts`

**Shared Types (`/shared/`)**
- Maintains type consistency between Go backend and TypeScript frontend
- Contains JSON schemas for OpenAI structured responses
- Key types: `Suggestion`, `Detail`, `Route`, `Position`

### OpenAI Integration
- Uses structured JSON responses with `response_format: json_schema`
- Schemas defined in `/shared/schemas/` ensure consistent AI responses
- No natural language responses - all AI output is structured JSON
- Model selection via `OPENAI_MODEL` environment variable (defaults to `gpt-4o-mini`)

### Environment Variables
- `OPENAI_API_KEY` - Required for AI functionality
- `OPENAI_MODEL` - AI model selection (optional, defaults to `gpt-4o-mini`)
- `NEXT_PUBLIC_API_URL` - Frontend API server URL configuration

## Development Guidelines

### AI Prompt Engineering
- Always use JSON schema for OpenAI responses
- Pass user profile context to AI for personalized recommendations
- Default prompts are in Japanese for cycling course suggestions

### Map Integration
- Use dynamic imports with `{ ssr: false }` for React-Leaflet components
- MapClient components must be client-side only due to Leaflet's DOM dependencies

### Type Safety
- Import types from `/shared/types.ts` in frontend
- Use corresponding Go structs from `/shared/types.go` in backend
- Maintain type consistency across the full stack

### Branch Naming
- Use English lowercase with hyphens only
- Format: `feature/description-of-change`
- No Japanese or non-ASCII characters

### PR Format
- Title: `[Potarin] <descriptive title>`
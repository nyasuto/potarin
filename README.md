# ポタりん V2 — AIエージェント駆動リアーキテクチャ (Goバックエンド版)

## プロジェクト概要

ポタりん は、ユーザーの希望に応じてAIが散歩・サイクリングコースを提案し、地図上にルートを可視化するアプリケーションです。
このバージョンでは、AIエージェント（Codex等）を積極活用しながら、バックエンドをGo言語で完全再設計します。

## 利用技術スタック

採用技術

### フロントエンド

Next.js 15 (App Router, TypeScript, React 19)
bun

### バックエンド

Golang + Fiber

### AI連携

OpenAI GPT-4o / GPT-4o-mini API

### 地図描画

React-Leaflet + Leaflet

### スタイリング

Tailwind CSS

### ビルド・開発

Bun (フロント用) / Go modules

### デプロイ

Vercel (フロント), Fly.io or Railway (Goバックエンド)

### 状態管理

React Context or Zustand

### API通信

REST + JSON

### CI/CD

GitHub Actions

### AI開発支援

ChatGPT (Codex/Custom GPT)

## ディレクトリ構成

- /frontend (Next.js, TypeScript, React)
- /backend (Go + Fiber)
- /shared (APIスキーマ定義, JSON Schema)

`/frontend`, `/backend`, `/shared` ディレクトリを作成し、最小構成のアプリケーションを配置しています。

## クイックスタート

1. `backend` ディレクトリで Go サーバーを起動

```bash
cd backend
go run .
```

2. 別ターミナルで `frontend` を起動

```bash
cd frontend
bun install
bun run dev
```

環境変数 `NEXT_PUBLIC_API_URL` を設定すると、フロントエンドが参照する API サーバー URL を変更できます。
バックエンドで OpenAI を利用する場合は、プロジェクトルートの `.env.local` に `OPENAI_API_KEY` を設定してください。

## 現在の機能

- `/api/v1/suggestions` で AI によるコース提案を取得
- `/api/v1/details` でコースの詳細 (summary とルート情報) を取得
- Next.js フロントエンドで提案一覧と詳細ページを表示

## 今後の予定

- 地図描画機能 (React-Leaflet)
- ユーザー現在地連携や履歴保存などの拡張

### 設計・実装の注意点

#### AIプロンプト設計上の注意

- JSONスキーマを必ずAIに渡す (response_format: json_schema)
- 一切自然文レスポンスさせない (Goのパース強靭性が最大限活きる)

#### Goバックエンド設計上の注意

- Fiber を使った軽量APIサーバー
- 全APIは JSON POST/GETのみ（OpenAPIスキーマで定義可）
- OpenAIエラーハンドリングはpanicせず安全に戻す
- 型安全のために明示的な構造体定義を使う

#### フロント実装注意

- MapClientは必ずdynamic(..., { ssr: false })でSSR除外
- API通信時の型はsharedディレクトリからimportして厳密に管理
- Codexを使ってAPIスキーマ→型定義生成も可能

## 完成イメージ

- AIが実時間で適切なサイクリングコースを提案
- 地図上にルートを動的に描画
- ユーザーは距離・気分・天候を指定可能
- バックエンドはGoが堅牢にAPI連携を処理
- 開発はCodexエージェントが随時補助

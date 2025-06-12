package schemas

import _ "embed"

// SuggestionSchema is the JSON schema for shared.Suggestion.
//
//go:embed suggestion.json
var SuggestionSchema string

// SuggestionsSchema is the JSON schema for []shared.Suggestion wrapped in an object.
//
//go:embed suggestions.json
var SuggestionsSchema string

// DetailSchema is the JSON schema for shared.Detail.
//
//go:embed detail.json
var DetailSchema string

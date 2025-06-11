package schemas

import _ "embed"

// SuggestionSchema is the JSON schema for shared.Suggestion.
//
//go:embed suggestion.json
var SuggestionSchema string

// DetailSchema is the JSON schema for shared.Detail.
//
//go:embed detail.json
var DetailSchema string

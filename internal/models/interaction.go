package models

// InteractionRequest represents the payload sent to the Gemini interactions API.
type InteractionRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

// InteractionContent represents content parts in an interaction step.
type InteractionContent struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

// InteractionStep represents a step (e.g., thought, model_output) in the interaction.
type InteractionStep struct {
	Type    string               `json:"type"`
	Content []InteractionContent `json:"content,omitempty"`
}

// InteractionResponse represents the response payload from the Gemini interactions API.
type InteractionResponse struct {
	ID     string            `json:"id"`
	Status string            `json:"status"`
	Steps  []InteractionStep `json:"steps"`
	Model  string            `json:"model"`
}

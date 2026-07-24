package ui

import (
	"testing"

	"github.com/asheshgoplani/agent-deck/internal/session"
)

func TestRemoteSessionWarpLabelIncludesAgentAndModel(t *testing.T) {
	got := remoteSessionWarpLabel(session.RemoteSessionInfo{
		Title: "codextest",
		Tool:  "codex",
		Model: "gpt-5.6-luna",
	})
	if want := "Agent Deck · codextest · codex · gpt-5.6-luna"; got != want {
		t.Fatalf("remoteSessionWarpLabel() = %q, want %q", got, want)
	}
}

func TestRemoteSessionWarpLabelFallsBackToSessionName(t *testing.T) {
	got := remoteSessionWarpLabel(session.RemoteSessionInfo{Title: "shell"})
	if want := "Agent Deck · shell"; got != want {
		t.Fatalf("remoteSessionWarpLabel() = %q, want %q", got, want)
	}
}

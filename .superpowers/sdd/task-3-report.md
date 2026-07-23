# Task 3 implementation report

Implemented the AgentBox Agent picker in the private Agent Deck fork.

Changes:

- Added a fixed picker for `claude-code`, `codex`, and `pi-fireworks`.
- Added keyboard navigation, selection, dismissal, and Home-level routing.
- Prevented arbitrary text input for the AgentBox Agent field.
- Reused the selected agent to filter the compatible model catalog.
- Added regression tests for rendering, keyboard behavior, selection, and invalid text.
- Preserved the Agent Deck/AgentBox private integration boundary.

Verification:

- `go test ./internal/ui -run 'Test(NewDialog_AgentboxAgentPicker|Home_AgentboxAgentPicker)' -count=1` — passed.
- `go test ./internal/ui -run 'Test(Issue1353_AgentboxRemoteDialog|NewDialog_AgentboxAgentPicker|Home_AgentboxAgentPicker)' -count=1` — passed.
- `git diff --check` — passed.

The broader `go test ./internal/ui -count=1` run was stopped after it entered unrelated long-running Agent Deck tests; the focused picker and existing AgentBox dialog coverage passed.

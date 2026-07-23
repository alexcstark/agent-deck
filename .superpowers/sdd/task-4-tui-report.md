# Task 4 TUI implementation report

Added the requested compatible model catalogs to the private Agent Deck TUI:

- Claude: `claude-fable-5` and `claude-fable-5[1m]`.
- Codex: `gpt-5.6`, `gpt-5.6-sol`, `gpt-5.6-terra`, and `gpt-5.6-luna`.
- Pi / Fireworks: `accounts/fireworks/models/glm-5p2`.

The AgentBox picker tests now verify the current Claude, Codex, and Pi model memberships and use GPT-5.6 as the current Codex catalog head.

Verification:

- `go test ./internal/ui -run 'Test(NewDialog_Agentbox|Home_Agentbox|Issue1353_AgentboxRemoteDialog)' -count=1` — passed.
- `git diff --check` — passed.

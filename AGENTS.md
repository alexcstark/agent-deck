# Agent instructions

This repository is a public fork. Treat every committed file, branch, PR, and
build artifact as publicly visible.

- Never commit secrets, credentials, access tokens, private keys, or auth
  material.
- Never commit private endpoints, internal hostnames, infrastructure details,
  local machine paths, deployment commands, or customer data.
- Keep integrations generic and configurable. Internal AgentBox deployment and
  operational details belong in the private AgentBox repository.
- Before pushing or opening a PR, inspect the diff for accidental internal
  information and run the repository's available secret checks.

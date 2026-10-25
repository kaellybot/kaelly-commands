# Kaelly-commands

[![Golangci-lint](https://github.com/kaellybot/kaelly-commands/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/kaellybot/kaelly-commands/actions/workflows/golangci-lint.yml)
[![Test](https://github.com/kaellybot/kaelly-commands/actions/workflows/test.yml/badge.svg)](https://github.com/kaellybot/kaelly-commands/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/kaellybot/kaelly-commands/branch/main/graph/badge.svg)](https://codecov.io/gh/kaellybot/kaelly-commands) 

Discord commands to interact written in Go 

## Usage

Each game registers its own command set on its own Discord application:

| Function | Commands |
| --- | --- |
| `GetDofusCommands()` | about, align, almanax, config (get, almanax, rss, server, twitter), help, item, job, map, set |
| `GetTouchCommands()` | about, align, almanax, config (get, almanax, rss, server, twitter), help, job (level up to 100) |
| `GetRetroCommands()` | about, align, config (get, server), help, job (level up to 100) |

The command, option and custom ID names are shared by every game.

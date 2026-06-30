# yze-go-stdlog

A [`yze`](https://github.com/gomatic/yze) analyzer (category `data`) enforcing the gomatic Go logging standard: the standard library `log` package is forbidden in favor of structured logging with `log/slog`.

- **Rule:** `yze/stdlog`
- **Binary:** `cmd/yze-go-stdlog` runs it standalone.

Built on the [`go-yze`](https://github.com/gomatic/go-yze) framework.

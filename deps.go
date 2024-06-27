//go:build deps
// +build deps

package bot

import (
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
)

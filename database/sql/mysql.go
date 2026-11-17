//go:build mysql

// Package sql provides database helpers for storing Who's On First records in SQL.
package sql

import (
	_ "github.com/go-sql-driver/mysql"
)

package model

import (
	"database/sql"
	"io"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(db *sql.DB, pathToRoot string) error {
	// db, err := sql.Open("postgres", "postgres://localhost:5432/database?sslmode=enable")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	files, err := filepath.Glob(pathToRoot + "pkg/*/schema/*.sql")
	if err != nil {
		return err
	}

	fileInfo, err := os.Lstat(pathToRoot)
	if err != nil {
		return err
	}

	fileMode := fileInfo.Mode()
	unixPerms := fileMode & os.ModePerm
	if err := os.MkdirAll(pathToRoot+"tmp/migrations", unixPerms); err != nil {
		return err
	}

	for _, file := range files {
		err := CopyMigration(file, pathToRoot)
		if err != nil {
			return err
		}
	}

	// 実行rootからの相対pathになっている。全体実行はproject rootだが、テストの場合はtest対象のディレクトリ
	m, err := migrate.NewWithDatabaseInstance("file://"+pathToRoot+"tmp/migrations", "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		return err
	}

	return nil
}

func CopyMigration(file string, pathToRoot string) error {
	from, err := os.Open(file)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.Create(pathToRoot + "tmp/migrations/" + filepath.Base(file))
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		return err
	}

	return nil
}

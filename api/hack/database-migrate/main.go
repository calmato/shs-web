// usage: go run ./main.go -db-host=127.0.0.1 -db-port=3316 -db-password=1234567
package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"database/sql"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/pkg/set"
	_ "github.com/go-sql-driver/mysql"
)

const (
	migrateDB   = "migrations"
	schemaTable = "schemas"
)

var (
	databases = []string{
		"classrooms",
		"lessons",
		"users",
	}
	skipDDLs = []string{
		"2021120901-setup.sql",
	}
)

func main() {
	start := time.Now()
	fmt.Println("Start..")
	if err := run(); err != nil {
		panic(err)
	}
	fmt.Printf("Done: %s\n", time.Since(start))
}

type app struct {
	db *sql.DB
}

type schema struct {
	version  string
	database string
	filename string
	path     string
}

//nolint:funlen
func run() error {
	app := app{}
	host := flag.String("db-host", "mysql", "target mysql host")
	port := flag.String("db-port", "3306", "target mysql port")
	username := flag.String("db-username", "root", "target mysql username")
	password := flag.String("db-password", "12345678", "target mysql password")
	srcDir := flag.String("src", "./../../../infra/mysql/schema", "ddl source directory")
	flag.Parse()

	set := set.New(len(skipDDLs))
	set.AddStrings(skipDDLs...)

	/**
	 * -------------------------
	 * データベースの管理
	 * -------------------------
	 */
	err := app.setup(*host, *port, *username, *password)
	if err != nil {
		return err
	}

	isExists := app.checkMigrateDB()
	fmt.Println("recreate database:", !isExists)
	if !isExists {
		tx, err := app.begin()
		if err != nil {
			return err
		}
		defer app.close(tx)

		// マイクロサービス用のDBを再作成
		if err := app.dropDBIfExists(tx, databases...); err != nil {
			return app.rollback(tx, err)
		}
		if err := app.createDBIfNotExists(tx, databases...); err != nil {
			return app.rollback(tx, err)
		}

		// DDL管理用のDBを作成
		if err := app.createDBIfNotExists(tx, migrateDB); err != nil {
			return app.rollback(tx, err)
		}

		// DDL管理用のTableを作成
		if err := app.createSchemaTable(tx); err != nil {
			return app.rollback(tx, err)
		}

		if err := tx.Commit(); err != nil {
			return app.rollback(tx, err)
		}
	}

	/**
	 * -------------------------
	 * テーブルの管理
	 * -------------------------
	 */
	// DDLの取得
	schemas, err := app.newSchemas(*srcDir)
	if err != nil {
		return err
	}

	tx, err := app.begin()
	if err != nil {
		return err
	}
	defer app.close(tx)

	// DDLの適用
	for i := range schemas {
		if set.Contains(schemas[i].filename) {
			fmt.Printf("%s is skip ddl -> skip\n", schemas[i].filename)
			continue
		}

		isApplied, err := app.getSchema(tx, schemas[i])
		if err != nil {
			fmt.Println("debug: err=", err)
			return app.rollback(tx, err)
		}
		if isApplied {
			fmt.Printf("%s is already applied -> skip\n", schemas[i].filename)
			continue
		}

		fmt.Printf("%s is applying...", schemas[i].filename)
		if err := app.applySchema(tx, schemas[i]); err != nil {
			fmt.Println("debug: err=", err)
			return app.rollback(tx, err)
		}
		fmt.Printf(" -> succeeded!!\n")
	}

	if err := tx.Commit(); err != nil {
		fmt.Println("debug: err=", err)
		return app.rollback(tx, err)
	}
	return nil
}

/**
 * instance
 */
func (a *app) setup(host, port, username, password string) error {
	conf := fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port)
	conn, err := sql.Open("mysql", conf)
	if err != nil {
		return err
	}
	a.db = conn
	return nil
}

func (a *app) begin() (*sql.Tx, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

//nolint:unparam
func (a *app) close(tx *sql.Tx) func() {
	return func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			return
		}
	}
}

func (a *app) rollback(tx *sql.Tx, err error) error {
	return fmt.Errorf("%w: %s", err, tx.Rollback().Error())
}

func (a *app) checkMigrateDB() bool {
	stmt := fmt.Sprintf("USE `%s`", migrateDB)
	rs, _ := a.db.Exec(stmt)
	return rs != nil
}

func (a *app) createDBIfNotExists(tx *sql.Tx, dbs ...string) error {
	const format = "CREATE SCHEMA IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4"
	for i := range dbs {
		stmt := fmt.Sprintf(format, dbs[i])
		if _, err := tx.Exec(stmt); err != nil {
			return err
		}
	}
	return nil
}

func (a *app) dropDBIfExists(tx *sql.Tx, dbs ...string) error {
	const format = "DROP DATABASE IF EXISTS `%s`"
	for i := range dbs {
		stmt := fmt.Sprintf(format, dbs[i])
		if _, err := tx.Exec(stmt); err != nil {
			return err
		}
	}
	return nil
}

func (a *app) createSchemaTable(tx *sql.Tx) error {
	//nolint:lll
	const format = "CREATE TABLE `%s`.`%s` (`version` VARCHAR(10) NOT NULL, `database` VARCHAR(256) NOT NULL, `filename` VARCHAR(256) NOT NULL, `created_at` INT NOT NULL, `updated_at` INT NOT NULL, PRIMARY KEY(`version`)) ENGINE = InnoDB"
	stmt := fmt.Sprintf(format, migrateDB, schemaTable)
	_, err := tx.Exec(stmt)
	return err
}

func (a *app) getSchema(tx *sql.Tx, schema *schema) (bool, error) {
	const format = "SELECT `version` FROM `%s`.`%s` WHERE `version` = '%s' LIMIT 1"
	stmt := fmt.Sprintf(format, migrateDB, schemaTable, schema.version)
	rs, err := tx.Query(stmt)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}
	defer rs.Close()
	return rs.Next(), nil
}

func (a *app) applySchema(tx *sql.Tx, schema *schema) error {
	bytes, err := ioutil.ReadFile(schema.path)
	if err != nil {
		return err
	}

	// 1つのファイルに複数定義が書いてある場合はエラーになるため、分割して適用
	sqls := strings.Split(string(bytes), ";")
	for _, sql := range sqls {
		if sql == "" || sql == "\n" {
			continue // split時、配列の最後に空文字が入るため
		}
		if _, err := tx.Exec(sql); err != nil {
			return err
		}
	}

	now := jst.Now().Unix()
	//nolint:lll
	const format = "INSERT INTO `%s`.`%s` (`version`, `database`, `filename`, `created_at`, `updated_at`) VALUES ('%s', '%s', '%s', '%d', '%d')"
	stmt := fmt.Sprintf(format, migrateDB, schemaTable, schema.version, schema.database, schema.filename, now, now)
	if _, err := tx.Exec(stmt); err != nil {
		return err
	}
	return nil
}

func (a *app) newSchemas(srcDir string) ([]*schema, error) {
	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return nil, err
	}

	schemas := make([]*schema, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		// filename: {version}-{databases}-{message}.sql
		// e.g.) 2021120902-users-create-table-teachers.sql
		filename := file.Name()
		strs := strings.Split(filename, "-")

		schema := &schema{
			version:  strs[0],
			database: strs[1],
			filename: filename,
			path:     filepath.Join(srcDir, filename),
		}
		schemas = append(schemas, schema)
	}
	return schemas, nil
}

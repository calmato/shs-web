package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	centity "github.com/calmato/shs-web/api/internal/classroom/entity"
	uentity "github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/pkg/uuid"
	"github.com/calmato/shs-web/api/proto/user"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
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
	user      *database.Client
	classroom *database.Client
}

func run() error {
	app := &app{}
	host := flag.String("db-host", "mysql", "target mysql host")
	port := flag.String("db-port", "3306", "target mysql port")
	username := flag.String("db-username", "root", "target mysql username")
	password := flag.String("db-password", "12345678", "target mysql password")
	isDelete := flag.Bool("is-delete", true, "if true, delete the existing record")
	uid := flag.String("uid", "cngxK2YbQkiUfRUcp8zSet", "admin's id to be created")
	flag.Parse()

	if *uid == "" {
		return fmt.Errorf("uid must be required")
	}
	userDB, err := app.setup(*host, *port, "users", *username, *password)
	if err != nil {
		return err
	}
	app.user = userDB
	classroomDB, err := app.setup(*host, *port, "classrooms", *username, *password)
	if err != nil {
		return err
	}
	app.classroom = classroomDB

	ctx := context.Background()
	eg, _ := errgroup.WithContext(ctx)

	/**
	 * -------------------------
	 * users
	 * -------------------------
	 */
	eg.Go(func() error {
		if *isDelete {
			tables := []string{"teachers"}
			if err := app.delete(app.user.DB, tables...); err != nil {
				return err
			}
		}

		tx, err := app.user.Begin()
		if err != nil {
			return err
		}
		defer app.user.Close(tx)

		// 管理者の登録
		err = app.insertAdmin(tx, *uid)
		if err != nil {
			tx.Rollback()
			return err
		}

		// 講師の登録
		err = app.insertTeachers(tx, 10)
		if err != nil {
			tx.Rollback()
			return err
		}

		return tx.Commit().Error
	})

	/**
	 * -------------------------
	 * classrooms
	 * -------------------------
	 */
	eg.Go(func() error {
		if *isDelete {
			tables := []string{"subjects"}
			if err := app.delete(app.classroom.DB, tables...); err != nil {
				return err
			}
		}

		tx, err := app.classroom.Begin()
		if err != nil {
			return err
		}
		defer app.classroom.Close(tx)

		// 授業科目の登録
		err = app.insertSubjects(tx)
		if err != nil {
			tx.Rollback()
			return err
		}

		return tx.Commit().Error
	})

	return eg.Wait()
}

func (a *app) setup(host, port, db, username, password string) (*database.Client, error) {
	params := &database.Params{
		Socket:   "tcp",
		Host:     host,
		Port:     port,
		Database: db,
		Username: username,
		Password: password,
	}
	return database.NewClient(params)
}

func (a *app) delete(tx *gorm.DB, tables ...string) error {
	for _, table := range tables {
		sql := fmt.Sprintf("TRUNCATE TABLE %s", table)
		if err := tx.Exec(sql).Error; err != nil {
			return err
		}
	}
	return nil
}

func (a *app) insertAdmin(tx *gorm.DB, uid string) error {
	now := jst.Now()
	teacher := &uentity.Teacher{
		ID:            uid,
		LastName:      "開発用",
		FirstName:     "管理者",
		LastNameKana:  "かいはつよう",
		FirstNameKana: "かんりしゃ",
		Mail:          "admin@calmato.jp",
		Role:          int32(user.Role_ROLE_ADMINISTRATOR),
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	return tx.Create(&teacher).Error
}

func (a *app) insertTeachers(tx *gorm.DB, size int) error {
	now := jst.Now()
	teachers := make(uentity.Teachers, size)
	for i := 0; i < size; i++ {
		uid := uuid.Base58Encode(uuid.New())
		teacher := &uentity.Teacher{
			ID:            uid,
			LastName:      "開発用",
			FirstName:     fmt.Sprintf("講師%03d", i),
			LastNameKana:  "かいはつよう",
			FirstNameKana: fmt.Sprintf("こうし%03d", i),
			Mail:          fmt.Sprintf("teacher%03d@calmato.jp", i),
			Role:          int32(user.Role_ROLE_TEACHER),
			CreatedAt:     now,
			UpdatedAt:     now,
		}
		teachers[i] = teacher
	}
	return tx.Create(&teachers).Error
}

func (a *app) insertSubjects(tx *gorm.DB) error {
	names := []string{"国語", "数学", "社会", "理科", "英語"}
	colors := []string{"#F8BBD0", "#BBDEFB", "#FFE0B2", "#E8F5E9", "#FEE6C9"}
	now := jst.Now()
	subjects := make(centity.Subjects, len(names))
	for i := range names {
		subject := &centity.Subject{
			Name:      names[i],
			Color:     colors[i],
			CreatedAt: now,
			UpdatedAt: now,
		}
		subjects[i] = subject
	}
	return tx.Create(&subjects).Error
}

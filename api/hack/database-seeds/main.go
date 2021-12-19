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
	"github.com/calmato/shs-web/api/proto/classroom"
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
	subjectsMap := map[classroom.SchoolType]centity.Subjects{
		classroom.SchoolType_SCHOOL_TYPE_ELEMENTARY_SCHOOL: {
			{Name: "国語", Color: "#F8BBD0"},
			{Name: "算数", Color: "#BBDEFB"},
		},
		classroom.SchoolType_SCHOOL_TYPE_JUNIOR_HIGH_SCHOOL: {
			{Name: "国語", Color: "#F8BBD0"},
			{Name: "数学", Color: "#BBDEFB"},
			{Name: "社会", Color: "#FFE0B2"},
			{Name: "理科", Color: "#E8F5E9"},
			{Name: "英語", Color: "#FEE6C9"},
		},
		classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL: {
			{Name: "国語", Color: "#F8BBD0"},
			{Name: "数学Ⅰ", Color: "#BBDEFB"},
			{Name: "数学Ⅱ", Color: "#BBDEFB"},
			{Name: "数学A", Color: "#BBDEFB"},
			{Name: "数学B", Color: "#BBDEFB"},
			{Name: "英語", Color: "#FEE6C9"},
		},
	}

	now := jst.Now()
	subjects := make(centity.Subjects, 0)
	for schoolType, ss := range subjectsMap {
		for _, subject := range ss {
			subject.SchoolType = int32(schoolType)
			subject.CreatedAt = now
			subject.UpdatedAt = now
		}
		subjects = append(subjects, ss...)
	}
	return tx.Create(&subjects).Error
}

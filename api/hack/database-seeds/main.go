// usage: go run ./main.go -db-host=127.0.0.1 -db-port=3316 -db-password=1234567
package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"time"

	centity "github.com/calmato/shs-web/api/internal/classroom/entity"
	uentity "github.com/calmato/shs-web/api/internal/user/entity"
	"github.com/calmato/shs-web/api/pkg/database"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/pkg/uuid"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	adminID     = "cngxK2YbQkiUfRUcp8zSet"
	teacherSize = 10
	studentSize = 10
	roomSize    = 4
)

var (
	truncateUserTables = []string{
		"teachers",
		"students",
	}
	truncateClassroomTables = []string{
		"rooms",
		"teacher_subjects",
		"subjects",
	}
	weekdayLessons = centity.Lessons{
		{StartTime: "1700", EndTime: "1830"},
		{StartTime: "1830", EndTime: "2000"},
		{StartTime: "2000", EndTime: "2130"},
	}
	holidayLessons = centity.Lessons{
		{StartTime: "1530", EndTime: "1700"},
		{StartTime: "1700", EndTime: "1830"},
		{StartTime: "1830", EndTime: "2000"},
		{StartTime: "2000", EndTime: "2130"},
	}
	subjectsMap = map[centity.SchoolType]centity.Subjects{
		centity.SchoolTypeElementarySchool: {
			{Name: "国語", Color: "#F8BBD0"},
			{Name: "算数", Color: "#BBDEFB"},
			{Name: "社会", Color: "#FFE0B2"},
			{Name: "理科", Color: "#E8F5E9"},
			{Name: "英語", Color: "#FEE6C9"},
		},
		centity.SchoolTypeJuniorHighSchool: {
			{Name: "国語", Color: "#F8BBD0"},
			{Name: "数学", Color: "#BBDEFB"},
			{Name: "社会", Color: "#FFE0B2"},
			{Name: "地理", Color: "#FFE0B2"},
			{Name: "歴史", Color: "#FFE0B2"},
			{Name: "公民", Color: "#FFE0B2"},
			{Name: "現代社会", Color: "#FFE0B2"},
			{Name: "理科", Color: "#E8F5E9"},
			{Name: "英語", Color: "#FEE6C9"},
		},
		centity.SchoolTypeHighSchool: {
			{Name: "国語", Color: "#F8BBD0"},
			{Name: "現代文", Color: "#F8BBD0"},
			{Name: "古典", Color: "#F8BBD0"},
			{Name: "数学Ⅰ", Color: "#BBDEFB"},
			{Name: "数学Ⅱ", Color: "#BBDEFB"},
			{Name: "数学Ⅲ", Color: "#BBDEFB"},
			{Name: "数学A", Color: "#BBDEFB"},
			{Name: "数学B", Color: "#BBDEFB"},
			{Name: "数学活用", Color: "#BBDEFB"},
			{Name: "地理", Color: "#FFE0B2"},
			{Name: "歴史", Color: "#FFE0B2"},
			{Name: "世界史", Color: "#FFE0B2"},
			{Name: "日本史", Color: "#FFE0B2"},
			{Name: "公民", Color: "#FFE0B2"},
			{Name: "現代社会", Color: "#FFE0B2"},
			{Name: "倫理", Color: "#FFE0B2"},
			{Name: "政治・経済", Color: "#FFE0B2"},
			{Name: "物理", Color: "#E8F5E9"},
			{Name: "化学", Color: "#E8F5E9"},
			{Name: "生物", Color: "#E8F5E9"},
			{Name: "地学", Color: "#E8F5E9"},
			{Name: "英語", Color: "#FEE6C9"},
		},
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
	user      *database.Client
	classroom *database.Client
}

//nolint:funlen
func run() error {
	app := &app{}
	host := flag.String("db-host", "mysql", "target mysql host")
	port := flag.String("db-port", "3306", "target mysql port")
	username := flag.String("db-username", "root", "target mysql username")
	password := flag.String("db-password", "12345678", "target mysql password")
	isDelete := flag.Bool("is-delete", false, "if true, delete the existing record")
	flag.Parse()

	/**
	 * -------------------------
	 * setup
	 * -------------------------
	 */
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
			err := app.delete(app.user.DB, truncateUserTables...)
			if err != nil {
				return err
			}
		}

		tx, err := app.user.Begin()
		if err != nil {
			return err
		}
		defer app.user.Close(tx)

		// 管理者の登録
		err = app.upsertAdmin(tx, adminID)
		if err != nil {
			tx.Rollback()
			return err
		}

		// 開発用生徒の登録
		err = app.upsertDeveloper(tx, adminID)
		if err != nil {
			tx.Rollback()
			return err
		}

		// 講師の登録
		err = app.upsertTeachers(tx, teacherSize)
		if err != nil {
			tx.Rollback()
			return err
		}

		// 生徒の登録
		err = app.upsertStudents(tx, studentSize)
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
			err := app.delete(app.classroom.DB, truncateClassroomTables...)
			if err != nil {
				return err
			}
		}

		tx, err := app.classroom.Begin()
		if err != nil {
			return err
		}
		defer app.classroom.Close(tx)

		// ブースの登録
		err = app.upsertRooms(tx, roomSize)
		if err != nil {
			tx.Rollback()
			return err
		}

		// 授業科目の登録
		err = app.upsertSubjects(tx)
		if err != nil {
			tx.Rollback()
			return err
		}

		// 授業スケジュールの登録
		err = app.upsertSchedules(tx)
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
	const (
		before = "SET foreign_key_checks = 0"
		after  = "SET foreign_key_checks = 1"
	)
	defer tx.Exec(after)

	if err := tx.Exec(before).Error; err != nil {
		return err
	}

	for _, table := range tables {
		sql := fmt.Sprintf("TRUNCATE TABLE %s", table)
		if err := tx.Exec(sql).Error; err != nil {
			return err
		}
	}
	return nil
}

func (a *app) upsertAdmin(tx *gorm.DB, uid string) error {
	now := jst.Now()
	teacher := &uentity.Teacher{
		ID:            uid,
		LastName:      "開発用",
		FirstName:     "管理者",
		LastNameKana:  "かいはつよう",
		FirstNameKana: "かんりしゃ",
		Mail:          "admin@calmato.jp",
		Role:          uentity.RoleAdministrator,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&teacher).Error
}

func (a *app) upsertDeveloper(tx *gorm.DB, uid string) error {
	now := jst.Now()
	student := &uentity.Student{
		ID:            uid,
		LastName:      "開発用",
		FirstName:     "生徒",
		LastNameKana:  "かいはつよう",
		FirstNameKana: "せいと",
		Mail:          "admin@calmato.jp",
		BirthYear:     1996,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&student).Error
}

func (a *app) upsertTeachers(tx *gorm.DB, size int) error {
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
			Role:          uentity.RoleTeacher,
			CreatedAt:     now,
			UpdatedAt:     now,
		}
		teachers[i] = teacher
	}
	return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&teachers).Error
}

func (a *app) upsertStudents(tx *gorm.DB, size int) error {
	now := jst.Now()
	years := newRandomBirthYear(now)
	students := make(uentity.Students, size)
	for i := 0; i < size; i++ {
		uid := uuid.Base58Encode(uuid.New())
		student := &uentity.Student{
			ID:            uid,
			LastName:      "開発用",
			FirstName:     fmt.Sprintf("生徒%03d", i),
			LastNameKana:  "かいはつよう",
			FirstNameKana: fmt.Sprintf("せいと%03d", i),
			Mail:          fmt.Sprintf("student%03d@calmato.jp", i),
			BirthYear:     getRandomBirthYear(years),
			CreatedAt:     now,
			UpdatedAt:     now,
		}
		students[i] = student
	}
	return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&students).Error
}

// 年度計算
// - 01月01日 ~ 04月01日: 去年の年を返す
// - 04月02日 ~ 12月31日: 現在の年を返す
func getFiscalYear(now time.Time) int64 {
	month, day := now.Month(), now.Day()
	if 0 < month && month < 3 {
		return int64(now.Year() - 1)
	} else if month == 4 && day == 1 {
		return int64(now.Year() - 1)
	}
	return int64(now.Year())
}

func newRandomBirthYear(now time.Time) []int64 {
	const (
		maxSize         = 12 // 小学校1~6年, 中学校1~3年, 高校1~3年
		firstYear int64 = 7  // 小学校1年 = 今年から7年前生まれ
		finalYear int64 = 18 // 高校3年 = 今年から18年前生まれ
	)
	current := getFiscalYear(now)
	years := make([]int64, 0, maxSize)
	for i := firstYear; i <= finalYear; i++ {
		year := current - i
		years = append(years, year)
	}
	return years
}

func getRandomBirthYear(years []int64) int64 {
	num := rand.Intn(len(years))
	return years[num]
}

func (a *app) upsertRooms(tx *gorm.DB, size int) error {
	now := jst.Now()
	rooms := make(centity.Rooms, size)
	for i := 0; i < size; i++ {
		room := &centity.Room{
			ID:        int32(i + 1),
			CreatedAt: now,
			UpdatedAt: now,
		}
		rooms[i] = room
	}
	return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&rooms).Error
}

func (a *app) upsertSubjects(tx *gorm.DB) error {
	now := jst.Now()
	subjects := make(centity.Subjects, 0)
	for schoolType, ss := range subjectsMap {
		for _, subject := range ss {
			subject.SchoolType = schoolType
			subject.CreatedAt = now
			subject.UpdatedAt = now
		}
		subjects = append(subjects, ss...)
	}
	return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&subjects).Error
}

func (a *app) upsertSchedules(tx *gorm.DB) error {
	const weekdays = 7
	now := jst.Now()
	schedules := make(centity.Schedules, weekdays)
	for i := 0; i < weekdays; i++ {
		weekday := time.Weekday(i)
		schedule := &centity.Schedule{
			Weekday:   weekday,
			IsClosed:  false,
			Lessons:   newLessons(weekday),
			CreatedAt: now,
			UpdatedAt: now,
		}
		if err := schedule.FillJSON(); err != nil {
			return err
		}
		schedules[i] = schedule
	}
	return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&schedules).Error
}

func newLessons(weekday time.Weekday) centity.Lessons {
	if weekday == time.Sunday || weekday == time.Saturday {
		return holidayLessons
	}
	return weekdayLessons
}

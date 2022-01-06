package entity

import (
	"testing"
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestStudent(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 04, 01, 12, 0, 0, 0)
	tests := []struct {
		name          string
		now           time.Time
		lastName      string
		firstName     string
		lastNameKana  string
		firstNameKana string
		mail          string
		password      string
		schoolType    SchoolType
		grade         int64
		expect        *Student
	}{
		{
			name:          "success to school unknown",
			now:           now,
			lastName:      "浜田",
			firstName:     "直志",
			lastNameKana:  "はまだ",
			firstNameKana: "ただし",
			mail:          "student-test002@calmato.jp",
			password:      "12345678",
			schoolType:    SchoolTypeUnknown,
			grade:         0,
			expect: &Student{
				ID:            "ThisIsTestIdForStudent",
				LastName:      "浜田",
				FirstName:     "直志",
				LastNameKana:  "はまだ",
				FirstNameKana: "ただし",
				Mail:          "student-test002@calmato.jp",
				BirthYear:     2021,
				Password:      "12345678",
				SchoolType:    SchoolTypeUnknown,
				Grade:         0,
			},
		},
		{
			name:          "success to elementary school",
			now:           now,
			lastName:      "浜田",
			firstName:     "直志",
			lastNameKana:  "はまだ",
			firstNameKana: "ただし",
			mail:          "student-test002@calmato.jp",
			password:      "12345678",
			schoolType:    SchoolTypeElementarySchool,
			grade:         1,
			expect: &Student{
				ID:            "ThisIsTestIdForStudent",
				LastName:      "浜田",
				FirstName:     "直志",
				LastNameKana:  "はまだ",
				FirstNameKana: "ただし",
				Mail:          "student-test002@calmato.jp",
				BirthYear:     2014,
				Password:      "12345678",
				SchoolType:    SchoolTypeElementarySchool,
				Grade:         1,
			},
		},
		{
			name:          "success to junior high school",
			now:           now,
			lastName:      "浜田",
			firstName:     "直志",
			lastNameKana:  "はまだ",
			firstNameKana: "ただし",
			mail:          "student-test002@calmato.jp",
			password:      "12345678",
			schoolType:    SchoolTypeJuniorHighSchool,
			grade:         1,
			expect: &Student{
				ID:            "ThisIsTestIdForStudent",
				LastName:      "浜田",
				FirstName:     "直志",
				LastNameKana:  "はまだ",
				FirstNameKana: "ただし",
				Mail:          "student-test002@calmato.jp",
				BirthYear:     2008,
				Password:      "12345678",
				SchoolType:    SchoolTypeJuniorHighSchool,
				Grade:         1,
			},
		},
		{
			name:          "success to high school",
			now:           now,
			lastName:      "浜田",
			firstName:     "直志",
			lastNameKana:  "はまだ",
			firstNameKana: "ただし",
			mail:          "student-test002@calmato.jp",
			password:      "12345678",
			schoolType:    SchoolTypeHighSchool,
			grade:         1,
			expect: &Student{
				ID:            "ThisIsTestIdForStudent",
				LastName:      "浜田",
				FirstName:     "直志",
				LastNameKana:  "はまだ",
				FirstNameKana: "ただし",
				Mail:          "student-test002@calmato.jp",
				BirthYear:     2005,
				Password:      "12345678",
				SchoolType:    SchoolTypeHighSchool,
				Grade:         1,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewStudent(
				tt.lastName, tt.firstName, tt.lastNameKana, tt.firstNameKana,
				tt.mail, tt.password, tt.schoolType, tt.grade, tt.now)
			actual.ID = "ThisIsTestIdForStudent"
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestStudent_Fill(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 04, 01, 12, 0, 0, 0)
	tests := []struct {
		name         string
		now          time.Time
		student      *Student
		expectSchool SchoolType
		expectGrade  int64
	}{
		{
			name: "success to school unknown",
			now:  now,
			student: &Student{
				ID:            "lKEnoE6LetnMs5Byk2a7Zx",
				LastName:      "浜田",
				FirstName:     "直志",
				LastNameKana:  "はまだ",
				FirstNameKana: "ただし",
				Mail:          "student-test002@calmato.jp",
				BirthYear:     2015,
				CreatedAt:     now,
				UpdatedAt:     now,
			},
			expectSchool: SchoolTypeUnknown,
			expectGrade:  0,
		},
		{
			name: "success to elementary school",
			now:  now,
			student: &Student{
				ID:            "lKEnoE6LetnMs5Byk2a7Zx",
				LastName:      "浜田",
				FirstName:     "直志",
				LastNameKana:  "はまだ",
				FirstNameKana: "ただし",
				Mail:          "student-test002@calmato.jp",
				BirthYear:     2014,
				CreatedAt:     now,
				UpdatedAt:     now,
			},
			expectSchool: SchoolTypeElementarySchool,
			expectGrade:  1,
		},
		{
			name: "success to junior high school",
			now:  now,
			student: &Student{
				ID:            "lKEnoE6LetnMs5Byk2a7Zx",
				LastName:      "浜田",
				FirstName:     "直志",
				LastNameKana:  "はまだ",
				FirstNameKana: "ただし",
				Mail:          "student-test002@calmato.jp",
				BirthYear:     2008,
				CreatedAt:     now,
				UpdatedAt:     now,
			},
			expectSchool: SchoolTypeJuniorHighSchool,
			expectGrade:  1,
		},
		{
			name: "success to high school",
			now:  now,
			student: &Student{
				ID:            "lKEnoE6LetnMs5Byk2a7Zx",
				LastName:      "浜田",
				FirstName:     "直志",
				LastNameKana:  "はまだ",
				FirstNameKana: "ただし",
				Mail:          "student-test002@calmato.jp",
				BirthYear:     2005,
				CreatedAt:     now,
				UpdatedAt:     now,
			},
			expectSchool: SchoolTypeHighSchool,
			expectGrade:  1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.student.Fill(tt.now)
			assert.Equal(t, tt.expectSchool, tt.student.SchoolType)
			assert.Equal(t, tt.expectGrade, tt.student.Grade)
		})
	}
}

func TestStudent_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 04, 01, 12, 0, 0, 0)
	tests := []struct {
		name    string
		now     time.Time
		student *Student
		expect  *user.Student
	}{
		{
			name: "success",
			now:  now,
			student: &Student{
				ID:            "lKEnoE6LetnMs5Byk2a7Zx",
				LastName:      "浜田",
				FirstName:     "直志",
				LastNameKana:  "はまだ",
				FirstNameKana: "ただし",
				Mail:          "student-test001@calmato.jp",
				BirthYear:     2005,
				SchoolType:    SchoolTypeHighSchool,
				Grade:         1,
				CreatedAt:     now,
				UpdatedAt:     now,
			},
			expect: &user.Student{
				Id:            "lKEnoE6LetnMs5Byk2a7Zx",
				LastName:      "浜田",
				FirstName:     "直志",
				LastNameKana:  "はまだ",
				FirstNameKana: "ただし",
				Mail:          "student-test001@calmato.jp",
				BirthYear:     2005,
				SchoolType:    user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
				Grade:         1,
				CreatedAt:     now.Unix(),
				UpdatedAt:     now.Unix(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.student.Proto())
		})
	}
}

func TestStudents_Fill(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 04, 01, 12, 0, 0, 0)
	tests := []struct {
		name     string
		now      time.Time
		students Students
		expect   Students
	}{
		{
			name: "success",
			now:  now,
			students: Students{
				{
					ID:            "lKEnoE6LetnMs5Byk2a7Zx",
					LastName:      "浜田",
					FirstName:     "直志",
					LastNameKana:  "はまだ",
					FirstNameKana: "ただし",
					Mail:          "student-test001@calmato.jp",
					BirthYear:     2005,
					SchoolType:    SchoolTypeJuniorHighSchool,
					Grade:         2,
					CreatedAt:     now,
					UpdatedAt:     now,
				},
			},
			expect: Students{
				{
					ID:            "lKEnoE6LetnMs5Byk2a7Zx",
					LastName:      "浜田",
					FirstName:     "直志",
					LastNameKana:  "はまだ",
					FirstNameKana: "ただし",
					Mail:          "student-test001@calmato.jp",
					BirthYear:     2005,
					SchoolType:    SchoolTypeHighSchool,
					Grade:         1,
					CreatedAt:     now,
					UpdatedAt:     now,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.students.Fill(tt.now)
			assert.Equal(t, tt.expect, tt.students)
		})
	}
}
func TestStudents_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 04, 01, 12, 0, 0, 0)
	tests := []struct {
		name     string
		now      time.Time
		students Students
		expect   []*user.Student
	}{
		{
			name: "success",
			students: Students{
				{
					ID:            "lKEnoE6LetnMs5Byk2a7Zx",
					LastName:      "浜田",
					FirstName:     "直志",
					LastNameKana:  "はまだ",
					FirstNameKana: "ただし",
					Mail:          "student-test001@calmato.jp",
					BirthYear:     2005,
					SchoolType:    SchoolTypeHighSchool,
					Grade:         1,
					CreatedAt:     now,
					UpdatedAt:     now,
				},
			},
			expect: []*user.Student{
				{
					Id:            "lKEnoE6LetnMs5Byk2a7Zx",
					LastName:      "浜田",
					FirstName:     "直志",
					LastNameKana:  "はまだ",
					FirstNameKana: "ただし",
					Mail:          "student-test001@calmato.jp",
					BirthYear:     2005,
					SchoolType:    user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
					Grade:         1,
					CreatedAt:     now.Unix(),
					UpdatedAt:     now.Unix(),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.students.Proto())
		})
	}
}

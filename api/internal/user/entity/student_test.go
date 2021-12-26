package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestStudent_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name    string
		student *Student
		expect  *user.Student
	}{
		{
			name: "success",
			student: &Student{
				ID:            "lKEnoE6LetnMs5Byk2a7Zx",
				LastName:      "浜田",
				FirstName:     "直志",
				LastNameKana:  "はまだ",
				FirstNameKana: "ただし",
				Mail:          "student-test001@calmato.jp",
				BirthYear:     2021,
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
				BirthYear:     2021,
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

func TestStudents_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
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
					BirthYear:     2021,
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
					BirthYear:     2021,
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

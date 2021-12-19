package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestStudent(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name    string
		student *user.Student
		expect  *Student
	}{
		{
			name: "success",
			student: &user.Student{
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
			expect: &Student{
				Student: &user.Student{
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
			assert.Equal(t, tt.expect, NewStudent(tt.student))
		})
	}
}

func TestStudents(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
		students []*user.Student
		expect   Students
	}{
		{
			name: "success",
			students: []*user.Student{
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
			expect: Students{
				{
					Student: &user.Student{
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
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudents(tt.students))
		})
	}
}

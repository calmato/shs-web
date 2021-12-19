package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
)

func TestSubject(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name    string
		subject *classroom.Subject
		expect  *Subject
	}{
		{
			name: "success",
			subject: &classroom.Subject{
				Id:        1,
				Name:      "国語",
				Color:     "#f8bbd0",
				CreatedAt: now.Unix(),
				UpdatedAt: now.Unix(),
			},
			expect: &Subject{
				Subject: &classroom.Subject{
					Id:        1,
					Name:      "国語",
					Color:     "#f8bbd0",
					CreatedAt: now.Unix(),
					UpdatedAt: now.Unix(),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewSubject(tt.subject))
		})
	}
}

func TestSubjects(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
		subjects []*classroom.Subject
		expect   Subjects
	}{
		{
			name: "success",
			subjects: []*classroom.Subject{
				{
					Id:        1,
					Name:      "国語",
					Color:     "#f8bbd0",
					CreatedAt: now.Unix(),
					UpdatedAt: now.Unix(),
				},
			},
			expect: Subjects{
				{
					Subject: &classroom.Subject{
						Id:        1,
						Name:      "国語",
						Color:     "#f8bbd0",
						CreatedAt: now.Unix(),
						UpdatedAt: now.Unix(),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewSubjects(tt.subjects))
		})
	}
}

package mailer

import (
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
)

type TemplateDataBuilder struct {
	data map[string]string
}

func NewTemplateDataBuilder() *TemplateDataBuilder {
	return &TemplateDataBuilder{
		data: map[string]string{},
	}
}

func (b *TemplateDataBuilder) Build() map[string]string {
	return b.data
}

func (b *TemplateDataBuilder) Data(data map[string]string) *TemplateDataBuilder {
	if data != nil {
		b.data = data
	}
	return b
}

func (b *TemplateDataBuilder) YearMonth(yearMonth time.Time) *TemplateDataBuilder {
	b.data["年月"] = jst.Format(yearMonth, "2006年01月")
	return b
}

func (b *TemplateDataBuilder) Name(name string) *TemplateDataBuilder {
	b.data["氏名"] = name
	return b
}

func (b *TemplateDataBuilder) WebURL(url string) *TemplateDataBuilder {
	b.data["サイトURL"] = url
	return b
}

func (b *TemplateDataBuilder) TeacherID(teacherID string) *TemplateDataBuilder {
	b.data["actor"] = "teacher"
	b.data["teacherId"] = teacherID
	return b
}

func (b *TemplateDataBuilder) StudentID(studentID string) *TemplateDataBuilder {
	b.data["actor"] = "student"
	b.data["studentId"] = studentID
	return b
}

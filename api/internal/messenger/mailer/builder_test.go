package mailer

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/stretchr/testify/assert"
)

func TestTemplateBuilder(t *testing.T) {
	builder := NewTemplateDataBuilder().
		Data(map[string]string{"key": "value"}).
		YearMonth(jst.Date(2022, 1, 2, 18, 30, 0, 0)).
		Name("中村 広大").
		TeacherID("teacherid").
		StudentID("studentid")
	data := builder.Build()
	assert.Equal(t, "value", data["key"])
	assert.Equal(t, "2022年01月", data["yearMonth"])
	assert.Equal(t, "中村 広大", data["name"])
	assert.Equal(t, "生徒", data["actor"])
	assert.Equal(t, "teacherid", data["teacherId"])
	assert.Equal(t, "studentid", data["studentId"])
}

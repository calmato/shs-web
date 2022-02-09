package mailer

import (
	"net/url"
	"strings"
)

/**
 * --------------------------
 * 講師関連URL生成用
 * --------------------------
 */
type TeacherURLMaker struct {
	url *url.URL
}

func NewTeacherURLMaker(url *url.URL) *TeacherURLMaker {
	return &TeacherURLMaker{url: url}
}

// SignIn - サインイン
func (m *TeacherURLMaker) SignIn() string {
	// e.g.) /signin
	paths := []string{"signin"}
	webURL := *m.url // copy
	webURL.Path = strings.Join(paths, "/")
	return webURL.String()
}

/**
 * --------------------------
 * 生徒関連URL生成用
 * --------------------------
 */
type StudentURLMaker struct {
	url *url.URL
}

func NewStudentURLMaker(url *url.URL) *StudentURLMaker {
	return &StudentURLMaker{url: url}
}

// SignIn - サインイン
func (m *StudentURLMaker) SignIn() string {
	// e.g.) /signin
	paths := []string{"signin"}
	webURL := *m.url // copy
	webURL.Path = strings.Join(paths, "/")
	return webURL.String()
}

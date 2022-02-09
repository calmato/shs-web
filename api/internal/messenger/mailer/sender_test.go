package mailer

import (
	"testing"
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 2, 2, 18, 0, 0, 0)
	type args struct {
		templateID       string
		fromName         string
		fromAddress      string
		subject          string
		contents         []*Content
		personalizations []*Personalization
	}
	tests := []struct {
		name   string
		args   args
		expect *mail.SGMailV3
	}{
		{
			name: "success",
			args: args{
				templateID:  "template-id",
				fromName:    "test user",
				fromAddress: "test-user@calmato.jp",
				subject:     "テストタイトル",
				contents: []*Content{{
					ContentType: "type",
					Value:       "value",
				}},
				personalizations: []*Personalization{
					{
						Name:          "test to user",
						Address:       "test-to@calmato.jp",
						Type:          AddressTypeTo,
						Substitutions: map[string]interface{}{"key": "value"},
					},
					{
						Name:          "test cc user",
						Address:       "test-cc@calmato.jp",
						Type:          AddressTypeCC,
						Substitutions: map[string]interface{}{"key": "value"},
					},
					{
						Name:          "test bcc user",
						Address:       "test-bcc@calmato.jp",
						Type:          AddressTypeBCC,
						Substitutions: map[string]interface{}{"key": "value"},
					},
				},
			},
			expect: &mail.SGMailV3{
				From: &mail.Email{
					Name:    "test user",
					Address: "test-user@calmato.jp",
				},
				Subject: "テストタイトル",
				Personalizations: []*mail.Personalization{
					{
						To: []*mail.Email{{
							Name:    "test to user",
							Address: "test-to@calmato.jp",
						}},
						CC:            []*mail.Email{},
						BCC:           []*mail.Email{},
						Headers:       map[string]string{},
						Substitutions: map[string]string{},
						CustomArgs:    map[string]string{},
						DynamicTemplateData: map[string]interface{}{
							"key": "value",
						},
						Categories: []string{},
					},
					{
						To: []*mail.Email{},
						CC: []*mail.Email{{
							Name:    "test cc user",
							Address: "test-cc@calmato.jp",
						}},
						BCC:           []*mail.Email{},
						Headers:       map[string]string{},
						Substitutions: map[string]string{},
						CustomArgs:    map[string]string{},
						DynamicTemplateData: map[string]interface{}{
							"key": "value",
						},
						Categories: []string{},
					},
					{
						To: []*mail.Email{},
						CC: []*mail.Email{},
						BCC: []*mail.Email{{
							Name:    "test bcc user",
							Address: "test-bcc@calmato.jp",
						}},
						Headers:       map[string]string{},
						Substitutions: map[string]string{},
						CustomArgs:    map[string]string{},
						DynamicTemplateData: map[string]interface{}{
							"key": "value",
						},
						Categories: []string{},
					},
				},
				Content: []*mail.Content{{
					Type:  "type",
					Value: "value",
				}},
				Attachments: []*mail.Attachment{},
				Categories:  []string{"template-id"},
				TemplateID:  "template-id",
				SendAt:      int(now.Unix()),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			client := &client{now: func() time.Time { return now }}
			assert.Equal(t, tt.expect, client.newMessage(
				tt.args.templateID,
				tt.args.fromName,
				tt.args.fromAddress,
				tt.args.subject,
				tt.args.contents,
				tt.args.personalizations,
			))
		})
	}
}

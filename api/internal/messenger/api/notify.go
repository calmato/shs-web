package api

import (
	"context"
	"strconv"

	"github.com/calmato/shs-web/api/internal/messenger/entity"
	"github.com/calmato/shs-web/api/internal/messenger/mailer"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/pkg/pubsub"
	"github.com/calmato/shs-web/api/pkg/uuid"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/messenger"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
)

// メッセージ登録 (授業スケジュール確定)
func (s *messengerService) NotifyLessonDecided(
	ctx context.Context, req *messenger.NotifyLessonDecidedRequest,
) (*messenger.NotifyLessonDecidedResponse, error) {
	if err := s.validator.NotifyLessonDecided(req); err != nil {
		return nil, gRPCError(err)
	}

	eg, ectx := errgroup.WithContext(ctx)
	var summary *lesson.ShiftSummary
	eg.Go(func() error {
		in := &lesson.GetShiftSummaryRequest{Id: req.ShiftSummaryId}
		out, err := s.lesson.GetShiftSummary(ectx, in)
		if err != nil {
			return err
		}
		summary = out.Summary
		return nil
	})
	var lessons entity.Lessons
	eg.Go(func() error {
		in := &lesson.ListLessonsRequest{ShiftSummaryId: req.ShiftSummaryId}
		out, err := s.lesson.ListLessons(ectx, in)
		if err != nil {
			return err
		}
		lessons = entity.NewLessons(out.Lessons)
		return nil
	})
	if err := eg.Wait(); err != nil {
		return nil, gRPCError(err)
	}
	if len(lessons) == 0 {
		return &messenger.NotifyLessonDecidedResponse{}, nil
	}

	yearMonth, err := jst.ParseFromYYYYMM(strconv.FormatInt(int64(summary.YearMonth), 10))
	if err != nil {
		return nil, gRPCError(err)
	}
	builder := mailer.NewTemplateDataBuilder().YearMonth(yearMonth)
	params := &messenger.NotifierRequest{
		Key:        uuid.Base58Encode(uuid.New()),
		TeacherIds: lessons.TeacherIDs(),
		StudentIds: lessons.StudentIDs(),
		Email: &messenger.EmailConfig{
			EmailId:       mailer.EmailIDLessonDecided,
			Substitutions: builder.Build(),
		},
	}
	if err := s.publish(ctx, params); err != nil {
		return nil, gRPCError(err)
	}
	return &messenger.NotifyLessonDecidedResponse{}, nil
}

func (s *messengerService) publish(ctx context.Context, pb proto.Message) error {
	data, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	msg := &pubsub.Message{Data: data}
	_, err = s.publisher.Publish(ctx, msg)
	return err
}

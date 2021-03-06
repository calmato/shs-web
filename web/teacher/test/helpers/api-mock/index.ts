import * as AuthStore from './auth'
import * as LessonStore from './lesson'
import * as ShiftStore from './shift'
import * as SubmissionStore from './submission'
import * as UserStore from './user'
import * as ClassroomStore from './classroom'
import { ErrorResponse } from '~/types/api/exception'

const err: { response: { data: ErrorResponse } } = {
  response: {
    data: {
      status: 400,
      message: 'api error',
      details: 'some error',
    },
  },
}

export default {
  get: {
    ...AuthStore.showAuth,
    ...LessonStore.listSubjects,
    ...ShiftStore.listShiftSummaries,
    ...ShiftStore.listShiftDetails,
    ...SubmissionStore.listTeacherSubmissions,
    ...SubmissionStore.listTeacherShifts,
    ...UserStore.listTeachers,
    ...UserStore.showTeacher,
    ...ClassroomStore.getTotalRoomsByApi,
    ...ClassroomStore.getSchedulesByApi,
  },
  post: {
    ...ShiftStore.createShifts,
    ...SubmissionStore.submitTeacherShifts,
    ...UserStore.createTeacher,
  },
  patch: {
    ...ShiftStore.updateShiftSummarySchedule,
    ...UserStore.updateTeacherSubjects,
    ...UserStore.updateTeacherRole,
  },
  put: {},
  delete: {
    ...ShiftStore.deleteShifts,
    ...UserStore.deleteTeacher,
  },
  error: err,
}

import * as AuthStore from './auth'
import * as ShiftStore from './shift'
import * as UserStore from './user'
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
    ...ShiftStore.listShiftSummaries,
    ...ShiftStore.listShiftDetails,
    ...UserStore.listTeachers,
  },
  post: {
    ...ShiftStore.createShifts,
    ...UserStore.createTeacher,
  },
  patch: {},
  put: {},
  delete: {},
  error: err,
}

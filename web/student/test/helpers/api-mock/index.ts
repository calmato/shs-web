import * as AuthStore from './auth'
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
  },
  post: {},
  patch: {},
  put: {},
  delete: {},
  error: err,
}

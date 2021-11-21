import * as UserStore from './user'
import { ErrorResponse } from '~/types/api/exception'

const err: { response: ErrorResponse } = {
  response: {
    status: 400,
    message: 'api error',
    details: 'some error',
  },
}

export default {
  get: {
  },
  post: {
    ...UserStore.hello,
  },
  patch: {
  },
  put: {},
  delete: {
  },
  error: err,
}

import { ErrorResponse } from '~/types/api/exception'

export interface IErrorResponse {
  status: number
  data: ErrorResponse
}

export class ApiError extends Error {
  constructor(public status: number, public message: string, public data?: ErrorResponse) {
    super(message)
  }
}

import { ShiftSummariesResponse } from '~/types/api/v1'

export const listShiftSummaries: { [key: string]: ShiftSummariesResponse } = {
  '/v1/shifts': {
    summaries: [
      {
        id: 1,
        year: 2021,
        month: 2,
        status: 3,
        openAt: '2021-01-01T00:00:00+09:00',
        endAt: '2021-01-15T00:00:00+09:00',
        createdAt: '2021-12-30T19:25:57+09:00',
        updatedAt: '2021-12-30T19:25:57+09:00',
      },
      {
        id: 2,
        year: 2022,
        month: 2,
        status: 1,
        openAt: '2021-01-01T00:00:00+09:00',
        endAt: '2021-01-15T00:00:00+09:00',
        createdAt: '2021-12-30T19:25:57+09:00',
        updatedAt: '2021-12-30T19:25:57+09:00',
      },
    ],
  },
  '/v1/shifts?limit=20&offset=0&status=1': {
    summaries: [
      {
        id: 2,
        year: 2022,
        month: 2,
        status: 1,
        openAt: '2021-01-01T00:00:00+09:00',
        endAt: '2021-01-15T00:00:00+09:00',
        createdAt: '2021-12-30T19:25:57+09:00',
        updatedAt: '2021-12-30T19:25:57+09:00',
      },
    ],
  },
}

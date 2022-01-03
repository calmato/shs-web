import { TeacherShiftsResponse, TeacherSubmissionsResponse } from '~/types/api/v1'

export const listTeacherSubmissions: { [key: string]: TeacherSubmissionsResponse } = {
  '/v1/teachers/teacherid/submissions': {
    summaries: [
      {
        id: 1,
        year: 2021,
        month: 2,
        shiftStatus: 3,
        submissionStatus: 2,
        openAt: '2021-01-01T00:00:00+09:00',
        endAt: '2021-01-15T00:00:00+09:00',
        createdAt: '2021-12-30T19:25:57+09:00',
        updatedAt: '2021-12-30T19:25:57+09:00',
      },
      {
        id: 2,
        year: 2022,
        month: 2,
        shiftStatus: 1,
        submissionStatus: 1,
        openAt: '2021-01-01T00:00:00+09:00',
        endAt: '2021-01-15T00:00:00+09:00',
        createdAt: '2021-12-30T19:25:57+09:00',
        updatedAt: '2021-12-30T19:25:57+09:00',
      },
    ],
  },
}

export const listTeacherShifts: { [key: string]: TeacherShiftsResponse } = {
  '/v1/teachers/teacherid/submissions/1': {
    summary: {
      id: 1,
      year: 2022,
      month: 2,
      shiftStatus: 3,
      submissionStatus: 2,
      openAt: '2021-01-01T00:00:00+09:00',
      endAt: '2021-01-15T00:00:00+09:00',
      createdAt: '2021-12-30T19:25:57+09:00',
      updatedAt: '2021-12-30T19:25:57+09:00',
    },
    shifts: [
      {
        date: '20210201',
        isClosed: false,
        lessons: [
          { id: 1, enabled: true, startTime: '1700', endTime: '1830' },
          { id: 2, enabled: false, startTime: '1830', endTime: '2000' },
        ],
      },
      {
        date: '20210202',
        isClosed: true,
        lessons: [],
      },
    ],
  },
}

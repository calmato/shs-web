import { ShiftDetailsResponse, ShiftSummariesResponse } from '~/types/api/v1'

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

export const updateShiftSummarySchedule: { [key: string]: {} } = {
  '/v1/shifts/1/schedule': {},
}

export const createShifts: { [key: string]: ShiftDetailsResponse } = {
  '/v1/shifts': {
    summary: {
      id: 1,
      year: 2022,
      month: 2,
      status: 3,
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
          { id: 1, startTime: '1700', endTime: '1830' },
          { id: 2, startTime: '1830', endTime: '2000' },
        ],
      },
      {
        date: '20210202',
        isClosed: true,
        lessons: [],
      },
    ],
    rooms: 4,
    teachers: [],
    students: [],
    lessons: [],
  },
}

export const deleteShifts: { [key: string]: {} } = {
  '/v1/shifts/1': {},
}

export const listShiftDetails: { [key: string]: ShiftDetailsResponse } = {
  '/v1/shifts/1': {
    summary: {
      id: 1,
      year: 2022,
      month: 2,
      status: 3,
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
          { id: 1, startTime: '1700', endTime: '1830' },
          { id: 2, startTime: '1830', endTime: '2000' },
        ],
      },
      {
        date: '20210202',
        isClosed: true,
        lessons: [],
      },
    ],
    rooms: 4,
    teachers: [
      {
        teacher: {
          id: '000000000000000000001',
          lastName: '中村',
          firstName: '太郎',
          lastNameKana: 'なかむら',
          firstNameKana: 'たろう',
          mail: 'teacher-001@calmato.jp',
          role: 1,
          subjects: [],
          createdAt: '2021-12-02T18:30:00+09:00',
          updatedAt: '2021-12-02T18:30:00+09:00',
        },
        lessonTotal: 0,
      },
      {
        teacher: {
          id: '000000000000000000002',
          lastName: '西山',
          firstName: '幸子',
          lastNameKana: 'にしやま',
          firstNameKana: 'さちこ',
          mail: 'teacher-002@calmato.jp',
          role: 1,
          subjects: [],
          createdAt: '2021-12-02T18:30:00+09:00',
          updatedAt: '2021-12-02T18:30:00+09:00',
        },
        lessonTotal: 2,
      },
      {
        teacher: {
          id: '000000000000000000003',
          lastName: '鈴木',
          firstName: '小太郎',
          lastNameKana: 'すずき',
          firstNameKana: 'こたろう',
          mail: 'teacher-003@calmato.jp',
          role: 2,
          subjects: [],
          createdAt: '2021-12-02T18:30:00+09:00',
          updatedAt: '2021-12-02T18:30:00+09:00',
        },
        lessonTotal: 1,
      },
    ],
    students: [
      {
        student: {
          id: '100000000000000000001',
          lastName: '中村',
          firstName: '太郎',
          lastNameKana: 'なかむら',
          firstNameKana: 'たろう',
          mail: 'student-001@calmato.jp',
          schoolType: 3,
          grade: 2,
          subjects: [],
          createdAt: '2021-12-02T18:30:00+09:00',
          updatedAt: '2021-12-02T18:30:00+09:00',
        },
        lessonTotal: 0,
        suggestedLessonsTotal: 0,
      },
      {
        student: {
          id: '100000000000000000002',
          lastName: '西山',
          firstName: '幸子',
          lastNameKana: 'にしやま',
          firstNameKana: 'さちこ',
          mail: 'student-002@calmato.jp',
          schoolType: 1,
          grade: 6,
          subjects: [],
          createdAt: '2021-12-02T18:30:00+09:00',
          updatedAt: '2021-12-02T18:30:00+09:00',
        },
        lessonTotal: 2,
        suggestedLessonsTotal: 0,
      },
    ],
    lessons: [
      {
        id: 1,
        shiftId: 1,
        subjectId: 1,
        room: 1,
        teacherId: '000000000000000000001',
        studentId: '100000000000000000001',
        startAt: '2022-02-01T17:00:00+09:00',
        endAt: '2022-02-01T18:30:00+09:00',
        notes: '',
        createdAt: '2021-12-02T18:30:00+09:00',
        updatedAt: '2021-12-02T18:30:00+09:00',
      },
    ],
  },
}

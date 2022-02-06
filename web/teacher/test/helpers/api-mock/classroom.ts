import { ScheduleResponse, TotalRoomsResponse } from '~/types/api/v1'

export const getTotalRoomsByApi: { [key: string]: TotalRoomsResponse } = {
  '/v1/rooms': {
    total: 5,
  },
}

export const getSchedulesByApi: { [key: string]: ScheduleResponse } = {
  '/v1/schedules': {
    schedules: [
      {
        weekday: 0,
        isClosed: false,
        lessons: [
          { startTime: '1530', endTime: '1700' },
          { startTime: '1700', endTime: '1830' },
          { startTime: '1830', endTime: '2000' },
          { startTime: '2000', endTime: '2130' },
        ],
      },
      {
        weekday: 1,
        isClosed: false,
        lessons: [
          { startTime: '1700', endTime: '1830' },
          { startTime: '1830', endTime: '2000' },
          { startTime: '2000', endTime: '2130' },
        ],
      },
      {
        weekday: 2,
        isClosed: false,
        lessons: [
          { startTime: '1700', endTime: '1830' },
          { startTime: '1830', endTime: '2000' },
          { startTime: '2000', endTime: '2130' },
        ],
      },
      {
        weekday: 3,
        isClosed: false,
        lessons: [
          { startTime: '1700', endTime: '1830' },
          { startTime: '1830', endTime: '2000' },
          { startTime: '2000', endTime: '2130' },
        ],
      },
      {
        weekday: 4,
        isClosed: false,
        lessons: [
          { startTime: '1700', endTime: '1830' },
          { startTime: '1830', endTime: '2000' },
          { startTime: '2000', endTime: '2130' },
        ],
      },
      {
        weekday: 5,
        isClosed: false,
        lessons: [
          { startTime: '1700', endTime: '1830' },
          { startTime: '1830', endTime: '2000' },
          { startTime: '2000', endTime: '2130' },
        ],
      },
      {
        weekday: 6,
        isClosed: false,
        lessons: [
          { startTime: '1530', endTime: '1700' },
          { startTime: '1700', endTime: '1830' },
          { startTime: '1830', endTime: '2000' },
          { startTime: '2000', endTime: '2130' },
        ],
      },
    ],
  },
}

export interface LessonTime {
  startTime: string
  endTime: string
}

export interface Schedule {
  weekday: 0 | 1 | 2 | 3 | 4 | 5 | 6
  isClosed: boolean
  lessons: LessonTime[]
}

export interface ScheduleResponse {
  schedules: Schedule[]
}

export interface TotalRoomsResponse {
  total: number
}

export interface ScheduleRequest {
  schedules: Schedule[]
}

export interface TotalRoomsRequest {
  total: number
}

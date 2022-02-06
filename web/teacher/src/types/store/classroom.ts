import { HourForm } from '~/types//form'
import { Schedule } from '~/types//api/v1/classroom'

export interface ClassroomState {
  totalRooms: number
  schedules: Schedule[]
}

export interface UpdateSchedulesPayload {
  regularHoliday: number[]
  weekdayHourForm: HourForm[]
  holidayHourForm: HourForm[]
}

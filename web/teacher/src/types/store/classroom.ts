import { Schedule } from '~/types//api/v1/classroom'

export interface ClassroomState {
  totalRooms: number
  schedules: Schedule[]
}

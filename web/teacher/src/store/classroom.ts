import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators'
import axios from 'axios'
import { $axios } from '~/plugins/axios'
import {
  LessonTime,
  Schedule,
  ScheduleRequest,
  ScheduleResponse,
  TotalRoomsRequest,
  TotalRoomsResponse,
} from '~/types/api/v1'
import { ClassroomState, UpdateSchedulesPayload } from '~/types/store'
import { ErrorResponse } from '~/types/api/exception'
import { ApiError } from '~/types/exception'
import { HourForm } from '~/types/form'
import dayjs from '~/plugins/dayjs'

const initialState: ClassroomState = {
  totalRooms: 0,
  schedules: [],
}

function extractHourFormBySchedule(schedule: Schedule): HourForm[] {
  return schedule.lessons.map((item) => {
    return {
      startAt: dayjs(item.startTime, 'HHmm').format('HH:mm'),
      endAt: dayjs(item.endTime, 'HHmm').format('HH:mm'),
    }
  })
}

function hourForm2LessonTime(hourForm: HourForm): LessonTime {
  return {
    startTime: dayjs(hourForm.startAt, 'HH:mm').format('HHmm'),
    endTime: dayjs(hourForm.endAt, 'HH:mm').format('HHmm'),
  }
}

@Module({
  name: 'classroom',
  stateFactory: true,
  namespaced: true,
})
export default class ClassroomModule extends VuexModule {
  private totalRooms: ClassroomState['totalRooms'] = initialState.totalRooms
  private schedules: ClassroomState['schedules'] = initialState.schedules

  public get getTotalRooms(): number {
    return this.totalRooms
  }

  public get getSchedules(): Schedule[] {
    return this.schedules
  }

  public get weekdayHourFormValue(): HourForm[] {
    const weekdays: number[] = [1, 2, 3, 4, 5] // 1: 月 ~ 5:金
    const schedules: Schedule[] = this.getSchedules.filter((item) => weekdays.includes(item.weekday))
    const target: Schedule | undefined = schedules.find((item): boolean => !item.isClosed)
    return target ? extractHourFormBySchedule(target) : []
  }

  public get holidayHourFormValue(): HourForm[] {
    const holidays: number[] = [0, 6] // 0: 日, 6: 土
    const schedules: Schedule[] = this.getSchedules.filter((item) => holidays.includes(item.weekday))
    const target: Schedule | undefined = schedules.find((item): boolean => !item.isClosed)
    return target ? extractHourFormBySchedule(target) : []
  }

  public get regularHoliday(): number[] {
    return this.getSchedules.filter((item) => item.isClosed).map((item) => item.weekday)
  }

  @Mutation
  private setTotalRooms(totalRooms: number): void {
    this.totalRooms = totalRooms
  }

  @Mutation
  private setSchedules(schedules: Schedule[]): void {
    this.schedules = schedules
  }

  @Action({})
  public factory(): void {
    this.setTotalRooms(initialState.totalRooms)
    this.setSchedules(initialState.schedules)
  }

  @Action({ rawError: true })
  public async getTotalRoomsByApi(): Promise<void> {
    try {
      const res: TotalRoomsResponse = await $axios.$get('/v1/rooms')
      this.setTotalRooms(res.total)
    } catch (err) {
      if (axios.isAxiosError(err)) {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      }
      throw new Error('internal server error')
    }
  }

  @Action({ rawError: true })
  public async getSchedulesByApi(): Promise<void> {
    try {
      const res: ScheduleResponse = await $axios.$get('/v1/schedules')
      this.setSchedules(res.schedules)
    } catch (err) {
      if (axios.isAxiosError(err)) {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      }
      throw new Error('internal server error')
    }
  }

  @Action({ rawError: true })
  public async updateTotalRooms(payload: TotalRoomsRequest): Promise<void> {
    try {
      await $axios.$patch('/v1/rooms', payload)
      await this.getTotalRoomsByApi()
    } catch (err) {
      if (axios.isAxiosError(err)) {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      }
      throw new Error('internal server error')
    }
  }

  @Action({ rawError: true })
  public async updateSchedules(payload: UpdateSchedulesPayload): Promise<void> {
    try {
      const HOLIDAY: number[] = [0, 6]
      const { regularHoliday, weekdayHourForm, holidayHourForm } = payload

      const updateSchedules: ScheduleRequest = {
        schedules: this.schedules.map((item) => {
          return {
            weekday: item.weekday,
            isClosed: regularHoliday.includes(item.weekday),
            lessons: HOLIDAY.includes(item.weekday)
              ? holidayHourForm.map((i) => hourForm2LessonTime(i))
              : weekdayHourForm.map((i) => hourForm2LessonTime(i)),
          }
        }),
      }

      await $axios.$patch('/v1/schedules', updateSchedules)
      await this.getSchedulesByApi()
    } catch (err) {
      if (axios.isAxiosError(err)) {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      }
      throw new Error('internal server error')
    }
  }
}

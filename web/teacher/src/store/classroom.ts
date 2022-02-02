import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators'
import axios from 'axios'
import { $axios } from '~/plugins/axios'
import { Schedule, ScheduleResponse, TotalRoomsResponse } from '~/types/api/v1'
import { ClassroomState } from '~/types/store'
import { ErrorResponse } from '~/types/api/exception'
import { ApiError } from '~/types/exception'

const initialState: ClassroomState = {
  totalRooms: 0,
  schedules: [],
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

  @Mutation
  private setTotalRooms(totalRooms: number): void {
    this.totalRooms = totalRooms
  }

  @Mutation
  private setSchedules(schedules: Schedule[]): void {
    this.schedules = schedules
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
}

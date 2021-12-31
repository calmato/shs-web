import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { AxiosError } from 'axios'
import { $axios } from '~/plugins/axios'
import {
  ShiftSummariesResponse,
  ShiftSummary as ShiftSummaryResponse,
  ShiftDetail as ShiftDetailResponse,
  ShiftDetailLesson as LessonResponse,
} from '~/types/api/v1'
import { ShiftDetail, ShiftStatus, ShiftState, ShiftSummary, ShiftDetailLesson } from '~/types/store'
import { ErrorResponse } from '~/types/api/exception'
import { ApiError } from '~/types/exception'
import { ShiftsNewForm } from '~/types/form'
import { CreateShiftsRequest, ShiftDetailsResponse } from '~/types/api/v1/shift'

const initialState: ShiftState = {
  summary: {
    id: 0,
    year: 0,
    month: 0,
    status: ShiftStatus.UNKNOWN,
    openAt: '',
    endAt: '',
    createdAt: '',
    updatedAt: '',
  },
  summaries: [],
  details: new Map<string, ShiftDetail>(),
}

@Module({
  name: 'shift',
  stateFactory: true,
  namespaced: true,
})
export default class ShiftModule extends VuexModule {
  private summary: ShiftState['summary'] = initialState.summary
  private summaries: ShiftState['summaries'] = initialState.summaries
  private details: ShiftState['details'] = initialState.details

  public get getSummary(): ShiftSummary {
    return this.summary
  }

  public get getSummaries(): ShiftSummary[] {
    return this.summaries
  }

  public get getDetails(): Map<string, ShiftDetail> {
    return this.details
  }

  @Mutation
  private setSummaries({ summaries }: { summaries: ShiftSummary[] }): void {
    this.summaries = summaries
  }

  @Mutation
  private setDetails({ summary, details }: { summary: ShiftSummary; details: Map<string, ShiftDetail> }): void {
    this.summary = summary
    this.details = details
  }

  @Action({})
  public factory(): void {
    this.setSummaries({ summaries: initialState.summaries })
    this.setDetails({ summary: initialState.summary, details: initialState.details })
  }

  @Action({ rawError: true })
  public async listShiftSummaries({
    limit,
    offset,
    status,
  }: {
    limit: number
    offset: number
    status: ShiftStatus
  }): Promise<void> {
    let query: string = ''
    if (limit !== 0 || offset !== 0 || status !== ShiftStatus.UNKNOWN) {
      query = `?limit=${limit}&offset=${offset}&status=${status}`
    }

    await $axios
      .$get('/v1/shifts' + query)
      .then((res: ShiftSummariesResponse) => {
        const summaries: ShiftSummary[] = res.summaries.map((data: ShiftSummaryResponse): ShiftSummary => {
          return { ...data }
        })
        this.setSummaries({ summaries })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async createShifts({ form }: { form: ShiftsNewForm }): Promise<void> {
    const closedDates = form.params.closedDates.map((closedDate: string): string => {
      return replaceDate(closedDate, '-', '')
    })

    const req: CreateShiftsRequest = {
      yearMonth: replaceDate(form.params.yearMonth, '-', ''),
      openDate: replaceDate(form.params.openDate, '-', ''),
      endDate: replaceDate(form.params.endDate, '-', ''),
      closedDates,
    }

    await $axios
      .$post('/v1/shifts', req)
      .then((res: ShiftDetailsResponse) => {
        const summary: ShiftSummary = { ...res.summary }
        const details = new Map<string, ShiftDetail>()

        Object.keys(res.shifts).forEach((key: string) => {
          const shift: ShiftDetailResponse = res.shifts[key]
          const lessons: ShiftDetailLesson[] = shift.lessons.map((lesson: LessonResponse): ShiftDetailLesson => {
            return { ...lesson }
          })
          details.set(key, { isClosed: shift.isClosed, lessons })
        })

        this.setDetails({ summary, details })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }
}

function replaceDate(date: string, oldVal: string, newVal: string): string {
  return date.replaceAll(oldVal, newVal)
}

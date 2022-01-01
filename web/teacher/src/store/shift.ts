import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { AxiosError } from 'axios'
import { $axios } from '~/plugins/axios'
import dayjs from '~/plugins/dayjs'
import {
  CreateShiftsRequest,
  ShiftDetailsResponse,
  ShiftSummariesResponse,
  ShiftSummary as ShiftSummaryResponse,
  ShiftDetail as ShiftDetailResponse,
  ShiftDetailLesson as LessonResponse,
  UpdateShiftSummaryScheduleRequest,
} from '~/types/api/v1'
import { ShiftDetail, ShiftStatus, ShiftState, ShiftSummary, ShiftDetailLesson } from '~/types/store'
import { ErrorResponse } from '~/types/api/exception'
import { ApiError } from '~/types/exception'
import { ShiftsNewForm, ShiftSummaryEditScheduleForm } from '~/types/form'

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
  details: [],
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

  public get getDetails(): ShiftDetail[] {
    return this.details
  }

  @Mutation
  private setSummaries({ summaries }: { summaries: ShiftSummary[] }): void {
    this.summaries = summaries
  }

  @Mutation
  private setDetails({ summary, details }: { summary: ShiftSummary; details: ShiftDetail[] }): void {
    this.summary = summary
    this.details = details
  }

  @Mutation
  private addSummaries({ summary }: { summary: ShiftSummary }): void {
    this.summaries.unshift(summary)
  }

  @Mutation
  private replaceSummariesSchedule({
    summaryId,
    openAt,
    endAt,
  }: {
    summaryId: number
    openAt: string
    endAt: string
  }): void {
    const index: number = this.summaries.findIndex((val: ShiftSummary) => {
      return val.id === summaryId
    })
    if (index === -1) {
      return
    }
    this.summaries.splice(index, 1, { ...this.summaries[index], openAt, endAt })
  }

  @Mutation
  private removeSummaries({ summaryId }: { summaryId: number }): void {
    const index: number = this.summaries.findIndex((val: ShiftSummary) => {
      return val.id === summaryId
    })
    if (index === -1) {
      return
    }
    this.summaries.splice(index, 1)
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
  public async updateShiftSummarySchedule({ form }: { form: ShiftSummaryEditScheduleForm }): Promise<void> {
    const summaryId: number = form.params.summaryId
    const req: UpdateShiftSummaryScheduleRequest = {
      openDate: replaceDate(form.params.openDate, '-', ''),
      endDate: replaceDate(form.params.endDate, '-', ''),
    }

    await $axios
      .$patch(`/v1/shifts/${summaryId}/schedule`, req)
      .then(() => {
        const format: string = 'YYYY-MM-DDThh:mm:ss'
        const openAt: string = dayjs(form.params.openDate).tz().format(format)
        const endAt: string = dayjs(form.params.endDate).tz().format(format)

        this.replaceSummariesSchedule({ summaryId, openAt, endAt })
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

        const details: ShiftDetail[] = res.shifts.map((shift: ShiftDetailResponse): ShiftDetail => {
          const lessons: ShiftDetailLesson[] = shift.lessons.map((lesson: LessonResponse): ShiftDetailLesson => {
            return { ...lesson }
          })
          return { ...shift, lessons }
        })

        this.addSummaries({ summary })
        this.setDetails({ summary, details })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async deleteShifts({ summaryId }: { summaryId: number }): Promise<void> {
    await $axios
      .$delete(`/v1/shifts/${summaryId}`)
      .then(() => {
        this.removeSummaries({ summaryId })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async listShiftDetails({ summaryId }: { summaryId: number }): Promise<void> {
    await $axios
      .$get(`/v1/shifts/${summaryId}`)
      .then((res: ShiftDetailsResponse) => {
        const summary: ShiftSummary = { ...res.summary }

        const details: ShiftDetail[] = res.shifts.map((shift: ShiftDetailResponse): ShiftDetail => {
          const lessons: ShiftDetailLesson[] = shift.lessons.map((lesson: LessonResponse): ShiftDetailLesson => {
            return { ...lesson }
          })
          return { ...shift, lessons }
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

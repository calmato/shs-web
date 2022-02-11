import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators'
import { AxiosError } from 'axios'
import {
  SubmissionRequest,
  SubmissionResponse,
  SubmissionsResponse,
  SubmissionDetail as v1SubmissionDetail,
  SubmissionDetailLesson as v1SubmissionDetailLesson,
  SubmissionSummary as v1SubmissionSummary,
} from '~/types/api/v1'
import {
  ShiftStatus,
  SubmissionState,
  SubmissionStatus,
  SubmissionDetail,
  SubmissionDetailLesson,
  SubmissionSummary,
} from '~/types/store'
import { $axios } from '~/plugins/axios'
import { ErrorResponse } from '~/types/api/exception'
import { ApiError } from '~/types/exception'

const initialState: SubmissionState = {
  summary: {
    id: 0,
    year: 0,
    month: 0,
    shiftStatus: ShiftStatus.UNKNOWN,
    submissionStatus: SubmissionStatus.UNKNOWN,
    openAt: '',
    endAt: '',
    createdAt: '',
    updatedAt: '',
  },
  summaries: [],
  shifts: [],
}

@Module({
  name: 'submission',
  stateFactory: true,
  namespaced: true,
})
export default class SubmissionModule extends VuexModule {
  private summary: SubmissionState['summary'] = initialState.summary
  private summaries: SubmissionState['summaries'] = initialState.summaries
  private shifts: SubmissionState['shifts'] = initialState.shifts

  public get getSummary(): SubmissionSummary {
    return this.summary
  }

  public get getSummaries(): SubmissionSummary[] {
    return this.summaries
  }

  public get getShifts(): SubmissionDetail[] {
    return this.shifts
  }

  @Mutation
  private setSummaries({ summaries }: { summaries: SubmissionSummary[] }): void {
    this.summaries = summaries
  }

  @Mutation
  private setShifts({ summary, shifts }: { summary: SubmissionSummary; shifts: SubmissionDetail[] }): void {
    this.summary = summary
    this.shifts = shifts
  }

  @Action({})
  public factory(): void {
    this.setSummaries({ ...initialState })
    this.setShifts({ ...initialState })
  }

  @Action({ rawError: true })
  public async listStudentSubmissions(): Promise<void> {
    await $axios
      .$get(`/v1/submissions`)
      .then((res: SubmissionsResponse) => {
        const summaries: SubmissionSummary[] = res.summaries.map((data: v1SubmissionSummary): SubmissionSummary => {
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
  public async listStudentShifts({ studentId, shiftId }: { studentId: string; shiftId: number }): Promise<void> {
    await $axios
      .$get(`/v1/students/${studentId}/submissions/${shiftId}`)
      .then((res: SubmissionResponse) => {
        const summary: SubmissionSummary = { ...res.summary }
        const shifts: SubmissionDetail[] = res.shifts.map((shift: v1SubmissionDetail): SubmissionDetail => {
          const lessons: SubmissionDetailLesson[] = shift.lessons.map(
            (lesson: v1SubmissionDetailLesson): SubmissionDetailLesson => {
              return { ...lesson }
            }
          )
          return { ...shift, lessons }
        })
        this.setShifts({ summary, shifts })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async submitStudentShifts({
    studentId,
    shiftId,
    lessonIds,
  }: {
    studentId: string
    shiftId: number
    lessonIds: number[]
  }): Promise<void> {
    const req: SubmissionRequest = {
      shiftIds: lessonIds,
    }

    await $axios.$post(`/v1/students/${studentId}/submissions/${shiftId}`, req).catch((err: AxiosError) => {
      const res: ErrorResponse = { ...err.response?.data }
      throw new ApiError(res.status, res.message, res)
    })
  }
}

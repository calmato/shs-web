import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators'
import { AxiosError } from 'axios'
import {
  TeacherSubmissionsResponse,
  TeacherShiftSummary as SummaryResponse,
  TeacherShiftsResponse,
  TeacherShiftDetail as ShiftDetailResponse,
  TeacherShiftDetailLesson as LessonResponse,
  SubmitTeacherShiftRequest,
} from '~/types/api/v1'
import {
  ShiftStatus,
  SubmissionState,
  SubmissionStatus,
  TeacherShiftDetail,
  TeacherShiftDetailLesson,
  TeacherShiftSummary,
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

  public get getSummary(): TeacherShiftSummary {
    return this.summary
  }

  public get getSummaries(): TeacherShiftSummary[] {
    return this.summaries
  }

  public get getShifts(): TeacherShiftDetail[] {
    return this.shifts
  }

  @Mutation
  private setSummaries({ summaries }: { summaries: TeacherShiftSummary[] }): void {
    this.summaries = summaries
  }

  @Mutation
  private setShifts({ summary, shifts }: { summary: TeacherShiftSummary; shifts: TeacherShiftDetail[] }): void {
    this.summary = summary
    this.shifts = shifts
  }

  @Action({})
  public factory(): void {
    this.setSummaries({ ...initialState })
    this.setShifts({ ...initialState })
  }

  @Action({ rawError: true })
  public async listTeacherSubmissions({ teacherId }: { teacherId: string }): Promise<void> {
    await $axios
      .$get(`/v1/teachers/${teacherId}/submissions`)
      .then((res: TeacherSubmissionsResponse) => {
        const summaries: TeacherShiftSummary[] = res.summaries.map((data: SummaryResponse): TeacherShiftSummary => {
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
  public async listTeacherShifts({ teacherId, shiftId }: { teacherId: string; shiftId: number }): Promise<void> {
    await $axios
      .$get(`/v1/teachers/${teacherId}/submissions/${shiftId}`)
      .then((res: TeacherShiftsResponse) => {
        const summary: TeacherShiftSummary = { ...res.summary }
        const shifts: TeacherShiftDetail[] = res.shifts.map((shift: ShiftDetailResponse): TeacherShiftDetail => {
          const lessons: TeacherShiftDetailLesson[] = shift.lessons.map(
            (lesson: LessonResponse): TeacherShiftDetailLesson => {
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
  public async submitTeacherShifts({
    teacherId,
    shiftId,
    lessonIds,
  }: {
    teacherId: string
    shiftId: number
    lessonIds: number[]
  }): Promise<void> {
    const req: SubmitTeacherShiftRequest = {
      shiftIds: lessonIds,
    }

    await $axios.$post(`/v1/teachers/${teacherId}/submissions/${shiftId}`, req).catch((err: AxiosError) => {
      const res: ErrorResponse = { ...err.response?.data }
      throw new ApiError(res.status, res.message, res)
    })
  }
}

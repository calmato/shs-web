import { Action, Module, Mutation, VuexModule } from 'vuex-module-decorators'
import { AxiosError } from 'axios'
import {
  SubmissionResponse,
  SubmissionDetail as v1SubmissionDetail,
  SubmissionDetailLesson as v1SubmissionDetailLesson,
  SubmissionLesson as v1SubmissionLesson,
  SubmissionRequest,
  SubmissionSummary as v1SubmissionSummary,
  SubmissionsResponse,
  SubmissionTemplate as v1SubmissionTemplate,
  SubmissionTemplateLesson as v1SubmissionTemplateLesson,
  SubmissionTemplateResponse,
  SubmissionTemplateRequest,
} from '~/types/api/v1'
import {
  ShiftStatus,
  SubmissionState,
  SubmissionStatus,
  SubmissionDetail,
  SubmissionDetailLesson,
  SubmissionLesson,
  SubmissionSummary,
  SubmissionTemplate,
  SubmissionTemplateLesson,
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
  templates: [],
  suggestedLessons: [],
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
  private templates: SubmissionState['templates'] = initialState.templates
  private lessons: SubmissionState['suggestedLessons'] = initialState.suggestedLessons

  public get getSummary(): SubmissionSummary {
    return this.summary
  }

  public get getSummaries(): SubmissionSummary[] {
    return this.summaries
  }

  public get getShifts(): SubmissionDetail[] {
    return this.shifts
  }

  public get getTemplates(): SubmissionTemplate[] {
    return this.templates
  }

  public get getLessons(): SubmissionLesson[] {
    return this.lessons
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

  @Mutation
  private setTemplates(templates: SubmissionTemplate[]): void {
    this.templates = templates
  }

  @Mutation
  private setLessons(lessons: SubmissionLesson[]): void {
    this.lessons = lessons
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
  public async listStudentLessons({ lessonId }: { lessonId: number }): Promise<void> {
    await $axios
      .$get(`/v1/submissions/${lessonId}`)
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
        const lessons: SubmissionLesson[] = res.suggestedLessons.map((lesson: v1SubmissionLesson): SubmissionLesson => {
          return { ...lesson }
        })
        this.setShifts({ summary, shifts })
        this.setLessons(lessons)
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async submitStudentShifts({
    shiftId,
    lessons,
    lessonIds,
  }: {
    shiftId: number
    lessons: SubmissionLesson[]
    lessonIds: number[]
  }): Promise<void> {
    const req: SubmissionRequest = {
      suggestedLessons: lessons,
      shiftIds: lessonIds,
    }

    await $axios.$post(`/v1/submissions/${shiftId}`, req).catch((err: AxiosError) => {
      const res: ErrorResponse = { ...err.response?.data }
      throw new ApiError(res.status, res.message, res)
    })
  }

  @Action({ rawError: true })
  public async getSubmissionTemplate(): Promise<void> {
    await $axios
      .$get('/v1/me/submission')
      .then((res: SubmissionTemplateResponse) => {
        const templates: SubmissionTemplate[] = res.schedules.map(
          (schedule: v1SubmissionTemplate): SubmissionTemplate => {
            const lessons: SubmissionTemplateLesson[] = schedule.lessons.map(
              (lesson: v1SubmissionTemplateLesson): SubmissionTemplateLesson => ({ ...lesson })
            )
            return { weekday: schedule.weekday, lessons }
          }
        )
        const lessons: SubmissionLesson[] = res.suggestedLessons.map((lesson: v1SubmissionLesson): SubmissionLesson => {
          return { ...lesson }
        })
        this.setTemplates(templates)
        this.setLessons(lessons)
      })
      .catch((err: AxiosError) => {
        console.log('debug', err)
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async upsertSubmissionTemplate(payload: {
    schedules: SubmissionTemplate[]
    lessons: SubmissionLesson[]
  }): Promise<void> {
    const req: SubmissionTemplateRequest = {
      schedules: payload.schedules.map(
        (schedule: SubmissionTemplate): v1SubmissionTemplate => ({
          weekday: schedule.weekday,
          lessons: schedule.lessons.map(
            (lesson: SubmissionTemplateLesson): v1SubmissionTemplateLesson => ({ ...lesson })
          ),
        })
      ),
      suggestedLessons: payload.lessons.map((lesson: SubmissionLesson): v1SubmissionLesson => ({ ...lesson })),
    }

    await $axios.$post('/v1/me/submission', req).catch((err: AxiosError) => {
      const res: ErrorResponse = { ...err.response?.data }
      throw new ApiError(res.status, res.message, res)
    })
  }
}

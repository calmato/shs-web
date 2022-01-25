import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { AxiosError } from 'axios'
import { $axios } from '~/plugins/axios'
import dayjs from '~/plugins/dayjs'
import {
  CreateShiftsRequest,
  UpdateShiftSummaryScheduleRequest,
  ShiftDetailsResponse,
  ShiftSummariesResponse,
  StudentShiftsResponse,
  TeacherShiftsResponse,
  Lesson as v1Lesson,
  ShiftSummary as V1ShiftSummary,
  ShiftDetail as V1ShiftDetail,
  ShiftDetailLesson as V1ShiftDetailLesson,
  StudentShift as v1StudentShift,
  SuggestedLesson as v1SuggestedLesson,
  TeacherShift as v1TeacherShift,
} from '~/types/api/v1'
import {
  Lesson,
  ShiftDetail,
  ShiftDetailLesson,
  ShiftStatus,
  ShiftState,
  ShiftSummary,
  StudentShift,
  StudentShiftSummary,
  StudentSubmissionDetail,
  SubmissionStatus,
  SuggestedLesson,
  TeacherShift,
  TeacherShiftSummary,
  TeacherSubmissionDetail,
} from '~/types/store'
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
  rooms: 4,
  teachers: [],
  students: [],
  lessons: [],
  teacherSubmission: {
    id: '',
    name: '',
    nameKana: '',
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
    shifts: [],
    submissionTotal: 0,
  },
  studentSubmission: {
    id: '',
    name: '',
    nameKana: '',
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
    shifts: [],
    suggestedLessons: [],
    submissionTotal: 0,
  },
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
  private teachers: ShiftState['teachers'] = initialState.teachers
  private students: ShiftState['students'] = initialState.students
  private rooms: ShiftState['rooms'] = initialState.rooms
  private lessons: ShiftState['lessons'] = initialState.lessons
  private teacherSubmission: ShiftState['teacherSubmission'] = initialState.teacherSubmission
  private studentSubmission: ShiftState['studentSubmission'] = initialState.studentSubmission

  public get getSummary(): ShiftSummary {
    return this.summary
  }

  public get getSummaries(): ShiftSummary[] {
    return this.summaries
  }

  public get getDetails(): ShiftDetail[] {
    return this.details
  }

  public get getRooms(): number {
    return this.rooms
  }

  public get getTeachers(): TeacherShift[] {
    return this.teachers
  }

  public get getStudents(): StudentShift[] {
    return this.students
  }

  public get getLessons(): Lesson[] {
    return this.lessons
  }

  public get getTeacherSubmission(): TeacherSubmissionDetail {
    return this.teacherSubmission
  }

  public get getStudentSubmission(): StudentSubmissionDetail {
    return this.studentSubmission
  }

  @Mutation
  private setSummaries({ summaries }: { summaries: ShiftSummary[] }): void {
    this.summaries = summaries
  }

  @Mutation
  private setDetails({
    summary,
    details,
    teachers,
    students,
    rooms,
    lessons,
  }: {
    summary: ShiftSummary
    details: ShiftDetail[]
    teachers: TeacherShift[]
    students: StudentShift[]
    rooms: number
    lessons: Lesson[]
  }): void {
    this.summary = summary
    this.details = details
    this.teachers = teachers
    this.students = students
    this.rooms = rooms
    this.lessons = lessons
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

  @Mutation
  private setTeacherSubmission(submission: TeacherSubmissionDetail): void {
    const teacher: TeacherShift | undefined = this.teachers.find(
      (val: TeacherShift): boolean => val.id === submission.id
    )
    if (!teacher) {
      return
    }
    submission.name = teacher.name
    submission.nameKana = teacher.nameKana
    this.teacherSubmission = submission
  }

  @Mutation
  private setStudentSubmission(submission: StudentSubmissionDetail): void {
    const student: StudentShift | undefined = this.students.find(
      (val: StudentShift): boolean => val.id === submission.id
    )
    if (!student) {
      return
    }
    submission.name = student.name
    submission.nameKana = student.nameKana
    this.studentSubmission = submission
  }

  @Action({})
  public factory(): void {
    this.setSummaries({ summaries: initialState.summaries })
    this.setDetails({
      summary: initialState.summary,
      details: initialState.details,
      teachers: initialState.teachers,
      students: initialState.students,
      rooms: initialState.rooms,
      lessons: initialState.lessons,
    })
    this.setTeacherSubmission(initialState.teacherSubmission)
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
        const summaries: ShiftSummary[] = res.summaries.map((data: V1ShiftSummary): ShiftSummary => {
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

        const details: ShiftDetail[] = res.shifts.map((shift: V1ShiftDetail): ShiftDetail => {
          const lessons: ShiftDetailLesson[] = shift.lessons.map((lesson: V1ShiftDetailLesson): ShiftDetailLesson => {
            return { ...lesson }
          })
          return { ...shift, lessons }
        })

        this.addSummaries({ summary })
        this.setDetails({
          summary,
          details,
          teachers: initialState.teachers,
          students: initialState.students,
          rooms: initialState.rooms,
          lessons: initialState.lessons,
        })
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
        const rooms: number = res.rooms

        const details: ShiftDetail[] = res.shifts.map((shift: V1ShiftDetail): ShiftDetail => {
          const lessons: ShiftDetailLesson[] = shift.lessons.map((lesson: V1ShiftDetailLesson): ShiftDetailLesson => {
            return { ...lesson }
          })
          return { ...shift, lessons }
        })
        const teachers: TeacherShift[] = res.teachers.map((val: v1TeacherShift): TeacherShift => {
          const name: string = getName(val.teacher.lastName, val.teacher.firstName)
          const nameKana: string = getName(val.teacher.lastNameKana, val.teacher.firstNameKana)
          return { id: val.teacher.id, name, nameKana, lessonTotal: val.lessonTotal }
        })
        const students: StudentShift[] = res.students.map((val: v1StudentShift): StudentShift => {
          const name: string = getName(val.student.lastName, val.student.firstName)
          const nameKana: string = getName(val.student.lastNameKana, val.student.firstNameKana)
          return {
            id: val.student.id,
            name,
            nameKana,
            lessonTotal: val.lessonTotal,
            suggestedLessonsTotal: val.suggestedLessonsTotal,
          }
        })
        const lessons: Lesson[] = res.lessons.map((val: v1Lesson): Lesson => {
          return { ...val }
        })

        this.setDetails({ summary, details, teachers, students, rooms, lessons })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async showTeacherSubmissions({
    summaryId,
    teacherId,
  }: {
    summaryId: number
    teacherId: string
  }): Promise<void> {
    await $axios
      .$get(`/v1/shifts/${summaryId}/teachers/${teacherId}`)
      .then((res: TeacherShiftsResponse) => {
        let submissionTotal: number = 0
        const summary: TeacherShiftSummary = { ...res.summary }
        const shifts: ShiftDetail[] = res.shifts.map((shift: V1ShiftDetail): ShiftDetail => {
          const lessons: ShiftDetailLesson[] = shift.lessons.map((lesson: V1ShiftDetailLesson): ShiftDetailLesson => {
            return { ...lesson }
          })
          submissionTotal += lessons.length
          return { ...shift, lessons }
        })

        this.setTeacherSubmission({ id: teacherId, name: '', nameKana: '', summary, shifts, submissionTotal })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async showStudentSubmissions({
    summaryId,
    studentId,
  }: {
    summaryId: number
    studentId: string
  }): Promise<void> {
    await $axios
      .$get(`/v1/shifts/${summaryId}/students/${studentId}`)
      .then((res: StudentShiftsResponse) => {
        let submissionTotal: number = 0
        const summary: StudentShiftSummary = { ...res.summary }
        const shifts: ShiftDetail[] = res.shifts.map((shift: V1ShiftDetail): ShiftDetail => {
          const lessons: ShiftDetailLesson[] = shift.lessons.map((lesson: V1ShiftDetailLesson): ShiftDetailLesson => {
            return { ...lesson }
          })
          submissionTotal += lessons.length
          return { ...shift, lessons }
        })
        const suggestedLessons: SuggestedLesson[] = res.suggestedLessons.map(
          (lesson: v1SuggestedLesson): SuggestedLesson => {
            return { ...lesson }
          }
        )

        this.setStudentSubmission({
          id: studentId,
          name: '',
          nameKana: '',
          summary,
          shifts,
          suggestedLessons,
          submissionTotal,
        })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }
}

function getName(lastName: string, firstName: string): string {
  return `${lastName} ${firstName}`
}

function replaceDate(date: string, oldVal: string, newVal: string): string {
  return date.replaceAll(oldVal, newVal)
}

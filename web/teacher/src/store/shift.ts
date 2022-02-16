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
  ShiftSummary as v1ShiftSummary,
  ShiftDetail as v1ShiftDetail,
  ShiftDetailLesson as v1ShiftDetailLesson,
  ShiftLesson as v1ShiftLesson,
  Student as v1Student,
  StudentShift as v1StudentShift,
  Subject as v1Subject,
  SuggestedLesson as v1SuggestedLesson,
  Teacher as v1Teacher,
  TeacherShift as v1TeacherShift,
  ShiftLessonResponse,
  ShiftSubmissionsResponse,
  ShiftLessonsResponse,
  CreateShiftLessonRequest,
  UpdateShiftLessonRequest,
  UpdateShiftSummaryDecidedRequest,
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
  ShiftLessonDetail,
  Teacher,
  SubjectsMap,
  Student,
  Subject,
  SchoolType,
  ShiftLesson,
  ShiftUserLesson,
} from '~/types/store'
import { ErrorResponse } from '~/types/api/exception'
import { ApiError } from '~/types/exception'
import { ShiftLessonForm, ShiftsNewForm, ShiftSummaryEditScheduleForm } from '~/types/form'
import { schoolTypeNum2schoolTypeString, subjectResponses2Subjects } from '~/lib'

const initialState: ShiftState = {
  summary: {
    id: 0,
    year: 0,
    month: 0,
    decided: false,
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
  teacherLessons: {
    current: '',
    lessons: [],
    teachers: {},
    students: {},
    total: 0,
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
  studentLessons: {
    current: '',
    lessons: [],
    teachers: {},
    students: {},
    total: 0,
  },
  lessonDetail: {
    lessonId: 0,
    summaryId: 0,
    shiftId: 0,
    room: 0,
    date: '',
    current: undefined,
    teachers: [],
    students: [],
    lessons: [],
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
  private teacherLessons: ShiftState['teacherLessons'] = initialState.teacherLessons
  private studentSubmission: ShiftState['studentSubmission'] = initialState.studentSubmission
  private studentLessons: ShiftState['studentLessons'] = initialState.studentLessons
  private lessonDetail: ShiftState['lessonDetail'] = initialState.lessonDetail

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

  public get getTeacherLessons(): ShiftUserLesson {
    return this.teacherLessons
  }

  public get getStudentSubmission(): StudentSubmissionDetail {
    return this.studentSubmission
  }

  public get getStudentLessons(): ShiftUserLesson {
    return this.studentLessons
  }

  public get getLessonDetail(): ShiftLessonDetail {
    return this.lessonDetail
  }

  @Mutation
  private setSummary(summary: ShiftSummary): void {
    this.summary = summary
  }

  @Mutation
  private updateSummaryDecided(decided: boolean): void {
    this.summary.decided = decided
  }

  @Mutation
  private setSummaries(summaries: ShiftSummary[]): void {
    this.summaries = summaries
  }

  @Mutation
  private addSummaries(summary: ShiftSummary): void {
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
    const index: number = this.summaries.findIndex((val: ShiftSummary): boolean => val.id === summaryId)
    if (index === -1) {
      return
    }
    this.summaries.splice(index, 1)
  }

  @Mutation
  private setDetails(details: ShiftDetail[]): void {
    this.details = details
  }

  @Mutation
  private setTeachers(teachers: TeacherShift[]): void {
    this.teachers = teachers
  }

  @Mutation
  private setStudents(students: StudentShift[]): void {
    this.students = students
  }

  @Mutation
  private setRooms(rooms: number): void {
    this.rooms = rooms
  }

  @Mutation
  private setLessons(lessons: Lesson[]): void {
    this.lessons = lessons
  }

  @Mutation
  private addLesson(lesson: Lesson): void {
    const teacherIndex: number = this.teachers.findIndex((val: TeacherShift): boolean => val.id === lesson.teacherId)
    if (teacherIndex >= 0) {
      const lessonTotal: number = this.teachers[teacherIndex].lessonTotal + 1
      this.teachers.splice(teacherIndex, 1, { ...this.teachers[teacherIndex], lessonTotal })
    }
    const studentIndex: number = this.students.findIndex((val: StudentShift): boolean => val.id === lesson.studentId)
    if (studentIndex >= 0) {
      const lessonTotal: number = this.students[studentIndex].lessonTotal + 1
      this.students.splice(studentIndex, 1, { ...this.students[studentIndex], lessonTotal })
    }
    this.lessons.push(lesson)
  }

  @Mutation
  private replaceLesson(lesson: Lesson): void {
    const index: number = this.lessons.findIndex((val: Lesson): boolean => val.id === lesson.id)
    if (index === -1) {
      return
    }
    this.lessons.splice(index, 1, {
      ...this.lessons[index],
      subjectId: lesson.subjectId,
      teacherId: lesson.teacherId,
      studentId: lesson.studentId,
    })
  }

  @Mutation
  private removeLesson({
    lessonId,
    teacherId,
    studentId,
  }: {
    lessonId: number
    teacherId: string
    studentId: string
  }): void {
    const teacherIndex: number = this.teachers.findIndex((val: TeacherShift): boolean => val.id === teacherId)
    if (teacherIndex >= 0) {
      const lessonTotal: number = this.teachers[teacherIndex].lessonTotal - 1
      this.teachers.splice(teacherIndex, 1, { ...this.teachers[teacherIndex], lessonTotal })
    }
    const studentIndex: number = this.students.findIndex((val: StudentShift): boolean => val.id === studentId)
    if (studentIndex >= 0) {
      const lessonTotal: number = this.students[studentIndex].lessonTotal - 1
      this.students.splice(studentIndex, 1, { ...this.students[studentIndex], lessonTotal })
    }
    const lessonIndex: number = this.lessons.findIndex((val: Lesson): boolean => val.id === lessonId)
    if (lessonIndex >= 0) {
      this.lessons.splice(lessonIndex, 1)
    }
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
  private setTeacherLessons({
    current,
    lessons,
    total,
  }: {
    current: string
    lessons: ShiftLesson[]
    total: number
  }): void {
    const teachers: { [key: string]: TeacherShift } = {}
    const students: { [key: string]: StudentShift } = {}
    lessons.forEach((lesson: ShiftLesson): void => {
      const student: StudentShift | undefined = this.students.find((val: StudentShift) => val.id === lesson.studentId)
      if (student) {
        students[student.id] = student
      }
    })
    const teacher: TeacherShift | undefined = this.teachers.find((val: TeacherShift): boolean => val.id === current)
    if (teacher) {
      teachers[current] = teacher
    }
    this.teacherLessons = { current, lessons, teachers, students, total }
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

  @Mutation
  private setStudentLessons({
    current,
    lessons,
    total,
  }: {
    current: string
    lessons: ShiftLesson[]
    total: number
  }): void {
    const teachers: { [key: string]: TeacherShift } = {}
    const students: { [key: string]: StudentShift } = {}
    lessons.forEach((lesson: ShiftLesson): void => {
      const teacher: TeacherShift | undefined = this.teachers.find((val: TeacherShift) => val.id === lesson.teacherId)
      if (teacher) {
        teachers[teacher.id] = teacher
      }
    })
    const student: StudentShift | undefined = this.students.find((val: StudentShift): boolean => val.id === current)
    if (student) {
      students[current] = student
    }
    this.studentLessons = { current, lessons, teachers, students, total }
  }

  @Mutation
  private setLessonDetail(lessonDetail: ShiftLessonDetail): void {
    let date: string = ''
    for (const detail of this.details) {
      for (const lesson of detail.lessons) {
        if (lesson.id === lessonDetail.shiftId) {
          date = detail.date
          break
        }
      }
      if (date !== '') break
    }
    lessonDetail.date = date
    lessonDetail.teachers = lessonDetail.teachers.map((teacher: Teacher): Teacher => {
      const name = getName(teacher.lastName, teacher.firstName)
      const nameKana = getName(teacher.lastNameKana, teacher.firstNameKana)
      return { ...teacher, name, nameKana }
    })
    lessonDetail.students = lessonDetail.students.map((student: Student): Student => {
      const name = getName(student.lastName, student.firstName)
      const nameKana = getName(student.lastNameKana, student.firstNameKana)
      return { ...student, name, nameKana }
    })
    this.lessonDetail = lessonDetail
  }

  @Action({})
  public factory(): void {
    this.setSummary(initialState.summary)
    this.setSummaries(initialState.summaries)
    this.setDetails(initialState.details)
    this.setRooms(initialState.rooms)
    this.setTeachers(initialState.teachers)
    this.setStudents(initialState.students)
    this.setLessons(initialState.lessons)
    this.setTeacherSubmission(initialState.teacherSubmission)
    this.setTeacherLessons(initialState.teacherLessons)
    this.setStudentSubmission(initialState.studentSubmission)
    this.setStudentLessons(initialState.studentLessons)
    this.setLessonDetail(initialState.lessonDetail)
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
        const summaries: ShiftSummary[] = res.summaries.map((val: v1ShiftSummary): ShiftSummary => ({ ...val }))
        this.setSummaries(summaries)
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
  public async updateShiftSummaryDecided({
    summaryId,
    decided,
  }: {
    summaryId: number
    decided: boolean
  }): Promise<void> {
    const req: UpdateShiftSummaryDecidedRequest = {
      decided,
    }

    await $axios
      .$patch(`/v1/shifts/${summaryId}/decided`, req)
      .then(() => {
        this.updateSummaryDecided(decided)
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
        const { details } = convertShiftDetails(res.shifts)

        this.addSummaries({ ...res.summary })
        this.setSummary({ ...res.summary })
        this.setDetails(details)
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
        const { details } = convertShiftDetails(res.shifts)

        this.setSummary({ ...res.summary })
        this.setDetails(details)
        this.setRooms(res.rooms)
        this.setTeachers(convertTeacherShifts(res.teachers))
        this.setStudents(convertStudentShifts(res.students))
        this.setLessons(res.lessons.map((val: v1Lesson): Lesson => ({ ...val })))
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
        const summary: TeacherShiftSummary = { ...res.summary }
        const { details, total } = convertShiftDetails(res.shifts)

        this.setTeacherSubmission({
          id: teacherId,
          name: '',
          nameKana: '',
          summary,
          shifts: details,
          submissionTotal: total,
        })
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
        const summary: StudentShiftSummary = { ...res.summary }
        const { details, total } = convertShiftDetails(res.shifts)

        this.setStudentSubmission({
          id: studentId,
          name: '',
          nameKana: '',
          summary,
          shifts: details,
          suggestedLessons: convertSuggestedLessons(res.suggestedLessons),
          submissionTotal: total,
        })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async listShiftLessons({
    summaryId,
    shiftId,
    lessonId,
    room,
  }: {
    summaryId: number
    shiftId: number
    lessonId: number
    room: number
  }): Promise<void> {
    await $axios
      .$get(`/v1/shifts/${summaryId}/submissions/${shiftId}`)
      .then((res: ShiftSubmissionsResponse) => {
        let current: ShiftLesson | undefined
        const lessons: ShiftLesson[] = res.lessons.map((lesson: v1ShiftLesson): ShiftLesson => {
          if (lesson.room === room) {
            current = lesson
          }
          return { ...lesson }
        })
        const teachers: Teacher[] = res.teachers.map((teacher: v1Teacher): Teacher => {
          const subjects = convertSubjectsMap(teacher.subjects)
          return { ...teacher, subjects }
        })
        const students: Student[] = res.students.map((student: v1Student): Student => {
          const schoolType: SchoolType = schoolTypeNum2schoolTypeString(student.schoolType)
          const subjects: Subject[] = subjectResponses2Subjects(student.subjects)
          return { ...student, schoolType, subjects }
        })

        this.setLessonDetail({
          lessonId,
          summaryId,
          shiftId,
          room,
          date: '',
          current,
          teachers,
          students,
          lessons,
        })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async listTeacherLessons({ summaryId, teacherId }: { summaryId: number; teacherId: string }): Promise<void> {
    if (teacherId === '') {
      return
    }
    await $axios
      .$get(`/v1/shifts/${summaryId}/lessons?teacherId=${teacherId}`)
      .then((res: ShiftLessonsResponse) => {
        const total: number = res.total
        const lessons: ShiftLesson[] = res.lessons.map((lesson: v1Lesson): ShiftLesson => ({ ...lesson }))

        this.setTeacherLessons({ current: teacherId, lessons, total })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async listStudentLessons({ summaryId, studentId }: { summaryId: number; studentId: string }): Promise<void> {
    if (studentId === '') {
      return
    }
    await $axios
      .$get(`/v1/shifts/${summaryId}/lessons?studentId=${studentId}`)
      .then((res: ShiftLessonsResponse) => {
        const total: number = res.total
        const lessons: ShiftLesson[] = res.lessons.map((lesson: v1Lesson): ShiftLesson => ({ ...lesson }))

        this.setStudentLessons({ current: studentId, lessons, total })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  private async createLesson({ summaryId, form }: { summaryId: number; form: ShiftLessonForm }): Promise<void> {
    const req: CreateShiftLessonRequest = {
      shiftId: form.params.shiftId,
      room: form.params.room,
      subjectId: form.params.subjectId,
      teacherId: form.params.teacherId,
      studentId: form.params.studentId,
    }

    await $axios
      .$post(`/v1/shifts/${summaryId}/lessons`, req)
      .then((res: ShiftLessonResponse) => {
        this.addLesson({ ...res })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  private async updateLesson({ summaryId, form }: { summaryId: number; form: ShiftLessonForm }): Promise<void> {
    const req: UpdateShiftLessonRequest = {
      shiftId: form.params.shiftId,
      room: form.params.room,
      subjectId: form.params.subjectId,
      teacherId: form.params.teacherId,
      studentId: form.params.studentId,
    }

    await $axios
      .$patch(`/v1/shifts/${summaryId}/lessons/${form.params.lessonId}`, req)
      .then((res: ShiftLessonResponse) => {
        this.replaceLesson({ ...res })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async upsertLesson({ summaryId, form }: { summaryId: number; form: ShiftLessonForm }): Promise<void> {
    if (form.params.lessonId === 0) {
      await this.createLesson({ summaryId, form })
    } else {
      await this.updateLesson({ summaryId, form })
    }
  }

  @Action({ rawError: true })
  public async deleteLesson({ summaryId, form }: { summaryId: number; form: ShiftLessonForm }): Promise<void> {
    await $axios
      .$delete(`/v1/shifts/${summaryId}/lessons/${form.params.lessonId}`)
      .then(() => {
        this.removeLesson({ ...form.params })
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

function convertSubjectsMap(subjects: { [key: number]: v1Subject[] }): SubjectsMap {
  return subjects
    ? {
        小学校: subjectResponses2Subjects(subjects[1]),
        中学校: subjectResponses2Subjects(subjects[2]),
        高校: subjectResponses2Subjects(subjects[3]),
        その他: [],
      }
    : {
        小学校: [],
        中学校: [],
        高校: [],
        その他: [],
      }
}

function convertTeacherShifts(teachers: v1TeacherShift[]): TeacherShift[] {
  return teachers.map(
    (val: v1TeacherShift): TeacherShift => ({
      id: val.teacher.id,
      name: getName(val.teacher.lastName, val.teacher.firstName),
      nameKana: getName(val.teacher.lastNameKana, val.teacher.firstNameKana),
      isSubmit: val.isSubmit,
      lessonTotal: val.lessonTotal,
    })
  )
}

function convertStudentShifts(students: v1StudentShift[]): StudentShift[] {
  return students.map(
    (val: v1StudentShift): StudentShift => ({
      id: val.student.id,
      name: getName(val.student.lastName, val.student.firstName),
      nameKana: getName(val.student.lastNameKana, val.student.firstNameKana),
      isSubmit: val.isSubmit,
      suggestedLessons: convertSuggestedLessons(val.suggestedLessons),
      suggestedLessonsTotal: val.suggestedLessonsTotal,
      lessonTotal: val.lessonTotal,
    })
  )
}

function convertSuggestedLessons(lessons: v1SuggestedLesson[]): SuggestedLesson[] {
  return lessons.map((lesson: v1SuggestedLesson): SuggestedLesson => ({ ...lesson }))
}

function convertShiftDetails(details: v1ShiftDetail[]): { details: ShiftDetail[]; total: number } {
  let submissionTotal = 0
  const res: ShiftDetail[] = details.map((detail: v1ShiftDetail): ShiftDetail => {
    const lessons: ShiftDetailLesson[] = detail.lessons.map(
      (lesson: v1ShiftDetailLesson): ShiftDetailLesson => ({ ...lesson })
    )
    submissionTotal += lessons.length
    return { ...detail, lessons }
  })
  return { details: res, total: submissionTotal }
}

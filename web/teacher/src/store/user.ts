import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { AxiosError } from 'axios'
import { $axios } from '~/plugins/axios'
import {
  TeachersResponse,
  Teacher as v1Teacher,
  CreateTeacherRequest,
  TeacherResponse,
  UpdateTeacherSubjectsRequest,
  UpdateTeacherRoleRequest,
  UpdateTeacherMailRequest,
  UpdateTeacherPasswordRequest,
  CreateStudentRequest,
  StudentResponse,
  Student as v1Student,
  StudentsResponse,
  UpdateStudentSubjectsRequest,
} from '~/types/api/v1'
import { Role, Student, StudentMap, SubjectsMap, Teacher, TeacherMap, UserState } from '~/types/store'
import { ErrorResponse } from '~/types/api/exception'
import { ApiError } from '~/types/exception'
import {
  StudentEditSubjectForm,
  StudentNewForm,
  TeacherEditRoleForm,
  TeacherEditSubjectForm,
  TeacherNewForm,
  TeacherUpdateMailForm,
  TeacherUpdatePasswordForm,
} from '~/types/form'
import {
  schoolTypeNum2schoolTypeString,
  schoolTypeString2schoolTypeNum,
  subjectResponse2Subject,
  subjectResponses2Subjects,
} from '~/lib'

const initialState: UserState = {
  student: {
    id: '',
    lastName: '',
    firstName: '',
    lastNameKana: '',
    firstNameKana: '',
    mail: '',
    schoolType: 'その他',
    grade: 1,
    subjects: [],
    createdAt: '',
    updatedAt: '',
  },
  students: [],
  studentsTotal: 0,
  teacher: {
    id: '',
    lastName: '',
    firstName: '',
    lastNameKana: '',
    firstNameKana: '',
    mail: '',
    role: Role.TEACHER,
    subjects: initializeSubjects(),
    createdAt: '',
    updatedAt: '',
  },
  teachers: [],
  teachersTotal: 0,
}

@Module({
  name: 'user',
  stateFactory: true,
  namespaced: true,
})
export default class UserModule extends VuexModule {
  private student: UserState['student'] = initialState.student
  private students: UserState['students'] = initialState.students
  private studentsTotal: UserState['studentsTotal'] = initialState.studentsTotal
  private teacher: UserState['teacher'] = initialState.teacher
  private teachers: UserState['teachers'] = initialState.teachers
  private teachersTotal: UserState['teachersTotal'] = initialState.teachersTotal

  public get getStudent(): Student {
    return this.student
  }

  public get getStudents(): Student[] {
    return this.students
  }

  public get getStudentMap(): StudentMap {
    const students: StudentMap = {}
    this.students.forEach((student: Student) => {
      students[student.id] = student
    })
    return students
  }

  public get getStudentsTotal(): number {
    return this.studentsTotal
  }

  public get getTeacher(): Teacher {
    return this.teacher
  }

  public get getTeachers(): Teacher[] {
    return this.teachers
  }

  public get getTeacherMap(): TeacherMap {
    const teachers: TeacherMap = {}
    this.teachers.forEach((teacher: Teacher) => {
      teachers[teacher.id] = teacher
    })
    return teachers
  }

  public get getTeachersTotal(): number {
    return this.teachersTotal
  }

  /**
   * 生徒関連
   */
  @Mutation
  private setStudent(student: Student): void {
    const name = getName(student.lastName, student.firstName)
    const nameKana = getName(student.lastNameKana, student.firstNameKana)
    this.student = { ...student, name, nameKana }
  }

  @Mutation
  private setStudents({ students, total }: { students: Student[]; total: number }): void {
    this.students = students.map((student: Student): Student => {
      const name = getName(student.lastName, student.firstName)
      const nameKana = getName(student.lastNameKana, student.firstNameKana)
      return { ...student, name, nameKana }
    })
    this.studentsTotal = total
  }

  @Mutation
  private addStudent(student: Student): void {
    const name = getName(student.lastName, student.firstName)
    const nameKana = getName(student.lastNameKana, student.firstNameKana)
    this.students.push({ ...student, name, nameKana })
    this.studentsTotal += 1
  }

  @Mutation
  private removeStudent({ studentId }: { studentId: string }): void {
    const index: number = this.students.findIndex((val: Student) => {
      return val.id === studentId
    })
    if (index === -1) {
      return
    }
    this.students.splice(index, 1)
    this.studentsTotal -= 1
  }

  /**
   * 講師関連
   */
  @Mutation
  private setTeacher(teacher: Teacher): void {
    const name = getName(teacher.lastName, teacher.firstName)
    const nameKana = getName(teacher.lastNameKana, teacher.firstNameKana)
    this.teacher = { ...teacher, name, nameKana }
  }

  @Mutation
  private setTeachers({ teachers, total }: { teachers: Teacher[]; total: number }): void {
    this.teachers = teachers.map((teacher: Teacher): Teacher => {
      const name = getName(teacher.lastName, teacher.firstName)
      const nameKana = getName(teacher.lastNameKana, teacher.firstNameKana)
      return { ...teacher, name, nameKana }
    })
    this.teachersTotal = total
  }

  @Mutation
  private addTeacher(teacher: Teacher): void {
    const name = getName(teacher.lastName, teacher.firstName)
    const nameKana = getName(teacher.lastNameKana, teacher.firstNameKana)
    this.teachers.push({ ...teacher, name, nameKana })
    this.teachersTotal += 1
  }

  @Mutation
  private removeTeacher({ teacherId }: { teacherId: string }): void {
    const index: number = this.teachers.findIndex((val: Teacher) => {
      return val.id === teacherId
    })
    if (index === -1) {
      return
    }
    this.teachers.splice(index, 1)
    this.teachersTotal -= 1
  }

  @Action({})
  public factory(): void {
    this.setStudents({ students: initialState.students, total: initialState.studentsTotal })
    this.setTeachers({ teachers: initialState.teachers, total: initialState.teachersTotal })
  }

  /**
   * 生徒関連
   */
  @Action({ rawError: true })
  public async listStudents({ limit, offset }: { limit: number; offset: number }): Promise<void> {
    let query: string = ''
    if (limit !== 0 || offset !== 0) {
      query = `?limit=${limit}&offset=${offset}`
    }

    await $axios
      .$get('/v1/students' + query)
      .then((res: StudentsResponse) => {
        const students: Student[] = res.students.map((student: v1Student) => {
          const subjects = subjectResponses2Subjects(student.subjects)
          const schoolType = schoolTypeNum2schoolTypeString(student.schoolType)
          return { ...student, subjects, schoolType }
        })
        this.setStudents({ students, total: res.total })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async showStudent({ studentId }: { studentId: string }): Promise<void> {
    await $axios
      .$get(`/v1/students/${studentId}`)
      .then((res: StudentResponse) => {
        const subjects = subjectResponses2Subjects(res.subjects)
        const schoolType = schoolTypeNum2schoolTypeString(res.schoolType)
        this.setStudent({ ...res, subjects, schoolType })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async createStudent({ form }: { form: StudentNewForm }): Promise<void> {
    const req: CreateStudentRequest = {
      ...form.params,
      schoolType: schoolTypeString2schoolTypeNum(form.params.schoolType),
      grade: Number(form.params.grade),
    }

    await $axios
      .$post('/v1/students', req)
      .then((res: StudentResponse) => {
        const student: Student = {
          ...res,
          schoolType: schoolTypeNum2schoolTypeString(res.schoolType),
          subjects: subjectResponses2Subjects(res.subjects),
        }
        this.addStudent(student)
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async updateStudentSubjects({
    studentId,
    form,
  }: {
    studentId: string
    form: StudentEditSubjectForm
  }): Promise<void> {
    const req: UpdateStudentSubjectsRequest = {
      schoolType: form.params.schoolType,
      subjectIds: form.params.subjectIds,
    }

    await $axios.$patch(`/v1/students/${studentId}/subjects`, req).catch((err: AxiosError) => {
      const res: ErrorResponse = { ...err.response?.data }
      throw new ApiError(res.status, res.message, res)
    })
  }

  @Action({ rawError: true })
  public async deleteStudent({ studentId }: { studentId: string }): Promise<void> {
    await $axios
      .$delete(`/v1/students/${studentId}`)
      .then(() => {
        this.removeStudent({ studentId })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  /**
   * 講師関連
   */
  @Action({ rawError: true })
  public async listTeachers({ limit, offset }: { limit: number; offset: number }): Promise<void> {
    let query: string = ''
    if (limit !== 0 || offset !== 0) {
      query = `?limit=${limit}&offset=${offset}`
    }

    await $axios
      .$get('/v1/teachers' + query)
      .then((res: TeachersResponse) => {
        const teachers: Teacher[] = res.teachers.map((data: v1Teacher): Teacher => {
          const subjects = initializeSubjects()
          return { ...data, subjects }
        })
        this.setTeachers({ teachers, total: res.total })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async showTeacher({ teacherId }: { teacherId: string }): Promise<void> {
    await $axios
      .$get(`/v1/teachers/${teacherId}`)
      .then((res: TeacherResponse) => {
        const subjects: SubjectsMap = {
          小学校: res.subjects[1].map((i) => subjectResponse2Subject(i)),
          中学校: res.subjects[2].map((i) => subjectResponse2Subject(i)),
          高校: res.subjects[3].map((i) => subjectResponse2Subject(i)),
          その他: [],
        }
        this.setTeacher({ ...res, subjects })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async createTeacher({ form }: { form: TeacherNewForm }): Promise<void> {
    const req: CreateTeacherRequest = { ...form.params }

    await $axios
      .$post('/v1/teachers', req)
      .then((res: TeacherResponse) => {
        const subjects = res.subjects
          ? {
              小学校: res.subjects[1].map((i) => subjectResponse2Subject(i)),
              中学校: res.subjects[2].map((i) => subjectResponse2Subject(i)),
              高校: res.subjects[3].map((i) => subjectResponse2Subject(i)),
              その他: [],
            }
          : initializeSubjects()
        this.addTeacher({ ...res, subjects })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async updateTeacherSubjects({
    teacherId,
    form,
  }: {
    teacherId: string
    form: TeacherEditSubjectForm
  }): Promise<void> {
    const req: UpdateTeacherSubjectsRequest = {
      schoolType: form.params.schoolType,
      subjectIds: form.params.subjectIds,
    }

    await $axios.$patch(`/v1/teachers/${teacherId}/subjects`, req).catch((err: AxiosError) => {
      const res: ErrorResponse = { ...err.response?.data }
      throw new ApiError(res.status, res.message, res)
    })
  }

  @Action({ rawError: true })
  public async updateTeacherRole({ teacherId, form }: { teacherId: string; form: TeacherEditRoleForm }): Promise<void> {
    const req: UpdateTeacherRoleRequest = {
      role: form.params.role,
    }

    await $axios.$patch(`/v1/teachers/${teacherId}/role`, req).catch((err: AxiosError) => {
      const res: ErrorResponse = { ...err.response?.data }
      throw new ApiError(res.status, res.message, res)
    })
  }

  @Action({ rawError: true })
  public async deleteTeacher({ teacherId }: { teacherId: string }): Promise<void> {
    await $axios
      .$delete(`/v1/teachers/${teacherId}`)
      .then(() => {
        this.removeTeacher({ teacherId })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  // TODO: AuthStoreへ移行
  @Action({ rawError: true })
  public async updateMail({ form }: { form: TeacherUpdateMailForm }): Promise<void> {
    const req: UpdateTeacherMailRequest = {
      mail: form.params.mail,
    }

    await $axios.$patch(`/v1/me/mail`, req).catch((err: AxiosError) => {
      const res: ErrorResponse = { ...err.response?.data }
      throw new ApiError(res.status, res.message, res)
    })
  }

  // TODO: AuthStoreへ移行
  @Action({ rawError: true })
  public async updatePassword({ form }: { form: TeacherUpdatePasswordForm }): Promise<void> {
    const req: UpdateTeacherPasswordRequest = {
      password: form.params.password,
      passwordConfirmation: form.params.passwordConfirmaion,
    }

    await $axios.$patch(`/v1/me/password`, req).catch((err: AxiosError) => {
      const res: ErrorResponse = { ...err.response?.data }
      throw new ApiError(res.status, res.message, res)
    })
  }
}

function initializeSubjects(): SubjectsMap {
  return {
    小学校: [],
    中学校: [],
    高校: [],
    その他: [],
  }
}

function getName(lastName: string, firstName: string): string {
  return `${lastName} ${firstName}`
}

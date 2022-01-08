import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { AxiosError } from 'axios'
import { $axios } from '~/plugins/axios'
import {
  TeachersResponse,
  Teacher as V1Teacher,
  CreateTeacherRequest,
  TeacherResponse,
  UpdateTeacherSubjectsRequest,
  UpdateTeacherRoleRequest,
} from '~/types/api/v1'
import { Role, SchoolType, Student, StudentMap, SubjectsMap, Teacher, TeacherMap, UserState } from '~/types/store'
import { ErrorResponse } from '~/types/api/exception'
import { ApiError } from '~/types/exception'
import { TeacherEditRoleForm, TeacherEditSubjectForm, TeacherNewForm } from '~/types/form'

const initialState: UserState = {
  students: [
    {
      id: '123456789012345678901',
      name: '浜田 二郎',
      nameKana: 'はまだ じろう',
      lastName: '浜田',
      firstName: '二郎',
      lastNameKana: 'はまだ',
      firstNameKana: 'じろう',
      mail: 'student-001@calmato.jp',
      type: 1,
      grade: 2,
      createdAt: '',
      updatedAt: '',
    },
  ],
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
  private students: UserState['students'] = initialState.students
  private teacher: UserState['teacher'] = initialState.teacher
  private teachers: UserState['teachers'] = initialState.teachers
  private teachersTotal: UserState['teachersTotal'] = initialState.teachersTotal

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

  @Mutation
  private setStudents(students: Student[]): void {
    this.students = students.map((student: Student): Student => {
      const name = getName(student.lastName, student.firstName)
      const nameKana = getName(student.lastNameKana, student.firstNameKana)
      return { ...student, name, nameKana }
    })
  }

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
    this.setStudents(initialState.students)
    this.setTeachers({ teachers: initialState.teachers, total: initialState.teachersTotal })
  }

  @Action({ rawError: true })
  public async listTeachers({ limit, offset }: { limit: number; offset: number }): Promise<void> {
    let query: string = ''
    if (limit !== 0 || offset !== 0) {
      query = `?limit=${limit}&offset=${offset}`
    }

    await $axios
      .$get('/v1/teachers' + query)
      .then((res: TeachersResponse) => {
        const teachers: Teacher[] = res.teachers.map((data: V1Teacher): Teacher => {
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
        this.setTeacher({ ...res })
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
        const subjects = res.subjects || initializeSubjects()
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
}

function initializeSubjects(): SubjectsMap {
  return {
    [SchoolType.ELEMENTARY_SCHOOL]: [],
    [SchoolType.JUNIOR_HIGH_SCHOOL]: [],
    [SchoolType.HIGH_SCHOOL]: [],
  }
}

function getName(lastName: string, firstName: string): string {
  return `${lastName} ${firstName}`
}

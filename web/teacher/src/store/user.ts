import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { AxiosError } from 'axios'
import { $axios } from '~/plugins/axios'
import { TeachersResponse, Teacher as TeacherResponse, CreateTeacherRequest } from '~/types/api/v1'
import { Student, StudentMap, Teacher, TeacherMap, UserState } from '~/types/store'
import { ErrorResponse } from '~/types/api/exception'
import { ApiError } from '~/types/exception'
import { TeacherNewForm } from '~/types/form'

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
      const name: string = `${student.lastName} ${student.firstName}`
      const nameKana: string = `${student.lastNameKana} ${student.firstNameKana}`
      return { ...student, name, nameKana }
    })
  }

  @Mutation
  private setTeachers({ teachers, total }: { teachers: Teacher[]; total: number }): void {
    this.teachers = teachers.map((teacher: Teacher): Teacher => {
      const name: string = `${teacher.lastName} ${teacher.firstName}`
      const nameKana: string = `${teacher.lastNameKana} ${teacher.firstNameKana}`
      return { ...teacher, name, nameKana }
    })
    this.teachersTotal = total
  }

  @Mutation
  private addTeacher(teacher: Teacher): void {
    const name: string = `${teacher.lastName} ${teacher.firstName}`
    const nameKana: string = `${teacher.lastNameKana} ${teacher.firstNameKana}`
    this.teachers.push({ ...teacher, name, nameKana })
    this.teachersTotal += 1
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
        const teachers: Teacher[] = res.teachers.map((data: TeacherResponse): Teacher => {
          return { ...data }
        })
        this.setTeachers({ teachers, total: res.total })
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
        this.addTeacher(res)
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }
}

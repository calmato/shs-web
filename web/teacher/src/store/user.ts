import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { AxiosError } from 'axios'
import { $axios } from '~/plugins/axios'
import { TeachersResponse, Teacher as TeacherResponse } from '~/types/api/v1'
import { Student, StudentMap, Teacher, TeacherMap, UserState } from '~/types/store'
import { ErrorResponse } from '~/types/api/exception'
import { ApiError } from '~/types/exception'

const initialState: UserState = {
  teachers: [],
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
}

@Module({
  name: 'user',
  stateFactory: true,
  namespaced: true,
})
export default class UserModule extends VuexModule {
  private students: UserState['students'] = initialState.students
  private teachers: UserState['teachers'] = initialState.teachers

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

  @Mutation
  private setStudents(students: Student[]): void {
    this.students = students.map((student: Student): Student => {
      const name: string = `${student.lastName} ${student.firstName}`
      const nameKana: string = `${student.lastNameKana} ${student.firstNameKana}`
      return { ...student, name, nameKana }
    })
  }

  @Mutation
  private setTeachers(teachers: Teacher[]): void {
    this.teachers = teachers.map((teacher: Teacher): Teacher => {
      const name: string = `${teacher.lastName} ${teacher.firstName}`
      const nameKana: string = `${teacher.lastNameKana} ${teacher.firstNameKana}`
      return { ...teacher, name, nameKana }
    })
  }

  @Action({})
  public factory(): void {
    this.setStudents(initialState.students)
    this.setTeachers(initialState.teachers)
  }

  @Action({ rawError: true })
  public async listTeachers(): Promise<void> {
    // TODO: limit, offset周りの対応
    await $axios
      .$get('/v1/teachers')
      .then((res: TeachersResponse) => {
        const teachers: Teacher[] = res.teachers.map((data: TeacherResponse): Teacher => {
          return { ...data }
        })
        this.setTeachers(teachers)
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }
}

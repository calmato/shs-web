import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { $axios } from '~/plugins/axios'
import { ErrorResponse } from '~/types/api/exception'
import { HelloRequest, HelloResponse } from '~/types/api/v1'
import { ApiError } from '~/types/exception'
import { Student, StudentMap, Teacher, TeacherMap, UserState } from '~/types/store'

const initialState: UserState = {
  message: '',
  teachers: [
    {
      id: '000000000000000000001',
      lastname: '中村',
      firstname: '太郎',
      createdAt: '',
      updatedAt: '',
    },
  ],
  students: [
    {
      id: '123456789012345678901',
      lastname: '浜田',
      firstname: '二郎',
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
  private message: UserState['message'] = initialState.message
  private students: UserState['students'] = initialState.students
  private teachers: UserState['teachers'] = initialState.teachers

  public get getMessage(): string {
    return this.message
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
  private setMessage(message: string): void {
    this.message = message
  }

  @Mutation
  private setStudents(students: Student[]): void {
    this.students = students
  }

  @Mutation
  private setTeachers(teachers: Teacher[]): void {
    this.teachers = teachers
  }

  @Action({})
  public factory(): void {
    this.setMessage(initialState.message)
    this.setStudents(initialState.students)
    this.setTeachers(initialState.teachers)
  }

  @Action({ rawError: true })
  public async hello(): Promise<void> {
    const req: HelloRequest = { name: 'test' }

    await $axios
      .$post('/v1/hello', req)
      .then((res: HelloResponse) => {
        this.setMessage(res.message)
      })
      .catch((err: any) => {
        const res: ErrorResponse = err.response
        throw new ApiError(res.status, res.message, res)
      })
  }
}

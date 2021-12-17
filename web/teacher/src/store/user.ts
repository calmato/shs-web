import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { Student, StudentMap, Teacher, TeacherMap, UserState } from '~/types/store'

const initialState: UserState = {
  teachers: [
    {
      id: '000000000000000000001',
      name: '中村 太郎',
      nameKana: 'なかむら たろう',
      lastName: '中村',
      firstName: '太郎',
      lastNameKana: 'なかむら',
      firstNameKana: 'たろう',
      mail: 'teacher-001@calmato.jp',
      role: 0,
      createdAt: '',
      updatedAt: '',
    },
    {
      id: '000000000000000000002',
      name: '西山 幸子',
      nameKana: 'にしやま さちこ',
      lastName: '西山',
      firstName: '幸子',
      lastNameKana: 'にしやま',
      firstNameKana: 'さちこ',
      mail: 'teacher-002@calmato.jp',
      role: 0,
      createdAt: '',
      updatedAt: '',
    },
    {
      id: '000000000000000000003',
      name: '鈴木 小太郎',
      nameKana: 'すずき こたろう',
      lastName: '鈴木',
      firstName: '小太郎',
      lastNameKana: 'すずき',
      firstNameKana: 'こたろう',
      mail: 'teacher-003@calmato.jp',
      role: 1,
      createdAt: '',
      updatedAt: '',
    },
  ],
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
}

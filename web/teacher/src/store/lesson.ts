import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import axios, { AxiosError } from 'axios'
import { schoolTypeNum2schoolTypeString, schoolTypeString2schoolTypeNum, subjectResponse2Subject } from '~/lib'
import { $axios } from '~/plugins/axios'
import { ErrorResponse } from '~/types/api/exception'
import {
  SubjectsResponse,
  Subject as v1Subject,
  LessonsResponse,
  Lesson as v1Lesson,
  Teacher as v1Teacher,
  Student as v1Student,
} from '~/types/api/v1'
import { ApiError } from '~/types/exception'
import { SubjectEditForm, SubjectNewForm } from '~/types/form'
import {
  Lesson,
  LessonState,
  Student,
  StudentMap,
  Subject,
  SubjectMap,
  SubjectsMap,
  Teacher,
  TeacherMap,
} from '~/types/store'

const initialState: LessonState = {
  subjects: [],
  lessons: [],
  teachers: [],
  students: [],
}

@Module({
  name: 'lesson',
  stateFactory: true,
  namespaced: true,
})
export default class LessonModule extends VuexModule {
  private subjects: LessonState['subjects'] = initialState.subjects
  private lessons: LessonState['lessons'] = initialState.lessons
  private teachers: LessonState['teachers'] = initialState.teachers
  private students: LessonState['students'] = initialState.students

  public get getSubjects(): Subject[] {
    return this.subjects
  }

  public get getSubjectMap(): SubjectMap {
    const subjects: SubjectMap = {}
    this.subjects.forEach((subject: Subject) => {
      subjects[subject.id] = subject
    })
    return subjects
  }

  public get getSubjectsMap(): SubjectsMap {
    const subjects: SubjectsMap = {
      小学校: [],
      中学校: [],
      高校: [],
      その他: [],
    }
    this.subjects.forEach((subject: Subject) => {
      subjects[subject.schoolType].push(subject)
    })
    return subjects
  }

  public get getLessons(): Lesson[] {
    return this.lessons
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

  @Mutation
  private setSubjects(subjects: Subject[]): void {
    this.subjects = subjects
  }

  @Mutation
  private setLessons(lessons: Lesson[]): void {
    this.lessons = lessons
  }

  @Mutation
  private setUsers({ teachers, students }: { teachers: Teacher[]; students: Student[] }): void {
    this.teachers = teachers.map((teacher: Teacher): Teacher => {
      const name = getName(teacher.lastName, teacher.firstName)
      const nameKana = getName(teacher.lastNameKana, teacher.firstNameKana)
      return { ...teacher, name, nameKana }
    })
    this.students = students.map((student: Student): Student => {
      const name = getName(student.lastName, student.firstName)
      const nameKana = getName(student.lastNameKana, student.firstNameKana)
      return { ...student, name, nameKana }
    })
  }

  @Action({})
  public factory(): void {
    this.setSubjects(initialState.subjects)
    this.setLessons(initialState.lessons)
    this.setUsers({ teachers: initialState.teachers, students: initialState.students })
  }

  @Action({ rawError: true })
  public async listLessons({ since, until }: { since: string; until: string }): Promise<void> {
    let query: string = ''
    if (since !== '' || until !== '') {
      query = `?since=${since}&until=${until}`
    }

    await $axios
      .$get(`/v1/lessons${query}`)
      .then((res: LessonsResponse) => {
        const lessons: Lesson[] = res.lessons.map((lesson: v1Lesson): Lesson => {
          return { ...lesson }
        })
        const teachers: Teacher[] = res.teachers.map((teacher: v1Teacher): Teacher => {
          const subjects = initializeSubjects()
          return { ...teacher, subjects }
        })
        const students: Student[] = res.students.map((student: v1Student): Student => {
          const schoolType = schoolTypeNum2schoolTypeString(student.schoolType)
          const subjects: Subject[] = []
          return { ...student, schoolType, subjects }
        })
        this.setLessons(lessons)
        this.setUsers({ teachers, students })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async getAllSubjects(): Promise<void> {
    try {
      const res: SubjectsResponse = await $axios.$get('/v1/subjects')
      const subjects: Subject[] = res.subjects?.map((subject: v1Subject): Subject => {
        return subjectResponse2Subject(subject)
      })
      this.setSubjects(subjects)
      return
    } catch (err) {
      if (axios.isAxiosError(err)) {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      }
      throw new Error('internal server error')
    }
  }

  @Action({ rawError: true })
  public async createSubject(payload: SubjectNewForm): Promise<void> {
    try {
      const request = { ...payload, schoolType: schoolTypeString2schoolTypeNum(payload.schoolType) }
      await $axios.$post('/v1/subjects', request)
      await this.getAllSubjects()
    } catch (err) {
      if (axios.isAxiosError(err)) {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      }
      throw new Error('internal server error')
    }
  }

  @Action({ rawError: true })
  public async editSubject(payload: SubjectEditForm): Promise<void> {
    try {
      const request = { ...payload, schoolType: schoolTypeString2schoolTypeNum(payload.schoolType) }
      await $axios.$patch(`/v1/subjects/${payload.subjectId}`, request)
      await this.getAllSubjects()
    } catch (err) {
      if (axios.isAxiosError(err)) {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      }
      throw new Error('internal server error')
    }
  }

  @Action({ rawError: true })
  public async deleteSubject(subjectId: number): Promise<void> {
    try {
      await $axios.$delete(`/v1/subjects/${subjectId}`)
      await this.getAllSubjects()
    } catch (err) {
      if (axios.isAxiosError(err)) {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      }
      throw new Error('internal server error')
    }
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

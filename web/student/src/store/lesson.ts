import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import axios, { AxiosError } from 'axios'
import { $axios } from '~/plugins/axios'
import { ErrorResponse } from '~/types/api/exception'
import { Lesson, LessonState, Subject, SubjectMap, SubjectsMap, Teacher, TeacherMap } from '~/types/store'
import {
  Lesson as v1Lesson,
  LessonsResponse,
  SubjectsResponse,
  Subject as v1Subject,
  Teacher as v1Teacher,
} from '~/types/api/v1'
import { ApiError } from '~/types/exception'
import { subjectResponse2Subject } from '~/lib'

const initialState: LessonState = {
  subjects: [],
  lessons: [],
  teachers: [],
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

  @Mutation
  private setSubjects(subjects: Subject[]): void {
    this.subjects = subjects
  }

  @Mutation
  private setLessons(lessons: Lesson[]): void {
    this.lessons = lessons
  }

  @Mutation
  private setTeachers(teachers: Teacher[]): void {
    this.teachers = teachers.map((teacher: Teacher): Teacher => {
      const name = getName(teacher.lastName, teacher.firstName)
      const nameKana = getName(teacher.lastNameKana, teacher.firstNameKana)
      return { ...teacher, name, nameKana }
    })
  }

  @Action({ rawError: true })
  private factory(): void {
    this.lessons = initialState.lessons
    this.teachers = initialState.teachers
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
          return { ...teacher }
        })
        this.setLessons(lessons)
        this.setTeachers(teachers)
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

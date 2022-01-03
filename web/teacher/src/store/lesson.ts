import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { $axios } from '~/plugins/axios'
import { ErrorResponse } from '~/types/api/exception'
import { ApiError } from '~/types/exception'
import { Lesson, LessonState, SchoolType, Subject, SubjectMap } from '~/types/store'

const initialState: LessonState = {
  subjects: [
    {
      id: 1,
      name: '国語',
      color: '#F8BBD0',
      schoolType: SchoolType.ELEMENTARY_SCHOOL,
      createdAt: '',
      updatedAt: '',
    },
    {
      id: 2,
      name: '数学',
      color: '#BBDEFB',
      schoolType: SchoolType.JUNIOR_HIGH_SCHOOL,
      createdAt: '',
      updatedAt: '',
    },
    {
      id: 3,
      name: '英語',
      color: '#FEE6C9',
      schoolType: SchoolType.HIGH_SCHOOL,
      createdAt: '',
      updatedAt: '',
    },
  ],
  lessons: [
    {
      id: 1,
      teacherId: '000000000000000000001',
      studentId: '123456789012345678901',
      subjectId: 1,
      startAt: '2021-12-10T18:30:00+09:00',
      endAt: '2021-12-10T20:00:00+09:00',
      createdAt: '',
      updatedAt: '',
    },
  ],
}

@Module({
  name: 'lesson',
  stateFactory: true,
  namespaced: true,
})
export default class LessonModule extends VuexModule {
  private subjects: LessonState['subjects'] = initialState.subjects
  private lessons: LessonState['lessons'] = initialState.lessons

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

  public get getLessons(): Lesson[] {
    return this.lessons
  }

  @Mutation
  private setSubjects(subjects: Subject[]): void {
    this.subjects = subjects
  }

  @Mutation
  private setLessons(lessons: Lesson[]): void {
    this.lessons = lessons
  }

  @Action({})
  public factory(): void {
    this.setSubjects(initialState.subjects)
    this.setLessons(initialState.lessons)
  }

  @Action({ rawError: true })
  public async getAllSubjects(): Promise<void> {
    try {
      const res: { subjects: Subject[] } = await $axios.$get('/v1/subjects')
      this.setSubjects(res.subjects)
      return
    } catch (err) {
      if ($axios.isAxiosError(err)) {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      }
      throw new Error('internal server error')
    }
  }
}

import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { subjectResponse2Subject } from '~/lib'
import { $axios } from '~/plugins/axios'
import { ErrorResponse } from '~/types/api/exception'
import { SubjectsResponse, Subject as v1Subject } from '~/types/api/v1'
import { ApiError } from '~/types/exception'
import { Lesson, LessonState, Subject, SubjectMap } from '~/types/store'

const initialState: LessonState = {
  subjects: [
    {
      id: 1,
      name: '国語',
      color: '#F8BBD0',
      schoolType: '小学校',
      createdAt: '',
      updatedAt: '',
    },
    {
      id: 2,
      name: '数学',
      color: '#BBDEFB',
      schoolType: '中学校',
      createdAt: '',
      updatedAt: '',
    },
    {
      id: 3,
      name: '英語',
      color: '#FEE6C9',
      schoolType: '高校',
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
      const res: SubjectsResponse = await $axios.$get('/v1/subjects')
      const subjects: Subject[] = res.subjects?.map((subject: v1Subject): Subject => {
        return subjectResponse2Subject(subject)
      })
      this.setSubjects(subjects)
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

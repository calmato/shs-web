import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import Axios from 'axios'
import { schoolTypeString2schoolTypeNum, subjectResponse2Subject } from '~/lib'
import { $axios } from '~/plugins/axios'
import { ErrorResponse } from '~/types/api/exception'
import { SubjectsResponse, Subject as v1Subject } from '~/types/api/v1'
import { ApiError } from '~/types/exception'
import { SubjectNewForm } from '~/types/form'
import { Lesson, LessonState, Subject, SubjectMap, SubjectsMap } from '~/types/store'

const initialState: LessonState = {
  subjects: [],
  lessons: [
    {
      id: 1,
      shiftId: 0,
      subjectId: 1,
      room: 1,
      teacherId: '000000000000000000001',
      studentId: '123456789012345678901',
      startAt: '2021-12-10T18:30:00+09:00',
      endAt: '2021-12-10T20:00:00+09:00',
      notes: '',
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
      if (Axios.isAxiosError(err)) {
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
      console.log(request)
      await $axios.$post('/v1/subjects', request)
      await this.getAllSubjects()
    } catch (err) {
      if (Axios.isAxiosError(err)) {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      }
      throw new Error('internal server error')
    }
  }
}

import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { Lesson, LessonState, Subject, SubjectMap } from '~/types/store/lesson'

const initialState: LessonState = {
  subjects: [
    {
      id: 1,
      name: '国語',
      color: '#F8BBD0',
      createdAt: '',
      updatedAt: '',
    },
    {
      id: 2,
      name: '数学',
      color: '#BBDEFB',
      createdAt: '',
      updatedAt: '',
    },
    {
      id: 3,
      name: '英語',
      color: '#FEE6C9',
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
      startAt: '2021-12-10 18:30:00',
      endAt: '2021-12-10 20:00:00',
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
}

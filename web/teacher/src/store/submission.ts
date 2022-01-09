import { Module, VuexModule } from 'vuex-module-decorators'
import { Submission, SubmissionEditState } from '~/types/store/submission'

const initialState: Submission[] = [
  {
    title: '1月',
    endDate: '20210125',
    submissionStatus: '未提出',
    editStatus: '入力する',
  },
  {
    title: '2月',
    endDate: '20230225',
    submissionStatus: '提出済み',
    editStatus: '編集する',
  },
]

const initialEditState: SubmissionEditState = {
  summary: {
    id: 1,
    year: 2022,
    month: 2,
    status: 2,
    openAt: '2021-12-10T18:30:00+09:00',
    endAt: '2021-12-10T18:30:00+09:00',
    createdAt: '2021-12-10T18:30:00+09:00',
    updatedAt: '2021-12-10T18:30:00+09:00',
  },
  shifts: [
    {
      date: '20220201',
      isClosed: false,
      lessons: [
        {
          id: 1,
          startTime: '1700',
          endTime: '1830',
        },
        {
          id: 2,
          startTime: '1830',
          endTime: '2000',
        },
        {
          id: 3,
          startTime: '2000',
          endTime: '2130',
        },
      ],
    },
    {
      date: '20220202',
      isClosed: true,
      lessons: [],
    },
    {
      date: '20220203',
      isClosed: false,
      lessons: [
        {
          id: 4,
          startTime: '1700',
          endTime: '1830',
        },
        {
          id: 5,
          startTime: '1830',
          endTime: '2000',
        },
      ],
    },
    {
      date: '20220204',
      isClosed: false,
      lessons: [
        {
          id: 6,
          startTime: '1700',
          endTime: '1830',
        },
        {
          id: 7,
          startTime: '1830',
          endTime: '2000',
        },
      ],
    },
  ],
}

@Module({
  name: 'submission',
  stateFactory: true,
  namespaced: true,
})
export default class SubmissionModule extends VuexModule {
  private submission: Submission[] = initialState
  private submissionEdit: SubmissionEditState = initialEditState

  public get getSubmissions(): Submission[] {
    return this.submission
  }

  public get getSubmissionsEdit(): SubmissionEditState {
    return this.submissionEdit
  }
}

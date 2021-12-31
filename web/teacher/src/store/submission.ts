import { Module, VuexModule } from 'vuex-module-decorators'
import { Submission } from '~/types/store/submission'

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

@Module({
  name: 'submission',
  stateFactory: true,
  namespaced: true,
})
export default class SubmissionModule extends VuexModule {
  private submission: Submission[] = initialState

  public get getSubmissions(): Submission[] {
    return this.submission
  }
}

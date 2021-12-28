export type SubmissionStatusType = '未提出' | '提出済み'

export interface Submission {
  id: number
  title: string
  endDate: string
  submissionStatus: SubmissionStatusType
  editStatus: '入力する' | '編集する'
}

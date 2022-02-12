// SchoolType 校種
export const SchoolTypeArray = ['小学校', '中学校', '高校', 'その他'] as const
export type SchoolType = typeof SchoolTypeArray[number]

// PromiseState 通信状態
export enum PromiseState {
  NONE = 0, // 未接続
  LOADING = 1, // 通信中
}

// ShiftStatus シフト募集状況
export enum ShiftStatus {
  UNKNOWN = 0, // 不明
  WAITING = 1, // 募集前
  ACCEPTING = 2, // 募集中
  FINISHED = 3, // 締切後
}

// SubmissionStatus シフト提出状況
export enum SubmissionStatus {
  UNKNOWN = 0, // 不明
  WAITING = 1, // 未提出
  SUBMITTED = 2, // 提出済
}

export interface CommonState {
  snackbarColor: string
  snackbarMessage: string
  promiseState: PromiseState
}

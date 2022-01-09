// PromiseState 通信状態
export enum PromiseState {
  NONE = 0, // 未接続
  LOADING = 1, // 通信中
}

// Role 権限
export enum Role {
  UNKNOWN = 0, // 不明
  TEACHER = 1, // 講師
  ADMINISTRATOR = 2, // 管理者
}

// SchoolType 校種
export const SchoolTypeArray = ['小学校', '中学校', '高校', 'その他'] as const
export type SchoolType = typeof SchoolTypeArray[number]

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

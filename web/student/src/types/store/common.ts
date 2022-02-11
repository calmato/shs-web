// PromiseState 通信状態
export enum PromiseState {
  NONE = 0, // 未接続
  LOADING = 1, // 通信中
}

// SchoolType 校種
export enum SchoolType {
  UNKNOWN = 0, // 不明
  ELEMENTARY_SCHOOL = 1, // 小学校
  JUNIOR_HIGH_SCHOOL = 2, // 中学校
  HIGH_SCHOOL = 3, // 高等学校
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

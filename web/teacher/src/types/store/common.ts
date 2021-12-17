// PromiseState 通信状態
export enum PromiseState {
  NONE = 0, // 未接続
  LOADING = 1, // 通信中
}

// Role 権限
export enum Role {
  TEACHER = 0, // 講師
  ADMINISTRATOR = 1, // 管理者
}

// SchoolType 校種
export enum SchoolType {
  ElementarySchool = 0, // 小学校
  JuniorHighSchool = 1, // 中学校
  HighSchool = 2, // 高等学校
}

export interface CommonState {
  promiseState: PromiseState
}

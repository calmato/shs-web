export enum PromiseState {
  NONE,
  LOADING,
}

export interface CommonState {
  promiseState: PromiseState
}

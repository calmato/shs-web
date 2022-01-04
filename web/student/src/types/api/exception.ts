/**
 * Error - エラーレスポンス
 */
export interface ErrorResponse {
  status: number // ステータスコード
  message: string // エラー概要
  details: string // エラー詳細
}

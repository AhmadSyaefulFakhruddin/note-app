export type ReturnType<T> =
  | {
      success: true
      data: T
    }
  | {
      success: false
      error: Error
    }

export type ApiResponse<T> = {
  status?: 'success' | 'error' | 'fail'
  data?: T
  message?: string
}

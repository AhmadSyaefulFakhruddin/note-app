// #region AppError
export class AppError extends Error {
  public readonly name: string
  public readonly statusCode: number
  public readonly isOperational: boolean
  public readonly errorCode?: string
  public readonly details?: any

  constructor(
    message: string,
    statusCode: number,
    isOperational: boolean = true,
    errorCode?: string,
    details?: any,
  ) {
    super(message)

    Object.setPrototypeOf(this, new.target.prototype) // restore prototype chain

    this.name = this.constructor.name
    this.statusCode = statusCode
    this.isOperational = isOperational
    this.errorCode = errorCode
    this.details = details

    Error.captureStackTrace(this, this.constructor)
  }
}
// #endregion AppError

export class NotFoundError extends AppError {
  constructor(message: string = 'Resource not found', details?: any) {
    super(message, 404, true, 'NOT_FOUND', details)
  }
}

export class ValidationError extends AppError {
  constructor(message: string = 'Validation failed', details?: any) {
    super(message, 400, true, 'VALIDATION_ERROR', details)
  }
}

export class UnauthorizedError extends AppError {
  constructor(message: string = 'Unauthorized access', details?: any) {
    super(message, 401, true, 'UNAUTHORIZED', details)
  }
}

export class InternalServerError extends AppError {
  constructor(message: string = 'Internal server error', details?: any) {
    super(message, 500, true, 'INTERNAL_SERVER_ERROR', details)
  }
}

export class ResponseNotOkError extends AppError {
  constructor(message: string = 'Response not OK', details?: any) {
    super(message, 502, true, 'RESPONSE_NOT_OK', details)
  }
}

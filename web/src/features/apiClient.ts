import type { ZodSchema } from 'zod'
import type { ApiResponse } from './type'
import { AppError, NotFoundError, UnauthorizedError } from './errors'
import { notFound } from '@tanstack/react-router'

interface FetchOptions extends RequestInit {
  data?: any
  method: 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE'
}

const API_BASE_URL = process.env.API_BASE_URL || ''

export async function apiClient<T>(
  endpoint: string,
  options: FetchOptions = { method: 'GET' },
  zodSchema?: ZodSchema<T>,
): Promise<T> {
  const { data, headers, method, ...fetchOptions } = options

  const response = await fetch(`${API_BASE_URL}${endpoint}`, {
    method,
    headers: {
      'Content-Type': 'application/json',
      ...headers,
    },
    body: data ? JSON.stringify(data) : undefined,
    ...fetchOptions,
  })

  if (!response.ok) {
    if (response.status === 401) {
      throw new UnauthorizedError(`Unauthorized access to ${endpoint}`)
    }

    if (response.status === 404) {
      const contentType = response.headers.get('Content-Type') || ''
      if (contentType.includes('application/json')) {
        const errorData: ApiResponse<T> = await response.json()
        if (errorData && errorData.status === 'fail') {
          throw notFound()
        }
      }

      throw new NotFoundError(`Resource not found at ${endpoint}`)
    }

    const error = await response.json().catch((err) => err)
    throw new AppError(
      `API request failed: ${response.status} ${response.statusText} - ${error.message || 'Api Request Failed'}`,
      response.status,
      true,
      'API_ERROR',
      error,
    )
  }

  const result: ApiResponse<T> = await response.json().catch(() => null)

  if (zodSchema) {
    const parsed = zodSchema.safeParse(result.data)
    if (!parsed.success) {
      throw new Error(
        `Validation failed for ${endpoint}: ${parsed.error.message}`,
      )
    }
    return parsed.data as T
  }

  return result.data as T
}

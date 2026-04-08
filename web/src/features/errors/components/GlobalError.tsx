import React from 'react'

import { ErrorPage } from './ErrorPage'

interface GlobalErrorProps {
  error: Error
  reset?: () => void // Provided by TanStack Router to retry the render/loader
  componentStack?: string // Optional stack trace for debugging
}

export const GlobalError: React.FC<GlobalErrorProps> = ({
  error,
  reset,
  componentStack,
}) => {
  return (
    <ErrorPage
      title={error.name || 'Something went wrong'}
      message={
        error.message ||
        'An unexpected error occurred. Our team has been notified.'
      }
      onReset={reset}
      componentStack={componentStack}
    />
  )
}

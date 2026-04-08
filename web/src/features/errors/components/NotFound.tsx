import { ErrorPage } from './ErrorPage'

export function NotFoundPage() {
  return (
    <ErrorPage
      statusCode={404}
      title="Page not found"
      message="The page you're looking for doesn't exist or has been moved."
    />
  )
}

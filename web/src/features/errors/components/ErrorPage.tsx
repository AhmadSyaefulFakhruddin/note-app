import { Button } from '#/components/ui/button'
import {
  Card,
  CardHeader,
  CardTitle,
  CardContent,
  CardFooter,
} from '#/components/ui/card'
import { useRouter } from '@tanstack/react-router'
import { AlertCircle, Home, RefreshCw } from 'lucide-react'

interface ErrorPageProps {
  title: string
  message: string
  statusCode?: number
  onReset?: () => void
  componentStack?: string
}

export function ErrorPage({
  title = 'Something went wrong',
  message = 'An unexpected error occurred. Please try again later.',
  statusCode,
  onReset,
  componentStack,
}: ErrorPageProps) {
  const router = useRouter()

  const handleReset = () => {
    if (onReset) {
      router.invalidate()
      onReset()
    }
  }

  return (
    <div className="flex min-h-[80vh] items-center justify-center p-4">
      <Card className="max-w-md w-full shadow-lg">
        <CardHeader className="text-center pb-2">
          <div className="mx-auto mb-4 flex h-16 w-16 items-center justify-center rounded-full bg-destructive/10">
            <AlertCircle className="h-8 w-8 text-destructive" />
          </div>
          {statusCode && (
            <p className="text-sm font-semibold text-muted-foreground uppercase tracking-wider">
              Error {statusCode}
            </p>
          )}
          <CardTitle className="text-2xl">{title}</CardTitle>
        </CardHeader>
        <CardContent className="text-center">
          <p className="text-muted-foreground">{message}</p>
          {process.env.NODE_ENV === 'development' && componentStack && (
            <div className="bg-muted p-4 rounded-md text-left overflow-auto border border-border">
              <code className="text-sm text-muted-foreground wrap-break-word">
                {componentStack}
              </code>
            </div>
          )}
        </CardContent>
        <CardFooter className="flex justify-center gap-4">
          <Button
            variant="outline"
            onClick={() => router.navigate({ to: '/' })}
          >
            <Home className="mr-2 h-4 w-4" />
            Go Home
          </Button>

          {onReset && (
            <Button onClick={handleReset}>
              <RefreshCw className="mr-2 h-4 w-4" />
              Try Again
            </Button>
          )}
        </CardFooter>
      </Card>
    </div>
  )
}

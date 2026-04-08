import { Button } from '#/components/ui/button'
import { ErrorPage } from '#/features/errors/components/ErrorPage'
import { Link } from '@tanstack/react-router'
import { ArrowLeft, FileQuestion } from 'lucide-react'

export function NoteNotFoundComponent() {
  return (
    <ErrorPage
      title="Note not found"
      message={`The note you are looking for doesn't exist, has been deleted, or you
        might not have permission to view it.`}
    />
  )

  return (
    <div className="container mx-auto px-4 py-20 flex flex-col items-center justify-center text-center min-h-[60vh]">
      <div className="bg-muted p-6 rounded-full mb-6">
        <FileQuestion className="h-12 w-12 text-muted-foreground" />
      </div>
      <h1 className="text-3xl font-extrabold tracking-tight text-foreground mb-2">
        Note not found
      </h1>
      <p className="text-muted-foreground max-w-md mb-8">
        The note you are looking for doesn't exist, has been deleted, or you
        might not have permission to view it.
      </p>
      <Button asChild>
        <Link to="/">
          <ArrowLeft className="mr-2 h-4 w-4" />
          Back to all notes
        </Link>
      </Button>
    </div>
  )
}

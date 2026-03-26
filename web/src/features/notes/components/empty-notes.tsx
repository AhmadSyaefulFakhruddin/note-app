import { FileText } from 'lucide-react'
import { Button } from '@/components/ui/button'

export function EmptyNotes() {
  return (
    <div className="flex flex-col items-center justify-center py-20 text-center space-y-4">
      <FileText className="h-10 w-10 text-muted-foreground" />

      <div>
        <h3 className="text-lg font-medium">No notes yet</h3>

        <p className="text-sm text-muted-foreground">
          Create your first note to get started.
        </p>
      </div>

      <Button>Create Note</Button>
    </div>
  )
}

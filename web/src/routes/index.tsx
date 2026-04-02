import { Button } from '#/components/ui/button'
import { Input } from '#/components/ui/input'
import { NoteGrid } from '#/features/notes/components/home-page/note-grid'
import { fetchNoteListOptions } from '#/features/notes/services/note.service'
import { createFileRoute } from '@tanstack/react-router'
import { Plus, Search } from 'lucide-react'

export const Route = createFileRoute('/')({
  loader: async ({ context }) => {
    await context.queryClient.ensureQueryData(fetchNoteListOptions)
  },
  component: App,
})

function App() {
  return (
    <main className="container mx-auto px-4 py-8 max-w-350">
      <header className="flex flex-col md:flex-row md:items-center justify-between gap-4 mb-8">
        <div>
          <h1 className="text-3xl font-extrabold tracking-tight text-foreground">
            All Notes
          </h1>
          <p className="text-muted-foreground mt-1 text-sm">
            Manage your thoughts, ideas, and tasks.
          </p>
        </div>

        {/* Toolbar: Search & Action */}
        <div className="flex items-center gap-3 w-full md:w-auto">
          <div className="relative flex-1 md:w-64">
            <Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
            <Input
              type="search"
              placeholder="Search notes..."
              className="pl-9 bg-background focus-visible:ring-primary"
            />
          </div>
          <Button>
            <Plus className="mr-2 h-4 w-4" />
            New Note
          </Button>
        </div>
      </header>

      {/* Grid Component (Implicitly suspended if data isn't ready) */}
      <section aria-label="Notes Grid">
        <NoteGrid />
      </section>
    </main>
  )
}

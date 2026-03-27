import { EmptyNotes } from '#/features/notes/components/home-page/empty-notes'
import { NoteCard } from '#/features/notes/components/home-page/note-card'
import { getRouteApi } from '@tanstack/react-router'

export function NotesList() {
  const apiRoute = getRouteApi('/')
  const notes = apiRoute.useLoaderData()

  const hasNotes = notes.length > 0

  return (
    <main className="page-wrap px-4 pb-8 pt-14">
      {/* Notes */}
      {hasNotes ? (
        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
          {notes.map((note) => (
            <NoteCard key={note.id} note={note} />
          ))}
        </div>
      ) : (
        <EmptyNotes />
      )}
    </main>
  )
}

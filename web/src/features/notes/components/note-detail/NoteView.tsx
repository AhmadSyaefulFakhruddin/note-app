// import { Button } from '@/components/ui/button'
// import { Edit2 } from 'lucide-react'
import { NoteBreadcrumbs } from '#/features/notes/components/note-detail/NoteBreadcrumbs'
import { NoteActions } from '#/features/notes/components/note-detail/NoteActions'
import { NoteMetadata } from '#/features/notes/components/note-detail/NoteMetadata'
import { useState } from 'react'
import { fetchNoteByIdOptions } from '../../services/note.service'
import { useLoaderData, useParams } from '@tanstack/react-router'
import { useSuspenseQuery } from '@tanstack/react-query'

export function NoteView() {
  //   const { noteId } = useParams({ from: '/notes/$noteId' })
  //   const { data: note } = useSuspenseQuery(fetchNoteByIdOptions(noteId))

  const note = useLoaderData({ from: '/notes/$noteId' })

  // Local state for toggling modes (doesn't need global store since it's page-specific)
  const [isEditing, setIsEditing] = useState(false)
  return (
    <article className="container mx-auto px-4 py-8 max-w-4xl min-h-screen flex flex-col">
      <header className="mb-6 space-y-4">
        <NoteBreadcrumbs note={note} />

        {!isEditing && (
          <div className="flex items-start justify-between gap-4">
            <h1 className="text-4xl font-extrabold tracking-tight text-foreground wrap-break-word flex-1">
              {note.title}
            </h1>
            <div className="flex items-center gap-2 mt-1">
              {/* <Button
                variant="outline"
                size="sm"
                onClick={() => setIsEditing(true)}
              >
                <Edit2 className="mr-2 h-4 w-4" />
                Edit
              </Button> */}
              <NoteActions isPinned={note.isPinned} />
            </div>
          </div>
        )}
      </header>

      <main className="flex-1 flex flex-col">
        {/* {isEditing ? (
          <NoteEditor note={note} onCancel={() => setIsEditing(false)} />
        ) : ( */}
        <section className="prose prose-neutral dark:prose-invert max-w-none text-lg leading-relaxed">
          {/* If using Markdown, a renderer like ReactMarkdown would go here */}
          <p className="whitespace-pre-wrap">{note.content}</p>
        </section>
        {/* )} */}
      </main>

      {!isEditing && <NoteMetadata note={note} />}
    </article>
  )
}

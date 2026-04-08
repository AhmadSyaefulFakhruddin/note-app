import { createFileRoute, notFound } from '@tanstack/react-router'
import { useSuspenseQuery } from '@tanstack/react-query'
import { useState } from 'react'
import { fetchNoteById } from '#/features/notes/server/note'
import { fetchNoteByIdOptions } from '#/features/notes/services/note.service'
import { NoteView } from '#/features/notes/components/note-detail/NoteView'
import { NoteNotFoundComponent } from '#/features/notes/components/note-detail/NoteNotFound'

export const Route = createFileRoute('/notes/$noteId')({
  // 1. Loader fetches data server-side to guarantee SEO and eliminate loading spinners
  loader: async ({ params: { noteId }, context: { queryClient } }) => {
    const note = await queryClient.ensureQueryData(fetchNoteByIdOptions(noteId))

    // const note = await fetchNoteById({ data: { id: noteId } })

    if (!note) {
      throw notFound()
    }

    return note
  },

  head: ({ loaderData }) => ({
    meta: [
      { title: `${loaderData?.title || 'The Note Not Found'} | Notes App` },
      { name: 'description', content: `View note: ${loaderData?.title}` },
      { property: 'og:title', content: loaderData?.title },
      { property: 'og:type', content: 'article' },
    ],
  }),

  component: NoteDetailPage,
  notFoundComponent: NoteNotFoundComponent,
})

function NoteDetailPage() {
  return <NoteView />
}

import React, { useMemo } from 'react'
import { FileText } from 'lucide-react'
import { NoteCard } from './note-card'
import { useSuspenseQuery } from '@tanstack/react-query'
import { fetchNoteListOptions } from '../../services/note.service'

export const NoteGrid: React.FC = () => {
  const { data: notes } = useSuspenseQuery(fetchNoteListOptions)

  if (!notes.length) {
    return (
      <div className="flex flex-col items-center justify-center py-20 text-muted-foreground border-2 border-dashed border-border rounded-xl">
        <FileText className="h-12 w-12 mb-4 opacity-20" />
        <p className="text-lg font-medium">No notes found</p>
        <p className="text-sm">Create your first note to get started.</p>
      </div>
    )
  }

  // Sort notes: Pinned first, then by most recently updated
  const sortedNotes = useMemo(() => {
    return [...notes].sort((a, b) => {
      if (a.isPinned === b.isPinned) {
        return new Date(b.updatedAt).getTime() - new Date(a.updatedAt).getTime()
      }
      return a.isPinned ? -1 : 1
    })
  }, [notes])

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4 md:gap-6">
      {sortedNotes.map((note) => (
        <NoteCard key={note.id} note={note} />
      ))}
    </div>
  )
}

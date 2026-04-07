import React from 'react'
import { Badge } from '@/components/ui/badge'
import type { Note } from '../../types/note.type'

export const NoteMetadata: React.FC<{ note: Note }> = ({ note }) => {
  const formatDate = (date: Date) =>
    new Intl.DateTimeFormat('en-US', {
      dateStyle: 'medium',
      timeStyle: 'short',
    }).format(date)

  return (
    <footer className="mt-12 pt-6 border-t border-border flex flex-col sm:flex-row sm:items-center justify-between gap-4 text-sm text-muted-foreground">
      <div className="flex flex-wrap gap-2">
        {note.tags.map((tag) => (
          <Badge key={tag.id} variant="secondary" className="font-normal">
            #{tag.name}
          </Badge>
        ))}
      </div>
      <div className="flex flex-col sm:text-right">
        <span>Created: {formatDate(note.createdAt)}</span>
        <span>Last modified: {formatDate(note.updatedAt)}</span>
      </div>
    </footer>
  )
}

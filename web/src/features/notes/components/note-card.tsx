import React from 'react'
import { Link } from '@tanstack/react-router'
import { Pin, Cloud, CloudOff, RefreshCw } from 'lucide-react'
import {
  Card,
  CardHeader,
  CardTitle,
  CardContent,
  CardFooter,
} from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import type { Note } from '../types/note.schema'

interface NoteCardProps {
  note: Note
}

export const NoteCard: React.FC<NoteCardProps> = ({ note }) => {
  // Determine sync icon
  const SyncIcon =
    note.syncStatus === 'synced'
      ? Cloud
      : note.syncStatus === 'syncing'
        ? RefreshCw
        : CloudOff

  // Format date nicely
  const formattedDate = new Intl.DateTimeFormat('en-US', {
    month: 'short',
    day: 'numeric',
    year: 'numeric',
  }).format(new Date(note.updatedAt))

  return (
    // <Link
    //   to="/notes/$noteId"
    //   params={{ noteId: note.id }}
    //   className="block outline-none focus-visible:ring-2 focus-visible:ring-primary rounded-xl transition-transform hover:-translate-y-1"
    // >
    <Card className="h-full flex flex-col bg-card hover:bg-accent/40 transition-colors border-border shadow-sm cursor-pointer relative overflow-hidden group">
      {/* Top-right corner indicators */}
      <div className="absolute top-3 right-3 flex items-center gap-2 text-muted-foreground opacity-70 group-hover:opacity-100 transition-opacity">
        {note.isPinned && (
          <Pin
            className="h-4 w-4 fill-foreground text-foreground"
            aria-label="Pinned"
          />
        )}
        <SyncIcon
          className={`h-4 w-4 ${note.syncStatus === 'syncing' ? 'animate-spin' : ''}`}
          aria-label={`Sync status: ${note.syncStatus}`}
        />
      </div>

      <CardHeader className="pb-3 pt-5 pr-14">
        <CardTitle className="text-lg font-bold line-clamp-1 leading-tight">
          {note.title}
        </CardTitle>
        <p className="text-xs font-medium text-muted-foreground mt-1">
          {note.folder}
        </p>
      </CardHeader>

      <CardContent className="flex-1 pb-4">
        <p className="text-sm text-muted-foreground line-clamp-3 leading-relaxed whitespace-pre-wrap">
          {note.content}
        </p>
      </CardContent>

      <CardFooter className="pt-0 pb-4 flex items-center justify-between mt-auto">
        <div className="flex gap-1.5 overflow-hidden">
          {note.tags.slice(0, 2).map((tag) => (
            <Badge
              key={tag}
              variant="secondary"
              className="text-[10px] px-1.5 py-0 h-5 font-medium truncate max-w-20"
            >
              {tag}
            </Badge>
          ))}
          {note.tags.length > 2 && (
            <Badge
              variant="outline"
              className="text-[10px] px-1.5 py-0 h-5 font-medium"
            >
              +{note.tags.length - 2}
            </Badge>
          )}
        </div>
        <span className="text-xs text-muted-foreground ml-2 whitespace-nowrap">
          {formattedDate}
        </span>
      </CardFooter>
    </Card>
    // </Link>
  )
}

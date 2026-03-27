import React from 'react'
import { ChevronRight, Cloud, CloudOff, RefreshCw } from 'lucide-react'
import type { Note } from '../../types/note.type'

export const NoteBreadcrumbs: React.FC<{ note: Note }> = ({ note }) => {
  const SyncIcon =
    note.syncStatus === 'synced'
      ? Cloud
      : note.syncStatus === 'syncing'
        ? RefreshCw
        : CloudOff

  return (
    <div className="flex items-center justify-between w-full text-sm text-muted-foreground pb-4 border-b border-border">
      <nav aria-label="Breadcrumb" className="flex items-center space-x-1">
        <span className="hover:text-foreground cursor-pointer transition-colors">
          Notes
        </span>
        <ChevronRight className="h-4 w-4" />
        <span className="hover:text-foreground cursor-pointer transition-colors">
          {note.folder}
        </span>
        <ChevronRight className="h-4 w-4" />
        <span className="truncate max-w-37.5 font-medium text-foreground">
          {note.title}
        </span>
      </nav>

      <div
        className="flex items-center space-x-2"
        title={`Status: ${note.syncStatus}`}
      >
        <SyncIcon
          className={`h-4 w-4 ${note.syncStatus === 'syncing' ? 'animate-spin' : ''}`}
        />
        <span className="sr-only">Sync status: {note.syncStatus}</span>
      </div>
    </div>
  )
}

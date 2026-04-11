import React from 'react'
import { ChevronRight, Cloud, CloudOff, RefreshCw } from 'lucide-react'
import type { Note } from '../../types/note.type'
import { Button } from '#/components/ui/button'
import { Link } from '@tanstack/react-router'

export const NoteBreadcrumbs: React.FC<{ note: Note }> = ({ note }) => {
  const SyncIcon =
    note.syncStatus === 'synced'
      ? Cloud
      : note.syncStatus === 'syncing'
        ? RefreshCw
        : CloudOff

  const folderParts = note.folder ? note.folder.split('/') : []

  return (
    <div className="flex items-center justify-between w-full text-sm text-muted-foreground pb-4 border-b border-border">
      <nav aria-label="Breadcrumb" className="flex items-center space-x-1">
        <Button size="xs" asChild variant={'none'}>
          <Link to="/">Notes</Link>
        </Button>
        <ChevronRight className="h-4 w-4" />
        {folderParts.map((part, index) => (
          <React.Fragment key={index}>
            <Button size="xs" asChild variant={'none'}>
              <Link to="/">{part}</Link>
            </Button>
            {index < folderParts.length - 1 && (
              <ChevronRight className="h-4 w-4" />
            )}
          </React.Fragment>
        ))}

        {folderParts.length !== 0 && <ChevronRight className="h-4 w-4" />}
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

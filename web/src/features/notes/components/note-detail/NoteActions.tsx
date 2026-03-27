import React from 'react'
import { MoreVertical, Pin, Share, Archive, Trash2 } from 'lucide-react'
import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'

export const NoteActions: React.FC<{ isPinned: boolean }> = ({ isPinned }) => {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost" size="icon" aria-label="Open note actions">
          <MoreVertical className="h-5 w-5" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-48">
        <DropdownMenuItem className="cursor-pointer">
          <Pin className="mr-2 h-4 w-4" />
          <span>{isPinned ? 'Unpin Note' : 'Pin Note'}</span>
        </DropdownMenuItem>
        <DropdownMenuItem className="cursor-pointer">
          <Share className="mr-2 h-4 w-4" />
          <span>Share / Export</span>
        </DropdownMenuItem>
        <DropdownMenuSeparator />
        <DropdownMenuItem className="cursor-pointer">
          <Archive className="mr-2 h-4 w-4" />
          <span>Archive</span>
        </DropdownMenuItem>
        <DropdownMenuItem className="cursor-pointer text-destructive focus:bg-destructive focus:text-destructive-foreground">
          <Trash2 className="mr-2 h-4 w-4" />
          <span>Delete</span>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}

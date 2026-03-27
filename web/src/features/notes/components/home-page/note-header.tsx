import { Plus, Tags, Trash2 } from 'lucide-react'
import { Button } from '../../../../components/ui/button'
import { NotesSearch } from './note-search'
import ThemeToggle from '../../../../components/ThemeToggle'
import { useState } from 'react'
import { TagsDialog } from './tags-dialog'
import { useStore } from '@tanstack/react-store'
import { setSelectionMode } from '../../stores/ui/ui-action'
import { uiStore } from '../../stores/ui/ui-store'
import { deleteNote } from '../../stores/notes/note-action'

export function NoteHeader() {
  const [tagsOpen, setTagsOpen] = useState(false)
  const { selectedNotes, selectionMode } = useStore(uiStore, (s) => s)

  const selectedCount = selectedNotes.length

  function deleteNotes() {
    selectedNotes.forEach((note) => {
      deleteNote(note.id)
    })
  }

  return (
    <div className="flex items-center justify-between">
      <div>
        <h1 className="text-3xl font-semibold tracking-tight">Notes</h1>
        <p className="text-muted-foreground text-sm">
          Capture your ideas quickly.
        </p>
      </div>

      <div className="flex gap-2 items-center ">
        <ThemeToggle />

        {/* Tags manager */}
        <Button variant="outline" size="sm" onClick={() => setTagsOpen(true)}>
          <Tags className="h-4 w-4 mr-2" />
          Tags
        </Button>

        {/* Toggle selection mode */}
        <Button
          variant={selectionMode ? 'secondary' : 'outline'}
          size="sm"
          onClick={() => setSelectionMode()}
        >
          <Trash2 className="h-4 w-4 mr-2" />
          Select
        </Button>

        {selectionMode && selectedCount > 0 && (
          <Button variant="destructive" size="sm" onClick={deleteNotes}>
            Delete ({selectedCount})
          </Button>
        )}

        <Button className="gap-2">
          <Plus className="h-4 w-4" />
          New Note
        </Button>

        <NotesSearch />

        <TagsDialog open={tagsOpen} onOpenChange={setTagsOpen} />
      </div>
    </div>
  )
}

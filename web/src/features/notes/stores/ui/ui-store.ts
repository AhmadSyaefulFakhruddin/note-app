import { Store } from '@tanstack/store'
import type { Note } from '#/features/notes/note-type'

export const uiStore = new Store<{
  selectedNotes: Note[]
  selectionMode: boolean
}>({
  selectedNotes: [],
  selectionMode: false,
})

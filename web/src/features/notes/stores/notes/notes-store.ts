import { Store } from '@tanstack/store'
import type { Note } from '../../types/note.schema'

export const notesStore = new Store<{ notes: Note[] }>({
  notes: [],
})

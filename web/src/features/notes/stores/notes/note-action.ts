import { uiStore } from '../ui/ui-store'
import { notesStore } from './notes-store'

export function deleteNote(id: string) {
  notesStore.setState((state) => ({
    ...state,
    notes: state.notes.filter((note) => note.id != id),
  }))
}

export function setSelectNote(id: string) {
  const selectedNote = uiStore.state.selectedNotes.find((sn) => sn.id === id)

  console.log(uiStore.state.selectedNotes)

  if (!selectedNote) {
    uiStore.setState((state) => ({
      ...state,
      selectedNotes: [
        ...state.selectedNotes,
        notesStore.state.notes.find((note) => note.id === id)!,
      ],
    }))
    return
  }

  uiStore.setState((state) => ({
    ...state,
    selectedNotes: state.selectedNotes.filter((note) => note.id !== id),
  }))
}

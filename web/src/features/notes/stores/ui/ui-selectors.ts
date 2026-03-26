import { uiStore } from './ui-store'

export function getSelectedCount() {
  const notes = uiStore.state.selectedNotes

  return notes.length
}

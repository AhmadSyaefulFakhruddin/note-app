import { uiStore } from './ui-store'

export function setSelectionMode() {
  uiStore.setState((state) => ({
    ...state,
    selectionMode: !state.selectionMode,
  }))
}

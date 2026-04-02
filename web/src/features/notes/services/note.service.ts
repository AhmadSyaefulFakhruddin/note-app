import { queryOptions } from '@tanstack/react-query'
import { fetchNoteList } from '../server/note'

export const fetchNoteListOptions = queryOptions({
  queryKey: ['notes'],
  queryFn: () => fetchNoteList(),
})

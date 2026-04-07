import { queryOptions } from '@tanstack/react-query'
import { fetchNoteById, fetchNoteList } from '../server/note'

export const fetchNoteListOptions = queryOptions({
  queryKey: ['notes'],
  queryFn: () => fetchNoteList(),
})

export const fetchNoteByIdOptions = (id: string) =>
  queryOptions({
    queryKey: ['notes', id],
    queryFn: () => fetchNoteById({ data: { id } }),
  })

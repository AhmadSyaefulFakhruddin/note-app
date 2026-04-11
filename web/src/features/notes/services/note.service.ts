import { queryOptions, useMutation } from '@tanstack/react-query'
import { fetchNoteById, fetchNoteList } from '../server/note'
import { getContext } from '#/integrations/tanstack-query/root-provider'

export const fetchNoteListOptions = queryOptions({
  queryKey: ['notes'],
  queryFn: () => fetchNoteList(),
})

export const fetchNoteByIdOptions = (id: string) =>
  queryOptions({
    queryKey: ['notes', id],
    queryFn: () => fetchNoteById({ data: { id } }),
  })

export const updateNoteOptions = (
  id: string,
  data: { title: string; content: string },
) =>
  useMutation({
    mutationKey: ['notes', id],
    mutationFn: async () => {
      // Simulate an API call to update the note
      await new Promise((resolve) => setTimeout(resolve, 500)) // Simulate network delay
      // In a real implementation, you would make an API request here to update the note on the server
      return { id, ...data } // Return the updated note data
    },
    onSuccess: () => {
      const queryClient = getContext().queryClient
      // Invalidate the query for this note to refetch the updated data
      queryClient.invalidateQueries({ queryKey: ['notes', id] })
    },
  })

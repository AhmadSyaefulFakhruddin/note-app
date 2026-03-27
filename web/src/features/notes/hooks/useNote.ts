import { queryOptions } from '@tanstack/react-query'

export const noteQueryOptions = (noteId: string) =>
  queryOptions({
    queryKey: ['note', noteId],
    enabled: !!noteId,
    refetchOnWindowFocus: false,
    refetchOnReconnect: false,
  })

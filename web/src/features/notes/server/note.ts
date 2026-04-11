import { createServerFn } from '@tanstack/react-start'
import type {
  ApiResponse,
  DefaultApiResponse,
  Note,
  Notes,
} from '../types/note.type'
import {
  CreateNoteInputSchema,
  NoteIdInputSchema,
  NoteSchema,
  NotesSchema,
  UpdateNoteInputSchema,
} from '../types/note.schema'
import { apiClient } from '#/features/apiClient'

export const fetchNoteList = createServerFn({ method: 'GET' }).handler(
  async () => {
    return apiClient<Notes>(`/notes`, { method: 'GET' }, NotesSchema)
  },
)

export const fetchNoteById = createServerFn({ method: 'GET' })
  .inputValidator(NoteIdInputSchema)
  .handler(async ({ data: { id } }) => {
    return apiClient<Note>(`/notes/${id}`, { method: 'GET' }, NoteSchema)
  })

// export const updateNote = createServerFn({ method: 'POST' })
//   .inputValidator(UpdateNoteInputSchema)
//   .handler(async ({ data }) => {
//     const { id, ...updateData } = data

//   })

// export const deleteNote = createServerFn({ method: 'POST' })
//   .inputValidator(NoteIdInputSchema)
//   .handler(async ({ data: { id } }) => {
//     const response = await fetch(`${API_URL}/notes/${id}`, {
//       method: 'DELETE',
//     })

//     if (!response.ok) {
//       throw new Error(`Failed to delete note: ${response.statusText}`)
//     }

//     const result: ApiResponse<DefaultApiResponse> = await response.json()

//     if (result.status !== 'success') {
//       throw new Error(`API error: ${result.message || 'Unknown error'}`)
//     }

//     return result.data
//   })

export const createNote = createServerFn({ method: 'POST' })
  .inputValidator(CreateNoteInputSchema)
  .handler(async ({ data }) => {
    return apiClient<Note>(`/notes`, { method: 'POST', data }, NoteSchema)
  })

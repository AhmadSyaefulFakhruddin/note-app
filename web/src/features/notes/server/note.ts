import { createServerFn } from '@tanstack/react-start'
import type {
  ApiResponse,
  DefaultApiResponse,
  Note,
  Notes,
} from '../types/note.type'
import {
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

//     const response = await fetch(`${API_URL}/notes/${id}`, {
//       method: 'PATCH',
//       headers: { 'Content-Type': 'application/json' },
//       body: JSON.stringify(updateData),
//     })

//     if (!response.ok) {
//       throw new Error(`Failed to update note: ${response.statusText}`)
//     }

//     const result: ApiResponse<DefaultApiResponse> = await response.json()

//     if (result.status !== 'success') {
//       throw new Error(`API error: ${result.message || 'Unknown error'}`)
//     }

//     const validatedNote = NoteSchema.safeParse(result.data)

//     if (!validatedNote.success) {
//       console.error('Validation errors:', validatedNote.error)
//       throw new Error('Received invalid note data from API')
//     }

//     return validatedNote.data
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

// export const createNote = createServerFn({ method: 'POST' })
//   .inputValidator(NoteSchema.pick({ title: true, content: true }))
//   .handler(async ({ data }) => {
//     const response = await fetch(`${API_URL}/notes`, {
//       method: 'POST',
//       headers: { 'Content-Type': 'application/json' },
//       body: JSON.stringify(data),
//     })

//     if (!response.ok) {
//       throw new Error(`Failed to create note: ${response.statusText}`)
//     }

//     const result: ApiResponse<DefaultApiResponse> = await response.json()

//     if (result.status !== 'success') {
//       throw new Error(`API error: ${result.message || 'Unknown error'}`)
//     }

//     return result.data
//   })

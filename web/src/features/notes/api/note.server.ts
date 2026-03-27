import { createServerFn } from '@tanstack/react-start'
import type { ApiResponse, DefaultApiResponse, Note } from '../types/note.type'
import {
  NoteIdInputSchema,
  NoteSchema,
  UpdateNoteInputSchema,
} from '../types/note.schema'

const API_BASE_ = import.meta.env.API_URL

export const fetchNoteList = createServerFn({ method: 'GET' }).handler(
  async () => {
    const response = await fetch(`${API_BASE_}/notes`)

    if (!response.ok) {
      throw new Error(`Failed to fetch notes: ${response.statusText}`)
    }

    const result: ApiResponse<Note[]> = await response.json()

    if (result.status !== 'success') {
      throw new Error(`API error: ${result.message || 'Unknown error'}`)
    }

    const validatedNotes = NoteSchema.safeParse(result.data)

    if (!validatedNotes.success) {
      console.error('Validation errors:', validatedNotes.error)
      throw new Error('Received invalid note data from API')
    }

    return validatedNotes.data
  },
)

export const updateNote = createServerFn({ method: 'POST' })
  .inputValidator(UpdateNoteInputSchema)
  .handler(async ({ data }) => {
    const { id, ...updateData } = data

    const response = await fetch(`${API_BASE_}/notes/${id}`, {
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(updateData),
    })

    if (!response.ok) {
      throw new Error(`Failed to update note: ${response.statusText}`)
    }

    const result: ApiResponse<DefaultApiResponse> = await response.json()

    if (result.status !== 'success') {
      throw new Error(`API error: ${result.message || 'Unknown error'}`)
    }

    const validatedNote = NoteSchema.safeParse(result.data)

    if (!validatedNote.success) {
      console.error('Validation errors:', validatedNote.error)
      throw new Error('Received invalid note data from API')
    }

    return validatedNote.data
  })

export const deleteNote = createServerFn({ method: 'POST' })
  .inputValidator(NoteIdInputSchema)
  .handler(async ({ data: { id } }) => {
    const response = await fetch(`${API_BASE_}/notes/${id}`, {
      method: 'DELETE',
    })

    if (!response.ok) {
      throw new Error(`Failed to delete note: ${response.statusText}`)
    }

    const result: ApiResponse<DefaultApiResponse> = await response.json()

    if (result.status !== 'success') {
      throw new Error(`API error: ${result.message || 'Unknown error'}`)
    }

    return result.data
  })

export const createNote = createServerFn({ method: 'POST' })
  .inputValidator(NoteSchema.pick({ title: true, content: true }))
  .handler(async ({ data }) => {
    const response = await fetch(`${API_BASE_}/notes`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    })

    if (!response.ok) {
      throw new Error(`Failed to create note: ${response.statusText}`)
    }

    const result: ApiResponse<DefaultApiResponse> = await response.json()

    if (result.status !== 'success') {
      throw new Error(`API error: ${result.message || 'Unknown error'}`)
    }

    return result.data
  })

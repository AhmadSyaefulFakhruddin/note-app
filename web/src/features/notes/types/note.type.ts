import type z from 'zod'
import type {
  CreateNoteInputSchema,
  NoteSchema,
  NotesSchema,
  UpdateNoteInputSchema,
} from './note.schema'

export type Note = z.infer<typeof NoteSchema>
export type Notes = z.infer<typeof NotesSchema>
export type UpdateNoteInput = z.infer<typeof UpdateNoteInputSchema>
export type CreateNoteInput = z.infer<typeof CreateNoteInputSchema>

export type ApiResponse<T> = {
  status: 'success' | 'error' | 'fail'
  data?: T
  message?: string
  meta?: {
    page: number
    totalItems: number
  }
}

export type DefaultApiResponse = {
  message: string
  id: string
}

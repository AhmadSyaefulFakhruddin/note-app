import { z } from 'zod'

export const StatusResponseSchema = z.enum(['success', 'error', 'fail'])

export const TagsSchema = z.array(
  z.object({
    id: z.string(),
    name: z.string(),
  }),
)

export const NoteSchema = z.object({
  id: z.string(),
  title: z.string().min(1, 'Title cannot be empty'),
  content: z.string(),
  folder: z.string(),
  tags: TagsSchema.default([]),
  isPinned: z.boolean().default(false),
  isArchived: z.boolean().default(false),
  syncStatus: z.enum(['synced', 'syncing', 'error']).default('synced'),
  createdAt: z.coerce.date(),
  updatedAt: z.coerce.date(),
})

export const NotesSchema = z.array(NoteSchema)

export const UpdateNoteInputSchema = NoteSchema.pick({
  title: true,
  content: true,
  tags: true,
  folder: true,
  isPinned: true,
  isArchived: true,
  id: true,
})

export const NoteIdInputSchema = NoteSchema.pick({ id: true })

export const CreateNoteInputSchema = NoteSchema.pick({
  title: true,
  content: true,
  tags: true,
})

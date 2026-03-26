import { z } from 'zod'

export const NoteSchema = z.object({
  id: z.uuid(),
  title: z.string().min(1, 'Title cannot be empty'),
  content: z.string(),
  folder: z.string().default('Personal'),
  tags: z.array(z.string()),
  isPinned: z.boolean().default(false),
  isArchived: z.boolean().default(false),
  syncStatus: z.enum(['synced', 'syncing', 'error']).default('synced'),
  createdAt: z.date(),
  updatedAt: z.date(),
})

export const UpdateNoteSchema = NoteSchema.pick({
  title: true,
  content: true,
})

export type Note = z.infer<typeof NoteSchema>
export type UpdateNotePayload = z.infer<typeof UpdateNoteSchema>

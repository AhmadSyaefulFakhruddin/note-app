import React from 'react'
import { useForm } from '@tanstack/react-form'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Check, X } from 'lucide-react'

interface NoteEditorProps {
  note: Note
  onCancel: () => void
}

export const NoteEditor: React.FC<NoteEditorProps> = ({ note, onCancel }) => {
  const { mutateAsync, isPending } = useUpdateNote(note.id)

  const form = useForm({
    defaultValues: { title: note.title, content: note.content },
    onSubmit: async ({ value }) => {
      await mutateAsync(value)
      onCancel()
    },
  })

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault()
        e.stopPropagation()
        form.handleSubmit()
      }}
      className="space-y-6 flex flex-col flex-1"
    >
      <div className="flex items-center justify-between gap-4">
        <form.Field
          name="title"
          validators={{ onChange: UpdateNoteSchema.shape.title }}
          children={(field) => (
            <div className="flex-1">
              <Input
                value={field.state.value}
                onBlur={field.handleBlur}
                onChange={(e) => field.handleChange(e.target.value)}
                className="text-3xl font-extrabold border-none px-0 focus-visible:ring-0 shadow-none bg-transparent"
                placeholder="Note Title..."
                autoFocus
              />
            </div>
          )}
        />
        <div className="flex items-center gap-2">
          <Button
            type="button"
            variant="ghost"
            size="icon"
            onClick={onCancel}
            disabled={isPending}
          >
            <X className="h-5 w-5" />
            <span className="sr-only">Cancel</span>
          </Button>
          <Button type="submit" size="sm" disabled={isPending}>
            <Check className="mr-2 h-4 w-4" />
            {isPending ? 'Saving...' : 'Save'}
          </Button>
        </div>
      </div>

      <form.Field
        name="content"
        children={(field) => (
          <Textarea
            value={field.state.value}
            onBlur={field.handleBlur}
            onChange={(e) => field.handleChange(e.target.value)}
            className="flex-1 min-h-100 resize-none text-lg leading-relaxed border-none px-0 focus-visible:ring-0 shadow-none bg-transparent"
            placeholder="Start typing..."
          />
        )}
      />
    </form>
  )
}

import { Badge } from '@/components/ui/badge'

const MAX_VISIBLE = 3

export function NoteTags({ tags }: { tags: string[] }) {
  const visible = tags.slice(0, MAX_VISIBLE)
  const hidden = tags.length - MAX_VISIBLE

  return (
    <div className="flex flex-wrap gap-1">
      {visible.map((tag) => (
        <Badge key={tag} variant="secondary">
          {tag}
        </Badge>
      ))}

      {hidden > 0 && <Badge variant="outline">+{hidden} more</Badge>}
    </div>
  )
}

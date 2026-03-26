import { Input } from '@/components/ui/input'
import { Search } from 'lucide-react'

export function NotesSearch() {
  return (
    <div className="relative max-w-sm">
      <Search className="absolute left-3 top-3 h-4 w-4 text-muted-foreground" />

      <Input placeholder="Search notes..." className="pl-9" />
    </div>
  )
}

export interface Event {
  name: string
  start: string
  end: string
  color?: string
}

export interface Date {
  date: string
  time: string
  year: number
  month: number
  weekday: number
  day: number
  hour: number
  minute: number
  future: boolean
  hasDay: boolean
  hasTime: boolean
  past: boolean
  present: boolean
}

export interface CalendarType {
  name: string
  value: string
}

export interface CalendarRef {
  checkChange: () => void
  prev: () => void
  next: () => void
}

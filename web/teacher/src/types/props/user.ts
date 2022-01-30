export interface Actor {
  name: string
  value: string
}

export interface TableHeader {
  text: string
  value: string
  align?: 'start' | 'center' | 'end'
  sortable?: boolean
}

export interface TableFooter {
  itemsPerPageText: string
  itemsPerPageOptions: number[]
}

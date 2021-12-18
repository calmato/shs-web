export interface Actor {
  name: string
  value: string
}

export interface TableHeader {
  text: string
  value: string
  sortable: boolean
}

export interface TableFooter {
  itemsPerPageText: string
  itemsPerPageOptions: number[]
}

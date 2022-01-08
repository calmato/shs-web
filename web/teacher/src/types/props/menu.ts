import { Role } from '~/types/store'

export interface Menu {
  name: string
  icon: string
  path: string
  filter: Role[] | 'all'
}

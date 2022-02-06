import { Context } from '@nuxt/types'

const excludedPaths: string[] = ['/signin']

export default async ({ route, store, redirect }: Context) => {
  if (excludedPaths.includes(route.path)) {
    return
  }

  await store
    .dispatch('auth/authentication')
    .then(() => {
      Promise.all([
        store.dispatch('auth/showAuth'),
        store.dispatch('lesson/getAllSubjects'),
        store.dispatch('classroom/getTotalRoomsByApi'),
        store.dispatch('classroom/getSchedulesByApi'),
      ])
    })
    .catch(() => {
      redirect('/signin')
    })
}

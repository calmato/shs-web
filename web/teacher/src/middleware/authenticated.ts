import { Context } from '@nuxt/types'

const excludedPaths: string[] = ['/signin']

export default async ({ route, store, redirect }: Context) => {
  if (excludedPaths.includes(route.path)) {
    return
  }

  await store
    .dispatch('auth/authentication')
    .then(async () => {
      await store.dispatch('auth/showAuth')
      await store.dispatch('lesson/getAllSubjects')
    })
    .catch(() => {
      redirect('/signin')
    })
}

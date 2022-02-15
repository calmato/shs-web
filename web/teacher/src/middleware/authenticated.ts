import { Context } from '@nuxt/types'
import { InitialState } from '~/types/store'

const excludedPaths: string[] = ['/signin']

export default async ({ route, store, redirect }: Context) => {
  if (excludedPaths.includes(route.path)) {
    return
  }

  await store
    .dispatch('auth/authentication')
    .then(() => {
      const state = store.getters['common/getInitialState']
      if (state === InitialState.FINISHED) {
        return
      }

      store.dispatch('common/startInitialize')
      Promise.all([store.dispatch('auth/showAuth'), store.dispatch('lesson/getAllSubjects')])
      store.dispatch('common/endInitialize')
    })
    .catch(() => {
      redirect('/signin')
    })
}

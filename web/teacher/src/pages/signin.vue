<template>
  <the-sign-in :message="message" @click="handleClick" />
</template>

<script lang="ts">
import { computed, defineComponent, useAsync, SetupContext } from '@nuxtjs/composition-api'
import { UserStore, CommonStore } from '~/store'
import TheSignIn from '~/components/templates/TheSignIn.vue'

export default defineComponent({
  components: {
    TheSignIn,
  },

  setup(_, { root }: SetupContext) {
    const store = root.$store
    const router = root.$router

    const message = computed(() => store.getters['user/getMessage'])

    useAsync(async () => {
      await hello()
    })

    async function hello(): Promise<void> {
      CommonStore.startConnection()
      await UserStore.hello()
        .catch((err: Error) => {
          console.log('failed to hello', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClick = () => {
      router.push('/')
    }

    return {
      message,
      handleClick,
    }
  },
})
</script>

<template>
  <the-sign-in :form.sync="form" :loading="loading" @click="handleClick" />
</template>

<script lang="ts">
import { defineComponent, SetupContext, reactive, computed } from '@nuxtjs/composition-api'
import { AuthStore, CommonStore } from '~/store'
import TheSignIn from '~/components/templates/TheSignIn.vue'
import { SignInForm } from '~/types/form'
import { PromiseState } from '~/types/store'

export default defineComponent({
  components: {
    TheSignIn,
  },

  setup(_, { root }: SetupContext) {
    const store = root.$store
    const router = root.$router

    const form = reactive<SignInForm>({
      mail: '',
      password: '',
    })

    const loading = computed<boolean>(() => {
      return store.getters['common/getPromiseState'] === PromiseState.LOADING
    })

    const handleClick = async () => {
      CommonStore.startConnection()
      await AuthStore.signIn(form)
        .then(() => {
          router.push('/')
        })
        .catch((err: Error) => {
          console.log('failed to sign in', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    return {
      form,
      loading,
      handleClick,
    }
  },
})
</script>

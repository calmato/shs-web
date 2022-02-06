<template>
  <the-student-new :form="form" :loading="loading" @submit="handleSubmit" @cancel="handleCancel" />
</template>

<script lang="ts">
import { computed, defineComponent, reactive, useRouter, useStore } from '@nuxtjs/composition-api'
import TheStudentNew from '~/components/templates/TheStudentNew.vue'
import { CommonStore, UserStore } from '~/store'
import { StudentNewForm, StudentNewOptions, StudentNewParams } from '~/types/form'
import { PromiseState } from '~/types/store'

export default defineComponent({
  components: {
    TheStudentNew,
  },

  setup() {
    const router = useRouter()
    const store = useStore()

    const form = reactive<StudentNewForm>({
      params: { ...StudentNewParams },
      options: { ...StudentNewOptions },
    })

    const loading = computed<boolean>(() => {
      return store.getters['common/getPromiseState'] === PromiseState.LOADING
    })

    const handleSubmit = async (): Promise<void> => {
      CommonStore.startConnection()
      await UserStore.createStudent({ form })
        .then(() => {
          router.push('/users')
          CommonStore.showSnackbar({ color: 'success', message: '生徒を新規登録しました。' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleCancel = (): void => {
      router.back()
    }

    return {
      form,
      loading,
      handleSubmit,
      handleCancel,
    }
  },
})
</script>

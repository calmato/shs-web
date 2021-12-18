<template>
  <the-teacher-new :form="form" :loading="loading" @submit="handleSubmit" @cancel="handleCancel" />
</template>

<script lang="ts">
import { computed, defineComponent, reactive, SetupContext } from '@nuxtjs/composition-api'
import TheTeacherNew from '~/components/templates/TheTeacherNew.vue'
import { CommonStore, UserStore } from '~/store'
import { TeacherNewForm, TeacherNewOptions, TeacherNewParams } from '~/types/form'
import { PromiseState } from '~/types/store'

export default defineComponent({
  components: {
    TheTeacherNew,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store

    const form = reactive<TeacherNewForm>({
      params: { ...TeacherNewParams },
      options: { ...TeacherNewOptions },
    })

    const loading = computed<boolean>(() => {
      return store.getters['common/getPromiseState'] === PromiseState.LOADING
    })

    const handleSubmit = async (): Promise<void> => {
      CommonStore.startConnection()
      await UserStore.createTeacher({ form })
        .then(() => {
          router.push('/users')
        })
        .catch((err: Error) => {
          console.log('failed to create teacher', err)
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

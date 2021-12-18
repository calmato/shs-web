<template>
  <the-user-top :loading="loading" :teachers="teachers" :students="students" />
</template>

<script lang="ts">
import { computed, defineComponent, ref, SetupContext, useAsync } from '@nuxtjs/composition-api'
import TheUserTop from '~/components/templates/TheUserTop.vue'
import { CommonStore, UserStore } from '~/store'

export default defineComponent({
  components: {
    TheUserTop,
  },

  setup(_, { root }: SetupContext) {
    const store = root.$store

    const loading = ref<boolean>(false)

    const teachers = computed(() => store.getters['user/getTeachers'])
    const students = computed(() => store.getters['user/getStudents'])

    useAsync(async () => {
      CommonStore.startConnection()
      await UserStore.listTeachers()
        .catch((err: Error) => {
          console.log('feiled to list teachers', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    })

    return {
      loading,
      teachers,
      students,
    }
  },
})
</script>

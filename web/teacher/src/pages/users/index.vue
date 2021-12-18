<template>
  <the-user-top
    :loading="loading"
    :students="students"
    :teachers="teachers"
    :teachers-total="teachersTotal"
    :teachers-page.sync="teachersPage"
    :teachers-items-per-page.sync="teachersItemsPerPage"
    @click:new="handleClickNew"
  />
</template>

<script lang="ts">
import { computed, defineComponent, ref, SetupContext, useAsync, watch } from '@nuxtjs/composition-api'
import TheUserTop from '~/components/templates/TheUserTop.vue'
import { CommonStore, UserStore } from '~/store'
import { PromiseState } from '~/types/store'

export default defineComponent({
  components: {
    TheUserTop,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store

    const teachersPage = ref<number>(1)
    const teachersItemsPerPage = ref<number>(10)

    const loading = computed(() => store.getters['common/getPromiseState'] === PromiseState.LOADING)
    const students = computed(() => store.getters['user/getStudents'])
    const teachers = computed(() => store.getters['user/getTeachers'])
    const teachersTotal = computed(() => store.getters['user/getTeachersTotal'])

    watch(teachersPage, async () => {
      await listTeachers()
    })

    watch(teachersItemsPerPage, async () => {
      await listTeachers()
    })

    useAsync(async () => {
      await listTeachers()
    })

    async function listTeachers(): Promise<void> {
      CommonStore.startConnection()

      const limit: number = teachersItemsPerPage.value
      const offset: number = (teachersPage.value - 1) * limit

      await UserStore.listTeachers({ limit, offset })
        .catch((err: Error) => {
          console.log('feiled to list teachers', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickNew = (actor: string): void => {
      router.push(`/users/${actor}/new`)
    }

    return {
      loading,
      students,
      teachers,
      teachersTotal,
      teachersPage,
      teachersItemsPerPage,
      handleClickNew,
    }
  },
})
</script>

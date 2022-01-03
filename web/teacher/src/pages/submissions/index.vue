<template>
  <the-submission-top :summaries="summaries" @click="handleClick" />
</template>

<script lang="ts">
import { computed, defineComponent, SetupContext, useAsync } from '@nuxtjs/composition-api'
import TheSubmissionTop from '~/components/templates/TheSubmissionTop.vue'
import { CommonStore, SubmissionStore } from '~/store'
import { TeacherShiftSummary } from '~/types/store'

export default defineComponent({
  components: {
    TheSubmissionTop,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store

    const teacherId = computed<string>(() => store.getters['auth/getUid'])
    const summaries = computed<TeacherShiftSummary[]>(() => store.getters['submission/getSummaries'])

    useAsync(async () => {
      await listTeacherSubmissions()
    })

    async function listTeacherSubmissions(): Promise<void> {
      CommonStore.startConnection()

      await SubmissionStore.listTeacherSubmissions({ teacherId: teacherId.value })
        .catch((err: Error) => {
          console.log('failed to list teacher submissions', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClick = (summary: TeacherShiftSummary): void => {
      router.push(`/submissions/${summary.id}`)
    }

    return {
      summaries,
      handleClick,
    }
  },
})
</script>

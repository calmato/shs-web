<template>
  <the-submission-top :summaries="summaries" @click:show="handleClickShow" @click:top="handleClickTop" />
</template>

<script lang="ts">
import { computed, defineComponent, useAsync, useRouter, useStore } from '@nuxtjs/composition-api'
import TheSubmissionTop from '~/components/templates/TheSubmissionTop.vue'
import { CommonStore, SubmissionStore } from '~/store'
import { SubmissionSummary } from '~/types/store'

export default defineComponent({
  components: {
    TheSubmissionTop,
  },

  setup() {
    const store = useStore()
    const router = useRouter()

    const summaries = computed<SubmissionSummary[]>(() => store.getters['submission/getSummaries'])

    useAsync(async () => {
      await listStudentSubmissions()
    })

    async function listStudentSubmissions(): Promise<void> {
      CommonStore.startConnection()

      await SubmissionStore.listStudentSubmissions()
        .catch((err: Error) => {
          console.log('failed to list teacher submissions', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickShow = (summary: SubmissionSummary): void => {
      router.push(`/submissions/${summary.id}`)
    }

    const handleClickTop = (): void => {
      router.push('/')
    }

    return {
      summaries,
      handleClickShow,
      handleClickTop,
    }
  },
})
</script>

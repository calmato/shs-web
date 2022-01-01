<template>
  <div>
    <h2>授業登録画面</h2>
    <ul>
      <li>summary: {{ summary }}</li>
      <li>details: {{ details }}</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, SetupContext, useAsync } from '@nuxtjs/composition-api'
import { CommonStore, ShiftStore } from '~/store'
import { ShiftDetail, ShiftSummary } from '~/types/store'

export default defineComponent({
  setup(_, { root }: SetupContext) {
    const route = root.$route
    const store = root.$store

    const summary = computed<ShiftSummary>(() => store.getters['shift/getSummary'])
    const details = computed<ShiftDetail[]>(() => store.getters['shift/getDetails'])

    useAsync(async () => {
      await listShiftDetails()
    })

    async function listShiftDetails(): Promise<void> {
      CommonStore.startConnection()

      const summaryId: string = route.params.id

      await ShiftStore.listShiftDetails({ summaryId: Number(summaryId) })
        .catch((err: Error) => {
          console.log('feiled to list shift details', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    return {
      summary,
      details,
    }
  },
})
</script>

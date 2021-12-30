<template>
  <div>
    <!-- PCレイアウト -->
    <pc-shift-top
      class="hidden-sm-and-down"
      :accepting-summaries="getAcceptingSummaries()"
      :finished-summaries="getFinishedSummaries()"
      :waiting-summaries="getWaitingSummaries()"
      @click:new-lesson="handleClickNewLesson"
      @click:new-shift="handleClickNewShift"
      @click:edit-shift="handleClickEditShift"
    />
    <!-- スマホレイアウト -->
    <mb-shift-top class="hidden-md-and-up" @click="handleClickTop" />
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, SetupContext, useAsync } from '@nuxtjs/composition-api'
import MbShiftTop from '~/components/templates/MbShiftTop.vue'
import PcShiftTop from '~/components/templates/PcShiftTop.vue'
import { CommonStore, ShiftStore } from '~/store'
import { ShiftStatus, ShiftSummary } from '~/types/store'

export default defineComponent({
  components: {
    MbShiftTop,
    PcShiftTop,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store

    const summaries = computed<ShiftSummary[]>(() => store.getters['shift/getSummaries'])

    useAsync(async () => {
      await listShiftSummaries()
    })

    async function listShiftSummaries(): Promise<void> {
      CommonStore.startConnection()

      const limit: number = 0
      const offset: number = 0
      const status: ShiftStatus = ShiftStatus.UNKNOWN

      await ShiftStore.listShiftSummaries({ limit, offset, status })
        .catch((err: Error) => {
          console.log('feiled to list shift summaries', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const getWaitingSummaries = (): ShiftSummary[] => {
      return summaries.value.filter((summary: ShiftSummary) => {
        return summary.status === ShiftStatus.WAITING
      })
    }

    const getAcceptingSummaries = (): ShiftSummary[] => {
      return summaries.value.filter((summary: ShiftSummary) => {
        return summary.status === ShiftStatus.ACCEPTING
      })
    }

    const getFinishedSummaries = (): ShiftSummary[] => {
      return summaries.value.filter((summary: ShiftSummary) => {
        return summary.status === ShiftStatus.FINISHED
      })
    }

    const handleClickNewShift = (): void => {
      console.log('click', 'new', 'shift')
    }

    const handleClickEditShift = (shift: ShiftSummary): void => {
      console.log('click', 'edit', 'sihft', shift)
    }

    const handleClickNewLesson = (shift: ShiftSummary): void => {
      console.log('click', 'new', 'lesson', shift)
    }

    const handleClickTop = (): void => {
      router.push('/')
    }

    return {
      getAcceptingSummaries,
      getFinishedSummaries,
      getWaitingSummaries,
      handleClickNewShift,
      handleClickEditShift,
      handleClickNewLesson,
      handleClickTop,
    }
  },
})
</script>

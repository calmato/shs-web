<template>
  <pc-shift-detail :summary="summary" :details="details" :rooms="rooms" :teachers="teachers" />
</template>

<script lang="ts">
import { computed, defineComponent, SetupContext, useAsync } from '@nuxtjs/composition-api'
import PcShiftDetail from '~/components/templates/PcShiftDetail.vue'
import { CommonStore, ShiftStore } from '~/store'
import { ShiftDetail, ShiftSummary, TeacherShift } from '~/types/store'

export default defineComponent({
  components: {
    PcShiftDetail,
  },

  setup(_, { root }: SetupContext) {
    const route = root.$route
    const store = root.$store

    const summary = computed<ShiftSummary>(() => store.getters['shift/getSummary'])
    const details = computed<ShiftDetail[]>(() => store.getters['shift/getDetails'])
    const rooms = computed<number>(() => store.getters['shift/getRooms'])
    const teachers = computed<TeacherShift[]>(() => store.getters['shift/getTeachers'])

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
      rooms,
      teachers,
    }
  },
})
</script>

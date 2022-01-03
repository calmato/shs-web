<template>
  <div>
    <h2>シフト登録画面</h2>
    <ul>
      <li>summary: {{ summary }}</li>
      <li>shifts: {{ shifts }}</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, SetupContext, useAsync } from '@nuxtjs/composition-api'
import { CommonStore, SubmissionStore } from '~/store'
import { TeacherShiftDetail, TeacherShiftSummary } from '~/types/store'

export default defineComponent({
  setup(_, { root }: SetupContext) {
    const route = root.$route
    const store = root.$store

    const teacherId = computed<string>(() => store.getters['auth/getUid'])
    const summary = computed<TeacherShiftSummary>(() => store.getters['submission/getSummary'])
    const shifts = computed<TeacherShiftDetail[]>(() => store.getters['submission/getShifts'])

    useAsync(async () => {
      await listTeacherShifts()
    })

    async function listTeacherShifts(): Promise<void> {
      CommonStore.startConnection()

      const shiftId: number = route.params.id

      await SubmissionStore.listTeacherShifts({ teacherId: teacherId.value, shiftId })
        .catch((err: Error) => {
          console.log('failed to list teacher shifts', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    return {
      summary,
      shifts,
    }
  },
})
</script>

<template>
  <the-submission-edit />
</template>

<script lang="ts">
import { computed, defineComponent, SetupContext, useAsync } from '@nuxtjs/composition-api'
import TheSubmissionEdit from '~/components/templates/TheSubmissionEdit.vue'
import { CommonStore, SubmissionStore } from '~/store'
import { TeacherShiftDetail, TeacherShiftSummary } from '~/types/store'

export default defineComponent({
  components: {
    TheSubmissionEdit,
  },

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

      const shiftId = Number(route.params.id)

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

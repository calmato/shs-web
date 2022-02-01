<template>
  <the-submission-detail
    :loading="loading"
    :summary="summary"
    :shifts="shifts"
    :enabled-lesson-ids="enabledLessonIds"
    @click:change-items="handleClickChangeEnabled"
    @click:submit="handleClickSubmit"
  />
</template>

<script lang="ts">
import { computed, defineComponent, ref, useAsync, useRoute, useRouter, useStore } from '@nuxtjs/composition-api'
import TheSubmissionDetail from '~/components/templates/TheSubmissionDetail.vue'
import { CommonStore, SubmissionStore } from '~/store'
import { PromiseState, SubmissionDetail, SubmissionDetailLesson, SubmissionSummary } from '~/types/store'

export default defineComponent({
  components: {
    TheSubmissionDetail,
  },

  setup() {
    const router = useRouter()
    const route = useRoute()
    const store = useStore()

    const enabledLessonIds = ref<number[]>([])

    const teacherId = computed<string>(() => store.getters['auth/getUid'])
    const summary = computed<SubmissionSummary>(() => store.getters['submission/getSummary'])
    const shifts = computed<SubmissionDetail[]>(() => store.getters['submission/getShifts'])
    const loading = computed<boolean>(() => {
      return store.getters['common/getPromiseState'] === PromiseState.LOADING
    })

    useAsync(async () => {
      await listTeacherShifts()
    })

    async function listTeacherShifts(): Promise<void> {
      CommonStore.startConnection()

      const shiftId = Number(route.value.params.id)

      await SubmissionStore.listTeacherShifts({ teacherId: teacherId.value, shiftId })
        .then(() => {
          shifts.value.forEach((shift: SubmissionDetail): void => {
            shift.lessons.forEach((lesson: SubmissionDetailLesson): void => {
              if (lesson.enabled) {
                enabledLessonIds.value.push(lesson.id)
              }
            })
          })
        })
        .catch((err: Error) => {
          console.log('failed to list teacher shifts', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickChangeEnabled = (lessonId: number): void => {
      const index: number = enabledLessonIds.value.indexOf(lessonId)
      if (index === -1) {
        enabledLessonIds.value.push(lessonId)
      } else {
        enabledLessonIds.value.splice(index, 1)
      }
    }

    const handleClickSubmit = async (): Promise<void> => {
      CommonStore.startConnection()

      const shiftId = Number(route.value.params.id)

      await SubmissionStore.submitTeacherShifts({
        teacherId: teacherId.value,
        shiftId,
        lessonIds: enabledLessonIds.value,
      })
        .then(() => {
          router.push('/submissions')
          CommonStore.showSnackbar({ color: 'success', message: 'シフト希望を提出しました' })
        })
        .catch((err: Error) => {
          console.log('failed to submit teacher shifts', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    return {
      loading,
      enabledLessonIds,
      summary,
      shifts,
      handleClickChangeEnabled,
      handleClickSubmit,
    }
  },
})
</script>

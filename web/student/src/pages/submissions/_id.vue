<template>
  <the-submission-detail
    :loading="loading"
    :summary="summary"
    :shifts="shifts"
    :lessons="suggestedLessons"
    :subjects="student.subjects"
    :enabled-lesson-ids="enabledLessonIds"
    @click:change-items="handleClickChangeEnabled"
    @click:submit="handleClickSubmit"
    @click:add-lesson="handleClickAddSuggestedLesson"
    @click:remove-lesson="handleClickRemoveSuggestedLesson"
  />
</template>

<script lang="ts">
import {
  computed,
  defineComponent,
  reactive,
  ref,
  useAsync,
  useRoute,
  useRouter,
  useStore,
} from '@nuxtjs/composition-api'
import TheSubmissionDetail from '~/components/templates/TheSubmissionDetail.vue'
import { CommonStore, SubmissionStore } from '~/store'
import { ISubmissionSuggestedLesson } from '~/types/form'
import {
  Auth,
  PromiseState,
  SubmissionDetail,
  SubmissionDetailLesson,
  SubmissionLesson,
  SubmissionSummary,
} from '~/types/store'

export default defineComponent({
  components: {
    TheSubmissionDetail,
  },

  setup() {
    const router = useRouter()
    const route = useRoute()
    const store = useStore()

    const enabledLessonIds = ref<number[]>([])
    const suggestedLessons = reactive<ISubmissionSuggestedLesson>([])

    const student = computed<Auth>(() => store.getters['auth/getAuth'])
    const summary = computed<SubmissionSummary>(() => store.getters['submission/getSummary'])
    const shifts = computed<SubmissionDetail[]>(() => store.getters['submission/getShifts'])
    const lessons = computed<SubmissionLesson[]>(() => store.getters['submission/getLessons'])
    const loading = computed<boolean>(() => {
      return store.getters['common/getPromiseState'] === PromiseState.LOADING
    })

    useAsync(async () => {
      await listStudentLessons()
    })

    async function listStudentLessons(): Promise<void> {
      CommonStore.startConnection()

      const lessonId = Number(route.value.params.id)

      await SubmissionStore.listStudentLessons({ lessonId })
        .then(() => {
          shifts.value.forEach((shift: SubmissionDetail): void => {
            shift.lessons.forEach((lesson: SubmissionDetailLesson): void => {
              if (lesson.enabled) enabledLessonIds.value.push(lesson.id)
            })
          })
          lessons.value.forEach((lesson: SubmissionDetail): void => {
            suggestedLessons.push({ subjectId: lesson.subjectId, total: String(lesson.total) })
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
      const lessons: SubmissionLesson[] = suggestedLessons.map(
        (lesson: ISubmissionSuggestedLesson): Submissionlesson => {
          return { subjectId: lesson.subjectId, total: Number(lesson.total) }
        }
      )

      await SubmissionStore.submitStudentShifts({
        shiftId,
        lessons,
        lessonIds: enabledLessonIds.value,
      })
        .then(() => {
          router.push('/submissions')
          CommonStore.showSnackbar({ color: 'success', message: 'シフト希望を提出しました' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickAddSuggestedLesson = (): void => {
      suggestedLessons.push({ subjectId: 0, total: 0 })
    }

    const handleClickRemoveSuggestedLesson = (index: number): void => {
      suggestedLessons.splice(index, 1)
    }

    return {
      student,
      loading,
      enabledLessonIds,
      suggestedLessons,
      summary,
      shifts,
      lessons,
      handleClickChangeEnabled,
      handleClickSubmit,
      handleClickAddSuggestedLesson,
      handleClickRemoveSuggestedLesson,
    }
  },
})
</script>

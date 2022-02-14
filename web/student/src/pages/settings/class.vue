<template>
  <the-setting-class
    :user="student"
    :subjects="student.subjects"
    :schedules="schedulesForm"
    :lessons="suggestedLessons"
    @click:add-lesson="handleClickAddSuggestedLesson"
    @click:remove-lesson="handleClickRemoveSuggestedLesson"
    @click:submit="handleClickSubmit"
  />
</template>

<script lang="ts">
import { computed, defineComponent, reactive, useAsync, useRouter, useStore } from '@nuxtjs/composition-api'
import TheSettingClass from '~/components/templates/TheSettingClass.vue'
import { CommonStore, SubmissionStore } from '~/store'
import { ISubmissionSuggestedLesson } from '~/types/form'
import { Auth, SubmissionLesson, SubmissionTemplate } from '~/types/store'

export default defineComponent({
  components: {
    TheSettingClass,
  },

  setup() {
    const store = useStore()
    const router = useRouter()

    const schedulesForm = reactive<SubmissionTemplate[]>([])
    const suggestedLessons = reactive<ISubmissionSuggestedLesson[]>([])

    const student = computed<Auth>(() => store.getters['auth/getAuth'])
    const templates = computed<SubmissionTemplate[]>(() => store.getters['submission/getTemplates'])
    const lessons = computed<SubmissionLesson[]>(() => store.getters['submission/getLessons'])

    useAsync(async () => {
      await getTemplate()
    })

    async function getTemplate(): Promise<void> {
      CommonStore.startConnection()

      await SubmissionStore.getSubmissionTemplate()
        .then(() => {
          templates.value.forEach((template: SubmissionTemplate): void => {
            const lessons: SubmissionLesson[] = []
            template.lessons.forEach((lesson: SubmissionTemplateLesson): void => lessons.push({ ...lesson }))
            schedulesForm.push({ weekday: template.weekday, lessons })
          })
          lessons.value.forEach((lesson: SubmissionLesson): void => {
            suggestedLessons.push({ subjectId: lesson.subjectId, total: String(lesson.total) })
          })
          if (suggestedLessons.length === 0) handleClickAddSuggestedLesson()
        })
        .catch((err: Error) => {
          console.log('failed to get submission template', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickAddSuggestedLesson = (): void => {
      suggestedLessons.push({ subjectId: 0, total: '0' })
    }

    const handleClickRemoveSuggestedLesson = (index: number): void => {
      suggestedLessons.splice(index, 1)
    }

    const handleClickSubmit = async (): Promise<void> => {
      CommonStore.startConnection()

      const lessons: SubmissionLesson[] = suggestedLessons.map(
        (lesson: ISubmissionSuggestedLesson): SubmissionLesson => {
          return { subjectId: lesson.subjectId, total: Number(lesson.total) }
        }
      )

      await SubmissionStore.upsertSubmissionTemplate({ schedules: schedulesForm, lessons })
        .then(() => {
          router.push('/settings')
          CommonStore.showSnackbar({ color: 'success', message: '設定を更新しました' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    return {
      student,
      schedulesForm,
      suggestedLessons,
      handleClickAddSuggestedLesson,
      handleClickRemoveSuggestedLesson,
      handleClickSubmit,
    }
  },
})
</script>

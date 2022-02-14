<template>
  <v-container>
    <p class="text-h6 my-2">授業希望のカスタム設定</p>
    <the-submission-lesson-list
      :subjects="subjects"
      :lessons="lessons"
      @click:add-lesson="onClickAddLesson"
      @click:remove-lesson="onClickRemoveLesson"
    />
    <the-submission-schedule-list :schedules="schedules" />
    <v-container>
      <v-row class="justify-end">
        <v-btn color="primary" class="right mt-4" large :disabled="loading" @click="onClickSubmit">保存</v-btn>
      </v-row>
    </v-container>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, SetupContext } from '@vue/composition-api'
import TheSubmissionLessonList from '~/components/organisms/TheSubmissionLessonList.vue'
import TheSubmissionScheduleList from '~/components/organisms/TheSubmissionScheduleList.vue'
import { UserProps } from '~/types/props/setting'
import { ISubmissionSuggestedLesson } from '~/types/form'
import { Subject, SubmissionTemplate } from '~/types/store'

export default defineComponent({
  components: {
    TheSubmissionLessonList,
    TheSubmissionScheduleList,
  },

  props: {
    loading: {
      type: Boolean,
      default: false,
    },
    user: {
      type: Object as PropType<UserProps>,
      default: () => {},
    },
    subjects: {
      type: Array as PropType<Subject[]>,
      default: () => [],
    },
    schedules: {
      type: Array as PropType<SubmissionTemplate[]>,
      default: () => [],
    },
    lessons: {
      type: Array as PropType<ISubmissionSuggestedLesson[]>,
      default: () => [],
    },
  },

  setup(_, { emit }: SetupContext) {
    const onClickAddLesson = (): void => {
      emit('click:add-lesson')
    }

    const onClickRemoveLesson = (index: number): void => {
      emit('click:remove-lesson', index)
    }

    const onClickSubmit = (): void => {
      emit('click:submit')
    }

    return {
      onClickAddLesson,
      onClickRemoveLesson,
      onClickSubmit,
    }
  },
})
</script>

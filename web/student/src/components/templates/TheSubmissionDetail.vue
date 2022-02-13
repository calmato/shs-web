<template>
  <v-container>
    <div class="text-h6 my-2">{{ summary.year }}年 {{ summary.month }}月</div>
    <the-submission-lesson-list
      :subjects="subjects"
      :lessons="lessons"
      @click:add-lesson="onClickAddLesson"
      @click:remove-lesson="onClickRemoveLesson"
    />
    <v-divider />
    <v-container class="d-flex">
      <v-row>
        <v-col cols="4">授業日</v-col>
        <v-col cols="8">授業希望</v-col>
      </v-row>
    </v-container>
    <v-divider />
    <the-submission-shift-list
      :loading="loading"
      :shifts="shifts"
      :selected-items="enabledLessonIds"
      @click:change-items="onChangeItems"
      @click:submit="onClickSave"
    />
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheSubmissionLessonList from '~/components/organisms/TheSubmissionLessonList.vue'
import TheSubmissionShiftList from '~/components/organisms/TheSubmissionShiftList.vue'
import { ISubmissionSuggestedLesson } from '~/types/form'
import { Subject, SubmissionDetail, SubmissionSummary } from '~/types/store'

export default defineComponent({
  components: {
    TheSubmissionLessonList,
    TheSubmissionShiftList,
  },

  props: {
    loading: {
      type: Boolean,
      default: false,
    },
    summary: {
      type: Object as PropType<SubmissionSummary>,
      default: () => {},
    },
    shifts: {
      type: Array as PropType<SubmissionDetail[]>,
      default: () => [],
    },
    lessons: {
      type: Array as PropType<ISubmissionSuggestedLesson[]>,
      default: () => [],
    },
    subjects: {
      type: Array as ProptType<Subject[]>,
      default: () => [],
    },
    enabledLessonIds: {
      type: Array as PropType<number[]>,
      default: () => [],
    },
  },

  setup(_, { emit }: SetupContext) {
    const onChangeItems = (lessonId: number): void => {
      emit('click:change-items', lessonId)
    }

    const onClickSave = (): void => {
      emit('click:submit')
    }

    const onClickAddLesson = (): void => {
      emit('click:add-lesson')
    }

    const onClickRemoveLesson = (index: number): void => {
      emit('click:remove-lesson', index)
    }

    return {
      onChangeItems,
      onClickSave,
      onClickAddLesson,
      onClickRemoveLesson,
    }
  },
})
</script>

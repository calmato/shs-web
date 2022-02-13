<template>
  <v-container>
    <div class="d-flex align-center my-2">
      <p class="text-subtitle-1 mr-4 mb-0">科目とコマ数を指定</p>
      <v-btn v-show="isAddable()" class="suffix-btn" color="primary" fab elevation="0" right>
        <v-icon @click="onClickAddLesson">mdi-plus</v-icon>
      </v-btn>
    </div>
    <div v-for="(lesson, index) in lessons" :key="index" class="d-flex h-stack align-center">
      <the-select-with-slot
        label="科目"
        :value.sync="lesson.subjectId"
        :items="getSubjects(lesson.subjectId)"
        item-text="fullname"
        item-value="id"
        chips
      >
        <template #default="{ item }">
          <v-chip label :color="item.color">{{ item.fullname }}</v-chip>
        </template>
      </the-select-with-slot>
      <the-text-field type="number" label="コマ数" :value.sync="lesson.total" height="42" />
      <v-btn class="suffix-btn" color="error" fab elevation="0" right>
        <v-icon @click="onClickRemoveLesson(index)">mdi-minus</v-icon>
      </v-btn>
    </div>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheTextField from '~/components/atoms/TheTextField.vue'
import TheSelectWithSlot from '~/components/atoms/TheSelectWithSlot.vue'
import { ISubmissionSuggestedLesson } from '~/types/form'
import { Subject, SubmissionLesson } from '~/types/store'

export default defineComponent({
  components: {
    TheTextField,
    TheSelectWithSlot,
  },

  props: {
    lessons: {
      type: Array as PropType<ISubmissionSuggestedLesson[]>,
      default: () => [],
    },
    subjects: {
      type: Array as ProptType<Subject[]>,
      default: () => [],
    },
  },

  setup(props, { emit }: SetupContext) {
    const getSubjects = (select: number): void => {
      const items: Subject[] = []
      props.subjects.forEach((subject: Subject): void => {
        const index = props.lessons.findIndex((val: SubmissionLesson): boolean => {
          return val.subjectId !== select && val.subjectId === subject.id
        })
        if (index === -1) items.push(subject)
      })
      return items
    }

    const isAddable = (): boolean => {
      return props.lessons?.length < props.subjects?.length
    }

    const onClickAddLesson = (): void => {
      emit('click:add-lesson')
    }

    const onClickRemoveLesson = (index: number): void => {
      emit('click:remove-lesson', index)
    }

    return {
      getSubjects,
      isAddable,
      onClickAddLesson,
      onClickRemoveLesson,
    }
  },
})
</script>

<style lang="scss" scoped>
.suffix-btn {
  height: 24px;
  width: 24px;
}

.h-stack {
  gap: 1rem;
}
</style>

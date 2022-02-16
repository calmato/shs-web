<template>
  <table>
    <tr>
      <th class="fixed">生徒名</th>
      <td v-for="student in students" :key="student.id">
        <a
          class="text-decoration-underline"
          :class="[student.isSubmit ? 'black--text' : 'red--text']"
          @click="onClickSubmissions(student)"
        >
          {{ student.name }}
        </a>
      </td>
      <th class="fixed-right secondary--text text--accent-4">合計</th>
    </tr>
    <tr>
      <th class="fixed">登録授業数</th>
      <td v-for="student in students" :key="student.id">
        <a class="black--text text-decoration-underline" @click="onClickLessons(student)">
          {{ student.lessonTotal }}
        </a>
      </td>
      <td class="fixed-right">{{ getLessonTotal() }}</td>
    </tr>
    <tr>
      <th class="fixed">残り授業数</th>
      <td
        v-for="student in students"
        :key="student.id"
        :class="[getRemainingLesson(student) > 0 ? 'secondary--text text--accent-4' : 'black--text']"
      >
        {{ getRemainingLesson(student) }}
      </td>
      <td :class="[getRemainingLessonTotal() > 0 ? 'secondary--text text--accent-4' : '']" class="fixed-right">
        {{ getRemainingLessonTotal() }}
      </td>
    </tr>
  </table>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { StudentShift } from '~/types/store'

export default defineComponent({
  props: {
    students: {
      type: Array as PropType<StudentShift[]>,
      default: () => [],
    },
  },

  setup(props, { emit }: SetupContext) {
    const onClickSubmissions = (student: StudentShift): void => {
      emit('click:show-submissions', student.id)
    }

    const onClickLessons = (student: StudentShift): void => {
      emit('click:show-lessons', student.id)
    }

    const getRemainingLesson = (student: StudentShift): number => {
      return student.suggestedLessonsTotal - student.lessonTotal
    }

    const getLessonTotal = (): number => {
      let total: number = 0
      props.students.forEach((student: StudentShift) => {
        total += student.lessonTotal
      })
      return total
    }

    const getRemainingLessonTotal = (): number => {
      let total: number = 0
      props.students.forEach((student: StudentShift) => {
        total += getRemainingLesson(student)
      })
      return total
    }

    return {
      onClickSubmissions,
      onClickLessons,
      getLessonTotal,
      getRemainingLesson,
      getRemainingLessonTotal,
    }
  },
})
</script>

<style scoped>
table {
  display: block;
  overflow-x: scroll;
  white-space: nowrap;
  -webkit-overflow-scrolling: touch;
  border-collapse: collapse;
  border-spacing: 0;
  width: 100%;
}

th,
td {
  vertical-align: middle;
  padding: 4px 16px;
  border: 1px solid #e5e5e5;
  text-align: center;
}

.fixed {
  position: sticky;
  left: 0;
  background-color: #f5f5f5;
}
.fixed:before {
  content: '';
  position: absolute;
  top: 0px;
  left: -1px;
  width: 100%;
  height: 100%;
  border: 1px solid #f5f5f5;
}

.fixed-right {
  position: sticky;
  right: -1px;
  background-color: #f5f5f5;
}
.fixed:before {
  content: '';
  position: absolute;
  top: 0px;
  right: -1px;
  width: 100%;
  height: 100%;
  border: 1px solid #f5f5f5;
}
</style>

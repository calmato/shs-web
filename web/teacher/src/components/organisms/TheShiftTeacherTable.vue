<template>
  <table>
    <tr>
      <th class="fixed">講師名</th>
      <td v-for="teacher in teachers" :key="teacher.id">
        <a
          class="text-decoration-underline"
          :class="[teacher.isSubmit ? 'black--text' : 'red--text']"
          @click="onClickSubmissions(teacher)"
        >
          {{ teacher.name }}
        </a>
      </td>
      <th class="fixed-right info--text text--lighten-1 text-decoration-underline">合計</th>
    </tr>
    <tr>
      <th class="fixed">担当授業数</th>
      <td v-for="teacher in teachers" :key="teacher.id">
        <a class="black--text text-decoration-underline" @click="onClickLessons(teacher)">
          {{ teacher.lessonTotal }}
        </a>
      </td>
      <td class="fixed-right">{{ getLessonTotal() }}</td>
    </tr>
  </table>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { TeacherShift } from '~/types/store'

export default defineComponent({
  props: {
    teachers: {
      type: Array as PropType<TeacherShift[]>,
      default: () => [],
    },
  },

  setup(props, { emit }: SetupContext) {
    const getLessonTotal = (): number => {
      let total: number = 0
      props.teachers.forEach((teacher: TeacherShift) => {
        total += teacher.lessonTotal
      })
      return total
    }

    const onClickSubmissions = (teacher: TeacherShift): void => {
      emit('click:show-submissions', teacher.id)
    }

    const onClickLessons = (teacher: TeacherShift): void => {
      emit('click:show-lessons', teacher.id)
    }

    return {
      getLessonTotal,
      onClickSubmissions,
      onClickLessons,
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

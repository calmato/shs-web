<template>
  <v-card>
    <v-toolbar color="primary" dark>授業登録</v-toolbar>
    <v-card-text class="py-4">
      <v-row class="py-2">
        <v-col cols="3" class="text-center">日にち</v-col>
        <v-col cols="3" class="text-center">生徒</v-col>
        <v-col cols="3" class="text-center">科目</v-col>
        <v-col cols="3" class="text-center">講師</v-col>
      </v-row>
      <v-divider />
      <v-row class="d-flex py-2">
        <v-col cols="3" align="center" class="mt-3">{{ getDay() }}</v-col>
        <v-col cols="3" align="center">
          <v-chip-group rounded color="success" tag="v-row" column :value="selectedStudent" @change="onClickStudent">
            <v-chip
              v-for="student in lesson.students"
              :key="student.id"
              :value="student.id"
              class="col col-12 justify-center"
            >
              {{ student.name }}
            </v-chip>
          </v-chip-group>
        </v-col>
        <v-col cols="3" align="center">
          <v-chip-group rounded color="success" tag="v-row" column :value="selectedSubject" @change="onClickSubject">
            <v-chip
              v-for="subject in getSubjects()"
              :key="subject.id"
              :value="subject.id"
              class="col col-12 justify-center"
            >
              {{ subject.name }}
            </v-chip>
          </v-chip-group>
        </v-col>
        <v-col cols="3" align="center">
          <v-chip-group rounded color="success" tag="v-row" column :value="selectedTeacher" @change="onClickTeacher">
            <v-chip
              v-for="teacher in getTeachers()"
              :key="teacher.id"
              :value="teacher.id"
              class="col col-12 justify-center"
            >
              {{ teacher.name }}
            </v-chip>
          </v-chip-group>
          <v-btn v-if="showAddTeacherButton()">講師を追加</v-btn>
        </v-col>
      </v-row>
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <v-btn color="secondary" @click="onClose">閉じる</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, PropType, ref, SetupContext } from '@nuxtjs/composition-api'
import dayjs from '~/plugins/dayjs'
import { ShiftLessonDetail, Student, Subject, Teacher, TeacherShift } from '~/types/store'

export default defineComponent({
  props: {
    lesson: {
      type: Object as PropType<ShiftLessonDetail>,
      default: () => {},
    },
    teachers: {
      type: Array as PropType<TeacherShift[]>,
      default: () => [],
    },
    subjects: {
      type: Array as PropType<Subject[]>,
      default: () => [],
    },
    selectedTeacher: {
      type: String,
      default: '',
    },
    selectedStudent: {
      type: String,
      default: '',
    },
    selectedSubject: {
      type: Number,
      default: 0,
    },
  },

  setup(props, { emit }: SetupContext) {
    const addedTeacherId = ref<string>('')

    const refresh = (): void => {
      addedTeacherId.value = ''
    }

    const getDay = (): string => {
      return dayjs(props.lesson.date).tz().format('DD(ddd)')
    }

    const getSubjects = (): Subject[] => {
      const student: Student | undefined = props.lesson.students.find(
        (student: Student): boolean => student.id === props.selectedStudent
      )
      return student ? student.subjects : []
    }

    const addableTeacher = (teachers: Teacher[], teacherId: string): TeacherShift | undefined => {
      if (teacherId === '') {
        return
      }
      const index: number = teachers.findIndex((teacher: Teacher): boolean => teacher.id === teacherId)
      if (index >= 0) {
        return
      }
      return props.teachers.find((teacher: TeacherShift): boolean => teacher.id === teacherId)
    }

    const getTeachers = (): Teacher[] => {
      const teachers: Teacher[] = []
      if (props.selectedSubject === 0) {
        return teachers
      }
      // 選択された科目に対して、授業可能な講師一覧を取得
      const subject: Subject = props.subjects.find((subject: Subject): boolean => subject.id === props.selectedSubject)
      props.lesson.teachers.forEach((teacher: Teacher): void => {
        const index: number = teacher.subjects[subject.schoolType].findIndex(
          (val: Subject): boolean => val.id === subject.id
        )
        if (index >= 0) {
          teachers.unshift(teacher)
        }
      })

      // シフト提出していないが、授業可能な講師一覧に追加された講師を取得
      let teacher: TeacherShift | undefined
      teacher = addableTeacher(teachers, addedTeacherId.value)
      if (teacher) {
        teachers.push({ ...teacher })
      }

      // シフト提出していないが、授業登録された講師を取得
      teacher = addableTeacher(teachers, props.selectedTeacher)
      if (teacher) {
        teachers.push({ ...teacher })
        addedTeacherId.value = teacher.id
      }

      return teachers
    }

    const showAddTeacherButton = (): boolean => {
      return props.selectedSubject !== 0 && addedTeacherId.value === ''
    }

    const onClickTeacher = (teacherId: string | undefined): void => {
      emit('update:selected-teacher', teacherId || '')
    }

    const onClickStudent = (studentId: string | undefined): void => {
      emit('update:selected-student', studentId || '')
    }

    const onClickSubject = (subjectId: number | undefined): void => {
      emit('update:selected-subject', subjectId || 0)
    }

    const onClose = (): void => {
      refresh()
      emit('click:close')
    }

    return {
      addedTeacherId,
      getDay,
      getSubjects,
      getTeachers,
      showAddTeacherButton,
      onClickTeacher,
      onClickStudent,
      onClickSubject,
      onClose,
    }
  },
})
</script>

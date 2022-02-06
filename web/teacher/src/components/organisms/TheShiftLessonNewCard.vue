<template>
  <v-card>
    <v-toolbar color="primary" dark elevation="0">
      <span>授業登録</span>
      <v-spacer />
      <v-icon v-if="lessonId !== 0" dark @click="onDelete">mdi-delete</v-icon>
    </v-toolbar>
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
              v-for="student in getStudents()"
              :key="student.id"
              :value="student.id"
              :disabled="!student.enabled"
              :outlined="student.enabled"
              class="col col-12 justify-center"
            >
              {{ student.name }}
            </v-chip>
          </v-chip-group>
        </v-col>
        <v-col cols="3" align="center">
          <v-chip-group rounded color="success" tag="v-row" column :value="selectedSubject" @change="onClickSubject">
            <v-chip v-if="lessonLoading" class="col col-12 justify-center">
              <v-progress-circular indeterminate />
            </v-chip>
            <v-chip
              v-for="subject in getSubjects()"
              v-else
              :key="subject.id"
              :value="subject.id"
              :disabled="!subject.enabled"
              :outlined="subject.enabled"
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
              :disabled="!teacher.enabled"
              :outlined="teacher.enabled"
              class="col col-12 justify-center"
            >
              {{ teacher.name }}
            </v-chip>
          </v-chip-group>
          <v-menu v-if="showAddTeacherButton()" offset-y>
            <template #activator="{ on, attrs }">
              <v-btn v-bind="attrs" v-on="on">講師を追加</v-btn>
            </template>
            <v-list>
              <v-list-item v-for="teacher in getAddableTeachers()" :key="teacher.name">
                <v-list-item-title @click="onClickAddTeacher(teacher.id)">{{ teacher.name }}</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
        </v-col>
      </v-row>
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <v-btn color="primary" outlined @click="onClose">閉じる</v-btn>
      <v-btn :loading="loading" :disabled="loading || selectedTeacher === ''" color="primary" @click="onSubmit">
        登録
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, PropType, ref, SetupContext } from '@nuxtjs/composition-api'
import dayjs from '~/plugins/dayjs'
import { LessonFormItemTeacher, LessonFormItemStudent, LessonFormItemSubject } from '~/types/props/shift'
import {
  ShiftLesson,
  ShiftLessonDetail,
  ShiftUserLesson,
  Student,
  StudentShift,
  Subject,
  SuggestedLesson,
  Teacher,
  TeacherShift,
} from '~/types/store'

export default defineComponent({
  props: {
    loading: {
      type: Boolean,
      default: false,
    },
    lessonLoading: {
      type: Boolean,
      default: false,
    },
    lesson: {
      type: Object as PropType<ShiftLessonDetail>,
      default: () => {},
    },
    studentLessons: {
      type: Object as PropType<ShiftUserLesson>,
      default: () => {},
    },
    teachers: {
      type: Array as PropType<TeacherShift[]>,
      default: () => [],
    },
    students: {
      type: Array as PropType<StudentShift[]>,
      default: () => [],
    },
    subjects: {
      type: Array as PropType<Subject[]>,
      default: () => [],
    },
    lessonId: {
      type: Number,
      default: 0,
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

    const addStudent = (students: LessonFormItemStudent[], student: LessonFormItemStudent): void => {
      student.enabled ? students.unshift(student) : students.push(student)
    }

    const getStudents = (): LessonFormItemStudent[] => {
      const students: LessonFormItemStudent[] = []
      // 授業を希望する生徒一覧を取得
      props.lesson.students.forEach((val: Student): void => {
        const student: StudentShift | undefined = props.students.find((s: StudentShift): boolean => s.id === val.id)
        if (!student) {
          return
        }
        // すでに希望している授業数になっている場合はボタンをdisabledにできるように
        if (student.suggestedLessonsTotal - student.lessonTotal <= 0) {
          addStudent(students, { id: student.id, name: student.name, enabled: false })
          return
        }
        // 同じ時間の他の教室の授業に含まれているか確認 (含まれていたらボタンをdisabledにできるように)
        const lesson: ShiftLesson | undefined = props.lesson.lessons.find(
          (val: ShiftLesson): boolean => val.studentId === student.id
        )
        if (lesson) {
          // 含まれている && 今の教室 -> enabled: true, 含まれている && 別の教室 -> enabled: false
          const disabled: boolean = props.lesson.current?.studentId !== lesson.studentId
          addStudent(students, { id: student.id, name: student.name, enabled: !disabled })
        } else {
          // 含まれていない -> enabled: true
          addStudent(students, { id: student.id, name: student.name, enabled: true })
        }
      })
      return students
    }

    const addSubject = (subjects: LessonFormItemSubject[], subject: LessonFormItemSubject): void => {
      subject.enabled ? subjects.unshift(subject) : subjects.push(subject)
    }

    const getSubjects = (): LessonFormItemSubject[] => {
      const subjects: LessonFormItemSubject[] = []
      if (props.selectedStudent === '') {
        return subjects
      }
      // 選択した生徒の情報を取得
      const student: StudentShift | undefined = props.students.find(
        (student: StudentShift): boolean => student.id === props.selectedStudent
      )
      if (!student) {
        return subjects
      }
      // 生徒が希望する授業科目情報一覧を取得
      student?.suggestedLessons.forEach((lesson: SuggestedLesson): void => {
        const subject: Subject | undefined = props.subjects.find((val: Subject): boolean => val.id === lesson.subjectId)
        if (!subject) {
          return
        }
        const lessons = props.studentLessons.lessons.filter(
          (val: ShiftLesson): boolean => val.studentId === student.id && val.subjectId === lesson.subjectId
        )
        const enabled: boolean = lesson.total - lessons.length > 0
        addSubject(subjects, { id: subject.id, name: subject.fullname, enabled })
      })
      return subjects
    }

    const addableTeacher = (teachers: LessonFormItemTeacher[], teacherId: string): TeacherShift | undefined => {
      if (teacherId === '') {
        return
      }
      const index: number = teachers.findIndex((teacher: LessonFormItemTeacher): boolean => teacher.id === teacherId)
      if (index >= 0) {
        return
      }
      return props.teachers.find((teacher: TeacherShift): boolean => teacher.id === teacherId)
    }

    const addTeacher = (teachers: LessonFormItemTeacher[], teacher: LessonFormItemTeacher): void => {
      teacher.enabled ? teachers.unshift(teacher) : teachers.push(teacher)
    }

    const getTeachers = (): LessonFormItemTeacher[] => {
      const teachers: LessonFormItemTeacher[] = []
      if (props.selectedSubject === 0) {
        return teachers
      }
      // 選択された授業科目情報を取得
      const subject: Subject | undefined = props.subjects.find(
        (subject: Subject): boolean => subject.id === props.selectedSubject
      )
      if (!subject) {
        return teachers
      }
      // 授業可能な(担当授業に選択された授業科目が含まれた)講師一覧を取得
      props.lesson.teachers.forEach((teacher: Teacher): void => {
        const index: number = teacher.subjects[subject.schoolType].findIndex(
          (val: Subject): boolean => val.id === subject.id
        )
        if (index < 0) {
          addTeacher(teachers, { id: teacher.id, name: teacher.name || '', enabled: false })
          return
        }
        // 同じ時間の他の教室の授業に含まれているか確認 (含まれていたらボタンをdisabledにできるように)
        const lesson: ShiftLesson | undefined = props.lesson.lessons.find(
          (val: ShiftLesson): boolean => val.teacherId === teacher.id
        )
        if (lesson) {
          // 含まれている && 今の教室 -> enabled: true, 含まれている && 別の教室 -> enabled: false
          const disabled: boolean = props.lesson.current?.teacherId !== lesson.teacherId
          addTeacher(teachers, { id: teacher.id, name: teacher.name || '', enabled: !disabled })
        } else {
          // 含まれていない -> enabled: true
          addTeacher(teachers, { id: teacher.id, name: teacher.name || '', enabled: true })
        }
      })
      // シフト提出していないが、授業可能な講師一覧に追加された講師を取得
      let teacher: TeacherShift | undefined
      teacher = addableTeacher(teachers, addedTeacherId.value)
      if (teacher) {
        addTeacher(teachers, { id: teacher.id, name: teacher.name, enabled: true })
      }
      // 選択された講師を取得
      teacher = addableTeacher(teachers, props.selectedTeacher)
      if (teacher) {
        addTeacher(teachers, { id: teacher.id, name: teacher.name, enabled: true })
        addedTeacherId.value = teacher.id
      }

      return teachers
    }

    const getAddableTeachers = (): LessonFormItemTeacher[] => {
      const teachers: LessonFormItemTeacher[] = []
      props.teachers.forEach((teacher: TeacherShift): void => {
        const index: number = props.lesson.teachers.findIndex((val: Teacher): boolean => val.id === teacher.id)
        if (index < 0) {
          teachers.push({ id: teacher.id, name: teacher.name, enabled: true })
        }
      })
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

    const onClickAddTeacher = (teacherId: string): void => {
      addedTeacherId.value = teacherId
      onClickTeacher(teacherId)
    }

    const onSubmit = (): void => {
      emit('click:submit')
    }

    const onDelete = (): void => {
      emit('click:delete')
    }

    const onClose = (): void => {
      refresh()
      emit('click:close')
    }

    return {
      addedTeacherId,
      getDay,
      getTeachers,
      getStudents,
      getSubjects,
      getAddableTeachers,
      showAddTeacherButton,
      onClickTeacher,
      onClickStudent,
      onClickSubject,
      onClickAddTeacher,
      onSubmit,
      onDelete,
      onClose,
    }
  },
})
</script>

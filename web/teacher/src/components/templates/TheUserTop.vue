<template>
  <v-container class="px-0 pt-0">
    <v-tabs v-model="selector" grow class="pb-4">
      <v-tab v-for="actor in actors" :key="actor.value" :href="`#tab-${actor.value}`">
        {{ actor.name }}
      </v-tab>
    </v-tabs>

    <v-tabs-items v-model="selector">
      <v-tab-item value="tab-teachers">
        <v-row>
          <v-dialog :value.sync="teacherEditDialog" width="600px" scrollable @click:outside="onCloseTeacherDialog">
            <the-teacher-edit-card
              :is-admin="isAdmin"
              :subjects="subjects"
              :teacher="teacher"
              :loading="loading"
              :delete-dialog="teacherDeleteDialog"
              :edit-teacher-elementary-school-form="editTeacherElementarySchoolForm"
              :edit-teacher-junior-high-school-form="editTeacherJuniorHighSchoolForm"
              :edit-teacher-high-school-form="editTeacherHighSchoolForm"
              :edit-teacher-role-form="editTeacherRoleForm"
              @click:close="onCloseTeacherDialog"
              @click:delete="onClickDeleteTeacher"
              @click:delete-accept="onClickDeleteTeacherAccept"
              @click:delete-cancel="onClickDeleteTeacherCancel"
              @submit:elementary-school="onSubmitTeacherElementarySchool"
              @submit:junior-high-school="onSubmitTeacherJuniorHighSchool"
              @submit:high-school="onSubmitTeacherHighSchool"
              @submit:role="onSubmitTeacherRole"
            />
          </v-dialog>
          <v-col class="my-4 mx-3">
            <v-btn v-show="isAdmin" color="primary" block outlined @click="onClickNew('teachers')">
              <v-icon>mdi-plus</v-icon>
              新規登録
            </v-btn>
          </v-col>
          <v-col cols="12">
            <the-teacher-list
              :items="teachers"
              :total="teachersTotal"
              :loading="loading"
              :page="teachersPage"
              :items-per-page="teachersItemsPerPage"
              @update:page="$emit('update:teachers-page', $event)"
              @update:items-per-page="$emit('update:teachers-items-per-page', $event)"
              @click="onClickShowTeacher"
            />
          </v-col>
        </v-row>
      </v-tab-item>
      <v-tab-item value="tab-students">
        <v-row>
          <v-dialog :value.sync="studentEditDialog" width="600px" scrollable @click:outside="onCloseStudentDialog">
            <the-student-edit-card
              :is-admin="isAdmin"
              :subjects="subjects"
              :student="student"
              :loading="loading"
              :delete-dialog="studentDeleteDialog"
              :edit-student-elementary-school-form="editStudentElementarySchoolForm"
              :edit-student-junior-high-school-form="editStudentJuniorHighSchoolForm"
              :edit-student-high-school-form="editStudentHighSchoolForm"
              @click:close="onCloseStudentDialog"
              @click:delete="onClickDeleteStudent"
              @click:delete-accept="onClickDeleteStudentAccept"
              @click:delete-cancel="onClickDeleteStudentCancel"
              @submit:elementary-school="onSubmitStudentElementarySchool"
              @submit:junior-high-school="onSubmitStudentJuniorHighSchool"
              @submit:high-school="onSubmitStudentHighSchool"
            />
          </v-dialog>
          <v-col class="my-4 mx-3">
            <v-btn v-show="isAdmin" color="primary" block outlined @click="onClickNew('students')">
              <v-icon>mdi-plus</v-icon>
              新規登録
            </v-btn>
          </v-col>
          <v-col cols="12">
            <the-student-list
              :items="students"
              :total="studentsTotal"
              :page="studentsPage"
              :items-per-page="studentsItemsPerPage"
              :loading="loading"
              @update:page="$emit('update:students-page', $event)"
              @update:items-per-page="$emit('update:students-items-per-page', $event)"
              @click="onClickShowStudent"
            />
          </v-col>
        </v-row>
      </v-tab-item>
    </v-tabs-items>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, ref, SetupContext } from '@nuxtjs/composition-api'
import TheStudentList from '~/components/organisms/TheStudentList.vue'
import TheStudentEditCard from '~/components/organisms/TheStudentEditCard.vue'
import TheTeacherList from '~/components/organisms/TheTeacherList.vue'
import TheTeacherEditCard from '~/components/organisms/TheTeacherEditCard.vue'
import {
  StudentEditSubjectForm,
  StudentEditSubjectForElementarySchoolParams,
  StudentEditSubjectForHighSchoolParams,
  StudentEditSubjectForJuniorHighSchoolParams,
  StudentEditSubjectForElementarySchoolOptions,
  StudentEditSubjectForJuniorHighSchoolOptions,
  StudentEditSubjectForHighSchoolOptions,
  TeacherEditSubjectForm,
  TeacherEditSubjectForElementarySchoolParams,
  TeacherEditSubjectForHighSchoolParams,
  TeacherEditSubjectForJuniorHighSchoolParams,
  TeacherEditRoleForm,
  TeacherEditRoleParams,
  TeacherEditRoleOptions,
  TeacherEditSubjectForElementarySchoolOptions,
  TeacherEditSubjectForJuniorHighSchoolOptions,
  TeacherEditSubjectForHighSchoolOptions,
} from '~/types/form'
import { Actor } from '~/types/props/user'
import { Role, Student, SubjectsMap, Teacher } from '~/types/store'

export default defineComponent({
  components: {
    TheStudentList,
    TheStudentEditCard,
    TheTeacherList,
    TheTeacherEditCard,
  },

  props: {
    loading: {
      type: Boolean,
      default: false,
    },
    isAdmin: {
      type: Boolean,
      default: false,
    },
    subjects: {
      type: Object as PropType<SubjectsMap>,
      default: () => ({
        小学校: [],
        中学校: [],
        高校: [],
        その他: [],
      }),
    },
    /**
     * 生徒関連
     */
    student: {
      type: Object as PropType<Student>,
      default: () => ({
        id: '',
        lastName: '',
        firstName: '',
        lastNameKana: '',
        firstNameKana: '',
        mail: '',
        schoolType: 'その他',
        grade: 1,
        subjects: [],
        createdAt: '',
        updatedAt: '',
      }),
    },
    students: {
      type: Array as PropType<Student[]>,
      default: () => [],
    },
    studentEditDialog: {
      type: Boolean,
      default: false,
    },
    studentsTotal: {
      type: Number,
      default: 0,
    },
    studentsPage: {
      type: Number,
      default: 1,
    },
    studentsItemsPerPage: {
      type: Number,
      default: 10,
    },
    editStudentElementarySchoolForm: {
      type: Object as PropType<StudentEditSubjectForm>,
      default: () => ({
        params: StudentEditSubjectForElementarySchoolParams,
        options: StudentEditSubjectForElementarySchoolOptions,
      }),
    },
    editStudentJuniorHighSchoolForm: {
      type: Object as PropType<StudentEditSubjectForm>,
      default: () => ({
        params: StudentEditSubjectForJuniorHighSchoolParams,
        options: StudentEditSubjectForJuniorHighSchoolOptions,
      }),
    },
    editStudentHighSchoolForm: {
      type: Object as PropType<StudentEditSubjectForm>,
      default: () => ({
        params: StudentEditSubjectForHighSchoolParams,
        options: StudentEditSubjectForHighSchoolOptions,
      }),
    },
    /**
     * 講師関連
     */
    teacher: {
      type: Object as PropType<Teacher>,
      default: () => ({
        id: '',
        lastName: '',
        firstName: '',
        lastNameKana: '',
        firstNameKana: '',
        mail: '',
        role: Role.TEACHER,
        subjects: {},
        createdAt: '',
        updatedAt: '',
      }),
    },
    teachers: {
      type: Array as PropType<Teacher[]>,
      default: () => [],
    },
    teacherEditDialog: {
      type: Boolean,
      default: false,
    },
    teachersTotal: {
      type: Number,
      default: 0,
    },
    teachersPage: {
      type: Number,
      default: 1,
    },
    teachersItemsPerPage: {
      type: Number,
      default: 10,
    },
    editTeacherElementarySchoolForm: {
      type: Object as PropType<TeacherEditSubjectForm>,
      default: () => ({
        params: TeacherEditSubjectForElementarySchoolParams,
        options: TeacherEditSubjectForElementarySchoolOptions,
      }),
    },
    editTeacherJuniorHighSchoolForm: {
      type: Object as PropType<TeacherEditSubjectForm>,
      default: () => ({
        params: TeacherEditSubjectForJuniorHighSchoolParams,
        options: TeacherEditSubjectForJuniorHighSchoolOptions,
      }),
    },
    editTeacherHighSchoolForm: {
      type: Object as PropType<TeacherEditSubjectForm>,
      default: () => ({
        params: TeacherEditSubjectForHighSchoolParams,
        options: TeacherEditSubjectForHighSchoolOptions,
      }),
    },
    editTeacherRoleForm: {
      type: Object as PropType<TeacherEditRoleForm>,
      default: () => ({
        params: TeacherEditRoleParams,
        options: TeacherEditRoleOptions,
      }),
    },
  },

  setup(_, { emit }: SetupContext) {
    const actors: Actor[] = [
      { name: '講師', value: 'teachers' },
      { name: '生徒', value: 'students' },
    ]

    const selector = ref<string>('teachers')
    const studentDeleteDialog = ref<boolean>(false)
    const teacherDeleteDialog = ref<boolean>(false)

    const onClickNew = (actor: string): void => {
      emit('click:new', actor)
    }

    /**
     * 生徒関連
     */
    const onClickShowStudent = (student: Student): void => {
      emit('click:show-student', student)
    }

    const onClickDeleteStudent = (): void => {
      studentDeleteDialog.value = true
    }

    const onClickDeleteStudentAccept = (): void => {
      emit('submit:student-delete')
      studentDeleteDialog.value = false
    }

    const onClickDeleteStudentCancel = (): void => {
      studentDeleteDialog.value = false
    }

    const onCloseStudentDialog = (): void => {
      emit('click:close-student')
      studentDeleteDialog.value = false
    }

    const onSubmitStudentElementarySchool = (): void => {
      emit('submit:student-elementary-school')
    }

    const onSubmitStudentJuniorHighSchool = (): void => {
      emit('submit:student-junior-high-school')
    }

    const onSubmitStudentHighSchool = (): void => {
      emit('submit:student-high-school')
    }

    /**
     * 講師関連
     */
    const onClickShowTeacher = (teacher: Teacher): void => {
      emit('click:show-teacher', teacher)
    }

    const onClickDeleteTeacher = (): void => {
      teacherDeleteDialog.value = true
    }

    const onClickDeleteTeacherAccept = (): void => {
      emit('submit:teacher-delete')
      teacherDeleteDialog.value = false
    }

    const onClickDeleteTeacherCancel = (): void => {
      teacherDeleteDialog.value = false
    }

    const onCloseTeacherDialog = (): void => {
      emit('click:close-teacher')
      teacherDeleteDialog.value = false
    }

    const onSubmitTeacherElementarySchool = (): void => {
      emit('submit:teacher-elementary-school')
    }

    const onSubmitTeacherJuniorHighSchool = (): void => {
      emit('submit:teacher-junior-high-school')
    }

    const onSubmitTeacherHighSchool = (): void => {
      emit('submit:teacher-high-school')
    }

    const onSubmitTeacherRole = (): void => {
      emit('submit:teacher-role')
    }

    return {
      actors,
      selector,
      studentDeleteDialog,
      teacherDeleteDialog,
      onClickNew,
      onClickShowStudent,
      onClickShowTeacher,
      onCloseStudentDialog,
      onCloseTeacherDialog,
      onClickDeleteStudent,
      onClickDeleteStudentAccept,
      onClickDeleteStudentCancel,
      onClickDeleteTeacher,
      onClickDeleteTeacherAccept,
      onClickDeleteTeacherCancel,
      onSubmitStudentElementarySchool,
      onSubmitStudentJuniorHighSchool,
      onSubmitStudentHighSchool,
      onSubmitTeacherElementarySchool,
      onSubmitTeacherJuniorHighSchool,
      onSubmitTeacherHighSchool,
      onSubmitTeacherRole,
    }
  },
})
</script>

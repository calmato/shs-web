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
          <v-col class="d-flex flex-column align-end px-8">
            <v-btn v-show="isAdmin" color="primary" @click="onClickNew('teachers')">新規登録</v-btn>
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
        <the-student-list :items="students" :loading="loading" />
      </v-tab-item>
    </v-tabs-items>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, ref, SetupContext } from '@nuxtjs/composition-api'
import TheStudentList from '~/components/organisms/TheStudentList.vue'
import TheTeacherList from '~/components/organisms/TheTeacherList.vue'
import TheTeacherEditCard from '~/components/organisms/TheTeacherEditCard.vue'
import {
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
    students: {
      type: Array as PropType<Student[]>,
      default: () => [],
    },
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
    teacherEditDialog: {
      type: Boolean,
      default: false,
    },
    teachers: {
      type: Array as PropType<Teacher[]>,
      default: () => [],
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
    const teacherDeleteDialog = ref<boolean>(false)

    const onClickNew = (actor: string): void => {
      emit('click:new', actor)
    }

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
      teacherDeleteDialog,
      onClickNew,
      onClickShowTeacher,
      onCloseTeacherDialog,
      onClickDeleteTeacher,
      onClickDeleteTeacherAccept,
      onClickDeleteTeacherCancel,
      onSubmitTeacherElementarySchool,
      onSubmitTeacherJuniorHighSchool,
      onSubmitTeacherHighSchool,
      onSubmitTeacherRole,
    }
  },
})
</script>

<template>
  <v-card>
    <v-toolbar color="primary" dark>講師詳細</v-toolbar>

    <v-card-text>
      <v-row class="py-8">
        <v-col cols="12">
          <h3 class="mb-2">講師名</h3>
          <span>{{ getTeacherName() }}</span>
          <v-divider />
        </v-col>
        <v-col cols="12">
          <h3 class="mb-2">連絡先</h3>
          <span>{{ teacher.mail }}</span>
          <v-divider />
        </v-col>
        <v-col cols="12">
          <the-select
            :label="editTeacherRoleForm.options.role.label"
            :value.sync="editTeacherRoleForm.params.role"
            :items="roleItems"
            @blur="onSubmitRole"
          />
        </v-col>
        <v-col cols="12">
          <the-select-with-slot
            :label="editTeacherElementarySchoolForm.options.subjectIds.label"
            :value.sync="editTeacherElementarySchoolForm.params.subjectIds"
            :items="getElementarySchoolSubjects()"
            item-text="name"
            item-value="id"
            chips
            multiple
            append-outer-icon="mdi-lead-pencil"
            @blur="onSubmitElementarySchool"
          >
            <template #default="{ item }">
              <v-chip label :color="item.color">{{ item.name }}</v-chip>
            </template>
          </the-select-with-slot>
        </v-col>
        <v-col cols="12">
          <the-select-with-slot
            :label="editTeacherJuniorHighSchoolForm.options.subjectIds.label"
            :value.sync="editTeacherJuniorHighSchoolForm.params.subjectIds"
            :items="getJuniorHighSchoolSubjects()"
            item-text="name"
            item-value="id"
            chips
            multiple
            append-outer-icon="mdi-lead-pencil"
            @blur="onSubmitJuniorHighSchool"
          >
            <template #default="{ item }">
              <v-chip label :color="item.color">{{ item.name }}</v-chip>
            </template>
          </the-select-with-slot>
        </v-col>
        <v-col cols="12">
          <the-select-with-slot
            :label="editTeacherHighSchoolForm.options.subjectIds.label"
            :value.sync="editTeacherHighSchoolForm.params.subjectIds"
            :items="getHighSchoolSubjects()"
            item-text="name"
            item-value="id"
            chips
            multiple
            append-outer-icon="mdi-lead-pencil"
            @blur="onSubmitHighSchool"
          >
            <template #default="{ item }">
              <v-chip label :color="item.color">{{ item.name }}</v-chip>
            </template>
          </the-select-with-slot>
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
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheSelect from '~/components/atoms/TheSelect.vue'
import TheSelectWithSlot from '~/components/atoms/TheSelectWithSlot.vue'
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
import { RoleItem } from '~/types/props/teacher'
import { Role, SchoolType, Subject, SubjectsMap, Teacher } from '~/types/store'

export default defineComponent({
  components: {
    TheSelect,
    TheSelectWithSlot,
  },

  props: {
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
        subjects: {
          [SchoolType.ELEMENTARY_SCHOOL]: [],
          [SchoolType.JUNIOR_HIGH_SCHOOL]: [],
          [SchoolType.HIGH_SCHOOL]: [],
        },
        createdAt: '',
        updatedAt: '',
      }),
    },
    subjects: {
      type: Object as PropType<SubjectsMap>,
      default: () => ({
        [SchoolType.ELEMENTARY_SCHOOL]: [],
        [SchoolType.JUNIOR_HIGH_SCHOOL]: [],
        [SchoolType.HIGH_SCHOOL]: [],
      }),
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

  setup(props, { emit }: SetupContext) {
    const roleItems: RoleItem[] = [
      { text: '講師', value: Role.TEACHER },
      { text: '管理者', value: Role.ADMINISTRATOR },
    ]

    const getTeacherName = (): string => {
      const name: string = `${props.teacher.lastName} ${props.teacher.firstName}`
      const nameKana: string = `${props.teacher.lastNameKana} ${props.teacher.firstNameKana}`
      return `${name} (${nameKana})`
    }

    const getElementarySchoolSubjects = (): Subject[] => {
      return props.subjects[SchoolType.ELEMENTARY_SCHOOL]
    }

    const getJuniorHighSchoolSubjects = (): Subject[] => {
      return props.subjects[SchoolType.JUNIOR_HIGH_SCHOOL]
    }

    const getHighSchoolSubjects = (): Subject[] => {
      return props.subjects[SchoolType.HIGH_SCHOOL]
    }

    const onClose = (): void => {
      emit('click:close')
    }

    const onSubmitRole = (): void => {
      emit('submit:role')
    }

    const onSubmitElementarySchool = (): void => {
      emit('submit:elementary-school')
    }

    const onSubmitJuniorHighSchool = (): void => {
      emit('submit:junior-high-school')
    }

    const onSubmitHighSchool = (): void => {
      emit('submit:high-school')
    }

    return {
      roleItems,
      getTeacherName,
      getElementarySchoolSubjects,
      getJuniorHighSchoolSubjects,
      getHighSchoolSubjects,
      onClose,
      onSubmitRole,
      onSubmitElementarySchool,
      onSubmitJuniorHighSchool,
      onSubmitHighSchool,
    }
  },
})
</script>

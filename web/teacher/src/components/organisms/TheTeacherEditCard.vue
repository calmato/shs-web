<template>
  <v-card>
    <v-toolbar color="primary" dark>
      <span>講師詳細</span>
      <v-spacer />
      <v-icon @click="onDelete">mdi-delete</v-icon>
    </v-toolbar>

    <v-card-text>
      <v-row v-if="deleteDialog" class="py-8">
        <h2>本当に削除しますか</h2>
      </v-row>

      <v-row v-else class="py-8">
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
            :disabled="!isAdmin"
            @blur="onSubmitRole"
          />
        </v-col>
        <v-col cols="12">
          <the-select-with-slot
            :label="editTeacherElementarySchoolForm.options.subjectIds.label"
            :value.sync="editTeacherElementarySchoolForm.params.subjectIds"
            :items="getElementarySchoolSubjects()"
            :append-outer-icon="isAdmin ? `mdi-lead-pencil` : undefined"
            :disabled="!isAdmin"
            item-text="name"
            item-value="id"
            chips
            multiple
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
            :disabled="!isAdmin"
            :append-outer-icon="isAdmin ? `mdi-lead-pencil` : undefined"
            item-text="name"
            item-value="id"
            chips
            multiple
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
            :disabled="!isAdmin"
            :append-outer-icon="isAdmin ? `mdi-lead-pencil` : undefined"
            item-text="name"
            item-value="id"
            chips
            multiple
            @blur="onSubmitHighSchool"
          >
            <template #default="{ item }">
              <v-chip label :color="item.color">{{ item.name }}</v-chip>
            </template>
          </the-select-with-slot>
        </v-col>
      </v-row>
    </v-card-text>

    <v-card-actions v-if="deleteDialog">
      <v-spacer />
      <v-btn @click="onDeleteCancel">キャンセル</v-btn>
      <v-btn color="error" :disabled="loading" @click="onDeleteAccept">削除する</v-btn>
    </v-card-actions>
    <v-card-actions v-else>
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
import { Role, Subject, SubjectsMap, Teacher } from '~/types/store'

export default defineComponent({
  components: {
    TheSelect,
    TheSelectWithSlot,
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
    deleteDialog: {
      type: Boolean,
      default: false,
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
    subjects: {
      type: Object as PropType<SubjectsMap>,
      default: () => ({
        小学校: [],
        中学校: [],
        高校: [],
        その他: [],
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
      return `${props.teacher.name} (${props.teacher.nameKana})`
    }

    const getElementarySchoolSubjects = (): Subject[] => {
      return props.subjects['小学校']
    }

    const getJuniorHighSchoolSubjects = (): Subject[] => {
      return props.subjects['中学校']
    }

    const getHighSchoolSubjects = (): Subject[] => {
      return props.subjects['高校']
    }

    const onClose = (): void => {
      emit('click:close')
    }

    const onDelete = (): void => {
      emit('click:delete')
    }

    const onDeleteAccept = (): void => {
      emit('click:delete-accept')
    }

    const onDeleteCancel = (): void => {
      emit('click:delete-cancel')
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
      onDelete,
      onDeleteAccept,
      onDeleteCancel,
      onSubmitRole,
      onSubmitElementarySchool,
      onSubmitJuniorHighSchool,
      onSubmitHighSchool,
    }
  },
})
</script>

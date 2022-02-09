<template>
  <v-card>
    <v-toolbar color="primary" dark elevation="0">
      <span>生徒詳細</span>
      <v-spacer />
      <v-icon @click="onDelete">mdi-delete</v-icon>
    </v-toolbar>

    <v-card-text>
      <v-row v-if="deleteDialog" class="py-8">
        <h2>本当に削除しますか</h2>
      </v-row>

      <v-row v-else class="py-8">
        <v-col cols="12">
          <h3 class="mb-2">生徒名</h3>
          <span>{{ getStudentName() }}</span>
          <v-divider />
        </v-col>
        <v-col cols="12">
          <h3 class="mb-2">連絡先</h3>
          <span>{{ student.mail }}</span>
          <v-divider />
        </v-col>
        <v-col cols="12">
          <h3 class="mb-2">校種・学年</h3>
          <span>{{ student.schoolType }} {{ student.grade }}年</span>
          <v-divider />
        </v-col>
        <v-col cols="12">
          <the-select-with-slot
            :label="editStudentElementarySchoolForm.options.subjectIds.label"
            :value.sync="editStudentElementarySchoolForm.params.subjectIds"
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
            :label="editStudentJuniorHighSchoolForm.options.subjectIds.label"
            :value.sync="editStudentJuniorHighSchoolForm.params.subjectIds"
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
            :label="editStudentHighSchoolForm.options.subjectIds.label"
            :value.sync="editStudentHighSchoolForm.params.subjectIds"
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
      <v-btn color="primary" outlined @click="onClose">閉じる</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheSelectWithSlot from '~/components/atoms/TheSelectWithSlot.vue'
import {
  StudentEditSubjectForm,
  StudentEditSubjectForElementarySchoolParams,
  StudentEditSubjectForHighSchoolParams,
  StudentEditSubjectForJuniorHighSchoolParams,
  StudentEditSubjectForElementarySchoolOptions,
  StudentEditSubjectForJuniorHighSchoolOptions,
  StudentEditSubjectForHighSchoolOptions,
} from '~/types/form'
import { Subject, SubjectsMap, Student } from '~/types/store'

export default defineComponent({
  components: {
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
    subjects: {
      type: Object as PropType<SubjectsMap>,
      default: () => ({
        小学校: [],
        中学校: [],
        高校: [],
        その他: [],
      }),
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
  },

  setup(props, { emit }: SetupContext) {
    const getStudentName = (): string => {
      return `${props.student.name} (${props.student.nameKana})`
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
      getStudentName,
      getElementarySchoolSubjects,
      getJuniorHighSchoolSubjects,
      getHighSchoolSubjects,
      onClose,
      onDelete,
      onDeleteAccept,
      onDeleteCancel,
      onSubmitElementarySchool,
      onSubmitJuniorHighSchool,
      onSubmitHighSchool,
    }
  },
})
</script>

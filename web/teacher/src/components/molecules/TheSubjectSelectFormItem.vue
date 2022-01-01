<template>
  <the-form-group>
    <div class="d-flex h-stack">
      <v-text-field class="flex-grow-1" label="氏名（姓）" :value="user.lastName" readonly />
      <v-text-field class="flex-grow-1" label="氏名（名）" :value="user.firstName" readonly />
    </div>
    <div class="d-flex h-stack">
      <v-text-field class="flex-grow-1" label="氏名（姓：かな）" :value="user.lastNameKana" readonly />
      <v-text-field class="flex-grow-1" label="氏名（名：かな）" :value="user.firstNameKana" readonly />
    </div>
    <v-select
      label="担当科目（小学校）"
      chips
      multiple
      append-outer-icon="mdi-lead-pencil"
      item-text="name"
      :items="elementarySchoolSubjects"
    />
    <v-select
      label="担当科目（中学校）"
      chips
      multiple
      append-outer-icon="mdi-lead-pencil"
      item-text="name"
      :items="juniorHighSchoolSubjects"
    />
    <v-select
      label="担当科目（高校）"
      chips
      multiple
      append-outer-icon="mdi-lead-pencil"
      item-text="name"
      :items="highSchoolSubjects"
    />
  </the-form-group>
</template>

<script lang="ts">
import { computed, defineComponent, PropType } from '@nuxtjs/composition-api'
import { Subject } from '~/types/store'

interface UserPoop {
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
}

export default defineComponent({
  components: {},

  props: {
    subjectForm: {
      type: Object,
      default: () => ({}),
    },
    user: {
      type: Object as PropType<UserPoop>,
      default: () => ({
        lastName: '',
        firstName: '',
        lastNameKana: '',
        firstNameKana: '',
      }),
    },
    elementarySchoolSubjects: {
      type: Array as PropType<Subject[]>,
      default: () => [],
    },
    juniorHighSchoolSubjects: {
      type: Array as PropType<Subject[]>,
      default: () => [],
    },
    highSchoolSubjects: {
      type: Array as PropType<Subject[]>,
      default: () => [],
    },
  },

  setup(props, { emit }) {
    const formData = computed({
      get: () => props.subjectForm,
      set: (val: object) => emit('update:form', val),
    })

    return {
      formData,
    }
  },
})
</script>

<style lang="scss">
.h-stack {
  gap: var(--space, 1rem);
}
</style>

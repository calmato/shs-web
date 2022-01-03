<template>
  <form>
    <div class="d-flex h-stack">
      <v-text-field class="flex-grow-1" label="氏名（姓）" :value="user.lastName" readonly />
      <v-text-field class="flex-grow-1" label="氏名（名）" :value="user.firstName" readonly />
    </div>
    <div class="d-flex h-stack">
      <v-text-field class="flex-grow-1" label="氏名（姓：かな）" :value="user.lastNameKana" readonly />
      <v-text-field class="flex-grow-1" label="氏名（名：かな）" :value="user.firstNameKana" readonly />
    </div>
    <v-select
      v-model="elementarySchoolSubjectsFormData"
      label="担当科目（小学校）"
      chips
      multiple
      append-outer-icon="mdi-lead-pencil"
      item-text="name"
      item-value="id"
      :items="elementarySchoolSubjects"
      @change="handleElementarySchoolSubjectsChange"
    >
      <template #selection="{ item }">
        <v-chip label :color="item.color">{{ item.name }}</v-chip>
      </template>
    </v-select>
    <v-select
      v-model="juniorHighSchoolSubjectsFormData"
      label="担当科目（中学校）"
      chips
      multiple
      append-outer-icon="mdi-lead-pencil"
      item-text="name"
      item-value="id"
      :items="juniorHighSchoolSubjects"
      @change="handleJuniorHighSchoolSubjectsChange"
    >
      <template #selection="{ item }">
        <v-chip label :color="item.color">{{ item.name }}</v-chip>
      </template>
    </v-select>
    <v-select
      v-model="highSchoolSubjectsFormData"
      label="担当科目（高校）"
      chips
      multiple
      append-outer-icon="mdi-lead-pencil"
      item-text="name"
      item-value="id"
      :items="highSchoolSubjects"
      @change="handleHighSchoolSubjectsChange"
    >
      <template #selection="{ item }">
        <v-chip label :color="item.color">{{ item.name }}</v-chip>
      </template>
    </v-select>
  </form>
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
    elementarySchoolSubjectsFormValue: {
      type: Array as PropType<number[]>,
      default: () => [],
    },
    juniorHighSchoolSubjectsFormValue: {
      type: Array as PropType<number[]>,
      default: () => [],
    },
    highSchoolSubjectsFormValue: {
      type: Array as PropType<number[]>,
      default: () => [],
    },
  },

  setup(props, { emit }) {
    const elementarySchoolSubjectsFormData = computed({
      get: () => props.elementarySchoolSubjectsFormValue,
      set: (val: object) => emit('update:elementarySchoolSubjectsFormValue', val),
    })

    const juniorHighSchoolSubjectsFormData = computed({
      get: () => props.juniorHighSchoolSubjectsFormValue,
      set: (val: object) => emit('update:juniorHighSchoolSubjectsFormValue', val),
    })
    const highSchoolSubjectsFormData = computed({
      get: () => props.highSchoolSubjectsFormValue,
      set: (val: object) => emit('update:highSchoolSubjectsFormValue', val),
    })

    const handleElementarySchoolSubjectsChange = (val: number[]) => {
      emit('handleElementarySchoolSubjectsChange', val)
    }

    const handleJuniorHighSchoolSubjectsChange = (val: number[]) => {
      emit('handleJuniorHighSchoolSubjectsChange', val)
    }

    const handleHighSchoolSubjectsChange = (val: number[]) => {
      emit('handleHighSchoolSubjectsChange', val)
    }

    return {
      elementarySchoolSubjectsFormData,
      juniorHighSchoolSubjectsFormData,
      highSchoolSubjectsFormData,
      handleElementarySchoolSubjectsChange,
      handleJuniorHighSchoolSubjectsChange,
      handleHighSchoolSubjectsChange,
    }
  },
})
</script>

<style lang="scss">
.h-stack {
  gap: var(--space, 1rem);
}
</style>

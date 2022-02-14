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
      label="受講科目"
      chips
      multiple
      item-text="name"
      item-value="id"
      :items="subjects"
      :value="getSubjectIds()"
      readonly
    >
      <template #selection="{ item }">
        <v-chip label :color="item.color">{{ item.fullname }}</v-chip>
      </template>
    </v-select>
  </form>
</template>

<script lang="ts">
import { defineComponent, PropType } from '@nuxtjs/composition-api'
import { UserProps } from '~/types/props/setting'
import { Subject } from '~/types/store'

export default defineComponent({
  props: {
    user: {
      type: Object as PropType<UserProps>,
      default: () => ({
        lastName: '',
        firstName: '',
        lastNameKana: '',
        firstNameKana: '',
      }),
    },
    subjects: {
      type: Array as PropType<Subject[]>,
      default: () => [],
    },
  },

  setup(props) {
    const getSubjectIds = (): void => {
      return props.subjects.map((subject: Subject) => subject.id)
    }

    return {
      getSubjectIds,
    }
  },
})
</script>

<style lang="scss">
.h-stack {
  gap: var(--space, 1rem);
}
</style>

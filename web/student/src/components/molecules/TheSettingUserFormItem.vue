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
  </form>
</template>

<script lang="ts">
import { defineComponent, PropType } from '@nuxtjs/composition-api'
import { UserProps } from '~/types/props/setting'

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
  },
})
</script>

<style lang="scss">
.h-stack {
  gap: var(--space, 1rem);
}
</style>

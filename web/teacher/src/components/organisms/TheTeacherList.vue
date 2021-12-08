<template>
  <v-data-table :mobile-breakpoint="0" :headers="headers" :items="items" :loading="loading">
    <template #[`item.role`]="{ item }">
      <v-chip :color="getRoleColor(item.role)" dark>
        {{ getRole(item.role) }}
      </v-chip>
    </template>
  </v-data-table>
</template>

<script lang="ts">
import { defineComponent, PropType } from '@nuxtjs/composition-api'
import { TableHeader } from '~/types/props/user'
import { Role, Teacher } from '~/types/store'

export default defineComponent({
  props: {
    items: {
      type: Array as PropType<Teacher[]>,
      default: () => [],
    },
    loading: {
      type: Boolean,
      default: false,
    },
  },

  setup() {
    const headers: TableHeader[] = [
      { text: '講師名', value: 'name', sortable: false },
      { text: '役職', value: 'role', sortable: false },
    ]

    const getRole = (role: Role): string => {
      switch (role) {
        case Role.TEACHER:
          return '講師'
        case Role.ADMINISTRATOR:
          return '管理者'
        default:
          return '不明'
      }
    }

    const getRoleColor = (role: Role): string => {
      switch (role) {
        case Role.TEACHER:
          return 'primary'
        case Role.ADMINISTRATOR:
          return 'secondary'
        default:
          return ''
      }
    }

    return {
      headers,
      getRole,
      getRoleColor,
    }
  },
})
</script>

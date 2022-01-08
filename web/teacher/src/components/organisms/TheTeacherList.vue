<template>
  <v-data-table
    :mobile-breakpoint="0"
    :headers="headers"
    :footer-props="footer"
    :items="items"
    :items-per-page="itemsPerPage"
    :loading="loading"
    :page="page"
    :server-items-length="total"
    @update:page="$emit('update:page', $event)"
    @update:items-per-page="$emit('update:items-per-page', $event)"
    @click:row="onClick"
  >
    <template #[`item.role`]="{ item }">
      <v-chip :color="getRoleColor(item.role)" dark>
        {{ getRole(item.role) }}
      </v-chip>
    </template>
  </v-data-table>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { TableFooter, TableHeader } from '~/types/props/user'
import { Role, Teacher } from '~/types/store'

export default defineComponent({
  props: {
    items: {
      type: Array as PropType<Teacher[]>,
      default: () => [],
    },
    itemsPerPage: {
      type: Number,
      default: 10,
    },
    loading: {
      type: Boolean,
      default: false,
    },
    page: {
      type: Number,
      default: 1,
    },
    total: {
      type: Number,
      default: 0,
    },
  },

  setup(_, { emit }: SetupContext) {
    const headers: TableHeader[] = [
      { text: '講師名', value: 'name', sortable: false },
      { text: '役職', value: 'role', sortable: false },
    ]
    const footer: TableFooter = {
      itemsPerPageText: '表示件数',
      itemsPerPageOptions: [10, 20, 30, 50],
    }

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

    const onClick = (teacher: Teacher): void => {
      emit('click', teacher)
    }

    return {
      headers,
      footer,
      getRole,
      getRoleColor,
      onClick,
    }
  },
})
</script>

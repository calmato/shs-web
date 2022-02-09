<template>
  <v-data-table
    :mobile-breakpoint="0"
    :headers="headers"
    :footer-props="footer"
    :page="page"
    :items="items"
    :items-per-page="itemsPerPage"
    :loading="loading"
    :server-items-length="total"
    @update:page="$emit('update:page', $event)"
    @update:items-per-page="$emit('update:items-per-page', $event)"
    @click:row="onClick"
  >
    <template #[`item.type`]="{ item }">
      <v-chip :color="getSchoolTypeColor(item.schoolType)" dark small>
        {{ item.schoolType }}
      </v-chip>
    </template>
    <template #[`item.grade`]="{ item }">{{ getSchoolYear(item.grade) }}</template>
  </v-data-table>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { TableFooter, TableHeader } from '~/types/props/user'
import { Student } from '~/types/store'

export default defineComponent({
  props: {
    items: {
      type: Array as PropType<Student[]>,
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
      { text: '生徒名', value: 'name', sortable: false },
      { text: '校種', value: 'type', sortable: false },
      { text: '学年', value: 'grade', sortable: false },
    ]
    const footer: TableFooter = {
      itemsPerPageText: '表示件数',
      itemsPerPageOptions: [10, 20, 30, 50],
    }

    const getSchoolTypeColor = (type: string): string => {
      switch (type) {
        case '小学校':
          return 'primary'
        case '中学校':
          return 'secondary'
        case '高校':
          return 'info'
        case 'その他':
        default:
          return ''
      }
    }

    const getSchoolYear = (grade: number): string => {
      return `${grade}年`
    }

    const onClick = (student: Student): void => {
      emit('click', student)
    }

    return {
      headers,
      footer,
      getSchoolTypeColor,
      getSchoolYear,
      onClick,
    }
  },
})
</script>

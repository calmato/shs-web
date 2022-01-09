<template>
  <v-data-table :mobile-breakpoint="0" :headers="headers" :items="items" :loading="loading">
    <template #[`item.type`]="{ item }">
      <v-chip :color="getSchoolTypeColor(item.type)" dark>
        {{ getSchoolType(item.type) }}
      </v-chip>
    </template>
    <template #[`item.grade`]="{ item }">{{ getSchoolYear(item.grade) }}</template>
  </v-data-table>
</template>

<script lang="ts">
import { defineComponent, PropType } from '@nuxtjs/composition-api'
import { schoolTypeNum2schoolTypeString } from '~/lib'
import { TableHeader } from '~/types/props/user'
import { SchoolType, Student } from '~/types/store'

export default defineComponent({
  props: {
    items: {
      type: Array as PropType<Student[]>,
      default: () => [],
    },
    loading: {
      type: Boolean,
      default: false,
    },
  },

  setup() {
    const headers: TableHeader[] = [
      { text: '生徒名', value: 'name', sortable: false },
      { text: '校種', value: 'type', sortable: false },
      { text: '学年', value: 'grade', sortable: false },
    ]

    const getSchoolType = schoolTypeNum2schoolTypeString

    const getSchoolTypeColor = (type: number): string => {
      const schoolType: SchoolType = schoolTypeNum2schoolTypeString(type)
      switch (schoolType) {
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

    return {
      headers,
      getSchoolType,
      getSchoolTypeColor,
      getSchoolYear,
    }
  },
})
</script>

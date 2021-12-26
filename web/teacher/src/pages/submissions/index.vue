<template>
  <v-row class="pt-6" justify="center">
    <v-col v-for="item in items" :key="item.id" cols="10" sm="6" md="4">
      <v-card outlined elevation="5">
        <v-card-title class="justify-center">{{ item.title }}シフト希望登録</v-card-title>
        <v-card-text class="text-center text-decoration-underline red--text text--lighten-1">
          提出期限: {{ getEnddate(item.endDate) }}まで
        </v-card-text>
        <v-container class="pl-6">
          <v-container class="d-flex">
            <v-chip text-color="white" :color="getStatusColor(item.status)">
              {{ getStatus(item.status) }}
            </v-chip>
            <v-row class="pt-3" justify="end">
              <v-btn color="primary">{{ getEditStatus(item.status) }}</v-btn>
            </v-row>
          </v-container>
        </v-container>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api'
import { Submission } from '~/types/props/submission'

export default defineComponent({
  setup(_) {
    const items: Submission[] = [
      {
        title: '1月',
        endDate: '20210125',
        status: 1,
      },
      {
        title: '2月',
        endDate: '20230225',
        status: 2,
      },
    ]

    const getEnddate = (endDate: string): string => {
      const year: string = endDate.slice(0, 4)
      const month: string = endDate.slice(5, 6)
      const date: string = endDate.slice(7, 8)
      return year + '/' + month + '/' + date
    }

    const getStatus = (status: number): string => {
      switch (status) {
        case 1:
          return '未提出'
        case 2:
          return '提出済み'
        default:
          return 'Unknown'
      }
    }

    const getEditStatus = (status: number): string => {
      switch (status) {
        case 1:
          return '入力する'
        case 2:
          return '編集する'
        default:
          return 'Unknown'
      }
    }

    const getStatusColor = (status: number): string => {
      switch (status) {
        case 1:
          return 'red'
        case 2:
          return 'light-green'
        default:
          return ''
      }
    }

    return {
      items,
      getEnddate,
      getStatus,
      getStatusColor,
      getEditStatus,
    }
  },
})
</script>

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
            <v-chip text-color="white" :color="getStatusColor(item.submissionStatus)">
              {{ item.submissionStatus }}
            </v-chip>
            <v-row class="pt-3" justify="end">
              <v-btn color="primary" @click="onClickEdit(`teachers`)">{{ item.editStatus }}</v-btn>
            </v-row>
          </v-container>
        </v-container>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { Submission, SubmissionStatusType } from '~/types/props/submission'

export default defineComponent({
  props: {
    items: {
      type: Array as PropType<Submission[]>,
      default: () => [],
    },
  },

  setup(_, { emit }: SetupContext) {
    const getEnddate = (endDate: string): string => {
      const year: string = endDate.slice(0, 4)
      const month: string = endDate.slice(5, 6)
      const date: string = endDate.slice(7, 8)
      return year + '/' + month + '/' + date
    }

    const getStatusColor = (status: SubmissionStatusType): string => {
      switch (status) {
        case '未提出':
          return 'red'
        case '提出済み':
          return 'light-green'
      }
    }

    const onClickEdit = (actor: string): void => {
      emit('click:edit', actor)
    }

    return {
      getEnddate,
      getStatusColor,
      onClickEdit,
    }
  },
})
</script>

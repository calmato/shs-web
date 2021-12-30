<template>
  <v-container>
    <v-row class="py-4">
      <v-col cols="12" align="center">
        <v-btn color="primary" @click="onClickNewShift">
          <v-icon>mdi-plus</v-icon>
          <span>シフトを募集する</span>
        </v-btn>
      </v-col>
      <v-col cols="12">
        <h3>募集中のシフト</h3>
        <v-row align="center" class="pa-4">
          <v-col v-for="summary in acceptingSummaries" :key="summary.id" cols="4">
            <the-shift-top-card :summary="summary" @click:new="onClickNewLesson" @click:edit="onClickEditShift" />
          </v-col>
        </v-row>
      </v-col>
      <v-col cols="12">
        <h3>募集終了後のシフト</h3>
        <v-row align="center" class="pa-4">
          <v-col v-for="summary in finishedSummaries" :key="summary.id" cols="4">
            <the-shift-top-card :summary="summary" @click:new="onClickNewLesson" @click:edit="onClickEditShift" />
          </v-col>
        </v-row>
      </v-col>
      <v-col cols="12">
        <h3>募集開始前のシフト</h3>
        <v-row align="center" class="pa-4">
          <v-col v-for="summary in waitingSummaries" :key="summary.id" cols="4">
            <the-shift-top-card :summary="summary" @click:new="onClickNewLesson" @click:edit="onClickEditShift" />
          </v-col>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import TheShiftTopCard from '~/components/organisms/TheShiftTopCard.vue'
import { ShiftSummary } from '~/types/store'

export default defineComponent({
  components: {
    TheShiftTopCard,
  },

  props: {
    acceptingSummaries: {
      type: Array as PropType<ShiftSummary[]>,
      default: () => [],
    },
    finishedSummaries: {
      type: Array as PropType<ShiftSummary[]>,
      default: () => [],
    },
    waitingSummaries: {
      type: Array as PropType<ShiftSummary[]>,
      default: () => [],
    },
  },

  setup(_, { emit }: SetupContext) {
    const onClickNewShift = (): void => {
      emit('click:new-shift')
    }

    const onClickEditShift = (shift: ShiftSummary): void => {
      emit('click:edit-shift', shift)
    }

    const onClickNewLesson = (shift: ShiftSummary): void => {
      emit('click:new-lesson', shift)
    }

    return {
      onClickNewShift,
      onClickEditShift,
      onClickNewLesson,
    }
  },
})
</script>

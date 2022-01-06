<template>
  <v-row>
    <v-col v-for="item in items" :key="item.id" cols="12">
      <v-container>
        {{ items.summary.month }}
      </v-container>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { defineComponent } from '@vue/composition-api'
import { SubmissionEditState } from '~/types/props/submission'

export default defineComponent({
  props: {
    items: {
      type: Object as PropType<SubmissionEditState>,
      default: () => ({}),
    },
  },
})
</script>

<template>
  <div class="stack d-flex align-center">
    <p>{{ id + 1 }}.</p>
    <v-text-field v-model="formData.startAt" type="time" label="開始時刻" step="300" min="17:00" max="21:30" />
    <span>~</span>
    <v-text-field v-model="formData.endAt" type="time" label="終了時刻" step="300" min="17:00" max="21:30" />
    <v-tooltip bottom>
      <template #activator="{ on, attrs }">
        <v-btn class="suffix-btn" color="error" fab elevation="0" v-bind="attrs" v-on="on" @click="handleClick"
          ><v-icon>mdi-minus</v-icon></v-btn
        >
      </template>
      <span>削除</span>
    </v-tooltip>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { HourForm } from '~/types/form'

export default defineComponent({
  props: {
    id: {
      type: Number,
      required: true,
    },
    item: {
      type: Object as PropType<HourForm>,
      default: () => ({
        startAt: '',
        endAt: '',
      }),
    },
  },

  setup(props, { emit }: SetupContext) {
    const formData = computed({
      get: () => props.item,
      set: (val) => emit('update:value', val),
    })

    const handleClick = () => {
      emit('click', props.id)
    }

    return {
      formData,
      handleClick,
    }
  },
})
</script>

<style lang="scss" scoped>
.stack {
  gap: var(--space, 1rem);
}

.suffix-btn {
  height: 24px;
  width: 24px;
}
</style>

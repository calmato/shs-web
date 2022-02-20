<template>
  <v-overlay :value="overlay" :absolute="absolute" class="align-start">
    <v-row align="start" class="py-8 ma-0 mt-4">
      <v-col v-for="item in items" :key="item.name" cols="4" md="3" align="center" class="py-8">
        <v-btn dark icon x-large @click="onClickItem(item)">
          <div class="d-flex flex-column align-center">
            <v-icon class="mb-4">{{ item.icon }}</v-icon>
            <span>{{ item.name }}</span>
          </div>
        </v-btn>
      </v-col>
    </v-row>
  </v-overlay>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { Menu } from '~/types/props/menu'

export default defineComponent({
  props: {
    absolute: {
      type: Boolean,
      default: false,
    },
    items: {
      type: Array as PropType<Menu[]>,
      default: () => [],
    },
    overlay: {
      type: Boolean,
      default: false,
    },
  },

  setup(_, { emit }: SetupContext) {
    const onClickItem = (item: Menu): void => {
      emit('click:item', item)
    }

    const onClickClose = (): void => {
      emit('click:close')
    }

    return {
      onClickItem,
      onClickClose,
    }
  },
})
</script>

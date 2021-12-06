<template>
  <v-row class="py-4">
    <v-col cols="12">
      <div class="text-subtitle-1">ユーザー設定</div>
      <v-card v-for="item in userItems" :key="`user-${item.title}`" elevation="0" class="my-1" @click="onClick(item)">
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
          <v-list-item-action>
            <v-icon>mdi-chevron-right</v-icon>
          </v-list-item-action>
        </v-list-item>
      </v-card>
    </v-col>
    <v-col>
      <div class="text-subtitle-1">教室設定</div>
      <v-card v-for="item in systemItems" :key="`sys-${item.title}`" elevation="0" class="my-1" @click="onClick(item)">
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
          <v-list-item-action>
            <v-icon>mdi-chevron-right</v-icon>
          </v-list-item-action>
        </v-list-item>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { Menu } from '~/types/props/setting'

export default defineComponent({
  props: {
    userItems: {
      type: Array as PropType<Menu[]>,
      default: () => [],
    },
    systemItems: {
      type: Array as PropType<Menu[]>,
      default: () => [],
    },
  },

  setup(_, { emit }: SetupContext) {
    const onClick = (item: Menu): void => {
      emit('click', item)
    }

    return {
      onClick,
    }
  },
})
</script>

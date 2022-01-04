<template>
  <v-navigation-drawer app clipped mobile-breakpoint="960">
    <!-- links -->
    <v-list dense nav shaped>
      <v-list-item-group v-model="selectedItem" color="primary">
        <v-list-item v-for="item in items" :key="item.path" link @click="onClick(item)">
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>{{ item.name }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list-item-group>
    </v-list>
    <!-- footer -->
    <template #append>
      <v-divider />
      <div align="center" class="pa-2 font-weight-thin">&copy; Calmato. All rights reserved.</div>
    </template>
  </v-navigation-drawer>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { Menu } from '~/types/props/menu'

export default defineComponent({
  props: {
    current: {
      type: String,
      default: '',
    },
    items: {
      type: Array as PropType<Menu[]>,
      default: () => [],
    },
  },

  setup(props, { emit }: SetupContext) {
    const target: Menu | undefined = props.items.filter((item: Menu) => item.path === props.current).shift()
    const selectedItem: number = target ? props.items.indexOf(target) : -1

    const onClick = (item: Menu) => {
      emit('click', item)
    }

    return {
      selectedItem,
      onClick,
    }
  },
})
</script>

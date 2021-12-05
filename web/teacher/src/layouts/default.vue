<template>
  <v-app>
    <the-header :overlay="overlay" @click="handleClickMenu" />
    <v-main>
      <nuxt />
      <the-menu
        :overlay="overlay"
        :absolute="true"
        :items="items"
        @click:item="handleClickMenuItem"
        @click:close="handleClickMenu"
      />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent, ref, SetupContext } from '@nuxtjs/composition-api'
import TheHeader from '~/components/organisms/TheHeader.vue'
import TheMenu from '~/components/organisms/TheMenu.vue'
import { Menu } from '~/types/props/menu'

export default defineComponent({
  components: {
    TheHeader,
    TheMenu,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router

    const items: Menu[] = [
      {
        name: 'トップ',
        icon: 'mdi-home',
        path: '/',
      },
      {
        name: 'シフト希望',
        icon: 'mdi-note-edit-outline',
        path: '#シフト希望', // TODO: update
      },
      {
        name: 'シフト作成',
        icon: 'mdi-chart-box-plus-outline',
        path: '/shifts',
      },
      {
        name: 'ユーザ一覧',
        icon: 'mdi-format-list-bulleted',
        path: '#ユーザ一覧', // TODO: update
      },
      {
        name: '設定',
        icon: 'mdi-cogs',
        path: '#設定', // TODO: update
      },
    ]

    const overlay = ref<boolean>(false)

    const handleClickMenu = (): void => {
      overlay.value = !overlay.value
    }

    const handleClickMenuItem = (item: Menu): void => {
      router.push(item.path)
      handleClickMenu()
    }

    return {
      items,
      overlay,
      handleClickMenu,
      handleClickMenuItem,
    }
  },
})
</script>
